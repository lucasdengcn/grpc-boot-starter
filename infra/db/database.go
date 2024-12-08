package db

import (
	"context"
	"database/sql"
	"fmt"
	"grpc-boot-starter/core/config"
	"grpc-boot-starter/core/logging"
	"os"
	"sync"

	"github.com/rs/zerolog/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	onceDB        sync.Once
	db            *gorm.DB
	dbScopedTxKey = "db.scopedTxKey"
)

// ConnectDB db connections
func ConnectDB() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}
	var err error = nil
	onceDB.Do(func() {
		cfg := config.GetConfig()
		conn, err := sql.Open(cfg.DataSource.Driver, cfg.DataSource.URL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		conn.SetMaxIdleConns(cfg.DataSource.PoolMin)
		conn.SetMaxOpenConns(cfg.DataSource.PoolMax)
		//
		err = conn.Ping()
		if err != nil {
			log.Fatal().Err(err).Msg("DB Ping Failed.")
			return
		}
		//
		db0, err := gorm.Open(postgres.New(postgres.Config{
			Conn:                 conn,
			PreferSimpleProtocol: false, // disables implicit prepared statement usage
		}), &gorm.Config{
			PrepareStmt:     true,
			CreateBatchSize: 100,
			Logger:          &GormLogger{},
		})
		log.Info().Msgf("GormLogger: %T", db0.Logger)
		//
		if err != nil {
			return
		}
		//
		db = db0
		log.Info().Msg("DB Connect Successfully.")
	})
	return db, err
}

// GetDBCon provider return db instance
func GetDBCon() *gorm.DB {
	if db == nil {
		dbCon, err := ConnectDB()
		if err != nil {
			log.Fatal().Err(err).Msg("DB Connect Failed.")
		}
		return dbCon
	}
	return db
}

// Close DB connection
func Close() {
	if db != nil {
		// TO DO
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal().Err(err).Msg("DB Close Failed.")
		}
		sqlDB.Close()
	}
	log.Info().Msg("DB Shutdown.")
}

// BeginTx return context
func BeginTx(ctx context.Context) context.Context {
	if db == nil {
		panic("DB not init yet.")
	}
	// attach db to current context, so that caller chain E2E
	tx := db.WithContext(ctx).Begin()
	if tx.Error != nil {
		logging.Panic(ctx).Msgf("Begin Tx Error: %v", tx.Error)
	}
	return context.WithValue(ctx, dbScopedTxKey, tx)
}

// CommitTx return context
func CommitTx(ctx context.Context) {
	dbTx := GetTx(ctx)
	if dbTx == nil {
		logging.Error(ctx).Msgf("tx Commit, but No Tx attached to the context. check the caller chain")
		return
	}
	err := dbTx.Commit()
	if err != nil {
		logging.Error(ctx).Msgf("tx Commit Error. %v", dbTx)
		panic(err)
	}
	logging.Debug(ctx).Msgf("tx Commit Success. %v", dbTx)
	dbTx = nil
}

// RollbackTx return error
func RollbackTx(ctx context.Context) {
	dbTx := GetTx(ctx)
	if dbTx == nil {
		logging.Error(ctx).Msgf("tx Rollback, but No Tx attached to the context. check the caller chain")
		return
	}
	err := dbTx.Rollback()
	if err != nil {
		logging.Error(ctx).Msgf("tx Rollback Error. %v", dbTx)
	} else {
		logging.Debug(ctx).Msgf("tx Rollback Success. %v", dbTx)
	}
}

// GetTx return sqlx.Tx
func GetTx(ctx context.Context) *gorm.DB {
	val := ctx.Value(dbScopedTxKey)
	if val == nil {
		// attach db to current context, so that caller chain E2E
		return db.WithContext(ctx)
	}
	dbTx, ok := val.(*gorm.DB)
	if !ok {
		logging.Panic(ctx).Msgf("Can't Convert Tx object from context. %v", dbTx)
	}
	return dbTx
}

// RecoverErrorHandle, to recover from panic.
func RecoverErrorHandle(ctx context.Context, r any) error {
	if r != nil {
		RollbackTx(ctx)
		if err, ok := r.(error); ok {
			logging.Error(ctx).Msgf("Recover Tx Err: %v", r)
			return err
		} else {
			logging.Error(ctx).Msgf("Recover Tx Err: %v", r)
			return fmt.Errorf("%v", r)
		}
	} else {
		CommitTx(ctx)
		return nil
	}
}

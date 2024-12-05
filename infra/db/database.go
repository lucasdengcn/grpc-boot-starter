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

var onceDB sync.Once

var db *gorm.DB

const txKey string = "txScoped"

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
		})
		//
		if err != nil {
			return
		}
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
func BeginTx(c context.Context) context.Context {
	if db == nil {
		panic("DB not init yet.")
	}
	tx := db.Begin()
	if tx.Error != nil {
		logging.Panic(c).Msgf("Begin Tx Error: %v", tx.Error)
	}
	return context.WithValue(c, txKey, tx)
}

// CommitTx return context
func CommitTx(c context.Context) {
	dbTx := GetTx(c)
	if dbTx == nil {
		return
	}
	err := dbTx.Commit()
	if err != nil {
		logging.Error(c).Msgf("tx Commit Error. %v", dbTx)
		panic(err)
	}
	logging.Debug(c).Msgf("tx Commit Success. %v", dbTx)
	dbTx = nil
}

// RollbackTx return error
func RollbackTx(c context.Context) {
	dbTx := GetTx(c)
	if dbTx == nil {
		return
	}
	err := dbTx.Rollback()
	if err != nil {
		logging.Error(c).Msgf("tx Rollback Error. %v", dbTx)
	} else {
		logging.Debug(c).Msgf("tx Rollback Success. %v", dbTx)
	}
}

// GetTx return sqlx.Tx
func GetTx(c context.Context) *gorm.DB {
	val := c.Value(txKey)
	if val == nil {
		return db
	}
	dbTx, ok := val.(*gorm.DB)
	if !ok {
		logging.Panic(c).Msgf("Can't Convert Tx object from context. %v", dbTx)
	}
	return dbTx
}

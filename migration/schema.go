package migration

import (
	"errors"
	"fmt"
	"grpc-boot-starter/config"

	"github.com/rs/zerolog/log"

	"github.com/golang-migrate/migrate/v4"
	// database driver
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate DB schema
func Migrate() {
	cfg := config.GetConfig()
	path := fmt.Sprintf("file://%s/migration/schemas", cfg.Application.WorkingPath)
	m, err := migrate.New(path, cfg.DataSource.URL)
	defer m.Close()
	if err != nil {
		log.Fatal().Err(err)
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return
		}
		if errors.Is(err, migrate.ErrNilVersion) {
			return
		}
		log.Fatal().Err(err)
	}
	log.Info().Msg("DB schema migrate successfully.")
}

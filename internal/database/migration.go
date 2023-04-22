package database

import (
	"net/http"

	"github.com/eikoshelev/go-telegram-bot-template/migrations"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // driver for migration tool
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type MigrationConfig struct {
	DatabaseURL string `yaml:"databaseUrl"`
}

func Migrate(cfg *MigrationConfig, logger zap.Logger) error {
	source, err := httpfs.New(http.FS(migrations.Static), ".")
	if err != nil {
		return errors.Wrap(err, "can't create file source")
	}

	m, err := migrate.NewWithSourceInstance("httpfs", source, cfg.DatabaseURL)
	if err != nil {
		return errors.Wrap(err, "can't create migration tool")
	}

	defer m.Close()

	ver, dirty, err := m.Version()
	if err != nil {
		logger.Error("can't get DB schema version")
	}

	logger.Info("db scheme state", zap.Any("version", ver), zap.Any("dirty", dirty))

	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Info("no db schema upgrade needed")
			return nil
		}
		return errors.Wrap(err, "can't up schema")
	}
	logger.Info("db schema upgraded")
	return nil
}

package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Config struct {
	ConnectionURI   string        `yaml:"connectionURI"`
	MaxConn         int32         `yaml:"maxConn"`
	MaxConnLifetime time.Duration `yaml:"maxConnLifetime"`
}

func NewPostgresDB(cfg Config) (pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), cfg.ConnectionURI)
	if err != nil {
		return pgxpool.Pool{}, errors.Wrap(err, "failed parse connectionURI for build pool")
	}

	if err := pool.Ping(context.Background()); err != nil {
		return pgxpool.Pool{}, errors.Wrap(err, "failed ping to database")
	}

	return *pool, nil
}

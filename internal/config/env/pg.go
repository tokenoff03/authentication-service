package env

import (
	"errors"
	"os"

	"github.com/tokenoff03/authentication-service/internal/config"
)

const dsnEnvName = "PG_DSN"

type pgConfig struct {
	dsn string
}

func NewPGConfig() (config.PgConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New(dsn)
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (c *pgConfig) DSN() string {
	return c.dsn
}

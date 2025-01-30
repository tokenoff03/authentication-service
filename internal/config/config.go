package config

import (
	"time"

	"github.com/joho/godotenv"
)

type GRPCConfig interface {
	Address() string
}

type PgConfig interface {
	DSN() string
}

type RedisConfig interface {
	DSN() string
}

type TokenConfig interface {
	AccessSecretKey() string
	RefreshSecretKey() string
	AccessTTL() time.Duration
	RefreshTTL() time.Duration
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

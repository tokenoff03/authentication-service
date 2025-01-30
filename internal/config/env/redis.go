package env

import (
	"errors"
	"net"
	"os"
)

const (
	redisHostEnvName = "REDIS_HOST"
	redisPortEnvName = "REDIS_PORT"
)

type redisConfig struct {
	host string
	port string
}

func NewRedisConfig() (*redisConfig, error) {
	host := os.Getenv(redisHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("redis host not found")
	}

	port := os.Getenv(redisPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("redis port not found")
	}

	return &redisConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *redisConfig) DSN() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

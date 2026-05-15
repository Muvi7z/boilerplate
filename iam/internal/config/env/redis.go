package env

import (
	"github.com/caarlos0/env/v11"
	"net"
	"time"
)

type redisEnvConfig struct {
	Host              string        `env:"REDIS_HOST, required"`
	Port              string        `env:"REDIS_PORT, required"`
	ConnectionTimeout time.Duration `env:"REDIS_CONNECTION_TIMEOUT, required"`
	MaxIdle           int           `env:"REDIS_MAX_IDLE, required"`
	IdleTimeout       time.Duration `env:"REDIS_IDLE_TIMEOUT, required"`
	CacheTTL          time.Duration `env:"REDIS_CACHE_TTL, required"`
}

type RedisConfig struct {
	raw redisEnvConfig
}

func NewRedisConfig() (*RedisConfig, error) {
	var raw redisEnvConfig
	err := env.Parse(&raw)
	if err != nil {
		return nil, err
	}

	return &RedisConfig{raw: raw}, nil
}

func (rc *RedisConfig) Address() string {
	return net.JoinHostPort(rc.raw.Host, rc.raw.Port)
}

func (rc *RedisConfig) ConnectionTimeout() time.Duration {
	return rc.raw.ConnectionTimeout
}

func (rc *RedisConfig) MaxIdle() int {
	return rc.raw.MaxIdle
}

func (rc *RedisConfig) IdleTimeout() time.Duration {
	return rc.raw.IdleTimeout
}

func (rc *RedisConfig) CacheTTL() time.Duration {
	return rc.raw.CacheTTL
}

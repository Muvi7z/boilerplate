package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type postgresEnvConfig struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     string `env:"EXTERNAL_REDIS_PORT" envDefault:"5432"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Database string `env:"POSTGRES_DB"`
	SSLMode  string `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
}

type postgresConfig struct {
	raw postgresEnvConfig
}

func NewPostgresConfig() (*postgresConfig, error) {
	var raw postgresEnvConfig
	err := env.Parse(&raw)
	if err != nil {
		return nil, err
	}

	return &postgresConfig{raw: raw}, nil
}

func (cfg *postgresConfig) URI() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.raw.Host,
		cfg.raw.Port,
		cfg.raw.User,
		cfg.raw.Database,
		cfg.raw.Password,
		cfg.raw.SSLMode,
	)
}

func (cfg *postgresConfig) DatabaseName() string {
	return cfg.raw.Database
}

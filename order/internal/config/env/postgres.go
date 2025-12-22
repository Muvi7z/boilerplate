package env

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type postgresEnvConfig struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Database string `env:"POSTGRES_DB"`
	Sslmode  string `env:"POSTGRES_SSL_MODE"`
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
		cfg.raw.Sslmode,
	)
}

func (cfg *postgresConfig) DatabaseName() string {
	return cfg.raw.Database
}

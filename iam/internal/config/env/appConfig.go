package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type appEnvConfig struct {
	Host string `env:"APP_HOST"`
	Port string `env:"APP_PORT"`

	MigrationsDir string `env:"MIGRATIONS_DIR"`
}

type AppConfig struct {
	raw appEnvConfig
}

func NewAppConfig() (*AppConfig, error) {
	var raw appEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &AppConfig{raw: raw}, nil
}

func (cfg *AppConfig) Address() string {
	return net.JoinHostPort(cfg.raw.Host, cfg.raw.Port)
}

func (cfg *AppConfig) MigrationsDir() string {
	return cfg.raw.MigrationsDir
}

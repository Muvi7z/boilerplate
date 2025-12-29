package env

import (
	"github.com/caarlos0/env/v11"
	"net"
)

type appEnvConfig struct {
	Host string `env:"APP_HOST"`
	Port string `env:"APP_PORT"`
}

type appConfig struct {
	raw appEnvConfig
}

func NewAppConfig() (*appConfig, error) {
	var raw appEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &appConfig{raw: raw}, nil
}

func (cfg *appConfig) Address() string {
	return net.JoinHostPort(cfg.raw.Host, cfg.raw.Port)
}

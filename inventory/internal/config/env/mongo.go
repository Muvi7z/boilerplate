package env

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type mongoEnvConfig struct {
	Host     string `env:"MONGO_HOST"`
	Port     string `env:"MONGO_PORT"`
	Database string `env:"MONGO_INITDB_DATABASE"`
	Username string `env:"MONGO_INITDB_ROOT_USERNAME"`
	Password string `env:"MONGO_INITDB_ROOT_PASSWORD"`
	AuthDB   string `env:"MONGO_AUTH_DB"`
}

type mongoConfig struct {
	raw mongoEnvConfig
}

func NewMongoConfig() (*mongoConfig, error) {
	var raw mongoEnvConfig
	err := env.Parse(&raw)
	if err != nil {
		return nil, err
	}

	return &mongoConfig{raw: raw}, nil
}

func (cfg *mongoConfig) URI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s",
		cfg.raw.Username,
		cfg.raw.Password,
		cfg.raw.Host,
		cfg.raw.Port,
		cfg.raw.Database,
		cfg.raw.AuthDB,
	)
}

func (cfg *mongoConfig) DatabaseName() string {
	return cfg.raw.Database
}

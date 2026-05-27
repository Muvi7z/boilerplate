package config

import (
	"os"

	"github.com/Muvi7z/boilerplate/iam/internal/config/env"
	"github.com/joho/godotenv"
)

var appConfig *Config

type Config struct {
	Logger          LoggerConfig
	IAMGRPC         IAMGRPCConfig
	Postgres        PostgresConfig
	Redis           RedisConfig
	AppServerConfig AppServerConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	loggerConfig, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	iamGRPCConfig, err := env.NewIamGRPCConfig()
	if err != nil {
		return err
	}

	postgresConfig, err := env.NewPostgresConfig()
	if err != nil {
		return err
	}

	redisConfig, err := env.NewRedisConfig()
	if err != nil {
		return err
	}

	appServerConfig, err := env.NewAppConfig()
	if err != nil {
		return err
	}

	appConfig = &Config{
		Logger:          loggerConfig,
		IAMGRPC:         iamGRPCConfig,
		Postgres:        postgresConfig,
		Redis:           redisConfig,
		AppServerConfig: appServerConfig,
	}

	return nil
}

func AppConfig() *Config {
	return appConfig
}

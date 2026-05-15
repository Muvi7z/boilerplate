package config

import (
	"github.com/Muvi7z/boilerplate/iam/internal/config/env"
	"github.com/joho/godotenv"
	"os"
)

var appConfig *config

type config struct {
	Logger   LoggerConfig
	IAMGRPC  IAMGRPCConfig
	Postgres PostgresConfig
	Redis    RedisConfig
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

	appConfig = &config{
		Logger:   loggerConfig,
		IAMGRPC:  iamGRPCConfig,
		Postgres: postgresConfig,
		Redis:    redisConfig,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}

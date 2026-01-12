package config

import (
	"github.com/Muvi7z/boilerplate/inventory/internal/config/env"
	"github.com/joho/godotenv"
)

var appConfig *config

type config struct {
	Logger        LoggerConfig
	InventoryGRPC InventoryGRPCConfig
	Mongo         MongoConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil {
		return err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	inventoryCfg, err := env.NewInventoryGRPCConfig()
	if err != nil {
		return err
	}

	mongoCfg, err := env.NewMongoConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:        loggerCfg,
		InventoryGRPC: inventoryCfg,
		Mongo:         mongoCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}

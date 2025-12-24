package config

import (
	"github.com/Muvi7z/boilerplate/order/internal/config/env"
	"github.com/joho/godotenv"
	"os"
)

var appConfig *config

type config struct {
	Logger        LoggerConfig
	Postgres      PostgresConfig
	PaymentGRPC   PaymentGRPCConfig
	InventoryGRPC InventoryGRPCConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	postgresCfg, err := env.NewPostgresConfig()
	if err != nil {
		return err
	}

	paymentGRPCCfg, err := env.NewPaymentGRPCConfig()
	if err != nil {
		return err
	}

	inventoryGRPCCfg, err := env.NewInventoryGRPCConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:        loggerCfg,
		Postgres:      postgresCfg,
		PaymentGRPC:   paymentGRPCCfg,
		InventoryGRPC: inventoryGRPCCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}

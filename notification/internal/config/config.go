package config

import (
	"github.com/Muvi7z/boilerplate/notification/internal/config/env"
	"github.com/joho/godotenv"
)

var appConfig *config

type config struct {
	OrderAssembledConsumerConfig OrderAssembledConsumerConfig
	OrderPaidConsumerConfig      OrderPaidConsumerConfig
	KafkaConfig                  KafkaConfig
	TelegramConfig               TelegramConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil {
		return err
	}

	orderAssembledConsumerConfig, err := env.NewOrderAssembledConsumerConfig()
	if err != nil {
		return err
	}

	orderPaidConsumerConfig, err := env.NewOrderPaidConsumerConfig()
	if err != nil {
		return err
	}

	kafkaConfig, err := env.NewKafkaConfig()
	if err != nil {
		return err
	}

	telegramConfig, err := env.NewTelegramConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		OrderAssembledConsumerConfig: orderAssembledConsumerConfig,
		OrderPaidConsumerConfig:      orderPaidConsumerConfig,
		KafkaConfig:                  kafkaConfig,
		TelegramConfig:               telegramConfig,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}

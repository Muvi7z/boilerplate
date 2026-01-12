package config

import (
	"github.com/Muvi7z/boilerplate/assembly/internal/config/env"
	"github.com/joho/godotenv"
)

var appConfig *config

type config struct {
	Logger                 LoggerConfig
	Kafka                  KafkaConfig
	OrderPaidConsumer      OrderPaidConsumerConfig
	OrderAssembledProducer OrderAssembledProducerConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil {
		return err
	}

	logger, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	kafka, err := env.NewKafkaConfig()
	if err != nil {
		return err
	}

	orderPaidConsumer, err := env.NewOrderPaidConsumerConfig()
	if err != nil {
		return err
	}

	orderAssembledProducer, err := env.NewOrderAssembledProducerConfig()

	appConfig = &config{
		Logger:                 logger,
		Kafka:                  kafka,
		OrderPaidConsumer:      orderPaidConsumer,
		OrderAssembledProducer: orderAssembledProducer,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}

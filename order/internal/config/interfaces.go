package config

import "github.com/IBM/sarama"

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type PostgresConfig interface {
	URI() string
	DatabaseName() string
}

type PaymentGRPCConfig interface {
	Address() string
}

type InventoryGRPCConfig interface {
	Address() string
}

type AppServerConfig interface {
	Address() string
	MigrationsDir() string
}

type KafkaConfig interface {
	Brokers() []string
}

type OrderAssembledConsumerConfig interface {
	Topic() string
	GroupId() string
	Config() *sarama.Config
}

type OrderPaidProducerConfig interface {
	Topic() string
	Config() *sarama.Config
}

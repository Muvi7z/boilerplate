package config

import "github.com/IBM/sarama"

type KafkaConfig interface {
	Brokers() []string
}

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type OrderPaidConsumerConfig interface {
	Topic() string
	GroupId() string
	Config() *sarama.Config
}

type OrderAssembledProducerConfig interface {
	Topic() string
	Config() *sarama.Config
}

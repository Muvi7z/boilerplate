package config

import "github.com/IBM/sarama"

type OrderPaidConsumerConfig interface {
	Topic() string
	GroupId() string
	Config() *sarama.Config
}

type OrderAssembledConsumerConfig interface {
	Topic() string
	GroupId() string
	Config() *sarama.Config
}

type KafkaConfig interface {
	Brokers() []string
}

type TelegramConfig interface {
	Token() string
}

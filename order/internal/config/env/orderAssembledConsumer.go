package env

import (
	"github.com/IBM/sarama"
	"github.com/caarlos0/env/v11"
)

type orderAssembledConsumerEnvConfig struct {
	Topic   string `env:"ORDER_ASSEMBLED_TOPIC_NAME"`
	GroupId string `env:"ORDER_ASSEMBLED_GROUP_ID"`
}

type orderAssembledConsumerConfig struct {
	raw *orderAssembledConsumerEnvConfig
}

func NewOrderAssembledConsumerConfig() (*orderAssembledConsumerConfig, error) {
	var raw orderAssembledConsumerEnvConfig

	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &orderAssembledConsumerConfig{raw: &raw}, nil
}

func (cfg *orderAssembledConsumerConfig) Topic() string {
	return cfg.raw.Topic
}

func (cfg *orderAssembledConsumerConfig) GroupId() string {
	return cfg.raw.GroupId
}

func (cfg *orderAssembledConsumerConfig) Config() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V4_0_0_0
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	return config
}

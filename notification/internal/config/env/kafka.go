package env

import "github.com/caarlos0/env/v11"

type kafkaEnvConfig struct {
	Brokers []string `env:"ORDER_KAFKA_BROKERS"`
}

type kafkaConfig struct {
	raw *kafkaEnvConfig
}

func NewKafkaConfig() (*kafkaConfig, error) {
	var raw kafkaEnvConfig
	err := env.Parse(&raw)
	if err != nil {
		return nil, err
	}

	return &kafkaConfig{
		raw: &raw,
	}, nil
}

func (cfg *kafkaConfig) Brokers() []string {
	return cfg.raw.Brokers
}

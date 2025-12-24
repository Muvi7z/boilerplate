package env

import (
	"github.com/caarlos0/env/v11"
	"net"
)

type paymentEnvConfig struct {
	Host string `env:"PAYMENT_GRPC_HOST, required"`
	Port string `env:"PAYMENT_GRPC_POST, required"`
}

type paymentGRPCConfig struct {
	raw paymentEnvConfig
}

func NewPaymentGRPCConfig() (*paymentGRPCConfig, error) {
	var raw paymentEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &paymentGRPCConfig{raw: raw}, nil
}

func (cfg *paymentGRPCConfig) Address() string {
	return net.JoinHostPort(cfg.raw.Host, cfg.raw.Port)
}

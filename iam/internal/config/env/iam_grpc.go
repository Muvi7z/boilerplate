package env

import (
	"github.com/caarlos0/env/v11"
	"net"
)

type iamEnvConfig struct {
	Host string `env:"IAM_GRPC_HOST"`
	Port string `env:"IAM_GRPC_PORT"`
}

type iamGrpcConfig struct {
	raw iamEnvConfig
}

func NewIamGRPCConfig() (*iamGrpcConfig, error) {
	var raw iamEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}
	return &iamGrpcConfig{raw: raw}, nil
}

func (cfg *iamGrpcConfig) Address() string {
	return net.JoinHostPort(cfg.raw.Host, cfg.raw.Port)
}

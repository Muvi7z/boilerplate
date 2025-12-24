package env

import (
	"github.com/caarlos0/env/v11"
	"net"
)

type inventoryEnvConfig struct {
	Host string `env:"INVENTORY_GRPC_HOST, required"`
	Port string `env:"INVENTORY_GRPC_POST, required"`
}

type inventoryGRPCConfig struct {
	raw inventoryEnvConfig
}

func NewInventoryGRPCConfig() (*inventoryGRPCConfig, error) {
	var raw inventoryEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &inventoryGRPCConfig{raw: raw}, nil
}

func (cfg *inventoryGRPCConfig) Address() string {
	return net.JoinHostPort(cfg.raw.Host, cfg.raw.Port)
}

package config

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type PostgresConfig interface {
}

type PaymentGRPCConfig interface {
	Address() string
}

type InventoryGRPCConfig interface {
	Address() string
}

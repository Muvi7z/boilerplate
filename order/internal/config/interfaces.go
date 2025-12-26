package config

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
}

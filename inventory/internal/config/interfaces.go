package config

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type MongoConfig interface {
	URI() string
	DatabaseName() string
}

type InventoryGRPCConfig interface {
	Address() string
}

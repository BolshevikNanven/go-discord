package config

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewConfig,
	NewDatabase,
	NewSnowflakeConfig,
	NewRedisConfig,
)

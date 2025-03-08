package config

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewConfig,
	NewEtcdConfig,
	NewRedisConfig,
	NewDatabaseConfig,
	NewRocketMQConfig,
	NewSnowflakeConfig,
)

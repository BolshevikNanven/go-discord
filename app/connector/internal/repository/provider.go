package repository

import (
	"discord/data"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewUserRepository,
	NewChannelRepository,
	data.NewRedis,
)

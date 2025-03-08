package repository

import (
	"discord/data"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewSpaceRepo,
	NewChannelRepo,
	NewSpaceUserRepo,
	NewChannelUserRepo,
	data.NewDatabase,
	data.NewRedis,
)

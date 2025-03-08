package repository

import (
	"discord/data"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewInboxRepository,
	NewUserRepository,
	NewChannelRepository,
	NewMessageRepository,
	data.NewRedis,
	data.NewDatabase,
)

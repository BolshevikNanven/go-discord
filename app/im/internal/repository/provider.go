package repository

import (
	"discord/data"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewMqRepository,
	NewMessageRepository,
	NewInboxRepository,
	data.NewDatabase,
	data.NewRedis,
	data.NewRocketMQProducer,
)

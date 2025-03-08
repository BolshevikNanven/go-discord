package repository

import (
	"discord/data"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	data.NewDatabase,
	NewUserRepo,
)

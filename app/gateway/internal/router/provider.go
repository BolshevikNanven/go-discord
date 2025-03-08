package router

import (
	v1 "discord/app/gateway/internal/router/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	v1.NewRouter,
	wire.Bind(new(Router), new(*v1.Router)),
)

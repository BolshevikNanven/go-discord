//go:build wireinject

package main

import (
	"discord/app/gateway/internal"
	"discord/app/gateway/internal/client"
	"discord/app/gateway/internal/config"
	"discord/app/gateway/internal/handler"
	"discord/app/gateway/internal/middleware"
	"discord/app/gateway/internal/router"

	"github.com/google/wire"
)

func initApp() chan<- struct{} {
	panic(wire.Build(
		internal.NewHTTPServer,
		router.ProviderSet,
		client.ProviderSet,
		handler.ProviderSet,
		middleware.ProviderSet,
		config.ProviderSet,
		internal.NewLogger,
		newApp,
	))
}

//go:build wireinject

package main

import (
	"discord/app/push/internal"
	"discord/app/push/internal/client"
	"discord/app/push/internal/config"
	"discord/app/push/internal/repository"
	"discord/data"
	"sync"

	"github.com/google/wire"
)

func runApp(wg *sync.WaitGroup) chan struct{} {
	panic(wire.Build(
		client.ProviderSet,
		config.ProviderSet,
		repository.ProviderSet,
		internal.NewServer,
		internal.NewLogger,
		data.NewRocketMQConsumer,
		newApp,
	))
}

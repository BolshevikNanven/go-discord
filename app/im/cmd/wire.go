//go:build wireinject

package main

import (
	"discord/app/im/internal"
	"discord/app/im/internal/client"
	"discord/app/im/internal/config"
	"discord/app/im/internal/repository"
	"discord/pkg/snowflakeutil"
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
		snowflakeutil.New,
		newApp,
	))
}

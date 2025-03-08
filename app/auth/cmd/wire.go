//go:build wireinject

package main

import (
	"discord/app/auth/internal"
	"discord/app/auth/internal/config"
	"discord/app/auth/internal/repository"
	"discord/pkg/snowflakeutil"
	"sync"

	"github.com/google/wire"
)

func runApp(wg *sync.WaitGroup) chan struct{} {
	panic(wire.Build(
		config.ProviderSet,
		snowflakeutil.New,
		repository.ProviderSet,
		internal.NewServer,
		internal.NewLogger,
		newApp,
	))
}

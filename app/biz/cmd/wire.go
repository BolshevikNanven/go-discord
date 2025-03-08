//go:build wireinject
// +build wireinject

package main

import (
	"discord/app/biz/internal"
	"discord/app/biz/internal/config"
	"discord/app/biz/internal/repository"
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
		newApp,
	))
}

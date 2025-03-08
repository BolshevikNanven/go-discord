//go:build wireinject
// +build wireinject

package main

import (
	"discord/app/connector/internal"
	"discord/app/connector/internal/client"
	"discord/app/connector/internal/config"
	"discord/app/connector/internal/hub"
	"discord/app/connector/internal/repository"
	"sync"

	"github.com/google/wire"
)

func runApp(wg *sync.WaitGroup) chan struct{} {
	panic(wire.Build(
		config.ProviderSet,
		repository.ProviderSet,
		hub.NewHub,
		internal.ProviderSet,
		client.ProviderSet,
		newApp,
	))
}

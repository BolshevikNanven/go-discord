// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"discord/app/auth/internal"
	"discord/app/auth/internal/config"
	"discord/app/auth/internal/repository"
	"discord/data"
	"discord/pkg/snowflakeutil"
	"sync"
)

// Injectors from wire.go:

func runApp(wg *sync.WaitGroup) chan struct{} {
	configConfig := config.NewConfig()
	databaseConfig := config.NewDatabase(configConfig)
	db := data.NewDatabase(databaseConfig)
	snowflakeutilConfig := config.NewSnowflakeConfig(configConfig)
	node := snowflakeutil.New(snowflakeutilConfig)
	userRepository := repository.NewUserRepo(db, node)
	jwtutilConfig := config.NewJwtConfig(configConfig)
	logger := internal.NewLogger()
	authServiceServer := internal.NewServer(userRepository, jwtutilConfig, logger)
	v := newApp(wg, authServiceServer, configConfig)
	return v
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"discord/app/gateway/internal"
	"discord/app/gateway/internal/client"
	"discord/app/gateway/internal/config"
	"discord/app/gateway/internal/handler"
	"discord/app/gateway/internal/middleware"
	"discord/app/gateway/internal/router/v1"
)

// Injectors from wire.go:

func initApp() chan<- struct{} {
	logger := internal.NewLogger()
	configConfig := config.NewConfig()
	authClientPool := client.NewAuthClientPool()
	authHandler := handler.NewAuthHandler(authClientPool)
	jwtutilConfig := config.NewJWTConfig(configConfig)
	authMiddleware := middleware.NewAuthMiddleware(jwtutilConfig)
	bizClientPool := client.NewBizClientPool()
	bizHandler := handler.NewBizHandler(bizClientPool)
	imClientPool := client.NewImClientPool()
	imHandler := handler.NewImHandler(imClientPool)
	router := v1.NewRouter(authHandler, authMiddleware, bizHandler, imHandler)
	app := internal.NewHTTPServer(logger, router)
	v := newApp(logger, configConfig, app)
	return v
}

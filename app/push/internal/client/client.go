package client

import (
	"discord/pkg/discovery"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/resolver"
)

var (
	keepaliveParams = keepalive.ClientParameters{
		Time:                90 * time.Second,
		Timeout:             3 * time.Second,
		PermitWithoutStream: true,
	}
)

func Register(logger *zap.Logger, config *discovery.EtcdConfig) {
	builder, _ := discovery.NewBuilder(logger, config.Address)
	resolver.Register(builder)
}

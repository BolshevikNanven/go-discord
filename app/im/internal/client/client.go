package client

import (
	"discord/pkg/discovery"
	"discord/pkg/sre"
	"fmt"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func newConn(name string) (*grpc.ClientConn, error) {
	slide := sre.NewGoogleSlide(6*time.Second, 2)
	conn, err := grpc.NewClient(
		fmt.Sprintf("etcd:///%s", name),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepaliveParams),
		grpc.WithChainUnaryInterceptor(slide.Sre()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

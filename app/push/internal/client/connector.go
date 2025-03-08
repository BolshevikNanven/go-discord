package client

import (
	"discord/api/connector"
	"discord/pkg/discovery"
	"discord/pkg/grpcpool"
	"fmt"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ConnectorClientPool struct {
	poolMap *sync.Map
}

func NewConnectorClientPool(logger *zap.Logger, config *discovery.EtcdConfig) *ConnectorClientPool {
	return &ConnectorClientPool{
		poolMap: &sync.Map{},
	}
}

func (c *ConnectorClientPool) Get(connectorId string) connector.ConnectorServiceClient {
	pool, ok := c.poolMap.Load(connectorId)
	if !ok {
		pool = grpcpool.New(func() (*grpc.ClientConn, error) {
			return newConn(connectorId)
		}, 3)
		c.poolMap.Store(connectorId, pool)
	}

	conn, err := pool.(*grpcpool.Pool).Get()
	if err != nil {
		return nil
	}

	return connector.NewConnectorServiceClient(conn)
}

func newConn(connector string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		fmt.Sprintf("etcd:///%s", connector),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepaliveParams),
	)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

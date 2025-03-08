package client

import (
	"discord/api/im"
	"discord/pkg/grpcpool"

	"google.golang.org/grpc"
)

type ImClientPool struct {
	pool *grpcpool.Pool
}

func NewImClientPool() *ImClientPool {
	pool := grpcpool.New(func() (*grpc.ClientConn, error) {
		return newConn("im")
	}, 3)

	return &ImClientPool{
		pool: pool,
	}
}

func (c *ImClientPool) Get() im.ImServiceClient {
	conn, _ := c.pool.Get()

	return im.NewImServiceClient(conn)
}

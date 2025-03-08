package client

import (
	"discord/api/biz"
	"discord/pkg/grpcpool"

	"google.golang.org/grpc"
)

type BizClientPool struct {
	pool *grpcpool.Pool
}

func NewBizClientPool() *BizClientPool {
	pool := grpcpool.New(func() (*grpc.ClientConn, error) {
		return newConn("biz")
	}, 3)

	return &BizClientPool{
		pool: pool,
	}
}

func (c *BizClientPool) Get() biz.BizServiceClient {
	conn, _ := c.pool.Get()

	return biz.NewBizServiceClient(conn)
}

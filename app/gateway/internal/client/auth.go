package client

import (
	"discord/api/auth"
	"discord/pkg/grpcpool"

	"google.golang.org/grpc"
)

type AuthClientPool struct {
	pool *grpcpool.Pool
}

func NewAuthClientPool() *AuthClientPool {
	pool := grpcpool.New(func() (*grpc.ClientConn, error) {
		return newConn("auth")
	}, 3)

	return &AuthClientPool{
		pool: pool,
	}
}

func (c *AuthClientPool) Get() auth.AuthServiceClient {
	conn, _ := c.pool.Get()

	return auth.NewAuthServiceClient(conn)
}

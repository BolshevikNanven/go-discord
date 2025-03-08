package grpcpool

import (
	"fmt"
	"sync"
	"sync/atomic"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type ConnectFunc func() (*grpc.ClientConn, error)

type Pool struct {
	mu   sync.RWMutex
	idx  atomic.Int64
	size int

	connFunc ConnectFunc
	conns    []*grpc.ClientConn
}

func New(connFunc ConnectFunc, size int) *Pool {
	pool := &Pool{
		mu:       sync.RWMutex{},
		connFunc: connFunc,
		conns:    make([]*grpc.ClientConn, size),
		size:     size,
	}

	return pool
}

func (p *Pool) Get() (*grpc.ClientConn, error) {
	p.mu.RLock()
	index := p.idx.Add(1) % int64(p.size)
	conn := p.conns[index]
	p.mu.RUnlock()

	if conn != nil && conn.GetState() != connectivity.Shutdown {
		return conn, nil
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	// 重复检查
	conn = p.conns[index]
	if conn != nil && conn.GetState() != connectivity.Shutdown {
		return conn, nil
	}

	newConn, err := p.connFunc()
	if err != nil {
		fmt.Println("create new conn failed", err)
		return nil, err
	}
	p.conns[index] = newConn

	return newConn, nil
}

package discovery

import (
	"context"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc/resolver"
)

type EtcdResolver struct {
	cli    *clientv3.Client
	logger *zap.Logger

	closeChan chan struct{}
	watchChan clientv3.WatchChan

	cc     resolver.ClientConn
	prefix string
}

func NewEtcdResolver(cli *clientv3.Client, logger *zap.Logger, cc resolver.ClientConn, prefix string) *EtcdResolver {
	return &EtcdResolver{
		cli:    cli,
		logger: logger,
		cc:     cc,
		prefix: prefix,
	}
}

func (et *EtcdResolver) ResolveNow(options resolver.ResolveNowOptions) {}
func (et *EtcdResolver) Close() {
	et.closeChan <- struct{}{}
}
func (et *EtcdResolver) Daemon() {
	et.sync()
	et.watchChan = et.cli.Watch(context.Background(), et.prefix, clientv3.WithPrefix())

	for {
		select {
		case <-et.closeChan:
			return
		case _, ok := <-et.watchChan:
			if ok {
				et.sync()
			}
		}
	}
}

func (et *EtcdResolver) sync() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	resp, err := et.cli.Get(ctx, et.prefix, clientv3.WithPrefix())
	if err != nil {
		et.logger.Error(err.Error())
		return
	}

	address := []resolver.Address{}
	for _, kv := range resp.Kvs {
		et.logger.Debug(kv.String())
		address = append(address, resolver.Address{Addr: string(kv.Value)})
	}

	et.cc.UpdateState(resolver.State{Addresses: address})
}

package discovery

import (
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc/resolver"
)

type Builder struct {
	cli    *clientv3.Client
	logger *zap.Logger
}

func NewBuilder(logger *zap.Logger, addr ...string) (*Builder, error) {
	var (
		err     error
		builder = &Builder{
			logger: logger,
		}
	)

	if builder.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 3 * time.Second,
	}); err != nil {
		return nil, err
	}

	return builder, nil
}

func (b *Builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	service := target.Endpoint()
	etcdResolver := NewEtcdResolver(b.cli, b.logger, cc, fmt.Sprintf("/%s/", service))

	go etcdResolver.Daemon()

	return etcdResolver, nil

}
func (et *Builder) Scheme() string {
	return "etcd"
}

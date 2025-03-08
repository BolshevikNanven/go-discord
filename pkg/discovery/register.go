package discovery

import (
	"context"
	"errors"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Register struct {
	etcdAddr    []string
	dialTimeout int
	ttl         int64

	srv       Server
	closeChan chan struct{}
	aliveChan <-chan *clientv3.LeaseKeepAliveResponse

	cli     *clientv3.Client
	leaseID clientv3.LeaseID
}

func NewRegister(addr ...string) *Register {
	return &Register{
		etcdAddr:    addr,
		dialTimeout: 3,
		ttl:         10,
	}
}

func (r *Register) Register(srv Server) error {
	var err error

	if strings.Split(srv.Addr, ":")[0] == "" {
		return errors.New("invaild server addr")
	}

	if r.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   r.etcdAddr,
		DialTimeout: time.Duration(r.dialTimeout) * time.Second,
	}); err != nil {
		return err
	}

	r.srv = srv

	if err = r.register(); err != nil {
		return err
	}

	r.closeChan = make(chan struct{})
	go r.daemon()

	return nil
}

func (r *Register) Stop() {
	r.closeChan <- struct{}{}
}

func (r *Register) register() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.dialTimeout)*time.Second)
	defer cancel()

	resp, err := r.cli.Grant(ctx, r.ttl)
	if err != nil {
		return err
	}
	r.leaseID = resp.ID

	if r.aliveChan, err = r.cli.KeepAlive(context.Background(), r.leaseID); err != nil {
		return err
	}

	_, err = r.cli.Put(context.Background(), r.srv.GetPath(), r.srv.Addr, clientv3.WithLease(r.leaseID))

	return err
}

func (r *Register) unregister() error {
	_, err := r.cli.Delete(context.Background(), r.srv.GetPath())
	return err
}

func (r *Register) daemon() {
	for {
		select {
		case <-r.closeChan:
			r.unregister()
			r.cli.Revoke(context.Background(), r.leaseID)
			return
		case alive := <-r.aliveChan:
			if alive == nil {
				r.register()
			}
		}
	}
}

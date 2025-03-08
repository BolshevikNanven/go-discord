package main

import (
	"bufio"
	"discord/api/im"
	"discord/app/im/internal/client"
	"discord/app/im/internal/config"
	"discord/pkg/discovery"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	closeChan := runApp(&wg)

	bufio.NewReader(os.Stdin).ReadString('\n')

	close(closeChan)

	fmt.Println("im service closed!")
}

func newApp(wg *sync.WaitGroup, logger *zap.Logger, service im.ImServiceServer, conf *config.Config) chan struct{} {
	client.Register(logger, conf.Etcd)

	address := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	register := discovery.NewRegister(conf.Etcd.Address)

	if err := register.Register(discovery.Server{
		Addr: address,
		Name: "im",
	}); err != nil {
		panic(fmt.Sprintf("failed to register service: %v", err))
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	srv := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             60 * time.Second,
			PermitWithoutStream: true,
		}),
	)
	im.RegisterImServiceServer(srv, service)

	closeChan := make(chan struct{})
	go func() {
		fmt.Println("im server started!")
		if err := srv.Serve(listener); err != nil {
			panic(fmt.Sprintf("failed to serve: %v", err))
		}
	}()

	go func() {
		<-closeChan
		register.Stop()
		listener.Close()
		srv.GracefulStop()
		wg.Done()
	}()

	return closeChan
}

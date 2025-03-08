package main

import (
	"bufio"
	"discord/api/auth"
	"discord/app/auth/internal/config"
	"discord/pkg/discovery"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	closeChan := runApp(&wg)

	bufio.NewReader(os.Stdin).ReadString('\n')

	close(closeChan)

	fmt.Println("auth service closed!")
}

func newApp(wg *sync.WaitGroup, service auth.AuthServiceServer, conf *config.Config) chan struct{} {
	address := fmt.Sprintf("%s:%s", conf.Host, conf.Port)

	register := discovery.NewRegister(conf.Etcd.Address)
	if err := register.Register(discovery.Server{
		Addr: address,
		Name: "auth",
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
	auth.RegisterAuthServiceServer(srv, service)

	closeChan := make(chan struct{})

	go func() {
		fmt.Println("auth server started!")
		if err := srv.Serve(listener); err != nil {
			fmt.Printf("server error: %v\n", err)
		}

		srv.GracefulStop()
		listener.Close()
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

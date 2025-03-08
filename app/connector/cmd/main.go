package main

import (
	"bufio"
	"discord/api/connector"
	"discord/app/connector/internal"
	"discord/app/connector/internal/client"
	"discord/app/connector/internal/config"
	"discord/app/connector/internal/hub"
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
	wg := &sync.WaitGroup{}
	wg.Add(1)
	closeChan := runApp(wg)

	bufio.NewReader(os.Stdin).ReadString('\n')

	close(closeChan)
	wg.Wait()
	fmt.Println("connector service closed!")
}

func newApp(
	wg *sync.WaitGroup,
	logger *zap.Logger,
	conf *config.Config,
	hub *hub.Hub,
	WebSocketServer *internal.WebSocketServer,
	rpcServer connector.ConnectorServiceServer,
) chan struct{} {
	address := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	websocketAddress := fmt.Sprintf("%s:%s", conf.Websocket.Host, conf.Websocket.Port)
	client.Register(logger, &conf.Etcd)

	register := discovery.NewRegister(conf.Etcd.Address)
	if err := register.Register(discovery.Server{
		Addr: address,
		Name: fmt.Sprintf("connector-%s", conf.Name),
	}); err != nil {
		panic(fmt.Sprintf("failed to register service: %v", err))
	}

	// 启动websocket服务
	go hub.Run()
	WebSocketServer.Start(websocketAddress)

	// 启动rpc服务
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
	connector.RegisterConnectorServiceServer(srv, rpcServer)
	go func() {
		if err := srv.Serve(listener); err != nil {
			panic(fmt.Sprintf("failed to serve: %v", err))
		}
	}()

	fmt.Printf("rpc server started at: %s\n", address)
	fmt.Printf("websocket server started at: %s\n", websocketAddress)
	closeChan := make(chan struct{})

	go func() {
		<-closeChan
		register.Stop()
		WebSocketServer.Stop()
		srv.Stop()

		wg.Done()
	}()

	return closeChan
}

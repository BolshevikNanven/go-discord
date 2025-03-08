package main

import (
	"bufio"
	"discord/app/push/internal"
	"discord/app/push/internal/client"
	"discord/app/push/internal/config"
	"discord/pkg/discovery"
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	closeChan := runApp(&wg)

	bufio.NewReader(os.Stdin).ReadString('\n')

	close(closeChan)

	fmt.Println("push service closed!")
}

func newApp(wg *sync.WaitGroup, logger *zap.Logger, server *internal.Server, conf *config.Config) chan struct{} {
	client.Register(logger, conf.Etcd)

	address := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	register := discovery.NewRegister(conf.Etcd.Address)

	if err := register.Register(discovery.Server{
		Addr: address,
		Name: "push",
	}); err != nil {
		panic(fmt.Sprintf("failed to register service: %v", err))
	}

	closeChan := make(chan struct{})

	go func() {
		fmt.Println("push server started!")
		server.Run()
	}()

	go func() {
		<-closeChan
		register.Stop()
		server.Stop()
		wg.Done()
	}()

	return closeChan
}

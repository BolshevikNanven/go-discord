package main

import (
	"bufio"
	"discord/app/gateway/internal/client"
	"discord/app/gateway/internal/config"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	closeChan := initApp()

	bufio.NewReader(os.Stdin).ReadString('\n')

	close(closeChan)

	fmt.Println("gateway service closed!")
}

func newApp(logger *zap.Logger, conf *config.Config, engine *fiber.App) chan<- struct{} {
	closeChan := make(chan struct{})
	client.Register(logger, conf.Etcd)

	go func() {
		if err := engine.Listen(fmt.Sprintf("%s:%s", conf.Host, conf.Port)); err != nil {
			fmt.Printf("服务器启动错误: %v\n", err)
			close(closeChan)
			return
		}
	}()

	go func() {
		<-closeChan

		if err := engine.Shutdown(); err != nil {
			fmt.Printf("服务器关闭错误: %v\n", err)
		} else {
			fmt.Println("服务器已成功关闭")
		}
	}()

	return closeChan
}

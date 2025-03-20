package hub

import (
	"discord/app/connector/internal/client"
	"discord/app/connector/internal/config"
	"discord/app/connector/internal/repository"
	"discord/pkg/epoller"
	"fmt"
	"net"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"go.uber.org/zap"
)

type ClientMessage struct {
	UserId  int64
	Message []byte
}

type Hub struct {
	epoll       epoller.Poller
	clients     map[int64]*Client
	connClients map[int]*Client

	connectorId string

	channelRepository repository.ChannelRepository
	userRepository    repository.UserRepository
	bizClientPool     *client.BizClientPool

	registerChan   chan *Client
	unregisterChan chan *Client
	sendChan       chan *ClientMessage

	closeChan chan struct{}

	logger *zap.Logger
}

func NewHub(
	conf *config.Config,
	userRepository repository.UserRepository,
	channelRepository repository.ChannelRepository,
	bizClientPool *client.BizClientPool,
	logger *zap.Logger,
) *Hub {
	ep, err := epoller.NewPoller(1024)
	if err != nil {
		panic("epoll new error!")
	}
	return &Hub{
		connectorId:       fmt.Sprintf("connector-%s", conf.Name),
		epoll:             ep,
		clients:           make(map[int64]*Client),
		connClients:       make(map[int]*Client),
		registerChan:      make(chan *Client),
		unregisterChan:    make(chan *Client),
		userRepository:    userRepository,
		channelRepository: channelRepository,
		bizClientPool:     bizClientPool,
		closeChan:         make(chan struct{}),
		sendChan:          make(chan *ClientMessage, 256),
		logger:            logger,
	}
}

func (h *Hub) Run() {
	go h.readPump()
	go h.writePump()
	go h.heartPump()

	for {
		select {
		case client := <-h.registerChan:
			h.clients[client.userId] = client
			h.connClients[client.fd] = client
			h.epoll.Add(client.conn)

		case client := <-h.unregisterChan:
			if _, ok := h.clients[client.userId]; ok {
				client.DeleteUserConnector()
				_ = h.epoll.Remove(client.conn)
				client.conn.Close()

				delete(h.connClients, client.fd)
				delete(h.clients, client.userId)
			}

		case <-h.closeChan:
			goto END
		}
	}
END:
	return
}

func (h *Hub) Serve(userId int64, conn net.Conn) {
	client := NewClient(userId, conn, h)

	h.registerChan <- client
	h.logger.Info(fmt.Sprintf("user %d connected, client %d", userId, client.fd))
}

func (h *Hub) SendMessage(userId int64, message []byte) bool {
	client, ok := h.clients[userId]
	if !ok {
		// redis记录不一致，删除redis记录
		_ = h.userRepository.DeleteUserConnector(client.spaceId, userId)
		return false
	}

	h.sendChan <- &ClientMessage{UserId: userId, Message: message}
	return true
}

func (h *Hub) Close() {
	for userId, client := range h.clients {
		// 直接删除
		_ = h.userRepository.DeleteUserConnector(client.spaceId, userId)
		delete(h.clients, userId)
	}

	close(h.sendChan)
	close(h.closeChan)
	close(h.registerChan)
	close(h.unregisterChan)
}

func (h *Hub) readPump() {
	for {
		connections, err := h.epoll.Wait(128)
		if err != nil {
			h.logger.Error("epoll wait error", zap.Error(err))
			continue
		}
		h.logger.Info(fmt.Sprintf("epoll wait %d connections", len(connections)))
		for _, conn := range connections {
			if conn == nil {
				break
			}

			fd := epoller.SocketFD(conn)
			if client, ok := h.connClients[fd]; ok {
				client.Receive()
			} else {
				h.logger.Error(fmt.Sprintf("client %d not found", fd))
				h.epoll.Remove(conn)
			}

		}
	}
}

func (h *Hub) writePump() {
	for {
		select {
		case msg := <-h.sendChan:
			if client, ok := h.clients[msg.UserId]; ok {
				if err := wsutil.WriteServerMessage(client.conn, ws.OpText, msg.Message); err != nil {
					h.unregisterChan <- client
				}
			}
		case <-h.closeChan:
			return
		}
	}
}

func (h *Hub) heartPump() {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		<-ticker.C
		for _, client := range h.clients {
			if time.Since(client.lastPong) > pingPeriod+pongWait {
				h.unregisterChan <- client
				continue
			}
			if err := wsutil.WriteServerMessage(client.conn, ws.OpPing, nil); err != nil {
				h.unregisterChan <- client
			}
		}
	}
}

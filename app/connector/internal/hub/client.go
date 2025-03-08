package hub

import (
	"discord/app/connector/internal/model"
	"encoding/json"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Client struct {
	spaceId int64
	userId  int64
	conn    *websocket.Conn

	sendChan chan []byte
	hub      *Hub
}

func ServeClient(userId int64, conn *websocket.Conn, hub *Hub) {
	client := &Client{
		userId:   userId,
		conn:     conn,
		hub:      hub,
		sendChan: make(chan []byte, 256),
	}

	client.hub.registerChan <- client
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregisterChan <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, originalMessage, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		var message model.Message
		if err := json.Unmarshal(originalMessage, &message); err != nil {
			continue
		}

		switch message.Type {
		case model.MessageTypeSwitchSpace:
			spaceId, err := strconv.ParseInt(message.Data, 10, 64)
			if err != nil {
				c.sendChan <- []byte(err.Error())
				continue
			}

			if err := c.hub.updateUserConnector(c, spaceId); err != nil {
				c.sendChan <- []byte(err.Error())
				continue
			}

			c.spaceId = spaceId

			resp := model.Message{
				Type: model.MessageTypeSwitchSpace,
				Data: strconv.FormatInt(spaceId, 10),
			}
			jsonResp, _ := json.Marshal(resp)
			c.sendChan <- jsonResp

		default:
			continue
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.sendChan:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))

			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}

	}
}

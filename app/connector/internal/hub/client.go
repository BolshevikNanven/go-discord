package hub

import (
	"context"
	"discord/api/biz"
	"discord/app/connector/internal/model"
	"discord/pkg/epoller"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

const (
	// Time allowed to read the next pong message from the peer.
	pongWait = 30 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
	spaceId int64
	userId  int64
	conn    net.Conn
	fd      int

	// 限流
	limiter *rate.Limiter
	// 心跳
	lastPong time.Time

	hub *Hub
}

func NewClient(userId int64, conn net.Conn, hub *Hub) *Client {
	return &Client{
		userId:   userId,
		conn:     conn,
		fd:       epoller.SocketFD(conn),
		hub:      hub,
		limiter:  rate.NewLimiter(rate.Limit(3), 128),
		lastPong: time.Now(),
	}
}

func (c *Client) Receive() {
	msg, op, err := wsutil.ReadClientData(c.conn)
	if err != nil {
		c.hub.logger.Error("read client data error", zap.Error(err))
		c.hub.unregisterChan <- c
		return
	}
	if !c.limiter.Allow() {
		c.hub.logger.Error("client sent too fast")
		c.hub.unregisterChan <- c
		return
	}

	c.hub.logger.Info(fmt.Sprintf("read client data, type: %d", op))
	switch op {
	case ws.OpPong:
		c.lastPong = time.Now()
	case ws.OpText:
		go c.handleMessage(msg)
	default:
		c.hub.unregisterChan <- c
	}
}

func (c *Client) handleMessage(data []byte) {
	var message model.Message
	if err := json.Unmarshal(data, &message); err != nil {
		return
	}

	switch message.Type {
	case model.MessageTypeSwitchSpace:
		spaceId, err := strconv.ParseInt(message.Data, 10, 64)
		if err != nil {
			c.hub.sendChan <- &ClientMessage{
				UserId:  c.userId,
				Message: []byte(err.Error()),
			}
			return
		}

		if err := c.UpdateUserConnector(); err != nil {
			c.hub.sendChan <- &ClientMessage{
				UserId:  c.userId,
				Message: []byte(err.Error()),
			}
			return
		}

		c.spaceId = spaceId

		resp := model.Message{
			Type: model.MessageTypeSwitchSpace,
			Data: strconv.FormatInt(spaceId, 10),
		}
		jsonResp, _ := json.Marshal(resp)
		c.hub.sendChan <- &ClientMessage{
			UserId:  c.userId,
			Message: jsonResp,
		}

	default:
		return
	}
}

func (c *Client) UpdateUserConnector() error {
	var (
		hub     = c.hub
		spaceId = c.spaceId
		userId  = c.userId
	)
	// 检查用户是否是空间成员
	resp, err := hub.bizClientPool.Get().IsSpaceMember(context.Background(), &biz.IsSpaceMemberRequest{
		SpaceId: spaceId,
		UserId:  userId,
	})
	if err != nil {
		return err
	}
	if !resp.IsMember {
		return nil
	}

	// 移动用户connector
	if err := hub.userRepository.MoveUserConnector(userId, hub.connectorId, spaceId, spaceId); err != nil {
		return err
	}

	// 首次连接
	if spaceId == 0 {
		resp, err := hub.bizClientPool.Get().GetChannelIds(context.Background(), &biz.GetChannelIdsRequest{
			SpaceId: spaceId,
			UserId:  userId,
		})
		if err != nil {
			return err
		}

		if len(resp.ChannelIds) > 0 {
			if err := hub.channelRepository.AddChannelConnectors(userId, resp.ChannelIds, hub.connectorId); err != nil {
				return err
			}
		}

		return nil
	}

	// 移动用户频道connector
	var (
		prevChannelIds []int64
		newChannelIds  []int64
	)
	if resp, err := hub.bizClientPool.Get().GetChannelIds(context.Background(), &biz.GetChannelIdsRequest{
		SpaceId: spaceId,
		UserId:  userId,
	}); err != nil {
		return err
	} else {
		prevChannelIds = resp.ChannelIds
	}
	if resp, err := hub.bizClientPool.Get().GetChannelIds(context.Background(), &biz.GetChannelIdsRequest{
		SpaceId: spaceId,
		UserId:  userId,
	}); err != nil {
		return err
	} else {
		newChannelIds = resp.ChannelIds
	}

	if err := hub.channelRepository.MoveChannelConnectors(userId, prevChannelIds, newChannelIds, hub.connectorId); err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteUserConnector() {
	var (
		hub     = c.hub
		spaceId = c.spaceId
		userId  = c.userId
	)
	_ = hub.userRepository.DeleteUserConnector(spaceId, userId)

	resp, err := hub.bizClientPool.Get().GetChannelIds(context.Background(), &biz.GetChannelIdsRequest{
		SpaceId: spaceId,
		UserId:  userId,
	})
	if err != nil {
		return
	}
	if len(resp.ChannelIds) == 0 {
		return
	}

	for _, channelId := range resp.ChannelIds {
		_ = hub.channelRepository.DeleteChannelConnector(channelId, userId, hub.connectorId)
	}

}

package hub

import (
	"context"
	"discord/api/biz"
	"discord/app/connector/internal/client"
	"discord/app/connector/internal/config"
	"discord/app/connector/internal/repository"
	"fmt"
)

type Hub struct {
	clients map[int64]*Client

	connectorId string

	channelRepository repository.ChannelRepository
	userRepository    repository.UserRepository
	bizClientPool     *client.BizClientPool

	registerChan   chan *Client
	unregisterChan chan *Client

	closeChan chan struct{}
}

func NewHub(
	conf *config.Config,
	userRepository repository.UserRepository,
	channelRepository repository.ChannelRepository,
	bizClientPool *client.BizClientPool,
) *Hub {
	return &Hub{
		connectorId:       fmt.Sprintf("connector-%s", conf.Name),
		clients:           make(map[int64]*Client),
		registerChan:      make(chan *Client),
		unregisterChan:    make(chan *Client),
		userRepository:    userRepository,
		channelRepository: channelRepository,
		bizClientPool:     bizClientPool,
		closeChan:         make(chan struct{}),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.registerChan:
			h.clients[client.userId] = client

			go client.readPump()
			go client.writePump()

		case client := <-h.unregisterChan:
			if _, ok := h.clients[client.userId]; ok {
				h.deleteUserConnector(client)

				close(client.sendChan)
				delete(h.clients, client.userId)
			}

		case <-h.closeChan:
			goto END
		}
	}
END:
	return
}

func (h *Hub) SendMessage(userId int64, message []byte) bool {
	client, ok := h.clients[userId]
	if !ok {
		// redis记录不一致，删除redis记录
		_ = h.userRepository.DeleteUserConnector(client.spaceId, userId)
		return false
	}

	client.sendChan <- message
	return true
}

func (h *Hub) Close() {
	for userId, client := range h.clients {
		_ = h.userRepository.DeleteUserConnector(client.spaceId, userId)
		delete(h.clients, userId)

		close(client.sendChan)
	}

	close(h.closeChan)
	close(h.registerChan)
	close(h.unregisterChan)
}

func (h *Hub) updateUserConnector(client *Client, spaceId int64) error {
	// 检查用户是否是空间成员
	resp, err := h.bizClientPool.Get().IsSpaceMember(context.Background(), &biz.IsSpaceMemberRequest{
		SpaceId: spaceId,
		UserId:  client.userId,
	})
	if err != nil {
		return err
	}
	if !resp.IsMember {
		return nil
	}

	// 移动用户connector
	if err := h.userRepository.MoveUserConnector(client.userId, h.connectorId, client.spaceId, spaceId); err != nil {
		return err
	}

	// 首次连接
	if client.spaceId == 0 {
		resp, err := h.bizClientPool.Get().GetChannelIds(context.Background(), &biz.GetChannelIdsRequest{
			SpaceId: spaceId,
			UserId:  client.userId,
		})
		if err != nil {
			return err
		}

		if len(resp.ChannelIds) > 0 {
			if err := h.channelRepository.AddChannelConnectors(client.userId, resp.ChannelIds, h.connectorId); err != nil {
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
	if resp, err := h.bizClientPool.Get().GetChannelIds(context.Background(), &biz.GetChannelIdsRequest{
		SpaceId: client.spaceId,
		UserId:  client.userId,
	}); err != nil {
		return err
	} else {
		prevChannelIds = resp.ChannelIds
	}
	if resp, err := h.bizClientPool.Get().GetChannelIds(context.Background(), &biz.GetChannelIdsRequest{
		SpaceId: spaceId,
		UserId:  client.userId,
	}); err != nil {
		return err
	} else {
		newChannelIds = resp.ChannelIds
	}

	if err := h.channelRepository.MoveChannelConnectors(client.userId, prevChannelIds, newChannelIds, h.connectorId); err != nil {
		return err
	}

	return nil
}

func (h *Hub) deleteUserConnector(client *Client) {
	_ = h.userRepository.DeleteUserConnector(client.spaceId, client.userId)

	resp, err := h.bizClientPool.Get().GetChannelIds(context.Background(), &biz.GetChannelIdsRequest{
		SpaceId: client.spaceId,
		UserId:  client.userId,
	})
	if err != nil {
		return
	}
	if len(resp.ChannelIds) == 0 {
		return
	}

	for _, channelId := range resp.ChannelIds {
		_ = h.channelRepository.DeleteChannelConnector(channelId, client.userId, h.connectorId)
	}

}

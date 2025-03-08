package internal

import (
	"context"
	"discord/api/connector"
	"discord/app/push/internal/client"
	"discord/app/push/internal/model"
	"discord/app/push/internal/repository"
	"encoding/json"
	"strconv"
	"time"

	rocketmq "github.com/apache/rocketmq-clients/golang/v5"
	"go.uber.org/zap"
)

var (
	_maxReceiveCount   int32 = 5
	_invisibleDuration       = 20 * time.Second
)

type Server struct {
	logger              *zap.Logger
	connectorClientPool *client.ConnectorClientPool
	consumer            rocketmq.SimpleConsumer

	inboxRepo   repository.InboxRepository
	userRepo    repository.UserRepository
	channelRepo repository.ChannelRepository
	messageRepo repository.MessageRepository
}

func NewServer(logger *zap.Logger,
	connectorClientPool *client.ConnectorClientPool,
	consumer rocketmq.SimpleConsumer,
	inboxRepo repository.InboxRepository,
	userRepo repository.UserRepository,
	channelRepo repository.ChannelRepository,
	messageRepo repository.MessageRepository,
) *Server {
	return &Server{
		logger:              logger,
		connectorClientPool: connectorClientPool,
		consumer:            consumer,
		inboxRepo:           inboxRepo,
		userRepo:            userRepo,
		channelRepo:         channelRepo,
		messageRepo:         messageRepo,
	}
}

func (s *Server) Run() {
	for {
		msgs, err := s.consumer.Receive(context.Background(), _maxReceiveCount, _invisibleDuration)
		if err != nil {
			// s.logger.Error("receive message failed", zap.Error(err))
			continue
		}

		var (
			tag         string
			chatMessage model.ChatMessage
		)
		for _, msg := range msgs {
			tag = *msg.GetTag()
			err = json.Unmarshal(msg.GetBody(), &chatMessage)
			if err != nil {
				s.logger.Error("unmarshal message failed", zap.Error(err))
				continue
			}

			switch tag {
			case "peer":
				err = s.handlePeerMessage(&chatMessage)
				if err != nil {
					s.logger.Error("handle peer message failed", zap.Error(err))
					continue
				}
			case "channel":
				err = s.handleChannelMessage(&chatMessage)
				if err != nil {
					s.logger.Error("handle channel message failed", zap.Error(err))
					continue
				}
			}

			s.consumer.Ack(context.Background(), msg)
		}

	}
}

func (s *Server) Stop() {
	s.consumer.GracefulStop()
}

func (s *Server) handlePeerMessage(msg *model.ChatMessage) error {
	err := s.messageRepo.Save(msg)
	if err != nil {
		s.logger.Error("save message failed", zap.Error(err))
		return err
	}
	err = s.inboxRepo.Save(msg.From, msg.SpaceId, msg.Id)
	if err != nil {
		s.logger.Error("save inbox failed", zap.Error(err))
		return err
	}

	connectorId := s.userRepo.GetUserConnector(msg.SpaceId, msg.From)
	if connectorId == "" {
		return nil
	}

	message, err := json.Marshal(msg)
	if err != nil {
		s.logger.Error("marshal message failed", zap.Error(err))
		return err
	}

	client := s.connectorClientPool.Get(connectorId)
	if client == nil {
		return nil
	}

	_, err = client.SendMessage(context.Background(), &connector.SendMessageRequest{
		UserId:  msg.From,
		Message: string(message),
	})
	if err != nil {
		s.logger.Error("send message failed", zap.Error(err))
		return err
	}

	return nil
}

func (s *Server) handleChannelMessage(msg *model.ChatMessage) error {
	err := s.messageRepo.Save(msg)
	if err != nil {
		return err
	}
	err = s.inboxRepo.UpdateChannelCurrent(msg.To, msg.Id)
	if err != nil {
		return err
	}

	connectors := s.channelRepo.GetUsersConnector(msg.To)
	if len(connectors) == 0 {
		return nil
	}
	message, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	for user, connectorId := range connectors {
		client := s.connectorClientPool.Get(connectorId)
		if client == nil {
			continue
		}

		userId, err := strconv.ParseInt(user, 10, 64)
		if err != nil {
			return err
		}
		_, err = client.SendMessage(context.Background(), &connector.SendMessageRequest{
			UserId:  userId,
			Message: string(message),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

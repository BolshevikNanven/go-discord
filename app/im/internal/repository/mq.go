package repository

import (
	"context"
	"discord/app/im/internal/model"
	"encoding/json"

	rocketmq "github.com/apache/rocketmq-clients/golang/v5"
	"go.uber.org/zap"
)

type MqRepository interface {
	SendChatMessage(message *model.ChatMessage, isChannel bool) error
}

type mqRepository struct {
	producer rocketmq.Producer
	logger   *zap.Logger
}

func NewMqRepository(producer rocketmq.Producer, logger *zap.Logger) MqRepository {
	return &mqRepository{
		producer: producer,
		logger:   logger,
	}
}

func (r *mqRepository) SendChatMessage(message *model.ChatMessage, isChannel bool) error {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := rocketmq.Message{
		Topic: "im",
		Body:  jsonMessage,
	}
	if isChannel {
		t := "channel"
		msg.Tag = &t
	} else {
		t := "peer"
		msg.Tag = &t
	}

	_, err = r.producer.Send(context.Background(), &msg)
	if err != nil {
		return err
	}

	return nil
}

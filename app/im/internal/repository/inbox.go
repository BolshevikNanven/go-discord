package repository

import (
	"context"
	"discord/app/im/internal/model"
	"discord/data"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type InboxRepository interface {
	RemoveInboxMessages(userId int64, spaceId int64, messageIds []int64) error
	UpdateChannelAckMessageId(channelId int64, userId int64, messageId int64) error
	GetInboxMessages(userId int64, spaceId int64, limit int) ([]*model.ChatMessage, error)
	GetChannelInbox(channelId int64, userId int64) (current int64, last int64, err error)
}

type inboxRepository struct {
	cache *redis.Client
	db    *gorm.DB
}

func NewInboxRepository(cache *redis.Client, db *gorm.DB) InboxRepository {
	return &inboxRepository{cache: cache, db: db}
}

func (r *inboxRepository) GetInboxMessages(userId int64, spaceId int64, limit int) ([]*model.ChatMessage, error) {
	key := fmt.Sprintf(data.KeyFormatInbox, spaceId, userId)

	values, err := r.cache.LRange(context.Background(), key, 0, int64(limit-1)).Result()
	if err == redis.Nil {
		return []*model.ChatMessage{}, nil
	}
	if err != nil {
		return nil, err
	}

	// 将字符串转换为int64
	var messageIds []int64
	for _, value := range values {
		messageId, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse message id error: %w", err)
		}
		messageIds = append(messageIds, messageId)
	}

	// 从数据库中批量查询消息
	var messages []*model.ChatMessage
	err = r.db.Where("id IN ?", messageIds).Find(&messages).Error
	if err != nil {
		return nil, fmt.Errorf("query messages error: %w", err)
	}

	return messages, nil
}

func (r *inboxRepository) RemoveInboxMessages(userId int64, spaceId int64, messageIds []int64) error {
	key := fmt.Sprintf(data.KeyFormatInbox, spaceId, userId)
	pipe := r.cache.TxPipeline()

	for _, messageId := range messageIds {
		pipe.LRem(context.Background(), key, 0, strconv.FormatInt(messageId, 10))
	}

	_, err := pipe.Exec(context.Background())
	return err

}

func (r *inboxRepository) UpdateChannelAckMessageId(channelId int64, userId int64, messageId int64) error {
	key := fmt.Sprintf(data.KeyFormatChannelAck, channelId, userId)
	strMessageId := strconv.FormatInt(messageId, 10)

	return r.cache.Set(context.Background(), key, strMessageId, 0).Err()
}

func (r *inboxRepository) GetChannelInbox(channelId int64, userId int64) (int64, int64, error) {
	key := fmt.Sprintf(data.KeyFormatChannelCurrent, channelId)
	currentStr, err := r.cache.Get(context.Background(), key).Result()
	current, _ := strconv.ParseInt(currentStr, 10, 64)
	if err == redis.Nil {
		return 0, 0, nil
	}
	if err != nil {
		return 0, 0, err
	}

	key = fmt.Sprintf(data.KeyFormatChannelAck, channelId, userId)
	lastStr, err := r.cache.Get(context.Background(), key).Result()
	last, _ := strconv.ParseInt(lastStr, 10, 64)
	if err == redis.Nil {
		return last, 0, nil
	}
	if err != nil {
		return 0, 0, err
	}

	return current, last, nil
}

package repository

import (
	"context"
	"discord/data"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type InboxRepository interface {
	Save(userId, spaceId, messageId int64) error
	UpdateChannelCurrent(channelId, messageId int64) error
}

type inboxRepository struct {
	cache *redis.Client
}

func NewInboxRepository(cache *redis.Client) InboxRepository {
	return &inboxRepository{
		cache: cache,
	}
}

func (r *inboxRepository) Save(userId, spaceId, messageId int64) error {
	key := fmt.Sprintf(data.KeyFormatInbox, spaceId, userId)
	strMessageId := strconv.FormatInt(messageId, 10)

	return r.cache.RPush(context.Background(), key, strMessageId).Err()
}

func (r *inboxRepository) UpdateChannelCurrent(channelId, messageId int64) error {
	key := fmt.Sprintf(data.KeyFormatChannelCurrent, channelId)
	strMessageId := strconv.FormatInt(messageId, 10)

	return r.cache.Set(context.Background(), key, strMessageId, 0).Err()
}

package repository

import (
	"context"
	"discord/data"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type ChannelRepository interface {
	GetUsersConnector(channelId int64) map[string]string
}

type channelRepository struct {
	cache *redis.Client
}

func NewChannelRepository(cache *redis.Client) ChannelRepository {
	return &channelRepository{
		cache: cache,
	}
}

func (r *channelRepository) GetUsersConnector(channelId int64) map[string]string {
	key := fmt.Sprintf(data.KeyFormatChannelUserConnector, channelId)

	users, err := r.cache.HGetAll(context.Background(), key).Result()
	if err != nil {
		return map[string]string{}
	}

	return users

}

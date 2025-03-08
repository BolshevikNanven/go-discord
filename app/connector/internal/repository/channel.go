package repository

import (
	"context"
	"discord/data"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type ChannelRepository interface {
	AddChannelConnectors(userId int64, channelIds []int64, connectorId string) error
	DeleteChannelConnector(channelId int64, userId int64, connectorId string) error
	MoveChannelConnectors(userId int64, prevChannelIds []int64, newChannelIds []int64, connectorId string) error
}

type channelRepository struct {
	cache *redis.Client
}

func NewChannelRepository(cache *redis.Client) ChannelRepository {
	return &channelRepository{
		cache: cache,
	}
}

func (r *channelRepository) AddChannelConnectors(userId int64, channelIds []int64, connectorId string) error {
	pipe := r.cache.TxPipeline()
	userIdStr := strconv.FormatInt(userId, 10)

	for _, channelId := range channelIds {
		pipe.HSet(
			context.Background(),
			fmt.Sprintf(data.KeyFormatChannelUserConnector, channelId),
			map[string]string{
				userIdStr: connectorId,
			},
		)

	}

	_, err := pipe.Exec(context.Background())
	return err
}

func (r *channelRepository) DeleteChannelConnector(channelId int64, userId int64, connectorId string) error {
	key := fmt.Sprintf(data.KeyFormatChannelUserConnector, channelId)
	userIdStr := strconv.FormatInt(userId, 10)

	return r.cache.HDel(
		context.Background(),
		key,
		userIdStr,
	).Err()
}

func (r *channelRepository) MoveChannelConnectors(userId int64, prevChannelIds []int64, newChannelIds []int64, connectorId string) error {
	pipe := r.cache.TxPipeline()
	userIdStr := strconv.FormatInt(userId, 10)

	for _, channelId := range prevChannelIds {
		pipe.HDel(
			context.Background(),
			fmt.Sprintf(data.KeyFormatChannelUserConnector, channelId),
			userIdStr,
		)
	}
	for _, channelId := range newChannelIds {
		pipe.HSet(
			context.Background(),
			fmt.Sprintf(data.KeyFormatChannelUserConnector, channelId),
			map[string]string{
				userIdStr: connectorId,
			},
		)

	}

	_, err := pipe.Exec(context.Background())
	return err
}

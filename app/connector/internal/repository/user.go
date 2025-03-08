package repository

import (
	"context"
	"discord/data"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type UserRepository interface {
	SetUserConnector(spaceId int64, userId int64, connectorId string) error
	DeleteUserConnector(spaceId int64, userId int64) error
	MoveUserConnector(userId int64, connectorId string, prevSpaceId int64, newSpaceId int64) error
}

type userRepository struct {
	cache *redis.Client
}

func NewUserRepository(cache *redis.Client) UserRepository {
	return &userRepository{
		cache: cache,
	}
}

func (r *userRepository) SetUserConnector(spaceId int64, userId int64, connectorId string) error {
	userIdStr := strconv.FormatInt(userId, 10)
	// 设置用户connector
	err := r.cache.HSet(
		context.Background(),
		fmt.Sprintf(data.KeyFormatUserConnector, spaceId),
		map[string]string{
			userIdStr: connectorId,
		},
	).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUserConnector(spaceId int64, userId int64) error {
	userIdStr := strconv.FormatInt(userId, 10)

	return r.cache.HDel(
		context.Background(),
		fmt.Sprintf(data.KeyFormatUserConnector, spaceId),
		userIdStr,
	).Err()
}

func (r *userRepository) MoveUserConnector(userId int64, connectorId string, prevSpaceId int64, newSpaceId int64) error {
	userIdStr := strconv.FormatInt(userId, 10)
	pipe := r.cache.TxPipeline()

	prevKey := fmt.Sprintf(data.KeyFormatUserConnector, prevSpaceId)
	newKey := fmt.Sprintf(data.KeyFormatUserConnector, newSpaceId)

	pipe.HDel(context.Background(), prevKey, userIdStr)
	pipe.HSet(context.Background(), newKey, map[string]string{
		userIdStr: connectorId,
	})

	_, err := pipe.Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

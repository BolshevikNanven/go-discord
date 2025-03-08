package repository

import (
	"context"
	"discord/data"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type UserRepository interface {
	GetUserConnector(spaceId int64, userId int64) string
}

type userRepository struct {
	cache *redis.Client
}

func NewUserRepository(cache *redis.Client) UserRepository {
	return &userRepository{
		cache: cache,
	}
}

func (r *userRepository) GetUserConnector(spaceId int64, userId int64) string {
	keys := fmt.Sprintf(data.KeyFormatUserConnector, spaceId)
	userIdStr := strconv.FormatInt(userId, 10)

	return r.cache.HGet(context.Background(), keys, userIdStr).Val()
}

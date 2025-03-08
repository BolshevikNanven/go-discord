package repository

import (
	"context"
	"discord/app/biz/internal/model"
	"discord/data"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const (
	channelCacheExpiration = 24 * time.Hour   // 缓存过期时间
	channelLockExpiration  = 10 * time.Second // 锁过期时间
)

type ChannelUserRepository interface {
	IsChannelMember(channelId int64, userId int64) (bool, error)
	Create(channelId int64, userId int64) error
	Delete(channelId int64, userId int64) error
	CacheUserChannels(userId int64) error
}

type channelUserRepository struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewChannelUserRepo(db *gorm.DB, cache *redis.Client) ChannelUserRepository {
	return &channelUserRepository{db: db, cache: cache}
}

func (r *channelUserRepository) IsChannelMember(channelId int64, userId int64) (bool, error) {
	key := fmt.Sprintf(data.KeyFormatUserChannel, userId)

	// 1. 查询缓存
	if r.cache.Exists(context.Background(), key).Val() > 0 {
		isMember, err := r.cache.SIsMember(context.Background(), key, channelId).Result()
		if err != nil && err != redis.Nil {
			return false, err
		}
		if err == nil && isMember {
			return true, nil
		}
	}

	// 2. 缓存不存在，尝试获取分布式锁
	lockKey := fmt.Sprintf(data.KeyFormatUserChannelLock, userId)
	locked, err := r.cache.SetNX(context.Background(), lockKey, "1", channelLockExpiration).Result()
	if err != nil {
		return false, err
	}

	// 3. 如果获取到锁，则重建缓存
	if locked {
		defer r.cache.Del(context.Background(), lockKey)
		if err := r.CacheUserChannels(userId); err != nil {
			return false, err
		}

		return r.cache.SIsMember(context.Background(), key, channelId).Result()
	} else {
		time.Sleep(100 * time.Millisecond)
	}

	// 4. 重新检查缓存
	if r.cache.Exists(context.Background(), key).Val() > 0 {
		isMember, err := r.cache.SIsMember(context.Background(), key, channelId).Result()
		if err != nil && err != redis.Nil {
			return false, err
		}
		if err == nil && isMember {
			return true, nil
		}
	}

	// 5. 直接查询数据库
	channelUser := model.ChannelUser{
		ChannelId: channelId,
		UserId:    userId,
	}
	err = r.db.Take(&channelUser).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}

func (r *channelUserRepository) Create(channelId int64, userId int64) error {
	channelUser := model.ChannelUser{
		ChannelId: channelId,
		UserId:    userId,
	}

	// 全量删除缓存
	key := fmt.Sprintf(data.KeyFormatUserChannel, userId)
	if err := r.cache.Del(context.Background(), key).Err(); err != nil {
		return err
	}

	if err := r.db.Create(&channelUser).Error; err != nil {
		return err
	}

	return nil
}

func (r *channelUserRepository) Delete(channelId int64, userId int64) error {
	// 全量删除缓存
	key := fmt.Sprintf(data.KeyFormatUserChannel, userId)
	if err := r.cache.Del(context.Background(), key).Err(); err != nil {
		return err
	}

	if err := r.db.Where("channel_id = ? AND user_id = ?", channelId, userId).Delete(&model.ChannelUser{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *channelUserRepository) CacheUserChannels(userId int64) error {
	ctx := context.Background()
	key := fmt.Sprintf(data.KeyFormatUserChannel, userId)

	var channelIds []int64
	if err := r.db.Model(&model.ChannelUser{}).
		Where("user_id = ?", userId).
		Pluck("channel_id", &channelIds).Error; err != nil {
		return err
	}

	// 使用pipeline重建缓存
	pipe := r.cache.Pipeline()
	pipe.Del(ctx, key)
	if len(channelIds) > 0 {
		strChannelIds := make([]string, len(channelIds))
		for i, channelId := range channelIds {
			strChannelIds[i] = strconv.FormatInt(channelId, 10)
		}
		pipe.SAdd(ctx, key, strChannelIds)
	}
	pipe.Expire(ctx, key, channelCacheExpiration)
	_, err := pipe.Exec(ctx)
	return err
}

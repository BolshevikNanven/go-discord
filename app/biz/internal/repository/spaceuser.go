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
	cacheExpiration = 24 * time.Hour   // 缓存过期时间
	lockExpiration  = 10 * time.Second // 锁过期时间
)

type SpaceUserRepository interface {
	IsSpaceMember(spaceId int64, userId int64) (bool, error)
	Create(spaceId int64, userId int64) error
	Delete(spaceId int64, userId int64) error
	CacheUserSpaceIds(userId int64) error
}

type spaceUserRepository struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewSpaceUserRepo(db *gorm.DB, cache *redis.Client) SpaceUserRepository {
	return &spaceUserRepository{db: db, cache: cache}
}

func (r *spaceUserRepository) IsSpaceMember(spaceId int64, userId int64) (bool, error) {
	key := fmt.Sprintf(data.KeyFormatUserSpace, userId)

	// 1. 查询缓存
	if r.cache.Exists(context.Background(), key).Val() > 0 {
		isMember, err := r.cache.SIsMember(context.Background(), key, spaceId).Result()
		if err != nil && err != redis.Nil {
			return false, err
		}
		if err == nil && isMember {
			return true, nil
		}
	}

	// 2. 缓存不存在，尝试获取分布式锁
	lockKey := fmt.Sprintf(data.KeyFormatUserSpaceLock, userId)
	locked, err := r.cache.SetNX(context.Background(), lockKey, "1", lockExpiration).Result()
	if err != nil {
		return false, err
	}

	// 3. 如果获取到锁，则重建缓存
	if locked {
		defer r.cache.Del(context.Background(), lockKey)
		if err := r.CacheUserSpaceIds(userId); err != nil {
			return false, err
		}

		return r.cache.SIsMember(context.Background(), key, spaceId).Result()
	} else {
		time.Sleep(100 * time.Millisecond)
	}

	// 4. 如果没有获取到锁，先重新检查缓存
	if r.cache.Exists(context.Background(), key).Val() > 0 {
		isMember, err := r.cache.SIsMember(context.Background(), key, spaceId).Result()
		if err != nil && err != redis.Nil {
			return false, err
		}
		if err == nil {
			return isMember, nil
		}
	}

	// 5. 缓存仍不存在，才查询数据库
	spaceUser := model.SpaceUser{
		SpaceId: spaceId,
		UserId:  userId,
	}
	err = r.db.Take(&spaceUser).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}

func (r *spaceUserRepository) Create(spaceId int64, userId int64) error {
	spaceUser := model.SpaceUser{
		SpaceId: spaceId,
		UserId:  userId,
	}

	// 全量删除缓存
	key := fmt.Sprintf(data.KeyFormatUserSpace, userId)
	if err := r.cache.Del(context.Background(), key).Err(); err != nil {
		return err
	}

	if err := r.db.Create(&spaceUser).Error; err != nil {
		return err
	}

	return nil
}

func (r *spaceUserRepository) Delete(spaceId int64, userId int64) error {
	// 全量删除缓存
	key := fmt.Sprintf(data.KeyFormatUserSpace, userId)
	if err := r.cache.Del(context.Background(), key).Err(); err != nil {
		return err
	}

	if err := r.db.Where("space_id = ? AND user_id = ?", spaceId, userId).Delete(&model.SpaceUser{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *spaceUserRepository) CacheUserSpaceIds(userId int64) error {
	ctx := context.Background()
	key := fmt.Sprintf(data.KeyFormatUserSpace, userId)

	var spaceIds []int64
	if err := r.db.Model(&model.SpaceUser{}).
		Where("user_id = ?", userId).
		Pluck("space_id", &spaceIds).Error; err != nil {
		return err
	}

	// 使用pipeline重建缓存
	pipe := r.cache.Pipeline()
	pipe.Del(ctx, key)
	if len(spaceIds) > 0 {
		strSpaceIds := make([]string, len(spaceIds))
		for i, spaceId := range spaceIds {
			strSpaceIds[i] = strconv.FormatInt(spaceId, 10)
		}
		pipe.SAdd(ctx, key, strSpaceIds)
	}
	pipe.Expire(ctx, key, cacheExpiration)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return err
}

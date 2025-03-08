package repository

import (
	"discord/app/biz/internal/model"
	"discord/pkg/snowflakeutil"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ChannelRepository interface {
	GetByID(channelId int64) (model.Channel, error)
	Create(channel model.Channel) (model.Channel, error)
	List(spaceId int64, userId int64) ([]model.Channel, error)
	All(spaceId int64, userId int64) ([]int64, error)
	Update(channel model.Channel) (model.Channel, error)
	Delete(channelId int64) error
}

type channelRepository struct {
	db        *gorm.DB
	snowflake *snowflakeutil.Node
	cache     *redis.Client
}

func NewChannelRepo(db *gorm.DB, snowflake *snowflakeutil.Node, cache *redis.Client) ChannelRepository {
	return &channelRepository{
		db:        db,
		snowflake: snowflake,
		cache:     cache,
	}
}

func (c *channelRepository) GetByID(channelId int64) (model.Channel, error) {
	var channel model.Channel
	return channel, c.db.Model(&model.Channel{}).Where("id = ?", channelId).First(&channel).Error
}

func (c *channelRepository) Create(channel model.Channel) (model.Channel, error) {
	channel.Id = c.snowflake.GenerateID()

	err := c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&channel).Error; err != nil {
			return err
		}

		if err := tx.Create(&model.ChannelUser{
			ChannelId: channel.Id,
			UserId:    channel.Owner,
		}).Error; err != nil {
			return err
		}

		return nil
	})

	return channel, err
}

// 获取用户空间下的已加入或公开的频道
func (c *channelRepository) List(spaceId int64, userId int64) ([]model.Channel, error) {
	var channels []model.Channel

	result := c.db.Model(&model.Channel{}).
		Joins("JOIN channel_user ON channel.id = channel_user.channel_id").
		Where("channel.space_id = ? AND (channel_user.user_id = ? OR channel.type = 'PUBLIC')", spaceId, userId).
		Find(&channels)

	if result.Error != nil {
		return nil, result.Error
	}

	return channels, nil
}

func (c *channelRepository) All(spaceId int64, userId int64) ([]int64, error) {
	var channelIds []int64

	result := c.db.Model(&model.Channel{}).
		Joins("JOIN channel_user ON channel.id = channel_user.channel_id").
		Where("channel.space_id = ? AND (channel_user.user_id = ? OR channel.type = 'PUBLIC')", spaceId, userId).
		Pluck("channel.id", &channelIds)

	if result.Error != nil {
		return nil, result.Error
	}

	return channelIds, nil
}

func (c *channelRepository) Update(channel model.Channel) (model.Channel, error) {
	return channel, c.db.Model(&model.Channel{}).Where("id = ?", channel.Id).Updates(channel).Error
}

func (c *channelRepository) Delete(channelId int64) error {
	return c.db.Model(&model.Channel{}).Where("id = ?", channelId).Delete(&model.Channel{}).Error
}

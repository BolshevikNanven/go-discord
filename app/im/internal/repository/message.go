package repository

import (
	"discord/app/im/internal/model"

	"gorm.io/gorm"
)

type MessageRepository interface {
	GetMessages(userId, spaceId, from int64, cursor int64, limit int) ([]*model.ChatMessage, error)
	GetChannelMessages(channelId, spaceId int64, cursor int64, limit int) ([]*model.ChatMessage, error)
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{
		db: db,
	}
}

func (r *messageRepository) GetMessages(userId, spaceId, from int64, cursor int64, limit int) ([]*model.ChatMessage, error) {
	var messages []*model.ChatMessage
	err := r.db.Model(&model.ChatMessage{}).
		Where("space_id = ?", spaceId).
		Where("id > ?", cursor).
		Where("(from = ? AND to = ? ) OR (from = ? AND to = ?)", from, userId, userId, from).
		Order("id DESC").
		Limit(limit).
		Find(&messages).Error
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *messageRepository) GetChannelMessages(channelId, spaceId int64, cursor int64, limit int) ([]*model.ChatMessage, error) {
	var messages []*model.ChatMessage
	err := r.db.Model(&model.ChatMessage{}).
		Where("space_id = ?", spaceId).
		Where("id > ?", cursor).
		Where("to = ?", channelId).
		Order("id DESC").
		Limit(limit).
		Find(&messages).Error
	if err != nil {
		return nil, err
	}

	return messages, nil
}

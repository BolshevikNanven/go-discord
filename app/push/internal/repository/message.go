package repository

import (
	"discord/app/push/internal/model"

	"gorm.io/gorm"
)

type MessageRepository interface {
	Save(message *model.ChatMessage) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{
		db: db,
	}
}

func (r *messageRepository) Save(message *model.ChatMessage) error {
	err := r.db.Create(message).Error
	if err != nil {
		return err
	}
	return nil
}

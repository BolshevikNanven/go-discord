package model

import "time"

type ChatMessage struct {
	Id        int64     `json:"id"`
	SpaceId   int64     `json:"space_id"`
	From      int64     `json:"from"`
	To        int64     `json:"to"`
	Type      string    `json:"type"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (m *ChatMessage) TableName() string {
	return "message"
}

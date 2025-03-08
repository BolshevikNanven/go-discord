package model

import "time"

type Channel struct {
	Id        int64 `gorm:"primaryKey"`
	SpaceId   int64
	Name      string
	Owner     int64
	Type      string
	CreatedAt time.Time
}

type ChannelUser struct {
	ChannelId        int64 `gorm:"primaryKey"`
	UserId           int64 `gorm:"primaryKey"`
	LastAckMessageId int64
	CreatedAt        time.Time
}

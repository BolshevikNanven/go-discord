package model

import "time"

type Space struct {
	Id        int64 `gorm:"primaryKey"`
	Name      string
	Avatar    string
	Owner     int64
	CreatedAt time.Time
}

type SpaceUser struct {
	SpaceId   int64 `gorm:"primaryKey"`
	UserId    int64 `gorm:"primaryKey"`
	Nickname  string
	CreatedAt time.Time
}

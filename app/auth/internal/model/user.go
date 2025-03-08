package model

import "time"

type User struct {
	Id        int64 `gorm:"primaryKey"`
	Username  string
	Password  string
	Phone     int
	CreatedAt time.Time
}

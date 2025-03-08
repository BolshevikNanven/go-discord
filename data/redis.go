package data

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	KeyFormatInbox         = "space:%d:user:%d:inbox" //List
	KeyFormatUserConnector = "space:%d:users"         //hashmap

	KeyFormatUserChannel     = "user:%d:channels" //Set
	KeyFormatUserChannelLock = "user:%d:lock:channels"
	KeyFormatUserSpace       = "user:%d:spaces" //Set
	KeyFormatUserSpaceLock   = "user:%d:lock:spaces"

	KeyFormatChannelCurrent       = "channel:%d:current"
	KeyFormatChannelAck           = "channel:%d:ack:%d"
	KeyFormatChannelUserConnector = "channel:%d:users" //hashmap
)

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

func NewRedis(config *RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
	})

	return rdb
}

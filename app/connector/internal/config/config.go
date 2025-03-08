package config

import (
	"discord/data"
	"discord/pkg/discovery"
	"discord/pkg/jwtutil"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Etcd      discovery.EtcdConfig `yaml:"etcd"`
	Jwt       jwtutil.Config       `yaml:"jwt"`
	Redis     data.RedisConfig     `yaml:"redis"`
	Websocket WebsocketConfig      `yaml:"websocket"`
	Host      string               `yaml:"host"`
	Port      string               `yaml:"port"`
	Name      string               `yaml:"name"`
}

type WebsocketConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func NewConfig() *Config {
	f, err := os.Open("./config.yaml")
	if err != nil {
		panic("load config failed")
	}
	defer func() {
		f.Close()
	}()

	config := &Config{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)
	if err != nil {
		panic("load config failed")
	}

	return config
}

func NewJwtConfig(config *Config) *jwtutil.Config {
	return &jwtutil.Config{
		SecretKey:          config.Jwt.SecretKey,
		AccessTokenExpiry:  config.Jwt.AccessTokenExpiry,
		RefreshTokenExpiry: config.Jwt.RefreshTokenExpiry,
	}
}

func NewRedisConfig(config *Config) *data.RedisConfig {
	return &data.RedisConfig{
		Host:     config.Redis.Host,
		Port:     config.Redis.Port,
		Password: config.Redis.Password,
	}
}

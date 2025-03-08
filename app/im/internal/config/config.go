package config

import (
	"discord/data"
	"discord/pkg/discovery"
	"discord/pkg/snowflakeutil"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Redis     *data.RedisConfig     `yaml:"redis"`
	Etcd      *discovery.EtcdConfig `yaml:"etcd"`
	Database  *data.DatabaseConfig  `yaml:"database"`
	Snowflake *snowflakeutil.Config `yaml:"snowflake"`
	MQ        *data.RocketMQConfig  `yaml:"mq"`
	Host      string                `yaml:"host"`
	Port      string                `yaml:"port"`
}

func NewConfig() *Config {
	f, err := os.Open("./config.yaml")
	if err != nil {
		panic("load config failed: " + err.Error())
	}
	defer func() {
		f.Close()
	}()

	config := &Config{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)
	if err != nil {
		panic("load config failed: " + err.Error())
	}

	return config
}

func NewRedisConfig(config *Config) *data.RedisConfig {
	return config.Redis
}

func NewEtcdConfig(config *Config) *discovery.EtcdConfig {
	return config.Etcd
}

func NewDatabaseConfig(config *Config) *data.DatabaseConfig {
	return config.Database
}

func NewRocketMQConfig(config *Config) *data.RocketMQConfig {
	return config.MQ
}

func NewSnowflakeConfig(config *Config) *snowflakeutil.Config {
	return config.Snowflake
}

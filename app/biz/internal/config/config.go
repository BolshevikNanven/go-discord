package config

import (
	"discord/data"
	"discord/pkg/discovery"
	"discord/pkg/snowflakeutil"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Etcd      discovery.EtcdConfig `yaml:"etcd"`
	Database  data.DatabaseConfig  `yaml:"database"`
	Snowflake snowflakeutil.Config `yaml:"snowflake"`
	Redis     data.RedisConfig     `yaml:"redis"`
	Host      string               `yaml:"host"`
	Port      string               `yaml:"port"`
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

func NewDatabase(config *Config) *data.DatabaseConfig {
	return &config.Database
}

func NewSnowflakeConfig(config *Config) *snowflakeutil.Config {
	return &snowflakeutil.Config{
		MachineID: config.Snowflake.MachineID,
	}
}

func NewRedisConfig(config *Config) *data.RedisConfig {
	return &config.Redis
}

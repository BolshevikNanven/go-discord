package config

import (
	"discord/data"
	"discord/pkg/discovery"
	"discord/pkg/jwtutil"
	"discord/pkg/snowflakeutil"
	"discord/pkg/tracer"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Etcd      discovery.EtcdConfig `yaml:"etcd"`
	Database  data.DatabaseConfig  `yaml:"database"`
	Jwt       jwtutil.Config       `yaml:"jwt"`
	Snowflake snowflakeutil.Config `yaml:"snowflake"`
	Tracer    tracer.Config        `yaml:"tracer"`
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

func NewJwtConfig(config *Config) *jwtutil.Config {
	return &jwtutil.Config{
		SecretKey:          config.Jwt.SecretKey,
		AccessTokenExpiry:  config.Jwt.AccessTokenExpiry,
		RefreshTokenExpiry: config.Jwt.RefreshTokenExpiry,
	}
}

func NewSnowflakeConfig(config *Config) *snowflakeutil.Config {
	return &snowflakeutil.Config{
		MachineID: config.Snowflake.MachineID,
	}
}

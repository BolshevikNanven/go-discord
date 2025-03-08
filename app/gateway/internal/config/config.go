package config

import (
	"discord/pkg/discovery"
	"discord/pkg/jwtutil"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Etcd *discovery.EtcdConfig `yaml:"etcd"`
	JWT  *jwtutil.Config       `yaml:"jwt"`
	Host string                `yaml:"host"`
	Port string                `yaml:"port"`
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

func NewJWTConfig(config *Config) *jwtutil.Config {
	return config.JWT
}

func NewEtcdConfig(config *Config) *discovery.EtcdConfig {
	return config.Etcd
}

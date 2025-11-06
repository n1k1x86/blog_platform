package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Mongo Mongo `yaml:"mongo"`
	App   App   `yaml:"app"`
}

type Mongo struct {
	URI      string `yaml:"uri"`
	DBName   string `yaml:"db_name"`
	CollName string `yaml:"coll_name"`
}

type App struct {
	Port string `yaml:"port"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}

	cfg := Config{}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

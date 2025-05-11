package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Server struct {
		Port string `env:"SERVER_PORT"`
	}
	MongoDB struct {
		URI        string `env:"MONGO_URI"`
		DB         string `env:"MONGO_DB"`
		Collection string `env:"MONGO_CONFIGS_COLLECTION"`
	}
}

func LoadYAML(path string) (*AppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func LoadFromEnv() *AppConfig {
	cfg := &AppConfig{}
	cfg.Server.Port = os.Getenv("SERVER_PORT")
	cfg.MongoDB.URI = os.Getenv("MONGO_URI")
	cfg.MongoDB.DB = os.Getenv("MONGO_DB")
	cfg.MongoDB.Collection = os.Getenv("MONGO_CONFIGS_COLLECTION")
	return cfg
}

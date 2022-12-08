package server

import "github.com/beeloon/go-rest-api/internal/app/store"

type Config struct {
	Port     string `toml:"port"`
	LogLeveL string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		Port:     ":8000",
		LogLeveL: "debug",
		Store:    store.NewConfig(),
	}
}

package config

import "os"

type AppConfig struct {
	Token string
}

func LoadConfig() *AppConfig {
	return &AppConfig{
		Token: os.Getenv("GH_TOKEN"),
	}
}

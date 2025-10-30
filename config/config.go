package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	GithubToken   string
	OpenAIAPIKey  string
	ProxyIP       string
	ProxyPort     string
	ProxyUsername string
	ProxyPassword string
}

func LoadConfig() *AppConfig {
	godotenv.Load()

	return &AppConfig{
		GithubToken:   os.Getenv("GH_TOKEN"),
		OpenAIAPIKey:  os.Getenv("OPENAI_API_KEY"),
		ProxyIP:       os.Getenv("PROXY_IP"),
		ProxyPort:     os.Getenv("PROXY_PORT"),
		ProxyUsername: os.Getenv("PROXY_USERNAME"),
		ProxyPassword: os.Getenv("PROXY_PASSWORD"),
	}
}

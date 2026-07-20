package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	DevToken     string
	RedisURL     string
}

func Load() *Config {
	godotenv.Load()

	cfg := &Config{
		ClientID:     os.Getenv("UPSTOX_API_KEY"),
		ClientSecret: os.Getenv("UPSTOX_API_SECRET"),
		RedirectURI:  os.Getenv("UPSTOX_REDIRECT_URI"),
		RedisURL:     os.Getenv("REDIS_URL"),
	}

	if cfg.ClientID == "" || cfg.ClientSecret == "" || cfg.RedirectURI == "" {
		log.Fatal("Missing environment variables")
	}

	return cfg
}

// Package config provides loading of configurations such as environment variables.
package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// Config represents the configurations required for the application to run.
type Config struct {
	APIToken string `env:"API_TOKEN"`
	APIURL   string `env:"API_URL" envDefault:"https://api.telegram.org"`
	BotName  string `env:"BOT_NAME"`
}

// Load loads the configuration required.
// It returns a Config struct and an 'ok' boolean value.
func Load() (config Config, ok bool) {
	var cfg Config

	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatalf("Can't load .env file. Error message: %v", err)
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Can't parse environment variables. Error message: %v", err)
	}

	ok = true

	return cfg, ok
}

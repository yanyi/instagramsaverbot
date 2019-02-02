// Package config provides loading of configurations such as environment variables.
package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config represents the configurations required for the application to run.
type Config struct {
	Version uint32 `yaml:"version"`
	Configs struct {
		AppEnv      string            `yaml:"app_env"`
		TelegramBot map[string]string `yaml:"telegram_bot"`
	} `yaml:"configs"`
}

// Load loads the configuration required.
// It returns a Config struct and an 'ok' boolean value.
func Load(configFile string) (config Config, ok bool) {
	var cfg Config

	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Can't load config.yml file. Error message: %v", err)
	}
	log.Println("Loaded config file. Will unmarshal config file now.")

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatalf("Unable to unmarshal config.yml file. Error message: %v", err)
	}

	log.Println("Unmarshalled config file.")
	ok = true

	return cfg, ok
}

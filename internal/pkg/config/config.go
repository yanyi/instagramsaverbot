// Package config provides loading of configurations such as environment variables.
package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/yanyi/instagramsaverbot/internal/pkg/applogger"
)

var (
	logger = applogger.New(os.Stderr)
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
		logger.Log(
			"event", "Can't load config.yml file",
			"error", err,
		)
	}
	logger.Log("event", "Loaded config file")

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		logger.Log(
			"event", "Unable to unmarshal config file",
			"error", err,
		)
	}

	logger.Log("event", "Unmarshalled config file")
	ok = true

	return cfg, ok
}

package main

import (
	"flag"
	"os"

	"github.com/yanyi/instagramsaverbot/internal/app/telegrambot"
	"github.com/yanyi/instagramsaverbot/internal/pkg/applogger"
	"github.com/yanyi/instagramsaverbot/internal/pkg/config"
)

const (
	configFileUsage = "Path to config file from root."
)

var (
	logger = applogger.New(os.Stderr)
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "file", "configs/config.yml", configFileUsage)
	flag.StringVar(&configFile, "f", "configs/config.yml", configFileUsage)
	flag.Parse()

	if _, err := os.Stat(configFile); err == nil {
		if cfg, ok := config.Load(configFile); ok {
			telegrambot.Start(cfg)
		}
	} else if os.IsNotExist(err) {
		logger.Log(
			"event", "Unable to load config file",
			"error", err,
		)
		os.Exit(1)
	}
}

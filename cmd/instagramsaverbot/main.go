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
	envUsage        = "Requires API_TOKEN and API_URL to be set. Especially for Heroku."
)

var (
	logger     = applogger.New(os.Stderr)
	configFile string
	useEnv     bool
)

func main() {
	flag.StringVar(&configFile, "file", "configs/config.yml", configFileUsage)
	flag.StringVar(&configFile, "f", "configs/config.yml", configFileUsage)
	flag.BoolVar(&useEnv, "env", false, envUsage)
	flag.Parse()

	if useEnv {
		startWithEnvVar()
	} else {
		startWithConfigFile()
	}

}

func startWithEnvVar() {
	telegramCfg := make(map[string]string)
	telegramCfg["api_token"] = os.Getenv("API_TOKEN")
	telegramCfg["api_url"] = os.Getenv("API_URL")
	telegramCfg["bot_name"] = os.Getenv("BOT_NAME")

	cfg := config.Config{
		Configs: config.Configs{
			TelegramBot: telegramCfg,
		},
	}

	telegrambot.Start(cfg)
}

func startWithConfigFile() {
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

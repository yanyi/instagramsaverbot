package main

import (
	"flag"
	"log"
	"os"

	"github.com/yanyi/instagramsaverbot/internal/app/telegrambot"
	"github.com/yanyi/instagramsaverbot/internal/pkg/config"
)

var (
	configFile = flag.String("file", "configs/config.yml", "Path to config file from root.")
)

func main() {
	flag.StringVar(configFile, "f", "configs/config.yml", "Path to config file from root.")
	flag.Parse()

	if _, err := os.Stat(*configFile); err == nil {
		if cfg, ok := config.Load(*configFile); ok {
			telegrambot.Start(cfg)
		}
	} else if os.IsNotExist(err) {
		log.Printf("File does not exist: '%v'. Can't start the app. Please check your file.", *configFile)
	}
}

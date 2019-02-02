package main

import (
	"github.com/yanyi/instagramsaverbot/internal/app/telegrambot"
	"github.com/yanyi/instagramsaverbot/internal/pkg/config"
)

func main() {
	if cfg, ok := config.Load(); ok {
		telegrambot.Start(cfg)
	}
}

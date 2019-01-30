package main

import (
	"github.com/yanyi/go-telegram-bot/internal/app/telegrambot"
	"github.com/yanyi/go-telegram-bot/internal/pkg/config"
)

func main() {
	if cfg, ok := config.Load(); ok {
		telegrambot.Start(cfg)
	}
}

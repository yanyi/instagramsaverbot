// Package telegrambot handles all the Telegram bot commands.
package telegrambot

import (
	"fmt"
	"os"
	"time"

	"github.com/yanyi/instagramsaverbot/internal/pkg/applogger"
	"github.com/yanyi/instagramsaverbot/internal/pkg/config"
	telebot "gopkg.in/tucnak/telebot.v2"
)

var (
	logger = applogger.New(os.Stderr)
)

// Start will take in the configurations required for the Telegram bot, and
// then it will start the Telegram bot.
func Start(cfg config.Config) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  cfg.Configs.TelegramBot["api_token"],
		URL:    cfg.Configs.TelegramBot["api_url"],
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		logger.Log("event", "Can't load Telegram bot", "error", err)
	}

	loadHandlers(bot)
	logger.Log("event", "Loaded Telegram message handlers")

	logger.Log("event", fmt.Sprintf("Bot %s started…", cfg.Configs.TelegramBot["bot_name"]))
	bot.Start()
}

// Package telegrambot handles all the Telegram bot commands.
package telegrambot

import (
	"log"
	"time"

	"github.com/yanyi/go-telegram-bot/internal/pkg/config"
	telebot "gopkg.in/tucnak/telebot.v2"
)

// Start will take in the configurations required for the Telegram bot, and
// then it will start the Telegram bot.
func Start(cfg config.Config) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  cfg.APIToken,
		URL:    cfg.APIURL,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal("Can't load the Telegram Bot. Error message: ", err)
	}

	bot.Handle("/hello", func(m *telebot.Message) {
		log.Printf("User: {%s} Message: {%s}", m.Sender.Username, m.Text)
		bot.Send(m.Sender, "Hello world")
	})

	log.Printf("Bot {%s} started…", cfg.BotName)
	bot.Start()
}

package telegrambot

import (
	"log"

	telebot "gopkg.in/tucnak/telebot.v2"
)

func loadHandlers(bot *telebot.Bot) {
	bot.Handle("/hello", func(m *telebot.Message) {
		log.Printf("User: {%s} Message Received: {%s}", m.Sender.Username, m.Text)
		go sendHelloWorld(bot, m)
	})

	bot.Handle("/save", func(m *telebot.Message) {
		log.Printf("User: {%s} Message Received: {%s}", m.Sender.Username, m.Text)
	})
}

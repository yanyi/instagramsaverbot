package telegrambot

import (
	telebot "gopkg.in/tucnak/telebot.v2"
)

func loadHandlers(bot *telebot.Bot) {

	bot.Handle("/hello", func(m *telebot.Message) {
		logger.Log(
			"event", "Received /hello command",
			"username", m.Sender.Username,
			"received", m.Text,
		)
		go sendHelloWorld(bot, m)
	})

	bot.Handle("/save", func(m *telebot.Message) {
		logger.Log(
			"event", "Received /save command",
			"username", m.Sender.Username,
			"received", m.Text,
		)
		go sendInstagramImage(bot, m)
	})
}

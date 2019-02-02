package telegrambot

import (
	telebot "gopkg.in/tucnak/telebot.v2"
)

func loadHandlers(bot *telebot.Bot) {

	bot.Handle("/start", func(m *telebot.Message) {
		logger.Log(
			"event", "Received /start",
			"username", m.Sender.Username,
			"payload", m.Payload,
		)
		go sendStartMessage(bot, m)
	})

	bot.Handle("/hello", func(m *telebot.Message) {
		logger.Log(
			"event", "Received /hello",
			"username", m.Sender.Username,
			"payload", m.Payload,
		)
		go sendHelloWorld(bot, m)
	})

	bot.Handle("/save", func(m *telebot.Message) {
		logger.Log(
			"event", "Received /save",
			"username", m.Sender.Username,
			"payload", m.Payload,
		)
		go sendInstagramImage(bot, m)
	})

	bot.Handle(telebot.OnText, func(m *telebot.Message) {
		logger.Log(
			"event", "Received message unhandled",
			"username", m.Sender.Username,
			"payload", m.Payload,
		)
	})
}

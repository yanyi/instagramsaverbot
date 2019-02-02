package telegrambot

import (
	telebot "gopkg.in/tucnak/telebot.v2"
)

func loadHandlers(bot *telebot.Bot) {

	bot.Handle("/start", func(m *telebot.Message) {
		logger.Log(
			"event", "Received /start",
			"sender", m.Sender,
			"payload", m.Payload,
		)
		go sendStartMessage(bot, m)
	})

	bot.Handle("/hello", func(m *telebot.Message) {
		logger.Log(
			"event", "Received /hello",
			"sender", m.Sender,
			"payload", m.Payload,
		)
		go sendHelloWorld(bot, m)
	})

	bot.Handle("/save", func(m *telebot.Message) {
		logger.Log(
			"event", "Received /save",
			"sender", m.Sender,
			"payload", m.Payload,
		)
		go sendInstagramImage(bot, m)
	})

	bot.Handle(telebot.OnText, func(m *telebot.Message) {
		logger.Log(
			"event", "Received message unhandled",
			"sender", m.Sender,
			"payload", m.Payload,
		)
	})
}

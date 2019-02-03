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

	bot.Handle("/help", func(m *telebot.Message) {
		logger.Log(
			"event", "Received /help",
			"sender", m.Sender,
			"payload", m.Payload,
		)
		go sendHelpMsg(bot, m)
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

		go func() {
			instagram, urls, err := checkIfContainsInstagram(bot, m)
			if err != nil {
				sendErrorMsg(bot, m, err)
			}

			if instagram {
				sendImageAlbum(bot, m, urls)
			}
		}()
	})

	bot.Handle(telebot.OnText, func(m *telebot.Message) {
		logger.Log(
			"event", "Received message unhandled",
			"sender", m.Sender,
			"userMessage", m.Text,
		)
		go func() {
			instagram, urls, err := checkIfContainsInstagram(bot, m)
			if err != nil {
				sendErrorMsg(bot, m, err)
			}

			if instagram {
				sendImageAlbum(bot, m, urls)
			}
		}()

	})
}

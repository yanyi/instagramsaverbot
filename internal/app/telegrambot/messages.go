package telegrambot

import (
	"fmt"

	"github.com/yanyi/instagramsaverbot/internal/app/scraper"
	telebot "gopkg.in/tucnak/telebot.v2"
)

func sendHelloWorld(bot *telebot.Bot, m *telebot.Message) {
	reply := fmt.Sprintf("Hello, %s 👋", m.Sender.FirstName)
	bot.Send(m.Sender, reply)
	logger.Log(
		"event", "Replied user",
		"username", m.Sender.Username,
		"reply", reply,
	)
}

func sendInstagramImage(bot *telebot.Bot, m *telebot.Message) {
	inputURL := m.Payload
	urls := []string{}
	err := scraper.Scrape(inputURL, &urls)
	if err != nil {
		errMsg := fmt.Sprintf("%s. Please try sending with just the Instagram link.", err.Error())
		bot.Send(m.Sender, errMsg)
		logger.Log(
			"event", "Can't scrape link",
			"error", err,
			"payload", m.Payload,
		)
		return
	}

	for _, url := range urls {
		bot.Send(m.Sender, url)
		logger.Log(
			"event", "Sent Instagram image URL",
			"username", m.Sender.Username,
			"reply", url,
		)
	}

	bot.Send(m.Sender, "Enjoy your links.")
}

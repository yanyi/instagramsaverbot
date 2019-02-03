package telegrambot

import (
	"fmt"

	"github.com/yanyi/instagramsaverbot/internal/app/scraper"
	telebot "gopkg.in/tucnak/telebot.v2"
)

func sendStartMessage(bot *telebot.Bot, m *telebot.Message) {
	bot.Notify(m.Chat, telebot.Typing)
	bot.Send(m.Chat, welcomeMsg)
	logger.Log(
		"event", "Welcomed user",
		"sender", m.Sender,
	)
}

func sendHelpMsg(bot *telebot.Bot, m *telebot.Message) {
	bot.Notify(m.Chat, telebot.Typing)
	bot.Send(m.Chat, helpMsg)
	logger.Log(
		"event", "Replied user with help commands",
		"sender", m.Sender,
	)
}

func sendHelloWorld(bot *telebot.Bot, m *telebot.Message) {
	bot.Notify(m.Chat, telebot.Typing)
	reply := fmt.Sprintf("Hello, %s 👋", m.Sender.FirstName)
	bot.Send(m.Chat, reply)
	logger.Log(
		"event", "Replied user with a greeting",
		"sender", m.Sender,
		"reply", reply,
	)
}

func sendInstagramImage(bot *telebot.Bot, m *telebot.Message) {
	inputURL := m.Payload
	urls := []string{}

	err := scraper.Scrape(inputURL, &urls)
	if err != nil {
		bot.Notify(m.Chat, telebot.Typing)
		msg := fmt.Sprintf(errMsg, err.Error())
		bot.Send(m.Chat, msg)
		logger.Log(
			"event", "Can't scrape link",
			"error", err,
			"payload", m.Payload,
		)
		return
	}

	sendImageAlbum(bot, m, urls)
}

func sendImageAlbum(bot *telebot.Bot, m *telebot.Message, urls []string) {
	album := telebot.Album{}
	logger.Log("event", "Start preparing image album", "sender", m.Sender)
	bot.Notify(m.Chat, telebot.UploadingPhoto)

	for _, url := range urls {
		photo := telebot.Photo{File: telebot.FromURL(url)}
		album = append(album, &photo)
		logger.Log(
			"event", "Got an Instagram image",
		)
	}

	bot.SendAlbum(m.Chat, album, telebot.Silent, telebot.NoPreview)
	logger.Log(
		"event", "Sent image album",
		"sender", m.Sender,
	)
	bot.Send(m.Chat, foundLinkMsg)
}

func sendErrorMsg(bot *telebot.Bot, m *telebot.Message, err error) {
	msgReply := fmt.Sprintf(errMsg, err.Error())
	bot.Send(m.Chat, msgReply)
}

package telegrambot

import (
	"fmt"

	"github.com/yanyi/instagramsaverbot/internal/app/scraper"
	telebot "gopkg.in/tucnak/telebot.v2"
	xurls "mvdan.cc/xurls/v2"
)

func sendStartMessage(bot *telebot.Bot, m *telebot.Message) {
	bot.Notify(m.Chat, telebot.Typing)
	msg := `Hello 👋! Hi!

To utilize me, send me a message like:
/save https://instagram.com/p/link
I will then return you an album of image(s).

If you need help, use the /help command. 😄
	`
	bot.Send(m.Chat, msg)
	logger.Log(
		"event", "Welcomed user",
		"sender", m.Sender,
		"reply", msg,
	)
}

func sendHelpMsg(bot *telebot.Bot, m *telebot.Message) {
	bot.Notify(m.Chat, telebot.Typing)
	helpMsg := `You looked for help!

The available commands I can handle are:
- /save https://instagram.com/p/link - Gives you an album of image(s)
- /help - You are viewing the help command now

Happy saving! 😄
	`

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
		"event", "Replied user",
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
		errMsg := fmt.Sprintf("%s. Please try sending an Instagram link that contains a photo. We currently do not support videos 🙇‍♂️.", err.Error())
		bot.Send(m.Chat, errMsg)
		logger.Log(
			"event", "Can't scrape link",
			"error", err,
			"payload", m.Payload,
		)
		return
	}

	logger.Log("event", "Start preparing image album", "sender", m.Sender)
	album := telebot.Album{}
	bot.Notify(m.Chat, telebot.UploadingPhoto)
	for _, url := range urls {
		photo := telebot.Photo{File: telebot.FromURL(url)}
		album = append(album, &photo)
		logger.Log(
			"event", "Gotten Instagram image",
			"sender", m.Sender,
			"photoURL", photo.FileURL,
		)
	}
	bot.SendAlbum(m.Chat, album, telebot.Silent, telebot.NoPreview)
	logger.Log("event", "Sent image album", "sender", m.Sender)
}

func checkIfInstagram(bot *telebot.Bot, m *telebot.Message) {
	message := m.Text
	logger.Log("event", "Checking if it's Instagram link", "message", message)
	urls := xurls.Relaxed().FindAllString(message, -1)

	urlsSlice := []string{}
	if len(urls) > 0 {
		for _, url := range urls {
			scraper.Scrape(url, &urlsSlice)
		}
	}

	// Found Instagram link(s)
	if len(urlsSlice) > 0 {
		logger.Log("event", "Found Instagram link(s)", "urls", urlsSlice)
		logger.Log("event", "Start preparing image album", "sender", m.Sender)
		album := telebot.Album{}
		bot.Notify(m.Chat, telebot.UploadingPhoto)
		for _, url := range urlsSlice {
			photo := telebot.Photo{File: telebot.FromURL(url)}
			album = append(album, &photo)
			logger.Log(
				"event", "Gotten Instagram image",
				"sender", m.Sender,
				"photoURL", photo.FileURL,
			)
		}
		bot.SendAlbum(m.Chat, album, telebot.Silent, telebot.NoPreview)
		logger.Log("event", "Sent image album", "sender", m.Sender)
	}
}

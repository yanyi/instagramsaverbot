package telegrambot

import (
	"errors"

	"github.com/yanyi/instagramsaverbot/internal/app/scraper"
	telebot "gopkg.in/tucnak/telebot.v2"
	xurls "mvdan.cc/xurls/v2"
)

// checkIfContainsInstagram checks a generic message sent by the user, if that
// contains any Instagram posts link. Returns a boolean value, a string slice
// and an error message.
func checkIfContainsInstagram(bot *telebot.Bot, m *telebot.Message) (bool, []string, error) {
	userMsg := m.Text
	urlsFromMsg := xurls.Relaxed().FindAllString(userMsg, -1)

	// Main logic to check for Instagram posts
	urls := []string{}
	var err error
	logger.Log("event", "Checking if there is an Instagram post", "userMessage", userMsg)
	if len(urlsFromMsg) > 0 {
		for _, url := range urlsFromMsg {
			err = scraper.Scrape(url, &urls)
		}
	}

	// Found no Instagram posts
	if err != nil && len(urls) == 0 {
		return false, nil, err
	}

	// Found Instagram post(s)
	if len(urls) > 0 {
		return true, urls, nil
	}

	return false, nil, errors.New("I did not find any links from your message")
}

package telegrambot

import (
	"log"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	telebot "gopkg.in/tucnak/telebot.v2"
)

type config struct {
	APIToken string `env:"API_TOKEN"`
	APIURL   string `env:"API_URL" envDefault:"https://api.telegram.org"`
	BotName  string `env:"BOT_NAME"`
}

func Start() {
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatalf("Can't load .env file. Error message: %v", err)
	}

	var cfg config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Can't parse environment variables. Error message: %v", err)
	}

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

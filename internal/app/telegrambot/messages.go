package telegrambot

import (
	"fmt"
	"log"

	telebot "gopkg.in/tucnak/telebot.v2"
)

func sendHelloWorld(bot *telebot.Bot, m *telebot.Message) {
	reply := fmt.Sprintf("Hello, %s 👋", m.Sender.FirstName)
	bot.Send(m.Sender, reply)
	log.Printf("User: {%s} Message Sent: {%s}", m.Sender.Username, reply)
}

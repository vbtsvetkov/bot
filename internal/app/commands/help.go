package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list products")
	if _, err := c.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func init() {
	registeredCommands["help"] = (*Commander).Help
}

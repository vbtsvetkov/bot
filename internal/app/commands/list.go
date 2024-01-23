package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsg := "Here all products: \n\n"
	products := c.productService.List()

	for _, p := range products {
		outputMsg += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
	if _, err := c.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func init() {
	registeredCommands["list"] = (*Commander).List
}

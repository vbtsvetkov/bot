package main

import (
	"github.com/joho/godotenv"
	"github.com/vbtsvetkov/bot/internal/sevice/product"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	_ = godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	productService := product.NewService()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)
		case "list":
			listCommand(bot, update.Message, productService)
		default:
			defaultBehavior(bot, update.Message)
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list products")
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productService *product.Service) {
	outputMsg := "Here all products: \n\n"
	products := productService.List()

	for _, p := range products {
		outputMsg += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

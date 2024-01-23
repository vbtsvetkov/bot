package main

import (
	"github.com/joho/godotenv"
	"github.com/vbtsvetkov/bot/internal/app/commands"
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
	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			commander.Help(update.Message)
		case "list":
			commander.List(update.Message)
		default:
			commander.Default(update.Message)
		}
	}
}

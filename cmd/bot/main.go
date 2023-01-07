package main

import (
	"log"
	"os"

	"github.com/andres-website/bot/cmd/bot/internal/app/commands"
	"github.com/andres-website/bot/cmd/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	token := os.Getenv("TOKEN")
	//token := "Your_token"

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Offset:  0,
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		if update.Message != nil { // If we got a message

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

} // END main

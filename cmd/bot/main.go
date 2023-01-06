package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const TOKEN = "YOUR_TOKEN"

func main() {
	bot, err := tgbotapi.NewBotAPI(TOKEN)
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

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You wrote: "+update.Message.Text)
			// msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

/* package main

import "fmt"

func main() {

	fmt.Println("Hello!")
}
*/

package main

import (
	"log"
	"os"
	"os/exec"

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

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You wrote: "+update.Message.Text)
			// msg.ReplyToMessageID = update.Message.MessageID

			// Запуск команды Windows (.Run дожидается завершения программы (закрытия калькулятора))
			/* 			cmd := exec.Command("calc")

			   			if err := cmd.Run(); err != nil {
			   				log.Fatal(err)
			   			} */

			// Запуск команды Windows (.Start запускает программу в фоне)
			/*
				https://stackoverflow.com/questions/48557810/execute-command-in-background
			*/
			cmd := exec.Command("calc")

			if err := cmd.Start(); err != nil {
				log.Printf("Failed to start cmd: %v", err)
				return
			}

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

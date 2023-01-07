package commands

import (
	"log"

	"github.com/andres-website/bot/cmd/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var registredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *product.Service,
) *Commander {
	return &Commander{

		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {

	// mutex := sync.Mutex{}
	// mutex.Lock()

	// defer mutex.Unlock()

	// defer callMe()

	defer func() {

		if panicValue := recover(); panicValue != nil {

			log.Printf("recovered from panic %s", panicValue)
		}

	}()

	// panic("Handle Panic")

	if update.Message != nil { // If we got a message

		switch update.Message.Command() {

		case "help":
			c.Help(update.Message)

		case "list":
			c.List(update.Message)

		case "get":
			c.Get(update.Message)

		default:
			c.Default(update.Message)
		}

	}
}

func callMe() {

}

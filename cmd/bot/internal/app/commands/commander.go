package commands

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/andres-website/bot/cmd/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Вариант с маршалингом анмаршалингом JSON для передачи данных из Инлайн клавиатуры
type CommandData struct {
	Offset int `json:"offset"`
}

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

	// recover after panic
	defer func() {

		if panicValue := recover(); panicValue != nil {

			log.Printf("recovered from panic %v", panicValue)
		}

	}()

	// Обработка Инлайн клавиатуры колбэк квэри
	if update.CallbackQuery != nil {

		// Вариант с распаршиванием строки Дата из Инлайн клавиатуры (split(var, "_"))
		/* 		args := update.CallbackQuery.Data
		   		split_data := strings.Split(args, "_")
		   		command := split_data[0]
		   		offset := split_data[1] */

		// Вариант с маршалингом анмаршалингом JSON для передачи данных из Инлайн клавиатуры
		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		// TODO: json.Unmarshal(...) - возвращает ошибку, т.е. здесь надо обработать её
		//		err := json.Unmarshal (...)

		// Откуда и как убрать часы на кнопках в telegram боте?
		// https://qna.habr.com/q/889107
		// https://core.telegram.org/bots/api#answercallbackquery
		// https://go-telegram-bot-api.dev/examples/inline-keyboard.html
		// Respond to the callback query, telling Telegram to show the user
		// a message with the data received.
		// Отвечаем на запрос обратного вызова, сообщая Telegram, что нужно показать пользователю
		// сообщение с полученными данными.
		// Сейчас вторым аргументом пустая строка, если там будет текст - он будет отображён
		// на несколько секунд на чёрном фоне аля поповер. ХЗ зачем может понадобится, а вдруг
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "" /* update.CallbackQuery.Data */)
		if _, err := c.bot.Request(callback); err != nil {
			panic(err)
		}

		// msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Data: "+update.CallbackQuery.Data)
		// # Вариант с распаршиванием строки Дата из Инлайн клавиатуры (split(var, "_"))
		/* 		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			"Data: command: "+command+"\noffset: "+offset,
		) */

		// # Вариант с маршалингом анмаршалингом JSON для передачи данных из Инлайн клавиатуры
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Parsed: %+v", parsedData),
		)

		c.bot.Send(msg)

		return
	}

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

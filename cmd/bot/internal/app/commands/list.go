package commands

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {

	outputText := "Here all the products:\n\n"

	products := c.productService.List()
	for _, p := range products {
		outputText += p.Title
		outputText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputText)

	// Simple Inline Keyboard

	serializedData, _ := json.Marshal(CommandData{
		Offset: 21,
	})
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(

			// # Вариант с распаршиванием строки Дата из Инлайн клавиатуры (split(var, "_"))
			/* tgbotapi.NewInlineKeyboardButtonData("Следующая страница", "list_10"), */

			// # Вариант с маршалингом анмаршалингом JSON для передачи данных из Инлайн клавиатуры
			tgbotapi.NewInlineKeyboardButtonData("Следующая страница", string(serializedData)),
		),
	)

	c.bot.Send(msg)
}

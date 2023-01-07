package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(inputMessage *tgbotapi.Message) {

	outputText := "Here all the products:\n\n"

	products := c.productService.List()
	for _, p := range products {
		outputText += p.Title
		outputText += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputText)

	c.bot.Send(msg)
}

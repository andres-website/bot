package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Help(inputMessage *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Тебе нужна помощь?\n"+
		"/help - Справка\n"+
		"/list - list products"+
		"/get IDX - getting product by idx (args - id product)")

	c.bot.Send(msg)
}

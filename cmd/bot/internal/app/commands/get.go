package commands

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		// log.Panic(err)
		log.Println("wrong args", args)
		return
	}

	product, err := c.productService.Get(idx)
	if err != nil {

		// TODO: Про ошибку можно сказать пользователю здесь или после этого где то отдельно
		log.Printf("Failed to get product with idx: %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		// fmt.Sprintf("successfully parsed arguments: %v", idx)
		product.Title,
	)

	c.bot.Send(msg)
}

package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outMessageText := "All products:\n\n"
	products := c.productService.List()

	for _, product := range products {
		outMessageText += product.Title
		outMessageText += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outMessageText)

	c.bot.Send(msg)
}

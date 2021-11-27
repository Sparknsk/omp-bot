package water

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WaterCommander) Default(inputMessage *tgbotapi.Message) {
	c.sendMessage(
		inputMessage.Chat.ID,
		"You wrote: "+inputMessage.Text,
	)
}

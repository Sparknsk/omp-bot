package water

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *WaterCommander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"You wrote: "+inputMessage.Text,
	)

	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("Ошибка Телеграм: %v", err)
	}
}

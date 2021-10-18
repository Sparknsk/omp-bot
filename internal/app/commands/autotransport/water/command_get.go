package water

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WaterCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	if args == "" {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Ошибка: не указан параметр ID"),
		)
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("Ошибка Телеграм: %v", err)
		}
		return
	}

	id, err := strconv.ParseUint(args, 10, 0)
	if err != nil {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Ошибка: параметр ID должен быть указан как число"),
		)
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("Ошибка Телеграм: %v", err)
		}
		return
	}

	p, err := c.service.Describe(id)
	if err != nil {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Ошибка: %v", err),
		)
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("Ошибка Телеграм: %v", err)
		}
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Сущность: %v", p),
	)
	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("Ошибка Телеграм: %v", err)
	}
}

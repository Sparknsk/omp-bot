package water

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WaterCommander) Delete(inputMessage *tgbotapi.Message) {
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

	ok, err := c.service.Remove(id)
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

	if !ok {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Не удалось удалить сущность с ID=%d", id),
		)
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("Ошибка Телеграм: %v", err)
		}
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Сущность с ID=%d успешно удалена", id),
	)
	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("Ошибка Телеграм: %v", err)
	}
}

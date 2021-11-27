package water

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WaterCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	if args == "" {
		c.sendMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Ошибка: не указан параметр ID"),
		)
		return
	}

	id, err := strconv.ParseUint(args, 10, 0)
	if err != nil {
		c.sendMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Ошибка: параметр ID должен быть указан как число"),
		)
		return
	}

	p, err := c.service.Describe(id)
	if err != nil {
		c.sendMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Ошибка: %v", err),
		)
		return
	}

	c.sendMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Сущность: %v", p),
	)
}

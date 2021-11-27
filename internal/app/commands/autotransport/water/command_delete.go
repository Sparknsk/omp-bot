package water

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WaterCommander) Delete(inputMessage *tgbotapi.Message) {
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

	ok, err := c.service.Remove(id)
	if err != nil {
		c.sendMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Ошибка: %v", err),
		)
		return
	}

	if !ok {
		c.sendMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Не удалось удалить сущность с ID=%d", id),
		)
		return
	}

	c.sendMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Сущность с ID=%d успешно удалена", id),
	)
}

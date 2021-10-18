package water

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/autotransport"
)

func (c *WaterCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	entityData := entityArgs{}
	if err := json.Unmarshal([]byte(args), &entityData); err != nil {
		c.sendMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Ошибка: параметры для создания сущности указаны неверно"),
		)
		return
	}

	entity := autotransport.NewWater(
		uint64(0),
		entityData.Name,
		entityData.Model,
		entityData.Manufacturer,
		entityData.Material,
		entityData.Speed,
	)

	id, err := c.service.Create(*entity)
	if err != nil {
		c.sendMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Ошибка: %v", err),
		)
		return
	}

	c.sendMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Сущность успешно добавлена, ID=%d", id),
	)
}

package water

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/autotransport"
	"regexp"
	"strconv"
)

func (c *WaterCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	re, _ := regexp.Compile(`^(?:\s+)?(\d+)\s+(.+)$`)
	if !re.MatchString(args) {
		c.sendMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Ошибка: параметры для редактирования сущности указаны неверно"),
		)
		return
	}

	result := re.FindAllStringSubmatch(args, -1)

	entityData := entityArgs{}
	if err := json.Unmarshal([]byte(result[0][2]), &entityData); err != nil {
		c.sendMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Ошибка: параметры для создания сущности указаны неверно"),
		)
		return
	}

	id, err := strconv.ParseUint(result[0][1], 10, 0)
	if err != nil {
		c.sendMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Ошибка: параметр ID должен быть указан как число"),
		)
		return
	}

	if entityData.Name == "" {
		c.sendMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Ошибка: необходимо заполнить поле name"),
		)
		return
	}

	entity := autotransport.NewWater(
		id,
		entityData.Name,
		entityData.Model,
		entityData.Manufacturer,
		entityData.Material,
		entityData.Speed,
	)

	if c.service.Update(id, *entity) != nil {
		c.sendMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Ошибка: %v", err),
		)
		return
	}

	c.sendMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Сущность с ID=%d успешно обновлена", id),
	)
}

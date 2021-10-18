package water

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/autotransport"
	"log"
	"regexp"
	"strconv"
)

func (c *WaterCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	re, _ := regexp.Compile(`^(\d)+\s(.+)$`)
	if !re.MatchString(args) {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Ошибка: параметры для редактирования сущности указаны неверно"),
		)
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("Ошибка Телеграм: %v", err)
		}
		return
	}

	result := re.FindAllStringSubmatch(args, -1)

	entityData := entityArgs{}
	if err := json.Unmarshal([]byte(result[0][2]), &entityData); err != nil {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Ошибка: параметры для создания сущности указаны неверно"),
		)
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("Ошибка Телеграм: %v", err)
		}
		return
	}

	id, err := strconv.ParseUint(result[0][1], 10, 0)
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

	entity := autotransport.NewWater(
		id,
		entityData.Name,
		entityData.Model,
		entityData.Manufacturer,
		entityData.Material,
		entityData.Speed,
	)

	if c.service.Update(id, *entity) != nil {
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
		fmt.Sprintf("Сущность с ID=%d успешно обновлена", id),
	)
	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("Ошибка Телеграм: %v", err)
	}
}

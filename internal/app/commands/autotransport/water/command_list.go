package water

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WaterCommander) List(inputMessage *tgbotapi.Message) {
	pageNumber := 1

	entitiesCount := c.service.Count()

	outputMsgText := fmt.Sprintf("Всего сущностей: %d\n\n", entitiesCount)
	if entitiesCount > 0 {
		entities := c.service.List(uint64(pageNumber), limitPerPage)
		for _, entity := range entities {
			outputMsgText += fmt.Sprintf("Сущность: %v", entity)+"\n\n"
		}
	} else {
		outputMsgText += fmt.Sprintf("Здесь пока ничего нет\n\nЧтобы добавить сущность, воспользуйтесь командой /new__autotransport__water {\"name\": \"entity name\"...}")+"\n"
	}

	var inlineKeyboardButtons []tgbotapi.InlineKeyboardButton
	if entitiesCount > 0 && entitiesCount > limitPerPage {
		callbackData, err := json.Marshal(CallbackListData{PageNumber: pageNumber+1})
		if err != nil {
			c.sendMessage(
				inputMessage.Chat.ID,
				fmt.Sprintf("Ошибка: не удалось получить список"),
			)
			return
		}

		inlineKeyboardButtons = append(inlineKeyboardButtons, tgbotapi.NewInlineKeyboardButtonData("Next", "autotransport__water__list__"+string(callbackData)))
	}

	c.sendMessageWithButtons(
		inputMessage.Chat.ID,
		outputMsgText,
		inlineKeyboardButtons,
	)
}

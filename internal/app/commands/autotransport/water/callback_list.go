package water

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	PageNumber int `json:"pageNumber"`
}

func (c *WaterCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	if err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData); err != nil {
		msg := tgbotapi.NewMessage(
			callback.Message.Chat.ID,
			fmt.Sprintf("Ошибка: номер страницы указан неверно"),
		)
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("Ошибка Телеграм: %v", err)
		}
		return
	}

	pageNumber := uint64(parsedData.PageNumber)

	entities := c.service.List(pageNumber, limitPerPage)
	outputMsgText := fmt.Sprintf("Всего сущностей: %d\n\n", c.service.Count())
	for _, entity := range entities {
		outputMsgText += fmt.Sprintf("Сущность: %v", entity)+"\n\n"
	}

	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		outputMsgText,
	)

	var inlineKeyboardButtons []tgbotapi.InlineKeyboardButton

	if pageNumber > 1 {
		callbackDataPrev, errPrev := json.Marshal(CallbackListData{PageNumber: int(pageNumber-1)})
		if errPrev != nil {
			msg = tgbotapi.NewMessage(
				callback.Message.Chat.ID,
				fmt.Sprintf("Ошибка: не удалось получить список"),
			)
			if _, err := c.bot.Send(msg); err != nil {
				log.Printf("Ошибка: %v", err)
			}
			return
		}
		inlineKeyboardButtons = append(inlineKeyboardButtons, tgbotapi.NewInlineKeyboardButtonData("Prev", "autotransport__water__list__"+string(callbackDataPrev)))
	}

	if pageNumber*limitPerPage < uint64(c.service.Count()) {
		callbackDataNext, errNext := json.Marshal(CallbackListData{PageNumber: int(pageNumber+1)})
		if errNext != nil {
			msg = tgbotapi.NewMessage(
				callback.Message.Chat.ID,
				fmt.Sprintf("Ошибка: не удалось получить список"),
			)
			if _, err := c.bot.Send(msg); err != nil {
				log.Printf("Ошибка: %v", err)
			}
			return
		}
		inlineKeyboardButtons = append(inlineKeyboardButtons, tgbotapi.NewInlineKeyboardButtonData("Next", "autotransport__water__list__"+string(callbackDataNext)))
	}

	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		inlineKeyboardButtons,
	)
	msg.ReplyMarkup = numericKeyboard

	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("Ошибка Телеграм: %v", err)
	}
}

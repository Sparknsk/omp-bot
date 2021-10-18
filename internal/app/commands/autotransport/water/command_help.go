package water

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *WaterCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"/help__autotransport__water - список доступных комманд\n"+
			"/list__autotransport__water ID - получить список сущностей\n"+
			"/get__autotransport__water ID - получить информацию о сущности по ID\n"+
			"/new__autotransport__water ENTITY - добавить новую сущность\n"+
			"/edit__autotransport__water ID ENTITY - изменить сущность по ID\n"+
			"/delete__autotransport__water ID - удалить сущность по ID\n",
	)

	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("Ошибка Телеграм: %v", err)
	}
}

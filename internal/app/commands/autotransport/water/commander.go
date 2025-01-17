package water

import (
	"github.com/ozonmp/omp-bot/internal/model/autotransport"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const limitPerPage = 2

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type WaterService interface {
	Describe(waterId uint64) (*autotransport.Water, error)
	List(cursor uint64, limit uint64) []autotransport.Water
	Create(water autotransport.Water) (uint64, error)
	Update(waterId uint64, water autotransport.Water) error
	Remove(waterId uint64) (bool, error)
	Count() int
}

type entityArgs struct {
	Name string `json:"name"`
	Model string `json:"model"`
	Manufacturer string `json:"manufacturer"`
	Material string `json:"material"`
	Speed uint `json:"speed"`
}

type WaterCommander struct {
	bot *tgbotapi.BotAPI
	service WaterService
}

func NewWaterCommander(bot *tgbotapi.BotAPI, service WaterService) *WaterCommander {
	return &WaterCommander{
		bot: bot,
		service: service,
	}
}

func (c *WaterCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("WaterCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *WaterCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}

func (c *WaterCommander) sendMessage(id int64, msg string) {
	newMsg := tgbotapi.NewMessage(id, msg)
	if _, err := c.bot.Send(newMsg); err != nil {
		log.Printf("Ошибка Телеграм: %v", err)
	}
}

func (c *WaterCommander) sendMessageWithButtons(id int64, msg string, buttons []tgbotapi.InlineKeyboardButton) {
	newMsg := tgbotapi.NewMessage(id, msg)

	if len(buttons) > 0 {
		newMsg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons)
	}

	if _, err := c.bot.Send(newMsg); err != nil {
		log.Printf("Ошибка Телеграм: %v", err)
	}
}
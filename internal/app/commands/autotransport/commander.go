package autotransport

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/autotransport/water"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type AutotransportCommander struct {
	bot *tgbotapi.BotAPI
	waterCommander Commander
}

func NewAutotransportCommander(
	bot *tgbotapi.BotAPI,
) *AutotransportCommander {
	return &AutotransportCommander{
		bot: bot,
		waterCommander: water.NewWaterCommander(bot),
	}
}

func (c *AutotransportCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "water":
		c.waterCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("AutotransportCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *AutotransportCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "water":
		c.waterCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("AutotransportCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}

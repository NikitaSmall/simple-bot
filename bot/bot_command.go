package bot

import (
	"github.com/nikitasmall/simple-bot/quoter"
	"gopkg.in/telegram-bot-api.v1"
	"log"
)

func (bot *Bot) ServeUpdates() {
	var msg tgbotapi.MessageConfig

	for update := range bot.Updates {
		command := update.Message.Text
		if bot.ApiBot.Debug == true {
			log.Printf("[%s] %s", update.Message.From.FirstName, command)
		}

		switch command {
		case "/quote":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, quoter.GetRandomQuote())
		default:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Don't understand "+update.Message.Text)
		}

		bot.ApiBot.Send(msg)
	}
}

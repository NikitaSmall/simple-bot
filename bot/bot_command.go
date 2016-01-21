package bot

import (
	"gopkg.in/telegram-bot-api.v1"
	"log"
)

func (bot *Bot) ServeUpdates() {
	for update := range bot.Updates {
		log.Printf("[%s] %s", update.Message.From.FirstName, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		bot.ApiBot.Send(msg)
	}
}

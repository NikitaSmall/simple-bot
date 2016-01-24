package bot

import (
	"gopkg.in/telegram-bot-api.v1"
)

// Function runs endless loop to listen user
// input from `bot.Updates` chan (chan of tgbotapi.Update).
// In case of new message starts to process it and send.
func (bot *Bot) ServeUpdates() {
	var msg tgbotapi.Chattable

	for update := range bot.Updates {
		msg = processCommand(update)
		bot.ApiBot.Send(msg)
	}
}

// func processes any user's input to bot.
// main handling function. It parses input and tries to process it.
func processCommand(update tgbotapi.Update) tgbotapi.Chattable {
	botCommand := newBotCommand(update)
	return botCommand.execute()
}

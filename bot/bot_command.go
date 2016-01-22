package bot

import (
	"github.com/nikitasmall/simple-bot/quoter"
	"gopkg.in/telegram-bot-api.v1"
	"regexp"
)

var commandRegexp = regexp.MustCompile(`/\w+|\w+|"[\w ]*"`)

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
	var msg tgbotapi.Chattable
	command, args := parseUserInput(update.Message.Text)

	switch command {
	case "/time":
		msg = tgbotapi.NewStickerUpload(update.Message.Chat.ID, "public/pic/at.jpg")
	case "/weather":
		if len(args) != 1 {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Wrong number of arguments! Should be one.")
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, quoter.GetCurrentWeather(args[0]))
		}
	case "/quote":
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, quoter.GetRandomQuote())
	default:
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Don't understand '"+update.Message.Text+"'")
	}

	return msg
}

func parseUserInput(input string) (string, []string) {
	args := commandRegexp.FindAllString(input, -1)
	return args[0], args[1:]
}

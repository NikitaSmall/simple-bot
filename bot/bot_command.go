package bot

import (
	"github.com/nikitasmall/simple-bot/quoter"
	"gopkg.in/telegram-bot-api.v1"
	"log"
	"regexp"
)

var commandRegexp = regexp.MustCompile(`/\w+|\w+|"[\w ]*"`)

func (bot *Bot) ServeUpdates() {
	var msg tgbotapi.MessageConfig

	for update := range bot.Updates {
		command, args := parseUserInput(update.Message.Text)

		if bot.ApiBot.Debug == true {
			log.Printf("[%s] %s", update.Message.From.FirstName, command)
		}

		switch command {
		case "/weather":
			if len(args) != 1 {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Wrong number of arguments! Should be one.")
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, quoter.GetCurrentWeather(args[0]))
			}
		case "/quote":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, quoter.GetRandomQuote())
		default:
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Don't understand "+update.Message.Text)
		}

		bot.ApiBot.Send(msg)
	}
}

func parseUserInput(input string) (string, []string) {
	args := commandRegexp.FindAllString(input, -1)
	return args[0], args[1:]
}

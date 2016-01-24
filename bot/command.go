package bot

import (
	"regexp"

	"gopkg.in/telegram-bot-api.v1"
)

type botCommand struct {
	command  string
	args     []string
	chatID   int
	fullText string
}

var commandRegexp = regexp.MustCompile(`/\w+|\w+|"[\w ]*"`)

func newBotCommand(update tgbotapi.Update) botCommand {
	command, args := parseUserInput(update.Message.Text)

	return botCommand{
		command:  command,
		args:     args,
		chatID:   update.Message.Chat.ID,
		fullText: update.Message.Text,
	}
}

func (bc botCommand) execute() tgbotapi.Chattable {
	var msg tgbotapi.Chattable

	switch {
	case bc.isQuoteRequest():
		msg = quoteHandler(bc)
	case bc.isWeatherRequest():
		msg = weatherHandler(bc)
	case bc.isAdventureTimeRequest():
		msg = adventureTimeHandler(bc)
	default:
		msg = defaultHandler(bc)
	}

	return msg
}

func parseUserInput(input string) (string, []string) {
	args := commandRegexp.FindAllString(input, -1)
	return args[0], args[1:]
}

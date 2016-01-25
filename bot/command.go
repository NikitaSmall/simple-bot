package bot

import (
	"regexp"

	"gopkg.in/telegram-bot-api.v1"
)

// botCommand represents data passed to bot.
// It contains all the needed info to execute command.
type botCommand struct {
	command  string
	args     []string
	chatID   int
	fullText string
}

// command parsed by this regexp.
var commandRegexp = regexp.MustCompile(`/\w+|\w+|"[\w ]*"`)

// function returns new botCommand created from parsed user input and
// chat data from tgbotapi.Update object
func newBotCommand(update tgbotapi.Update) botCommand {
	command, args := parseUserInput(update.Message.Text)

	return botCommand{
		command:  command,
		args:     args,
		chatID:   update.Message.Chat.ID,
		fullText: update.Message.Text,
	}
}

// function tries to execute botCommand and always send result message to user
// who send this command
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

// function parses user input and devide it to args and command
func parseUserInput(input string) (string, []string) {
	args := commandRegexp.FindAllString(input, -1)
	return args[0], args[1:]
}

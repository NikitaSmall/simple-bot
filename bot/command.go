package bot

import (
	"regexp"

	"github.com/nikitasmall/simple-bot/attachment"
	"github.com/nikitasmall/simple-bot/quoter"
	"gopkg.in/telegram-bot-api.v1"
)

type botCommand struct {
	command  string
	args     []string
	chatID   int
	fullText string
	update   tgbotapi.Update
}

var commandRegexp = regexp.MustCompile(`/\w+|\w+|"[\w ]*"`)

func newBotCommand(update tgbotapi.Update) botCommand {
	command, args := parseUserInput(update.Message.Text)

	return botCommand{
		command:  command,
		args:     args,
		chatID:   update.Message.Chat.ID,
		fullText: update.Message.Text,
		update:   update,
	}
}

func (bc botCommand) execute() tgbotapi.Chattable {
	var msg tgbotapi.Chattable

	switch {
	case bc.isAdventureTimeRequest():
		msg = tgbotapi.NewStickerUpload(bc.chatID, attachment.AdventureTimeStickers.GetAttachmentPath())
	case bc.isWeatherRequest():
		if len(bc.args) != 1 {
			msg = tgbotapi.NewMessage(bc.chatID, "Wrong number of arguments! Should be one.")
		} else {
			msg = tgbotapi.NewMessage(bc.chatID, quoter.GetCurrentWeather(bc.args[0]))
		}
	case bc.isQuoteRequest():
		msg = tgbotapi.NewMessage(bc.chatID, quoter.GetRandomQuote())
	default:
		msg = tgbotapi.NewMessage(bc.chatID, "Don't understand '"+bc.fullText+"'")
	}

	return msg
}

func parseUserInput(input string) (string, []string) {
	args := commandRegexp.FindAllString(input, -1)
	return args[0], args[1:]
}

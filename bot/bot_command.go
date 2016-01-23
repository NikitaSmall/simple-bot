package bot

import (
	"github.com/nikitasmall/simple-bot/quoter"
	"gopkg.in/telegram-bot-api.v1"
)

type botCommand struct {
	command string
	args    []string
	chatID  int
	update  tgbotapi.Update
}

func newBotCommand(update tgbotapi.Update) botCommand {
	command, args := parseUserInput(update.Message.Text)

	return botCommand{
		command: command,
		args:    args,
		chatID:  update.Message.Chat.ID,
		update:  update,
	}
}

func (bc botCommand) execute() tgbotapi.Chattable {
	var msg tgbotapi.Chattable

	switch bc.command {
	case "/time":
		msg = tgbotapi.NewStickerUpload(bc.chatID, "public/pic/at.jpg")
	case "/weather":
		if len(bc.args) != 1 {
			msg = tgbotapi.NewMessage(bc.chatID, "Wrong number of arguments! Should be one.")
		} else {
			msg = tgbotapi.NewMessage(bc.chatID, quoter.GetCurrentWeather(bc.args[0]))
		}
	case "/quote":
		msg = tgbotapi.NewMessage(bc.chatID, quoter.GetRandomQuote())
	default:
		msg = tgbotapi.NewMessage(bc.chatID, "Don't understand '"+bc.update.Message.Text+"'")
	}

	return msg
}

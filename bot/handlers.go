package bot

import (
	"github.com/nikitasmall/simple-bot/attachment"
	"github.com/nikitasmall/simple-bot/quoter"
	"gopkg.in/telegram-bot-api.v1"
)

// handler functions that handles incoming botCommand.
// in any case it should return tgbotapi.Chattable object.

func defaultHandler(bc botCommand) tgbotapi.Chattable {
	return tgbotapi.NewMessage(bc.chatID, "Don't understand '"+bc.fullText+"'")
}

func quoteHandler(bc botCommand) tgbotapi.Chattable {
	return tgbotapi.NewMessage(bc.chatID, quoter.GetRandomQuote())
}

func weatherHandler(bc botCommand) tgbotapi.Chattable {
	if len(bc.args) != 1 {
		return tgbotapi.NewMessage(bc.chatID, "Wrong number of arguments! Should be one.")
	} else {
		return tgbotapi.NewMessage(bc.chatID, quoter.GetCurrentWeather(bc.args[0]))
	}
}

func adventureTimeHandler(bc botCommand) tgbotapi.Chattable {
	return tgbotapi.NewStickerUpload(bc.chatID,
		attachment.AdventureTimeStickers.GetAttachmentPath())
}

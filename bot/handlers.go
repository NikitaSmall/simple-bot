package bot

import (
	"log"

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
	quote, err := quoter.GetRandomQuote()
	if err != nil || len(quote) == 0 {
		return tgbotapi.NewMessage(bc.chatID, "Can't get quote.")
	}

	return tgbotapi.NewMessage(bc.chatID, quote)
}

func weatherHandler(bc botCommand) tgbotapi.Chattable {
	if len(bc.args) != 1 {
		return tgbotapi.NewMessage(bc.chatID, "Wrong number of arguments! Should be one.")
	} else {
		weather, err := quoter.GetCurrentWeather(bc.args[0])

		if err != nil || len(weather) == 0 {
			return tgbotapi.NewMessage(bc.chatID, "Can't get weather for this city.")
		}
		return tgbotapi.NewMessage(bc.chatID, weather)
	}
}

func adventureTimeHandler(bc botCommand) tgbotapi.Chattable {
	msg, err := attachment.AdventureTimeStickers.GetAttachmentPath()
	if err != nil {
		log.Printf("Error on adventure time handling: %s", err.Error())
		return tgbotapi.NewMessage(bc.chatID, err.Error())
	}
	return tgbotapi.NewStickerUpload(bc.chatID, msg)
}

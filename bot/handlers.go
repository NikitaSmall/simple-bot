package bot

import (
	"log"

	"github.com/nikitasmall/simple-bot/attachment"
	"github.com/nikitasmall/simple-bot/quoter"
	"gopkg.in/telegram-bot-api.v1"
)

// Handler functions that handles incoming botCommand.
// In any case it should return tgbotapi.Chattable object.

// Default handler fires if none of matchers was fired.
// It should return a help message or `don't understand message`
func defaultHandler(bc botCommand) tgbotapi.Chattable {
	return tgbotapi.NewMessage(bc.chatID, "Don't understand '"+bc.fullText+"'")
}

// Quote handler returns a random joke from bash.io.
// If something goes wrong this handler will return error message.
func quoteHandler(bc botCommand) tgbotapi.Chattable {
	quote, err := quoter.GetRandomQuote()
	if err != nil || len(quote) == 0 {
		return tgbotapi.NewMessage(bc.chatID, "Can't get quote.")
	}

	return tgbotapi.NewMessage(bc.chatID, quote)
}

// Weahter handler returns current weather state for provided city-names.
// If something goes wrong this handler will return error message.
func weatherHandler(bc botCommand) tgbotapi.Chattable {
	if len(bc.args) != 1 {
		return tgbotapi.NewMessage(bc.chatID, "Wrong number of arguments! Should be one.")
	}

	weather, err := quoter.GetCurrentWeather(bc.args[0])
	if err != nil || len(weather) == 0 {
		return tgbotapi.NewMessage(bc.chatID, "Can't get weather for this city.")
	}

	return tgbotapi.NewMessage(bc.chatID, weather)
}

// Adventure time handler returns a sticker (if any exists in filesystem)
// or a message if none of stickers provided.
func adventureTimeHandler(bc botCommand) tgbotapi.Chattable {
	msg, err := attachment.AdventureTimeStickers.GetAttachmentPath()
	if err != nil {
		log.Printf("Error on adventure time handling: %s", err.Error())
		return tgbotapi.NewMessage(bc.chatID, err.Error())
	}

	return tgbotapi.NewStickerUpload(bc.chatID, msg)
}

func magicCardHandler(bc botCommand) tgbotapi.Chattable {
	var filePath string
	var err error

	if len(bc.args) > 1 {
		return tgbotapi.NewMessage(bc.chatID,
			"Wrong arguments number! Should be one or none. Format: '/m random' or '/m \"card name\"'")
	}

	if (len(bc.args) == 0) || bc.args[0] == "random" {
		filePath, err = attachment.GetRandomCard()
	} else {
		filePath, err = attachment.GetRandomCard() // attachment.GetCardByName(bc.args[0])
	}

	if err != nil {
		log.Printf("Error on magic card handling: %s", err.Error())
		return tgbotapi.NewMessage(bc.chatID, err.Error())
	}

	return tgbotapi.NewPhotoUpload(bc.chatID, filePath)
}

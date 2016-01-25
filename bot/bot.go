package bot

import (
	"log"

	"github.com/nikitasmall/simple-bot/config"
	"gopkg.in/telegram-bot-api.v1"
)

// bot struct represents API bot part and channel with updates for this bot.
type Bot struct {
	ApiBot  *tgbotapi.BotAPI
	Updates <-chan tgbotapi.Update
}

// function returns new bot with initialized update channel.
func CreateBot() *Bot {
	bot := newApiBot()

	return &Bot{
		ApiBot:  bot,
		Updates: newUpdatesChan(bot),
	}
}

// function creates new API bot
func newApiBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(config.Env["botToken"])
	if err != nil {
		log.Panic("Error on bot initializing! ", err.Error())
	}

	if config.Env["botMode"] == "debug" {
		bot.Debug = true
		log.Printf("Bot runned in debug mode.")
	}

	return bot
}

// function creates new update chan for provided bot
func newUpdatesChan(bot *tgbotapi.BotAPI) <-chan tgbotapi.Update {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic("Error on getting updates chan: ", err.Error())
	}

	return updates
}

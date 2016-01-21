package bot

import (
	"github.com/nikitasmall/simple-bot/config"
	"gopkg.in/telegram-bot-api.v1"
	"log"
)

type Bot struct {
	ApiBot  *tgbotapi.BotAPI
	Updates <-chan tgbotapi.Update
}

func CreateBot() *Bot {
	bot := newApiBot()

	return &Bot{
		ApiBot:  bot,
		Updates: newUpdatesChan(bot),
	}
}

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

func newUpdatesChan(bot *tgbotapi.BotAPI) <-chan tgbotapi.Update {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic("Error on getting updates chan: ", err.Error())
	}

	return updates
}

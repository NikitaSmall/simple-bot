package main

import (
	"github.com/nikitasmall/simple-bot/bot"
	"github.com/nikitasmall/simple-bot/config"
)

func main() {
	bot := bot.CreateBot(config.Env["botToken"])
	bot.ServeUpdates()
}

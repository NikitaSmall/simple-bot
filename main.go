package main

import (
	"github.com/nikitasmall/simple-bot/bot"
)

func main() {
	bot := bot.CreateBot()
	bot.ServeUpdates()
}

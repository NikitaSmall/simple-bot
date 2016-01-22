package quoter

import (
	"github.com/nikitasmall/simple-bot/config"
)

// bashQuoter is a Quoter instance which is able to get random jokes from http://bash.im
var bashQuoter = Quoter{
	url:          config.Env["quoteSource"],
	query:        ".quote .text",
	fromEncoding: "windows-1251",
}

// function returns the random joke from bash.im
func GetRandomQuote() string {
	return bashQuoter.getPageResult()
}

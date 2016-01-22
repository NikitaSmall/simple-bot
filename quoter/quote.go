package quoter

import (
	"github.com/nikitasmall/simple-bot/config"
	"io"
)

// simple interface to get some information from page in the internet
type PageParser interface {
	getPageResult() string
	getPage() io.Reader
}

// Quoter is a struct to get simple quotes from the page.
// Implements a PageParser interface and returns single string which
// could be found at the url via query string. fromEncoding should be a string
// that represents page original encoding.
type Quoter struct {
	url          string
	query        string
	fromEncoding string
}

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

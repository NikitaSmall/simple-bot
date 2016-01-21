package quoter

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"github.com/nikitasmall/simple-bot/config"
	"io"
	"log"
	"net/http"
)

func GetRandomQuote() string {
	page := getPage(config.Env["quoteSource"])

	doc, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		log.Print(err.Error())
	}

	quote := doc.Find(".quote .text").First()
	quote.Find("br").Each(func(i int, s *goquery.Selection) {
		s.ReplaceWithHtml("\n")
	})
	return quote.Text()
}

func getPage(url string) io.Reader {
	resp, err := http.Get(url)
	if err != nil {
		log.Print("Can't get the page %s with error: %s", url, err.Error())
	}

	page, err := iconv.NewReader(resp.Body, "windows-1251", "utf-8")
	if err != nil {
		log.Print("Can't read the page %s with error: %s", url, err.Error())
	}

	return page
}

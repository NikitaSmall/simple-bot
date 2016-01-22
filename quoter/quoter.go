package quoter

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"io"
	"log"
	"net/http"
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

// function parse a result page (provided by reader) and returns first single
// string contained in html-query (jQuery-like string).
// Part of PageParser interface Implementation for Quoter struct.
func (q Quoter) getPageResult() string {
	page := q.getPage()

	doc, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		log.Print(err.Error())
	}

	result := doc.Find(q.query).First()
	result.Find("br").Each(func(i int, s *goquery.Selection) {
		s.ReplaceWithHtml("\n")
	})
	return result.Text()
}

// function gets a page at the url provided by quoter instance
// and returns reader that represents body of the response page encoded in utf-8.
// Part of PageParser interface Implementation for Quoter struct.
func (q Quoter) getPage() io.Reader {
	resp, err := http.Get(q.url)
	if err != nil {
		log.Printf("Can't get the page %s with error: %s", q.url, err.Error())
	}

	page, err := iconv.NewReader(resp.Body, q.fromEncoding, "utf-8")
	if err != nil {
		log.Print("Can't read the page %s with error: %s", q.url, err.Error())
	}

	return page
}

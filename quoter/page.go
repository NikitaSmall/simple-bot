package quoter

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"io"
	"log"
	"net/http"
)

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

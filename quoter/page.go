package quoter

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"io"
	"log"
	"net/http"
)

func getPageResult(url string, query string) string {
	page := getPage(url)

	doc, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		log.Print(err.Error())
	}

	result := doc.Find(query).First()
	result.Find("br").Each(func(i int, s *goquery.Selection) {
		s.ReplaceWithHtml("\n")
	})
	return result.Text()
}

func getPage(url string) io.Reader {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Can't get the page %s with error: %s", url, err.Error())
	}

	page, err := iconv.NewReader(resp.Body, "windows-1251", "utf-8")
	if err != nil {
		log.Print("Can't read the page %s with error: %s", url, err.Error())
	}

	return page
}

package quoter

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"github.com/nikitasmall/simple-bot/config"
)

// Imager is a struct to get specific image from the page.
// Implements a PageParser interface and returns single string as a path to image which
// could be found at the url via query string. fromEncoding should be a string
// that represents page original encoding.
type Imager struct {
	url          string
	query        string
	fromEncoding string
}

func NewImager(url, query, fromEnc string) Imager {
	return Imager{
		url:          url,
		query:        query,
		fromEncoding: fromEnc,
	}
}

func (a Imager) SavePicture() (string, error) {
	imageUrl, err := a.GetPageResult()
	if err != nil {
		return "", err
	}

	resp, err := http.Get(imageUrl)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	imagePath := "public/pic/" + config.RandStringBytes(32) + ".jpg"
	file, err := os.Create(imagePath)
	defer file.Close()
	if err != nil {
		return "", err
	}

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		return "", err
	}

	return imagePath, nil
}

// function parse a result page (provided by reader) and returns an
// src of first single element contained in html-query (jQuery-like string) (it should be an image).
// Part of PageParser interface Implementation for Imager struct.
func (a Imager) GetPageResult() (string, error) {
	page, err := a.getPage()
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		log.Print(err.Error())
		return "", err
	}

	result := doc.Find(a.query).First()
	attr, ok := result.Attr("src")
	if !ok {
		log.Printf("Can't get image src from element '%s' in url %s", a.query, a.url)
		return "", errors.New("Can't get source of a picture at " + a.url)
	}

	return attr, nil
}

// function gets a page at the url provided by quoter instance
// and returns reader that represents body of the response page encoded in utf-8.
// Part of PageParser interface Implementation for Imager struct.
func (a Imager) getPage() (io.Reader, error) {
	resp, err := http.Get(a.url)
	if err != nil {
		log.Printf("Can't get the page %s with error: %s", a.url, err.Error())
		return nil, err
	}

	page, err := iconv.NewReader(resp.Body, a.fromEncoding, "utf-8")
	if err != nil {
		log.Print("Can't read the page %s with error: %s", a.url, err.Error())
		return nil, err
	}

	return page, nil
}

package quoter

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

// Imager is a struct to get specific image from the page.
// Implements a PageParser interface and returns single string as a path to image which
// could be found at the url via query string. fromEncoding should be a string
// that represents page original encoding. imageIndex is a number of image at the page.
type Imager struct {
	url          string
	query        string
	fromEncoding string
	imageIndex   int
}

// returns new imager instance.
func NewImager(url, query, fromEnc string, imageIndex int) Imager {
	return Imager{
		url:          url,
		query:        query,
		fromEncoding: fromEnc,
		imageIndex:   imageIndex,
	}
}

// function uploads a picture from provided by imager page and query
// and returns (opened!) resp.Body that implements ReadCloser interface.
func (a Imager) UploadPicture() (io.ReadCloser, error) {
	imageUrl, err := a.getPageResult()
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(imageUrl)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

// function parse a result page (provided by reader) and returns an
// src of first single element contained in html-query (jQuery-like string) (it should be an image).
// Part of PageParser interface Implementation for Imager struct.
func (a Imager) getPageResult() (string, error) {
	page, err := a.getPage()
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		log.Print(err.Error())
		return "", err
	}

	result := doc.Find(a.query).Eq(a.imageIndex)
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

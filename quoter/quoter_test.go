package quoter

import (
	"strings"
	"testing"
)

func TestGetCorrectPage(t *testing.T) {
	quoter := Quoter{
		url:          "https://golang.org/",
		query:        "div#heading-wide a",
		fromEncoding: "utf-8",
	}

	res, err := quoter.getPageResult()
	if err != nil {
		t.Errorf("Error on parse correct page: %s", err.Error())
	}

	if res != "The Go Programming Language" {
		t.Errorf("Wrong parse result: %s", res)
	}
}

func TestGetWrongPage(t *testing.T) {
	quoter := Quoter{
		url:          "https://fidelsnaff.net/",
		query:        "div",
		fromEncoding: "utf-8",
	}

	res, err := quoter.getPageResult()
	if err == nil {
		t.Error("Error doesn't rise on wrong page parse!")
	}

	if res != "" {
		t.Errorf("Some parse result, while should be empty: %s", res)
	}
}

func TestGetWrongQuery(t *testing.T) {
	quoter := Quoter{
		url:          "https://golang.org/",
		query:        "div#fidelsnaff",
		fromEncoding: "utf-8",
	}

	res, err := quoter.getPageResult()
	if err != nil {
		t.Error("Error on parse correct page: %s", err.Error())
	}

	if res != "" {
		t.Errorf("Some parse result, while should be empty: %s", res)
	}
}

func TestGetWrongEncoging(t *testing.T) {
	quoter := Quoter{
		url:          "https://golang.org/",
		query:        "div#fidelsnaff",
		fromEncoding: "win-1252",
	}

	_, err := quoter.getPageResult()
	if err == nil {
		t.Error("Error is omitted when work with wrong encoding.")
	}
}

func TestBashQuote(t *testing.T) {
	quote, err := GetRandomQuote()
	if err != nil {
		t.Errorf("Error on parse correct page: %s", err.Error())
	}

	if quote == "" {
		t.Errorf("Wrong parse result: %s", quote)
	}
}

func TestWeatherQuote(t *testing.T) {
	weather, err := GetCurrentWeather("london")
	if err != nil {
		t.Errorf("Error on parse correct page: %s", err.Error())
	}

	if weather == "" {
		t.Errorf("Wrong parse result: %s", weather)
	}
}

func TestImager(t *testing.T) {
	imager := NewImager(
		"https://github.com/alexyer", ".vcard-avatar img", "utf-8", 0)

	imageUrl, err := imager.getPageResult()
	if err != nil {
		t.Errorf("Error on image url getting: %s", err.Error())
	}

	if len(imageUrl) == 0 || !strings.Contains(imageUrl, "http") {
		t.Error("Wrong image url gathered.")
	}

	_, err = imager.UploadPicture()
	if err != nil {
		t.Errorf("Can't get image body with error: %s", err.Error())
	}
}

func TestWrongImager(t *testing.T) {
	imager := NewImager(
		"https://fidelsnaff", ".vcard-avatar img", "utf-8", 0)

	_, err := imager.UploadPicture()
	if err == nil {
		t.Error("Error is omitted when work with wrong url.")
	}

	imager = NewImager(
		"https://github.com/alexyer", "no image there", "utf-8", 0)

	_, err = imager.UploadPicture()
	if err == nil {
		t.Error("Error is omitted when work with wrong query.")
	}

	imager = NewImager(
		"https://github.com/alexyer", "no image there", "win-1252", 0)

	_, err = imager.UploadPicture()
	if err == nil {
		t.Error("Error is omitted when work with wrong encoding.")
	}
}

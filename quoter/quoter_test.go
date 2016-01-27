package quoter

import (
	"testing"
)

func TestGetCorrectPage(t *testing.T) {
	quoter := Quoter{
		url:          "https://golang.org/",
		query:        "div#heading-wide a",
		fromEncoding: "utf-8",
	}

	res, err := quoter.GetPageResult()
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

	res, err := quoter.GetPageResult()
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

	res, err := quoter.GetPageResult()
	if err != nil {
		t.Error("Error on parse correct page: %s", err.Error())
	}

	if res != "" {
		t.Errorf("Some parse result, while should be empty: %s", res)
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

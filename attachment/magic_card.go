package attachment

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/nikitasmall/simple-bot/config"
	"github.com/nikitasmall/simple-bot/quoter"
)

// prepared gatherer for random card.
// url and query for random card is always the same.
var RandomCardGatherer = quoter.NewImager(
	"http://magiccards.info/random.html",
	"img",
	"utf-8",
	1)

// function replaces randomGatherer by gatherer with specific cardName,
// after that it makes the same request and returns a filePath of new image or error.
func GetCardByName(cardName string) (string, error) {
	cardGatherer := newImagerWithCardName(cardName)

	return GetCard(cardGatherer)
}

// upload random image with provided Imager params and destroy it after 15 seconds,
// returns a path to uploaded file.
func GetCard(garherer quoter.Imager) (string, error) {
	body, err := garherer.UploadPicture()
	if err != nil {
		return "", err
	}

	imagePath, err := saveFile("public/pic/", body)

	go deleteFile(imagePath)
	return imagePath, err
}

// function takes dirPath and readCloser which trying to save first and close at the end.
// Returns a filePath to new saved image. FileName is random.
func saveFile(dirPath string, body io.ReadCloser) (string, error) {
	imagePath := dirPath + config.RandStringBytes(32) + ".jpg"
	defer body.Close()

	file, err := os.Create(imagePath)
	defer file.Close()

	if err != nil {
		return "", err
	}

	_, err = io.Copy(file, body)
	if err != nil {
		return "", err
	}

	return imagePath, nil
}

// function removes file provided by filePath after 15 seconds of calling.
// this time should be enougth to upload this file to server.
// use this in a side go-routine when working with a bulk random files.
func deleteFile(filePath string) {
	time.Sleep(15)
	os.Remove(filePath)
}

// function creates a new imager with specific url which will guide to seeked card.
// The same Imager in all other senses.
func newImagerWithCardName(cardName string) quoter.Imager {
	urlRequest := "http://magiccards.info/query?q=" + strings.Replace(cardName, " ", "+", -1)
	return quoter.NewImager(
		urlRequest,
		"img",
		"utf-8",
		1)
}

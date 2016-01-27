package attachment

import (
	"os"
	"time"

	"github.com/nikitasmall/simple-bot/quoter"
)

var randomCardGatherer = quoter.NewImager(
	"http://gatherer.wizards.com/Pages/Card/Details.aspx?action=random",
	".cardImage img",
	"utf-8")

// upload random image with provided Imager params and destroy it after 15 seconds,
// returns a path to uploaded file.
func GetRandomCard() (string, error) {
	imagePath, err := randomCardGatherer.SavePicture()

	go deleteFile(imagePath)
	return imagePath, err
}

// function removes file provided by filePath after 15 seconds of calling.
// this time should be enougth to upload this file to server.
// use this in a side go-routine when working with a bulk random files.
func deleteFile(filePath string) {
	time.Sleep(15)
	os.Remove(filePath)
}

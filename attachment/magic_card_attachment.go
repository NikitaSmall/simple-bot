package attachment

import (
	"strings"

	"github.com/nikitasmall/simple-bot/config"
	"github.com/nikitasmall/simple-bot/quoter"
)

// struct that holds imager with specific params and information about card:
// is it random or not.
type magicCardAttachment struct {
	imager        quoter.Imager
	temporaryCard bool
}

// function returns new magicCardAttachment struct for further work.
func newMagicCardAttachment(imager quoter.Imager, temporaryCard bool) magicCardAttachment {
	return magicCardAttachment{
		imager:        imager,
		temporaryCard: temporaryCard,
	}
}

// prepared gatherer for random card.
// url and query for random card is always the same.
var randomCardGatherer = newMagicCardAttachment(quoter.NewImager(
	"http://magiccards.info/random.html",
	"img",
	"utf-8",
	1),
	true)

// function replaces randomGatherer by gatherer with specific cardName,
// after that it makes the same request and returns a filePath of new image or error.
func GetCardByName(cardName string) (string, error) {
	cardGatherer := newMagicCardAttachment(newImagerWithCardName(cardName), true)
	return cardGatherer.GetAttachmentPath()
}

// function makes request to get a random card from online database
// and returns a filePath of new image or error.
func GetRandomCard() (string, error) {
	return randomCardGatherer.GetAttachmentPath()
}

// upload  and save image with provided Imager params,
// returns a path to uploaded file.
// delete image after short upload time if attachment marked as temporary.
func (a magicCardAttachment) GetAttachmentPath() (string, error) {
	body, err := a.imager.UploadPicture()
	if err != nil {
		return "", err
	}

	imagePath, err := saveFile(config.Env["magicCardPath"], body)

	if a.temporaryCard {
		go deleteFile(imagePath)
	}
	return imagePath, err
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

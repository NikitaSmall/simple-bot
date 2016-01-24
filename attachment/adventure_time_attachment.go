package attachment

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/nikitasmall/simple-bot/config"
)

// struct that represents stickers storage
// tries to find attachments at base path plus `/adventure_time`
// and counts stickers number.
type AdventureTimeAttachment struct {
	path          string
	stickersCount int
}

// storage
var AdventureTimeStickers *AdventureTimeAttachment

// initializing
func newAdventureTimeAttachment(basePath string) (*AdventureTimeAttachment, error) {
	path := basePath + "/adventure_time"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("Can't read adventure folder with error: %s" + err.Error())
		return nil, errors.New("Can't read adventure folder with error: " + err.Error())
	}

	if len(files) == 0 {
		log.Print("Adventure folder is empty!")
		return nil, errors.New("Stickers folder is empty!")
	}

	return &AdventureTimeAttachment{
		path:          path,
		stickersCount: len(files),
	}, nil
}

// function returns a path to upload random adventure time sticker
func (at *AdventureTimeAttachment) GetAttachmentPath() string {
	if at == nil {
		var err error
		at, err = newAdventureTimeAttachment(config.Env["attachmentAdventureTime"])
		if err != nil {
			return "Adventure time! Sorry, no stickers for today!"
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	stickerNumber := r.Intn(at.stickersCount)

	return fmt.Sprintf("%s/%d.jpg", at.path, stickerNumber)
}

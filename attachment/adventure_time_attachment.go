package attachment

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

// struct that represents stickers storage
// tries to find attachments at base path plus `/adventure_time`
// and counts stickers number.
type AdventureTimeAttachment struct {
	path          string
	stickersCount int
}

// storage
var AdventureTimeStickers = newAdventureTimeAttachment("public/pic")

// initializing
func newAdventureTimeAttachment(basePath string) AdventureTimeAttachment {
	path := basePath + "/adventure_time"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Panic("Can't read adventure folder with error: " + err.Error())
	}

	if len(files) == 0 {
		log.Panic("Adventure folder is empty!")
	}

	return AdventureTimeAttachment{
		path:          path,
		stickersCount: len(files),
	}
}

// function returns a path to upload random adventure time sticker
func (at AdventureTimeAttachment) GetAttachmentPath() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	stickerNumber := r.Int31n(int32(at.stickersCount))

	return fmt.Sprintf("%s/%d.jpg", at.path, stickerNumber)
}

package attachment

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/nikitasmall/simple-bot/config"
)

// folders for these tests are hidden by gitignore
func TestNewATAttachment(t *testing.T) {
	atStickers, err := newAdventureTimeAttachment("test_pics")
	if err != nil {
		t.Errorf("Error on atAttachment initialization: %s", err.Error())
	}

	if !strings.Contains(atStickers.path, "adventure_time") {
		t.Errorf("atAttachment created with wrong path: %s", atStickers.path)
	}

	stickers, err := ioutil.ReadDir(atStickers.path)
	if err != nil {
		t.Errorf("Error on stickers counting: %s", err.Error())
	}

	if atStickers.stickersCount != len(stickers) {
		t.Error("wrong number of stickers")
	}
}

func TestNewATAttachmentFail(t *testing.T) {
	_, err := newAdventureTimeAttachment("test_pics/pic")
	if err == nil {
		t.Errorf("Error doesn't rise on atAttachment initialization, when folder is empty.")
	}

	_, err = newAdventureTimeAttachment("test_pics/pic2")
	if err == nil {
		t.Errorf("Error doesn't rise on atAttachment initialization, when folder doesn't exist.")
	}
}

func TestATAttachmentGetPath(t *testing.T) {
	config.Env["attachmentAdventureTime"] = "test_pics"
	var AtStickers *AdventureTimeAttachment

	path, err := AtStickers.GetAttachmentPath()
	if err != nil {
		t.Errorf("Error on path getting: %s", err.Error())
	}

	if !strings.Contains(path, "adventure_time") || !strings.Contains(path, ".jpg") {
		t.Error(AtStickers)
		t.Errorf("wrong path format: %s", path)
	}
}

func TestATAttachmentGetPathWrong(t *testing.T) {
	var AtStickers *AdventureTimeAttachment
	config.Env["attachmentAdventureTime"] = "test_pics/pic"

	path, err := AtStickers.GetAttachmentPath()
	if err == nil {
		t.Errorf("Error on path getting: %s", err.Error())
	}

	if !strings.Contains(err.Error(), "Adventure time! Sorry, no stickers for today!") {
		t.Errorf("Object was initialized. Path: %s", path)
	}

	config.Env["attachmentAdventureTime"] = "test_pics/pic2"

	path, err = AtStickers.GetAttachmentPath()
	if err == nil {
		t.Errorf("Error on path getting: %s", err.Error())
	}

	if !strings.Contains(err.Error(), "Adventure time! Sorry, no stickers for today!") {
		t.Errorf("Object was initialized. Path: %s", path)
	}
}

func TestRandomMagicCardAttachment(t *testing.T) {
	config.Env["magicCardPath"] = "test_pics/pic"

	cardNameOne, err := randomCardGatherer.GetAttachmentPath()
	if err != nil {
		t.Errorf("Error on card getting, with error: %s", err.Error())
	}

	if len(cardNameOne) == 0 || !strings.Contains(cardNameOne, config.Env["magicCardPath"]) {
		t.Errorf("Wrong cardName format: %s", cardNameOne)
	}

	cardNameTwo, err := GetRandomCard()
	if err != nil {
		t.Errorf("Error on card getting, with error: %s", err.Error())
	}

	if cardNameOne == cardNameTwo {
		t.Errorf("cards has the same names: %s", cardNameOne)
	}
}

func TestDeterminatedMagicCardAttachment(t *testing.T) {
	config.Env["magicCardPath"] = "test_pics/pic"
	gatherer := newMagicCardAttachment(newImagerWithCardName("kor castigator"), true)

	imagePath, err := gatherer.GetAttachmentPath()
	if err != nil {
		t.Errorf("Error on image getting, with error: %s", err.Error())
	}

	if len(imagePath) == 0 || !strings.Contains(imagePath, config.Env["magicCardPath"]) {
		t.Errorf("Wrong cardName format: %s", imagePath)
	}

	_, err = GetCardByName("kor castigator")
	if err != nil {
		t.Errorf("Error on image getting, with error: %s", err.Error())
	}
}

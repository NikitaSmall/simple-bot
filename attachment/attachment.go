package attachment

import (
	"io"
	"os"
	"time"

	"github.com/nikitasmall/simple-bot/config"
)

// attachment is simple interface
// that can represent any type of data
// that could be uploaded with telegram bot api.
type Attachment interface {
	GetAttachmentPath() (string, error)
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
	time.Sleep(15 * time.Second)
	os.Remove(filePath)
}

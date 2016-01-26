package bot

import (
	"strings"
	"testing"

	"github.com/nikitasmall/simple-bot/config"
	"gopkg.in/telegram-bot-api.v1"
)

var testBot = CreateBot(config.Env["botToken"])

func testMessage(text string) tgbotapi.Update {
	return tgbotapi.Update{
		UpdateID: 123,
		Message: tgbotapi.Message{
			MessageID: 122,
			From:      tgbotapi.User{},
			Date:      123456,
			Chat: tgbotapi.Chat{
				ID:   123,
				Type: "personal",
			},
			Text: text,
		},
		InlineQuery: tgbotapi.InlineQuery{},
	}
}

func TestNewBot(t *testing.T) {
	newBot := newApiBot(config.Env["botToken"])
	if newBot.Debug != true {
		t.Error("Bot created in production mode")
	}

	if newBot.Self.UserName != "MapToGoBot" {
		t.Errorf("Wrong bot created. Username: %s", newBot.Self.UserName)
	}
}

func TestNewUpdatechan(t *testing.T) {
	newBot := newApiBot(config.Env["botToken"])
	updateChan := newUpdatesChan(newBot)

	select {
	case emptyMessage, ok := <-updateChan:
		if ok {
			t.Error("We have recieve some message!", emptyMessage)
		} else {
			t.Error("Error occured with channel")
		}
	default:
		// no message at testClient channel and all is ok
	}
}

func TestNewCommand(t *testing.T) {
	update := testMessage("hello")
	botCommand := newBotCommand(update)

	if botCommand.command != "hello" {
		t.Error("wrong command parsing")
	}

	if len(botCommand.args) != 0 {
		t.Error("wrong args parsing")
	}
}

func TestCommandExecute(t *testing.T) {
	update := testMessage("/weather odessa")
	botCommand := newBotCommand(update)

	msg := botCommand.execute()
	if msg.(tgbotapi.MessageConfig).Text == "" {
		t.Error("wrong command processing")
	}

	if msg.(tgbotapi.MessageConfig).Text == "Can't get weather for this city." {
		t.Error("Internal error on weather state getting")
	}
}

func TestProcessCommandWeatherWrongArgumentNumber(t *testing.T) {
	update := testMessage("/weather")
	msg := processCommand(update)

	text := msg.(tgbotapi.MessageConfig).Text
	if text != "Wrong number of arguments! Should be one." {
		t.Error("wrong command processing. ", text)
	}
}

func TestProcessCommand(t *testing.T) {
	update := testMessage("/joke")
	msg := processCommand(update)

	if msg.(tgbotapi.MessageConfig).Text == "" {
		t.Error("wrong command processing")
	}

	if msg.(tgbotapi.MessageConfig).Text == "Can't get quote." {
		t.Error("Internal error on joke getting")
	}
}

func TestProcessSimpleText(t *testing.T) {
	update := testMessage("What time is it?")
	msg := processCommand(update)

	text := msg.(tgbotapi.MessageConfig).Text
	if text != "Adventure time! Sorry, no stickers for today!" {
		t.Error("wrong command processing")
	}
}

func TestProcessSticker(t *testing.T) {
	config.Env["attachmentAdventureTime"] = "../public/pic"

	update := testMessage("What time is it?")
	msg := processCommand(update)

	filePath := msg.(tgbotapi.StickerConfig).File.(string)
	if !strings.Contains(filePath, "adventure_time") || !strings.Contains(filePath, ".jpg") {
		t.Error("wrong command processing filePath. ", filePath)
	}
}

func TestProcessBadCommand(t *testing.T) {
	update := testMessage("Fidelsnaff")
	msg := processCommand(update)

	text := msg.(tgbotapi.MessageConfig).Text
	if text != "Don't understand 'Fidelsnaff'" {
		t.Error("wrong command processing")
	}
}

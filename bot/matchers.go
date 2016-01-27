package bot

import (
	"strings"
)

// matchers that helps to select a proper handler for certain user's input

func (bc botCommand) isQuoteRequest() bool {
	return (bc.command == "/quote") || (bc.command == "/joke")
}

func (bc botCommand) isWeatherRequest() bool {
	return (bc.command == "/weather") || (bc.command == "/temp")
}

func (bc botCommand) isAdventureTimeRequest() bool {
	return (bc.command == "/time") ||
		(strings.Contains(strings.ToLower(bc.fullText), "what time is it"))
}

func (bc botCommand) isMagicCardRequest() bool {
	return (bc.command == "/magic") || (bc.command == "/m")
}

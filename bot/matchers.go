package bot

import (
	"strings"
)

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

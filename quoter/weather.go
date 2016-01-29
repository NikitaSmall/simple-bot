package quoter

import (
	"fmt"

	"github.com/nikitasmall/simple-bot/config"
)

// openWeatherQuoter is a Quoter instance which is able to get weather from openWeather API
var openWeatherQuoter = Quoter{
	url:          fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?mode=html&appid=%s&q=", config.Env["weatherApiKey"]),
	query:        "div[title='Current Temperature']",
	fromEncoding: "utf-8",
}

// GetCurrentWeather function returns a current temperature state for provided city.
// It returns empty string in case of wrong city-name.
func GetCurrentWeather(city string) (string, error) {
	w := openWeatherQuoter
	w.url += city

	return w.getPageResult()
}

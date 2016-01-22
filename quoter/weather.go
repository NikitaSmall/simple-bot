package quoter

import (
	"fmt"
	"github.com/nikitasmall/simple-bot/config"
)

var openWeatherQuoter = Quoter{
	url:          fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?mode=html&appid=%s&q=", config.Env["weatherApiKey"]),
	query:        "div[title='Current Temperature']",
	fromEncoding: "utf-8",
}

func GetCurrentWeather(city string) string {
	w := openWeatherQuoter
	w.url += city

	return w.getPageResult()
}

package config

import (
	"github.com/vrischmann/envconfig"
	"log"
)

var conf struct {
	Bot struct {
		Token string `envconfig:"default=150776353:AAH5JLFPP6-Qp0HBGXWFr5NOunyjcBDpxfA"`
		Mode  string `envconfig:"default=debug"`
	}
}

// initialized env configs
var Env = initializeConfig()

// initialize config and turns it to map
func initializeConfig() map[string]string {
	err := envconfig.Init(&conf)
	if err != nil {
		log.Panic("Error on env config initialize! ", err.Error())
	}

	return map[string]string{
		"botToken": conf.Bot.Token,
		"botMode":  conf.Bot.Mode,
	}
}

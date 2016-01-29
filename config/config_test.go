package config

import (
	"os"
	"testing"
)

func TestEnvironmentDefaultValues(t *testing.T) {
	defaultEnv := initializeConfig()

	if defaultEnv["botMode"] != "debug" {
		t.Error("Wrong botMode default env! ", defaultEnv["botMode"])
	}

	if defaultEnv["quoteSource"] != "http://bash.im/random" {
		t.Error("Wrong quoteSource default env! ", defaultEnv["quoteSource"])
	}
}

func TestEnvironmentCustomValues(t *testing.T) {
	os.Setenv("BOT_MODE", "prod")
	os.Setenv("QUOTE_SOURCE", "http://custom.com")

	customEnv := initializeConfig()
	if customEnv["botMode"] != "prod" {
		t.Error("Wrong botMode! ", customEnv["botMode"])
	}

	if customEnv["quoteSource"] != "http://custom.com" {
		t.Error("Wrong quoteSource env! ", customEnv["quoteSource"])
	}
}

func TestRandomName(t *testing.T) {
	length := 16
	stringName := RandStringBytes(length)

	if len(stringName) != length {
		t.Error("RandStringBytes returns string with wrong length!")
	}

	secondStringName := RandStringBytes(length)
	if stringName == secondStringName {
		t.Error("RandStringBytes creates same strings")
	}
}

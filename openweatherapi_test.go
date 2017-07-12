package openweatherapi

import (
	"io/ioutil"
	"testing"
)

const apiKeyFile = "openweather.key"
const cityBerlin = "Berlin,de"

func readAPIKey() string {
	key, err := ioutil.ReadFile(apiKeyFile)
	if err != nil {
		panic("cannot run test, you must provide openweathermap api key")
	}
	return string(key)
}

func TestDownloadWeatherData(t *testing.T) {
	resp, err := DownloadWeatherData(readAPIKey(), cityBerlin)
	if err != nil {
		t.Error("error while retrieving data: " + err.Error())
	} else if len(resp) == 0 {
		t.Error("received data is empty")
	}
}

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

func TestNewQuery(t *testing.T) {
	// arrange
	apiKeyFile := readAPIKey()
	location := cityBerlin
	unit := "imperial"

	// action
	q := NewQuery(apiKeyFile, location)

	// verify
	if q.APIKey != apiKeyFile || q.Location != location || q.Unit != "metric" {
		t.Error("query and query params do not match")
	}

	// action 2
	q = NewQuery(apiKeyFile, location, unit)

	// verify 2
	if q.APIKey != apiKeyFile || q.Location != location || q.Unit != unit {
		t.Error("query and query params do not match")
	}
}

func TestForecast(t *testing.T) {
	// arrange
	q := NewQuery(readAPIKey(), cityBerlin)

	// action
	resp, err := q.Forecast()

	// verify
	if err != nil {
		t.Error("error while retrieving data: " + err.Error())
	} else if len(resp) == 0 {
		t.Error("received data is empty")
	}
}

func TestWeather(t *testing.T) {
	// arrange
	q := NewQuery(readAPIKey(), cityBerlin)

	// action
	resp, err := q.Weather()

	// verify
	if err != nil {
		t.Error("error while retrieving data: " + err.Error())
	} else if len(resp) == 0 {
		t.Error("received data is empty")
	}
}

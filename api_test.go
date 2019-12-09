package openweather

import (
	"io/ioutil"
	"testing"

	"github.com/EricNeid/openweather/internal/test"
)

const apiKeyFile = "testdata/api.key"
const cityBerlin = "Berlin,de"

func readAPIKey() string {
	key, err := ioutil.ReadFile(apiKeyFile)
	if err != nil {
		panic(`
			Cannot run test, you must provide openweathermap api key. 
			Expected: testdata/api.key

			See https://home.openweathermap.org/users/sign_up
			for information how to obtain a key`)
	}
	return string(key)
}

func TestForecastRaw(t *testing.T) {
	// arrange
	q := NewQueryForCity(readAPIKey(), cityBerlin)
	// action
	resp, err := q.DailyForecast5Raw()
	// verify
	test.Ok(t, err)
	test.Assert(t, len(resp) > 0, "Received empty response")
}

func TestWeatherRaw(t *testing.T) {
	// arrange
	q := NewQueryForCity(readAPIKey(), cityBerlin)
	// action
	resp, err := q.WeatherRaw()
	// verify
	test.Ok(t, err)
	test.Assert(t, len(resp) > 0, "Received empty response")
}

func TestWeather(t *testing.T) {
	// arrange
	q := NewQueryForCity(readAPIKey(), cityBerlin)
	// action
	data, err := q.Weather()
	// verify
	test.Ok(t, err)
	test.Equals(t, "Berlin", data.Name)
}

func TestDailyForecast(t *testing.T) {
	// arrange
	q := NewQueryForCity(readAPIKey(), cityBerlin)
	// action
	data, err := q.DailyForecast5()
	// verify
	test.Ok(t, err)
	test.Equals(t, "Berlin", data.City.Name)
}

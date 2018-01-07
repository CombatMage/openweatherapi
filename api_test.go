package openweatherapi

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)
	assert.True(t, len(resp) > 0)
}

func TestWeatherRaw(t *testing.T) {
	// arrange
	q := NewQueryForCity(readAPIKey(), cityBerlin)
	// action
	resp, err := q.WeatherRaw()
	// verify
	assert.NoError(t, err)
	assert.True(t, len(resp) > 0)
}

func TestWeather(t *testing.T) {
	// arrange
	q := NewQueryForCity(readAPIKey(), cityBerlin)
	// action
	data, err := q.Weather()
	// verify
	assert.NoError(t, err)
	assert.Equal(t, "Berlin", data.Name)
}

func TestDailyForecast(t *testing.T) {
	// arrange
	q := NewQueryForCity(readAPIKey(), cityBerlin)
	// action
	data, err := q.DailyForecast5()
	// verify
	assert.NoError(t, err)
	assert.Equal(t, "Berlin", data.City.Name)
}

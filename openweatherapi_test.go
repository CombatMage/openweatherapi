package openweatherapi

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const apiKeyFile = "api.key"
const cityBerlin = "Berlin,de"

func readAPIKey() string {
	key, err := ioutil.ReadFile(apiKeyFile)
	if err != nil {
		panic(`cannot run test, you must provide openweathermap api key. 
			See https://home.openweathermap.org/users/sign_up`)
	}
	return string(key)
}

func TestNewQueryForCity(t *testing.T) {
	// arrange
	apiKey := readAPIKey()
	location := cityBerlin
	// action
	q := NewQueryForCity(apiKey, location)
	// verify
	assert.Equal(t, apiKey, q.APIKey)
	assert.Equal(t, location, q.Query)
	assert.Equal(t, "metric", q.Unit)

	// arrange
	unit := "imperial"
	// action
	q = NewQueryForCity(apiKey, location, unit)
	// verify
	assert.Equal(t, apiKey, q.APIKey)
	assert.Equal(t, location, q.Query)
	assert.Equal(t, unit, q.Unit)
}

func TestForecastRaw(t *testing.T) {
	// arrange
	q := NewQueryForCity(readAPIKey(), cityBerlin)
	// action
	resp, err := q.DailyForecastRaw()
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
	data, err := q.DailyForecast()
	// verify
	assert.NoError(t, err)
	assert.Equal(t, "Berlin", data.City.Name)
}

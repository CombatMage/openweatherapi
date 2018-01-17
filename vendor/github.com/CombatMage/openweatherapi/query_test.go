package openweatherapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueryForCity(t *testing.T) {
	// arrange
	apiKey := "testKey"
	location := cityBerlin
	// action
	q := NewQueryForCity(apiKey, location)
	// verify
	assert.Equal(t, apiKey, q.APIKey)
	assert.Equal(t, location, q.Query)
	assert.Equal(t, "metric", q.Unit)
	assert.Equal(t, queryTypeCity, q.queryType)

	// arrange
	unit := "imperial"
	// action
	q = NewQueryForCity(apiKey, location, unit)
	// verify
	assert.Equal(t, unit, q.Unit)
}

func TestNewQueryForZip(t *testing.T) {
	// arrange
	apiKey := "testKey"
	zip := "12345"
	// action
	q := NewQueryForZip(apiKey, zip)
	// verify
	assert.Equal(t, apiKey, q.APIKey)
	assert.Equal(t, zip, q.Query)
	assert.Equal(t, queryTypeZip, q.queryType)
}

func TestNewQueryForID(t *testing.T) {
	// arrange
	apiKey := "testKey"
	id := "42"
	// action
	q := NewQueryForID(apiKey, id)
	// verify
	assert.Equal(t, apiKey, q.APIKey)
	assert.Equal(t, id, q.Query)
	assert.Equal(t, queryTypeID, q.queryType)
}

func TestNewQueryForLocation(t *testing.T) {
	// arrange
	apiKey := "testKey"
	lat := "51"
	lon := "13"
	// action
	q := NewQueryForLocation(apiKey, lat, lon)
	// verify
	assert.Equal(t, apiKey, q.APIKey)
	assert.Equal(t, lat+"|"+lon, q.Query)
	assert.Equal(t, queryTypeGeo, q.queryType)
}

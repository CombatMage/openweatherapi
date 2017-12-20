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

	// arrange
	unit := "imperial"
	// action
	q = NewQueryForCity(apiKey, location, unit)
	// verify
	assert.Equal(t, apiKey, q.APIKey)
	assert.Equal(t, location, q.Query)
	assert.Equal(t, unit, q.Unit)
}

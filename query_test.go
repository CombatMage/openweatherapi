package openweather

import (
	"testing"

	"github.com/EricNeid/openweather/internal/test"
)

func TestNewQueryForCity(t *testing.T) {
	// arrange
	apiKey := "testKey"
	location := cityBerlin
	// action
	q := NewQueryForCity(apiKey, location)
	// verify
	test.Equals(t, apiKey, q.APIKey)
	test.Equals(t, location, q.Query)
	test.Equals(t, "metric", q.Unit)
	test.Equals(t, queryTypeCity, q.queryType)

	// arrange
	unit := "imperial"
	// action
	q = NewQueryForCity(apiKey, location, unit)
	// verify
	test.Equals(t, unit, q.Unit)
}

func TestNewQueryForZip(t *testing.T) {
	// arrange
	apiKey := "testKey"
	zip := "12345"
	// action
	q := NewQueryForZip(apiKey, zip)
	// verify
	test.Equals(t, apiKey, q.APIKey)
	test.Equals(t, zip, q.Query)
	test.Equals(t, queryTypeZip, q.queryType)
}

func TestNewQueryForID(t *testing.T) {
	// arrange
	apiKey := "testKey"
	id := "42"
	// action
	q := NewQueryForID(apiKey, id)
	// verify
	test.Equals(t, apiKey, q.APIKey)
	test.Equals(t, id, q.Query)
	test.Equals(t, queryTypeID, q.queryType)
}

func TestNewQueryForLocation(t *testing.T) {
	// arrange
	apiKey := "testKey"
	lat := "51"
	lon := "13"
	// action
	q := NewQueryForLocation(apiKey, lat, lon)
	// verify
	test.Equals(t, apiKey, q.APIKey)
	test.Equals(t, lat+"|"+lon, q.Query)
	test.Equals(t, queryTypeGeo, q.queryType)
}

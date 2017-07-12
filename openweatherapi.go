package openweatherapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Query represents a pending request to openweathermap
type Query struct {
	APIKey    string
	Unit      string
	Query     string
	queryType string
}

type weatherAPI interface {
	Forecast() (json string, err error)
	Weather() (json string, err error)
}

// NewQueryForCity creates a query for openweathermap from city name
func NewQueryForCity(apiKey string, city string, unit ...string) Query {
	u := "metric"
	if len(unit) > 0 {
		u = unit[0]
	}
	return Query{
		APIKey:    apiKey,
		Query:     city,
		queryType: "q",
		Unit:      u,
	}
}

// NewQueryForZip creates a query for openweathermap from zip code
func NewQueryForZip(apiKey string, zip string, unit ...string) Query {
	u := "metric"
	if len(unit) > 0 {
		u = unit[0]
	}
	return Query{
		APIKey:    apiKey,
		Query:     zip,
		queryType: "zip",
		Unit:      u,
	}
}

// NewQueryForID creates a query for openweathermap from city id
func NewQueryForID(apiKey string, id string, unit ...string) Query {
	u := "metric"
	if len(unit) > 0 {
		u = unit[0]
	}
	return Query{
		APIKey:    apiKey,
		Query:     id,
		queryType: "id",
		Unit:      u,
	}
}

// NewQueryForLocation creates a query for openweathermap from latitude and longitude
func NewQueryForLocation(apiKey string, lat string, lon string, unit ...string) Query {
	u := "metric"
	if len(unit) > 0 {
		u = unit[0]
	}
	return Query{
		APIKey:    apiKey,
		Query:     lat + "|" + lon,
		queryType: "lat|lon",
		Unit:      u,
	}
}

// Weather downloads current weather data from
// openweathermap and return them as string
func (query Query) Weather() (json string, err error) {
	return downloadString(weatherURL(query))
}

// Forecast downloads forecast data from
// openweathermap and return them as string
func (query Query) Forecast() (json string, err error) {
	return downloadString(forecastURL(query))
}

func downloadString(url string) (res string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func forecastURL(q Query) string {
	return "http://api.openweathermap.org/data/2.5/forecast/daily" + formatURLQuery(q)
}

func weatherURL(q Query) string {
	return "http://api.openweathermap.org/data/2.5/weather" + formatURLQuery(q)
}

func formatURLQuery(q Query) string {
	queryType := q.queryType
	queryValue := q.Query
	var query string

	if queryType == "lat|lon" {
		params := strings.Split(queryValue, "|") // expected format is lat|long
		lat := params[0]
		lon := params[1]
		query = fmt.Sprintf("?lat=%s&lon=%s", lat, lon)
	} else {
		query = fmt.Sprintf("?%s=%s", queryType, queryValue)
	}

	query = query + fmt.Sprintf("&appid=%s&units=%s", q.APIKey, q.Unit)
	return query
}

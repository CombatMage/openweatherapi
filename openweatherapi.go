package openweatherapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Query represents a pending request to openweathermap
type Query struct {
	APIKey   string
	Location string
	Unit     string
}

type weatherAPI interface {
	Forecast() (json string, err error)
	Weather() (json string, err error)
}

// NewQuery creates a query for openweathermap
func NewQuery(apiKey string, location string, unit ...string) Query {
	u := "metric"
	if len(unit) > 0 {
		u = unit[0]
	}

	return Query{
		APIKey:   apiKey,
		Location: location,
		Unit:     u,
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
	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/forecast/daily"+
			"?q=%s"+
			"&appid=%s"+
			"&units=%s", q.Location, q.APIKey, q.Unit)
	return url
}

func weatherURL(q Query) string {
	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather"+
			"?q=%s"+
			"&appid=%s"+
			"&units=%s", q.Location, q.APIKey, q.Unit)
	return url
}

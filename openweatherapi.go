package openweatherapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Query struct {
	APIKey   string
	Location string
}

func forecastURL(q Query) string {
	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/forecast/daily"+
			"?q=%s&appid=%s&units=metric", q.Location, q.APIKey)
	return url
}

// DownloadWeatherData downloads forecast data from
// openweathermap and return them as string
func DownloadWeatherData(query Query) (json string, err error) {
	resp, err := http.Get(forecastURL(query))
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

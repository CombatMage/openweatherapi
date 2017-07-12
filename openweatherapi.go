package openweatherapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// ForecastURL returns a string, used
// url to query openweathermap.
func ForecastURL(apiKey string, location string) string {
	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/forecast/daily"+
			"?q=%s&appid=%s&units=metric", location, apiKey)
	return url
}

// DownloadWeatherData downloads forecast data from
// openweathermap and return them as string
func DownloadWeatherData(apiKey string, location string) (json string, err error) {
	resp, err := http.Get(ForecastURL(apiKey, location))
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

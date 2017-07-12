package openweatherapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiKey = "a50b478e72ffeeaf8850a5b72bb68865"
const cityBerlin = "Berlin,de"

// ForecastURL returns a string, used
// url to query openweathermap.
func ForecastURL() string {
	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/forecast/daily"+
			"?q=%s&appid=%s&units=metric", cityBerlin, apiKey)
	return url
}

// DownloadWeatherData downloads forecast data from
// openweathermap and return them as string
func DownloadWeatherData() (json string, err error) {
	resp, err := http.Get(ForecastURL())
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

// Package openweatherapi contains helper functions to query 
// OpenWeatherMaps (http://openweathermap.org/) for weather information.
// Currently the current weather API (http://openweathermap.org/current) and the 
// 5 days forecast API (http://openweathermap.org/forecast5) are supported.
package openweatherapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// WeatherRaw downloads current weather data from openweathermap and return them as string.
func (query Query) WeatherRaw() (json string, err error) {
	bytes, err := download(weatherURL(query))
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Weather downloads current weather data from openweathermap and return them as WeatherData.
func (query Query) Weather() (data CurrentWeather, err error) {
	bytes, err := download(weatherURL(query))
	if err != nil {
		return CurrentWeather{}, err
	}

	data = CurrentWeather{}
	err = json.Unmarshal(bytes, &data)
	return data, err
}

// DailyForecastRaw downloads 5 days forecast data from openweathermap and return them as string.
func (query Query) DailyForecastRaw() (json string, err error) {
	bytes, err := download(dailyForecastURL(query))
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// DailyForecast downloads 5 days forecast data from openweathermap and return them as DailyForecast.
func (query Query) DailyForecast() (data DailyForecast, err error) {
	bytes, err := download(dailyForecastURL(query))
	if err != nil {
		return DailyForecast{}, err
	}
	data = DailyForecast{}
	err = json.Unmarshal(bytes, &data)
	return data, err
}

func download(url string) (res []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// WeatherIconURL returns an url to download matching icon for
// given weather id
func WeatherIconURL(iconID string) (url string) {
	return "http://openweathermap.org/img/w/" + iconID + ".png"
}

func dailyForecastURL(q Query) string {
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

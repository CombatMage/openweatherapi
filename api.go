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
	bytes, err := download(WeatherURL(query))
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Weather downloads current weather data from openweathermap and return them as WeatherData.
func (query Query) Weather() (data CurrentWeather, err error) {
	bytes, err := download(WeatherURL(query))
	if err != nil {
		return CurrentWeather{}, err
	}

	data = CurrentWeather{}
	err = json.Unmarshal(bytes, &data)
	return data, err
}

// DailyForecast5Raw downloads 5 days forecast data from openweathermap and return them as string.
func (query Query) DailyForecast5Raw() (json string, err error) {
	bytes, err := download(DailyForecast5URL(query))
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// DailyForecast5 downloads 5 days forecast data from openweathermap and return them as DailyForecast5.
func (query Query) DailyForecast5() (data DailyForecast5, err error) {
	bytes, err := download(DailyForecast5URL(query))
	if err != nil {
		return DailyForecast5{}, err
	}
	data = DailyForecast5{}
	err = json.Unmarshal(bytes, &data)
	return data, err
}

// DailyForecast16Raw downloads 16 days forecast data from openweathermap and return them as string.
// Warning: the 16 days forecast requires a paid account.
func (query Query) DailyForecast16Raw() (json string, err error) {
	bytes, err := download(DailyForecast16URL(query))
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// DailyForecast16 downloads 16 days forecast data from openweathermap and return them as DailyForecast16.
// Warning: the 16 days forecast requires a paid account.
func (query Query) DailyForecast16() (data DailyForecast16, err error) {
	bytes, err := download(DailyForecast16URL(query))
	if err != nil {
		return DailyForecast16{}, err
	}
	data = DailyForecast16{}
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
// given weather id.
func WeatherIconURL(iconID string) (url string) {
	return "http://openweathermap.org/img/w/" + iconID + ".png"
}

// DailyForecast5URL returns a matching url for the given query which can be used to obtain the 5 days forecast
// from openweathermap.org.
func DailyForecast5URL(q Query) string {
	return "http://api.openweathermap.org/data/2.5/forecast/daily" + formatURLQuery(q)
}

// DailyForecast16URL returns a matching url for the given query which can be used to obtain the 16 days forecast
// from openweathermap.org.
func DailyForecast16URL(q Query) string {
	return "http://api.openweathermap.org/data/2.5/forecast/daily" + formatURLQuery(q) + "&cnt=16"
}

// WeatherURL returns a matching url for the given query which can be used to obtain the current weather information
// from openweathermap.org.
func WeatherURL(q Query) string {
	return "http://api.openweathermap.org/data/2.5/weather" + formatURLQuery(q)
}

func formatURLQuery(q Query) string {
	queryType := q.queryType
	queryValue := q.Query
	var query string

	if queryType == queryTypeGeo {
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

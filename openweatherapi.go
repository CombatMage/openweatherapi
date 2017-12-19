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

// Query represents a pending request to openweathermap.
type Query struct {
	APIKey    string
	Unit      string
	Query     string
	queryType string
}

// CurrentWeather represents unmarshalled data from openweathermap
// for the current weather API (http://openweathermap.org/current).
type CurrentWeather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Rain struct {
		ThreeH int `json:"3h"`
	} `json:"rain"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

// DailyForecast represents unmarshalled data from openweathermap
// for the 5 days forecast weather API (http://openweathermap.org/forecast5).
type DailyForecast struct {
	Cod     string  `json:"cod"`
	Message float64 `json:"message"`
	City    struct {
		GeonameID  int     `json:"geoname_id"`
		Name       string  `json:"name"`
		Lat        float64 `json:"lat"`
		Lon        float64 `json:"lon"`
		Country    string  `json:"country"`
		Iso2       string  `json:"iso2"`
		Type       string  `json:"type"`
		Population int     `json:"population"`
	} `json:"city"`
	Cnt  int `json:"cnt"`
	List []struct {
		Dt   int `json:"dt"`
		Temp struct {
			Day   float64 `json:"day"`
			Min   float64 `json:"min"`
			Max   float64 `json:"max"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"temp"`
		Pressure float64 `json:"pressure"`
		Humidity int     `json:"humidity"`
		Weather  []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Speed  float64 `json:"speed"`
		Deg    int     `json:"deg"`
		Clouds int     `json:"clouds"`
		Snow   float64 `json:"snow,omitempty"`
	} `json:"list"`
}

// NewQueryForCity creates a query for openweathermap from city name.
// The unit is optional and defaults to metric.
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

// NewQueryForZip creates a query for openweathermap from zip code.
// The unit is optional and defaults to metric.
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

// NewQueryForID creates a query for openweathermap from city id.
// The unit is optional and defaults to metric.
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

// NewQueryForLocation creates a query for openweathermap from latitude and longitude.
// The unit is optional and defaults to metric.
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

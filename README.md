# go - openweatherapi

This Repo contains golang library to query OpenWetherMaps (<http://openweathermap.org/>) for weather information.

* current weather: http://openweathermap.org/current
* 5 days forecast: http://openweathermap.org/forecast5

## Install

```bash
go get github.com/EricNeid/openweather
```

## Documentation

Is available on ``godoc``:

<https://godoc.org/github.com/EricNeid/openweather>

## Examples

Consuming the library:

```go
import "github.com/EricNeid/openweather"

// create a query
q := openweather.NewQueryForCity(readAPIKey(), "Berlin,de")

// obtain data
resp, err := q.Weather()

// enjoy
fmt.Println(resp.Name) // Berlin
fmt.Println(resp.Weather[0].Description) // misc
fmt.Println(resp.Main.Temp) // 1
```

See the test files for more example.

A simple client for testing is also included:

```bash
go build cmd/openweatherclient
openweatherclient -key <OpenWeather API Key> -city Berlin,de
```

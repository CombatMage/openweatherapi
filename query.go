package openweatherapi

// Query represents a pending request to openweathermap.
type Query struct {
	APIKey    string
	Unit      string
	Query     string
	queryType string
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

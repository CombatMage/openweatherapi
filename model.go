package openweather

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

// DailyForecast5 represents unmarshalled data from openweathermap
// for the 5 days forecast weather API (http://openweathermap.org/forecast5).
type DailyForecast5 struct {
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

// DailyForecast16 represents unmarshalled data from openweathermap
// for the 16 days forecast weather API (http://openweathermap.org/forecast16).
type DailyForecast16 struct {
	Cod     string  `json:"cod"`
	Message float64 `json:"message"`
	City    struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		} `json:"coord"`
		Country string `json:"country"`
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
	} `json:"list"`
}

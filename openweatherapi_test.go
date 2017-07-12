package openweatherapi

import "testing"

func TestDownloadWeatherData(t *testing.T) {
	resp, err := DownloadWeatherData()
	if err != nil {
		t.Error("error while retrieving data: " + err.Error())
	} else if len(resp) == 0 {
		t.Error("received data is empty")
	}
}

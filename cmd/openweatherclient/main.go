package main

import (
	"flag"
	"fmt"

	"github.com/EricNeid/openweatherapi"
)

func main() {
	keyPtr := flag.String("key", "", "Supply valid api key, see https://home.openweathermap.org/users/sign_up for details.")
	cityPtr := flag.String("city", "", "City to look up")
	flag.Parse()

	if len(*keyPtr) == 0 || len(*cityPtr) == 0 {
		fmt.Println("Usage: simpleclient.exe -key <api-key> -city <city-to-query>")
		return
	}

	query := openweatherapi.NewQueryForCity(*keyPtr, *cityPtr, "metric")
	weather, err := query.WeatherRaw()

	if err != nil {
		fmt.Println("Error while retrieving data: " + err.Error())
		return
	}

	fmt.Println(weather)
}

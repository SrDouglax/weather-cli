package main

import (
	"fmt"
	"os"
	"strings"
	"weather-cli/requests"
)

func main() {

	// Get user query
	var local string = strings.Join(os.Args[1:], " ")

	// Hint if not provide query
	if len(os.Args) < 2 {
		fmt.Println("Use: go run main.go <query>")
		return
	}

	fmt.Println("Getting Coordinates...")

	// Get lat/lon from local query
	coord, err := requests.GetLocalCoordinates(local)
	if err != nil {
		fmt.Printf("[Main]: %v\n", err)
		return
	}

	fmt.Println("Getting weather...")

	// Get weather data based on lat/lon
	weather, err := requests.GetWeather(coord[0])
	if err != nil {
		fmt.Printf("[Main]: %v\n", err)
		return
	}

	// Print results
	fmt.Printf("\nTemperature: %v°\n", weather.Temperature)
}

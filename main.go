package main

import (
	"flag"
	"fmt"
	"strings"
	"weather-cli/api"
	"weather-cli/utils"
)

func main() {

	// Define args
	formatPtr := flag.String("format", "txt", "ex: json, txt")

	// Parse args
	flag.Parse()
	args := flag.Args()

	// Hint if not provide query
	if len(args) == 0 {
		fmt.Println("Use: go run main.go <query>")
		return
	}

	// Get user query and other params
	local := strings.Join(args, " ")
	format := *formatPtr

	fmt.Println("Getting Coordinates...")

	// Get lat/lon from local query
	coord, err := api.GetLocalCoordinates(local)
	if err != nil {
		fmt.Printf("[Main]: %v\n", err)
		return
	}

	if len(coord) == 0 {
		fmt.Printf("Place not found, try again with another name.")
		return
	}

	fmt.Println("Getting Weather...")

	// Get weather data based on lat/lon
	weather, err := api.GetWeather(coord[0])
	if err != nil {
		fmt.Printf("[Main]: %v\n", err)
		return
	}

	// Print results
	switch format {
	case "txt":
		fmt.Printf("\nTemperature: %v°\n", weather.Temperature)
	case "json":
		fmt.Printf("\n%v\n", utils.StructToString(weather))
	}
}

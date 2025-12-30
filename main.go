package main

import (
	"fmt"
	"os"
	"strings"
	"wheater-cli/requests"
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

	fmt.Println("Getting Wheater...")

	// Get wheater data based on lat/lon
	wheater, err := requests.GetWheater(coord[0])
	if err != nil {
		fmt.Printf("[Main]: %v\n", err)
		return
	}

	// Print results
	fmt.Printf("\nTemperature: %v°\n", wheater.Temperature)
}

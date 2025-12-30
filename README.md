# Weather CLI

A simple command-line interface tool written in Go to check the current weather conditions for a specific location.

## Description

This application uses the [Open-Meteo](https://open-meteo.com/) Geocoding and Weather APIs to resolve a location name to coordinates and fetch weather data.

## Prerequisites

- [Go](https://go.dev/dl/) 1.25 or higher

## Usage

To run the application, use the following command:

```bash
go run main.go <place>
```

### Example

```bash
go run main.go Recife
```

**Output:**
```
Getting Coordinates...
Getting weather...

Temperature: 25.3°
```

## Structure

- `main.go`: Entry point of the application. Handles CLI arguments and orchestration.
- `requests/localization.go`: Handles geocoding requests to convert city names to latitude/longitude.
- `requests/weather.go`: Handles weather requests to fetch weather data.
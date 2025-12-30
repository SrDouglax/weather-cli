package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherData struct {
	Temperature float32 `json:"temperature"`
}

func GetWeather(local Localization) (WeatherData, error) {
	apiUrl := fmt.Sprintf(`https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true&timezone=auto`, local.Latitude, local.Longitude)
	// fmt.Println("Weather URL Debug:", apiUrl)

	res, err := http.Get(apiUrl)
	if err != nil {
		return WeatherData{}, fmt.Errorf("falha na requisição: %v", err)
	}
	defer res.Body.Close()

	var response struct {
		CurrentWeather WeatherData `json:"current_weather"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return WeatherData{}, fmt.Errorf("falha ao decodificar JSON: %v", err)
	}

	return response.CurrentWeather, nil
}

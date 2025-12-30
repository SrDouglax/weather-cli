package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WheaterData struct {
	Temperature float32 `json:"temperature"`
}

func GetWheater(local Localization) (WheaterData, error) {
	apiUrl := fmt.Sprintf(`https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true&timezone=auto`, local.Latitude, local.Longitude)
	// fmt.Println("Weather URL Debug:", apiUrl)

	res, err := http.Get(apiUrl)
	if err != nil {
		return WheaterData{}, fmt.Errorf("falha na requisição: %v", err)
	}
	defer res.Body.Close()

	var response struct {
		CurrentWeather WheaterData `json:"current_weather"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return WheaterData{}, fmt.Errorf("falha ao decodificar JSON: %v", err)
	}

	return response.CurrentWeather, nil
}

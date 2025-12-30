package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Localization struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Name      string  `json:"name"`
	Country   string  `json:"country"`
}

func GetLocalCoordinates(query string) ([]Localization, error) {
	encodedQuery := url.QueryEscape(query)
	apiUrl := fmt.Sprintf(`https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=pt&format=json`, encodedQuery)
	// fmt.Println("URL Debug:", apiUrl)

	res, err := http.Get(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("falha na requisição: %v", err)
	}
	defer res.Body.Close()

	var response struct {
		Results []Localization `json:"results"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("falha ao decodificar JSON: %v", err)
	}

	if response.Results == nil {
		return []Localization{}, nil
	}

	return response.Results, nil
}

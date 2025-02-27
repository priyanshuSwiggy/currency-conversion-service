package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type APIResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func FetchRates() (map[string]float64, error) {
	apiURL := fmt.Sprintf("%s?app_id=%s", AppConfig.API.URL, AppConfig.API.Key)
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch rates")
	}
	var apiResponse APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	return apiResponse.Rates, nil
}

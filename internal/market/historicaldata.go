package market

import (
	"encoding/json"
	"net/http"
	"stockbroker/internal/auth"
	"stockbroker/internal/httpserver"
	"stockbroker/internal/models"
	"stockbroker/internal/providers"
)

func GetHistoricalData(instrument_key string, unit string, interval string, from string,
	to string) (models.CandleAPIResponse, error) {

	url := providers.FetchOHLCurl(instrument_key, unit, interval, from, to)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return models.CandleAPIResponse{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+auth.StoredAccessToken)

	resp, err := httpserver.HttpClient.Do(req)

	if err != nil {
		return models.CandleAPIResponse{}, err
	}

	defer resp.Body.Close()

	var result models.CandleAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return models.CandleAPIResponse{}, err
	}

	return result, nil
}

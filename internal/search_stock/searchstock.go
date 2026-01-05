package searchstock

import (
	"fmt"
	"io"
	"net/http"
)

func GetStockPrice(accessToken string, instrumentKey string) ([]byte, error) {
	url := fmt.Sprintf(
		"https://api.upstox.com/v2/market-quote/quotes?instrument_key=%s",
		instrumentKey,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK HTTP status: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

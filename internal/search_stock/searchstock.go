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
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	//  TEMP DEBUG
	fmt.Println("UPSTOX STATUS:", resp.StatusCode)
	fmt.Println("UPSTOX BODY:", string(body))

	return body, nil
}

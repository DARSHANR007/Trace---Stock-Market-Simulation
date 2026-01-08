package providers

import "fmt"

func FetchOHLCurl(
	instrumentKey string,
	unit string,
	interval string,
	fromDate string,
	toDate string,
) string {

	return fmt.Sprintf(
		"https://api.upstox.com/v3/historical-candle/%s/%s/%s/%s/%s",
		instrumentKey,
		unit,
		interval,
		fromDate,
		toDate,
	)
}

type CandleAPIResponse struct {
	Status string `json:"status"`
	Data   struct {
		Candles [][]interface{} `json:"candles"`
	} `json:"data"`
}

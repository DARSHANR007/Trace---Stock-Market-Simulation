package market

import (
	"net/http"
	"stockbroker/internal/auth"
	"stockbroker/internal/market"
	"stockbroker/internal/providers"
)

func GetHistoricalData(instrument_key string, unit string, interval string, from string,
	to string) (providers.CandleAPIResponse, error) {

	instrument, err := market.GetInstrumentKeyBySymbol(InstrumentDB, instrument_key)

	if err != nil {

		return nil, err
	}

	url := providers.FetchOHLCurl(instrument, unit, interval, from, to)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+auth.StoredAccessToken)

	resp, err := httpserver.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
}

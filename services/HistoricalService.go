package services

import (
	"fmt"
	"net/url"

	"github.com/tanmoyopenroot/upstox-sdk-go/endpoints"
	"github.com/tanmoyopenroot/upstox-sdk-go/interfaces"
)

// HistoricalService ... Services available for client
type HistoricalService struct {
	HTTPService interfaces.IHTTPService
}

// GetHistoricalData ...
func (service *HistoricalService) GetHistoricalData(params url.Values, headers map[string][]string, data interface{}) error {
	exchange := params.Get("exchange")
	params.Del("exchange")

	symbol := params.Get("symbol")
	params.Del("symbol")

	url := fmt.Sprintf(endpoints.URIOHLC, endpoints.URIhost, exchange, symbol, params.Encode())
	err := service.HTTPService.GET(url, headers, data)
	return err
}

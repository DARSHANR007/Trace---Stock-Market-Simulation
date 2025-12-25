package controllers

import (
	"net/http"
	"net/url"

	"github.com/tanmoyopenroot/upstox-sdk-go/interfaces"
	"github.com/tanmoyopenroot/upstox-sdk-go/models"
	"github.com/tanmoyopenroot/upstox-sdk-go/viewmodels"
)

// HistoricalController ...
type HistoricalController struct {
	HistoricalService interfaces.IHistoricalService
	Client            *models.ClientModel
}

// GetHistoricalData ...
func (controller *HistoricalController) GetHistoricalData(exchange, symbol, interval, startDate, endDate, format string) (*viewmodels.Historical, error) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	params := url.Values{}
	params.Set("exchange", exchange)
	params.Set("symbol", symbol)
	params.Set("interval", interval)
	params.Set("startDate", startDate)
	params.Set("endDate", endDate)
	params.Set("format", format)

	data := &viewmodels.Historical{}
	err := controller.HistoricalService.GetHistoricalData(params, headers, data)
	return data, err
}

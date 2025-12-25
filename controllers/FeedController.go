package controllers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/tanmoyopenroot/upstox-sdk-go/interfaces"
	"github.com/tanmoyopenroot/upstox-sdk-go/models"
	"github.com/tanmoyopenroot/upstox-sdk-go/viewmodels"
)

// FeedController ...
type FeedController struct {
	FeedService interfaces.IFeedService
	Client      *models.ClientModel
}

// GetLiveFeed ...
func (controller *FeedController) GetLiveFeed(exchange, symbol, requestType string) (*viewmodels.LiveFeed, error) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	params := url.Values{}
	params.Set("exchange", exchange)
	params.Set("symbol", symbol)
	params.Set("requestType", requestType)

	data := &viewmodels.LiveFeed{}
	err := controller.FeedService.GetLiveFeed(params, headers, data)
	return data, err
}

// SubscribeFeed ...
func (controller *FeedController) SubscribeFeed(symbol []string, exchange, requestType string) (*viewmodels.Subscribe, error) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	params := url.Values{}
	params.Set("symbol", strings.Join(symbol, ","))
	params.Set("exchange", exchange)
	params.Set("requestType", requestType)

	data := &viewmodels.Subscribe{}
	err := controller.FeedService.SubscribeFeed(params, headers, data)
	return data, err
}

// UnsubscribeFeed ...
func (controller *FeedController) UnsubscribeFeed(symbol []string, exchange, requestType string) (*viewmodels.Unsubscribe, error) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	params := url.Values{}
	params.Set("symbol", strings.Join(symbol, ","))
	params.Set("exchange", exchange)
	params.Set("requestType", requestType)

	data := &viewmodels.Unsubscribe{}
	err := controller.FeedService.UnsubscribeFeed(params, headers, data)
	return data, err
}

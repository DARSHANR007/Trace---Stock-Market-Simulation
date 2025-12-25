package services

import (
	"fmt"
	"net/url"

	"github.com/tanmoyopenroot/upstox-sdk-go/endpoints"
	"github.com/tanmoyopenroot/upstox-sdk-go/interfaces"
)

// FeedService ... Services available for client
type FeedService struct {
	HTTPService interfaces.IHTTPService
}

// GetLiveFeed ...
func (service *FeedService) GetLiveFeed(params url.Values, headers map[string][]string, data interface{}) error {
	exchange := params.Get("exchange")
	symbol := params.Get("symbol")
	requestType := params.Get("requestType")

	url := fmt.Sprintf(endpoints.URILiveFeed, endpoints.URIhost, exchange, symbol, requestType)
	err := service.HTTPService.GET(url, headers, data)
	return err
}

// SubscribeFeed ...
func (service *FeedService) SubscribeFeed(params url.Values, headers map[string][]string, data interface{}) error {
	exchange := params.Get("exchange")
	symbol := params.Get("symbol")
	requestType := params.Get("requestType")

	url := fmt.Sprintf(endpoints.URILiveFeedSubscribe, endpoints.URIhost, requestType, exchange, symbol)
	err := service.HTTPService.GET(url, headers, data)
	return err
}

// UnsubscribeFeed ...
func (service *FeedService) UnsubscribeFeed(params url.Values, headers map[string][]string, data interface{}) error {
	exchange := params.Get("exchange")
	symbol := params.Get("symbol")
	requestType := params.Get("requestType")

	url := fmt.Sprintf(endpoints.URILiveFeedSubscribe, endpoints.URIhost, requestType, exchange, symbol)
	err := service.HTTPService.GET(url, headers, data)
	return err
}

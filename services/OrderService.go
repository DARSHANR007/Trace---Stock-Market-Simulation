package services

import (
	"fmt"

	"github.com/tanmoyopenroot/upstox-sdk-go/endpoints"
	"github.com/tanmoyopenroot/upstox-sdk-go/interfaces"
)

// OrderService ... Services available for client
type OrderService struct {
	HTTPService interfaces.IHTTPService
}

// GetOrderHistory ...
func (service *OrderService) GetOrderHistory(headers map[string][]string, data interface{}) error {
	url := fmt.Sprintf(endpoints.URIGetOrders, endpoints.URIhost)
	err := service.HTTPService.GET(url, headers, data)
	return err
}

// GetOrderDetails ...
func (service *OrderService) GetOrderDetails(orderID string, headers map[string][]string, data interface{}) error {
	url := fmt.Sprintf(endpoints.URIGetOrdersInfo, endpoints.URIhost, orderID)
	err := service.HTTPService.GET(url, headers, data)
	return err
}

// GetTradeBook ...
func (service *OrderService) GetTradeBook(headers map[string][]string, data interface{}) error {
	url := fmt.Sprintf(endpoints.URITradeBook, endpoints.URIhost)
	err := service.HTTPService.GET(url, headers, data)
	return err
}

// GetTradeHistory ...
func (service *OrderService) GetTradeHistory(orderID string, headers map[string][]string, data interface{}) error {
	url := fmt.Sprintf(endpoints.URITradesInfo, endpoints.URIhost, orderID)
	err := service.HTTPService.GET(url, headers, data)
	return err
}

// PlaceOrder ...
func (service *OrderService) PlaceOrder(params interface{}, headers map[string][]string, data interface{}) error {
	url := fmt.Sprintf(endpoints.URIPlaceOrder, endpoints.URIhost)
	err := service.HTTPService.POST(url, params, headers, data)
	return err
}

// ModifyOrder ...
func (service *OrderService) ModifyOrder(orderID string, params interface{}, headers map[string][]string, data interface{}) error {
	url := fmt.Sprintf(endpoints.URIModifyOrder, endpoints.URIhost, orderID)
	err := service.HTTPService.PUT(url, params, headers, data)
	return err
}

// CancelOrder ...
func (service *OrderService) CancelOrder(orders string, headers map[string][]string, data interface{}) error {
	url := fmt.Sprintf(endpoints.URICancelOrder, endpoints.URIhost, orders)
	err := service.HTTPService.DELETE(url, headers, data)
	return err
}

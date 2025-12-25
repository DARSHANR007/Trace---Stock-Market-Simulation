package controllers

import (
	"net/http"
	"strings"

	"github.com/tanmoyopenroot/upstox-sdk-go/interfaces"
	"github.com/tanmoyopenroot/upstox-sdk-go/models"
	"github.com/tanmoyopenroot/upstox-sdk-go/viewmodels"
)

// OrderController ...
type OrderController struct {
	OrderService interfaces.IOrderService
	Client       *models.ClientModel
}

// GetOrderHistory ... Get user order history
func (controller *OrderController) GetOrderHistory() (*viewmodels.OrderHistoryModel, error) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.OrderHistoryModel{}
	err := controller.OrderService.GetOrderHistory(headers, data)
	return data, err
}

// GetOrderDetails ... Get user order details
func (controller *OrderController) GetOrderDetails(orderID string) (*viewmodels.OrderDetailsModel, error) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.OrderDetailsModel{}
	err := controller.OrderService.GetOrderDetails(orderID, headers, data)
	return data, err
}

// GetTradeBook ... Get user trade book
func (controller *OrderController) GetTradeBook() (*viewmodels.TradeBookModel, error) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.TradeBookModel{}
	err := controller.OrderService.GetTradeBook(headers, data)
	return data, err
}

// GetTradeHistory ... Get user trade history
func (controller *OrderController) GetTradeHistory(orderID string) (*viewmodels.TradeHistoryModel, error) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.TradeHistoryModel{}

	err := controller.OrderService.GetTradeHistory(orderID, headers, data)
	return data, err
}

// PlaceOrder ... place user order
func (controller *OrderController) PlaceOrder(params *viewmodels.PlaceOrderModel) (*viewmodels.PlaceOrderModel, error) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.PlaceOrderModel{}
	err := controller.OrderService.PlaceOrder(params, headers, data)
	return data, err
}

// ModifyOrder ... place user order
func (controller *OrderController) ModifyOrder(params *viewmodels.ModifyOrderModel) (*viewmodels.ModifyOrderModel, error) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.ModifyOrderModel{}
	err := controller.OrderService.ModifyOrder(data.OrderID, params, headers, data)
	return data, err
}

// CancelOrder ... cancel user order
func (controller *OrderController) CancelOrder(orders []string) (*viewmodels.CancelOrderModel, error) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.CancelOrderModel{}

	err := controller.OrderService.CancelOrder(strings.Join(orders, ","), headers, data)
	return data, err
}

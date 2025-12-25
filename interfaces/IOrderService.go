package interfaces

// IOrderService ... Interface for OrderController
type IOrderService interface {
	GetOrderHistory(headers map[string][]string, data interface{}) error
	GetOrderDetails(orderID string, headers map[string][]string, data interface{}) error
	GetTradeBook(headers map[string][]string, data interface{}) error
	GetTradeHistory(orderID string, headers map[string][]string, data interface{}) error
	PlaceOrder(params interface{}, headers map[string][]string, data interface{}) error
	ModifyOrder(orderID string, params interface{}, headers map[string][]string, data interface{}) error
	CancelOrder(orders string, headers map[string][]string, data interface{}) error
}

package viewmodels

import (
	"time"
)

// OrderHistoryModel ...
type OrderHistoryModel struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      []struct {
		Exchange          string `json:"exchange"`
		Token             int    `json:"token"`
		Symbol            string `json:"symbol"`
		Product           string `json:"product"`
		OrderType         string `json:"order_type"`
		Duration          string `json:"duration"`
		Price             int    `json:"price"`
		TriggerPrice      int    `json:"trigger_price"`
		Quantity          int    `json:"quantity"`
		DisclosedQuantity int    `json:"disclosed_quantity"`
		TransactionType   string `json:"transaction_type"`
		AveragePrice      int    `json:"average_price"`
		TradedQuantity    int    `json:"traded_quantity"`
		Message           string `json:"message"`
		ExchangeOrderID   string `json:"exchange_order_id"`
		ParentOrderID     string `json:"parent_order_id"`
		OrderID           string `json:"order_id"`
		ExchangeTime      string `json:"exchange_time"`
		TimeInMicro       string `json:"time_in_micro"`
		Status            string `json:"status"`
		IsAmo             bool   `json:"is_amo"`
		ValidDate         string `json:"valid_date"`
		OrderRequestID    string `json:"order_request_id"`
	} `json:"data,omitempty"`
}

// OrderDetailsModel ...
type OrderDetailsModel struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      []struct {
		Exchange          string  `json:"exchange"`
		Token             int     `json:"token"`
		Symbol            string  `json:"symbol"`
		Product           string  `json:"product"`
		OrderType         string  `json:"order_type"`
		Duration          string  `json:"duration"`
		Price             int     `json:"price"`
		TriggerPrice      int     `json:"trigger_price"`
		Quantity          int     `json:"quantity"`
		DisclosedQuantity int     `json:"disclosed_quantity"`
		TransactionType   string  `json:"transaction_type"`
		AveragePrice      float64 `json:"average_price"`
		TradedQuantity    int     `json:"traded_quantity"`
		Message           string  `json:"message"`
		ExchangeOrderID   string  `json:"exchange_order_id"`
		ParentOrderID     string  `json:"parent_order_id"`
		OrderID           string  `json:"order_id"`
		ExchangeTime      string  `json:"exchange_time"`
		TimeInMicro       string  `json:"time_in_micro"`
		Status            string  `json:"status"`
		IsAmo             bool    `json:"is_amo"`
		ValidDate         string  `json:"valid_date"`
		FillLeg           string  `json:"fill_leg"`
		OrderRequestID    string  `json:"order_request_id"`
		Report            string  `json:"report"`
		Text              string  `json:"text"`
	} `json:"data,omitempty"`
}

// TradeBookModel ...
type TradeBookModel struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      []struct {
		Exchange        string  `json:"exchange"`
		Token           int     `json:"token"`
		Symbol          string  `json:"symbol"`
		Product         string  `json:"product"`
		OrderType       string  `json:"order_type"`
		TransactionType string  `json:"transaction_type"`
		TradedQuantity  int     `json:"traded_quantity"`
		ExchangeOrderID string  `json:"exchange_order_id"`
		OrderID         string  `json:"order_id"`
		ExchangeTime    string  `json:"exchange_time"`
		TimeInMicro     string  `json:"time_in_micro"`
		TradedPrice     float64 `json:"traded_price"`
		TradeID         string  `json:"trade_id"`
	} `json:"data"`
}

// TradeHistoryModel ...
type TradeHistoryModel struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      []struct {
		Exchange        string `json:"exchange"`
		Token           int    `json:"token"`
		Symbol          string `json:"symbol"`
		Product         string `json:"product"`
		OrderType       string `json:"order_type"`
		TransactionType string `json:"transaction_type"`
		TradedQuantity  int    `json:"traded_quantity"`
		ExchangeOrderID string `json:"exchange_order_id"`
		OrderID         string `json:"order_id"`
		ExchangeTime    string `json:"exchange_time"`
		TimeInMicro     string `json:"time_in_micro"`
		TradedPrice     int    `json:"traded_price"`
		TradeID         string `json:"trade_id"`
	} `json:"data"`
}

// PlaceOrderModel ...
type PlaceOrderModel struct {
	TransactionType   string    `json:"transaction_type,omitempty"`
	Exchange          string    `json:"exchange,omitempty"`
	Symbol            string    `json:"symbol,omitempty"`
	Quantity          int       `json:"quantity,omitempty"`
	OrderType         string    `json:"order_type,omitempty"`
	Product           string    `json:"product,omitempty"`
	TriggerPrice      float64   `json:"trigger_price,omitempty"`
	Squareoff         float64   `json:"squareoff,omitempty"`
	IsAmo             bool      `json:"is_amo,omitempty"`
	DisclosedQuantity int       `json:"disclosed_quantity,omitempty"`
	Stoploss          float64   `json:"stoploss,omitempty"`
	Duration          string    `json:"duration,omitempty"`
	TrailingTicks     int       `json:"trailing_ticks,omitempty"`
	Price             float64   `json:"price,omitempty"`
	Code              int       `json:"code,omitempty"`
	Status            string    `json:"status,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
	Message           string    `json:"message,omitempty"`
	Data              struct {
		Exchange          string  `json:"exchange,omitempty"`
		Token             int     `json:"token,omitempty"`
		Symbol            string  `json:"symbol,omitempty"`
		Product           string  `json:"product,omitempty"`
		OrderType         string  `json:"order_type,omitempty"`
		Duration          string  `json:"duration,omitempty"`
		Price             float64 `json:"price,omitempty"`
		TriggerPrice      float64 `json:"trigger_price,omitempty"`
		Quantity          int     `json:"quantity,omitempty"`
		DisclosedQuantity int     `json:"disclosed_quantity,omitempty"`
		TransactionType   string  `json:"transaction_type,omitempty"`
		AveragePrice      float64 `json:"average_price,omitempty"`
		TradedQuantity    int     `json:"traded_quantity,omitempty"`
		Message           string  `json:"message,omitempty"`
		ExchangeOrderID   string  `json:"exchange_order_id,omitempty"`
		ParentOrderID     string  `json:"parent_order_id,omitempty"`
		OrderID           string  `json:"order_id,omitempty"`
		ExchangeTime      string  `json:"exchange_time,omitempty"`
		TimeInMicro       string  `json:"time_in_micro,omitempty"`
		Status            string  `json:"status,omitempty"`
		IsAmo             bool    `json:"is_amo,omitempty"`
		ValidDate         string  `json:"valid_date,omitempty"`
		OrderRequestID    string  `json:"order_request_id,omitempty"`
	} `json:"data,omitempty"`
}

// ModifyOrderModel ...
type ModifyOrderModel struct {
	OrderID           string    `json:"order_id,omitempty"`
	Quantity          int       `json:"quantity,omitempty"`
	OrderType         string    `json:"order_type,omitempty"`
	Price             float64   `json:"price,omitempty"`
	TriggerPrice      float64   `json:"trigger_price,omitempty"`
	DisclosedQuantity int       `json:"disclosed_quantity,omitempty"`
	Code              int       `json:"code,omitempty"`
	Status            string    `json:"status,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
	Message           string    `json:"message,omitempty"`
	Data              struct {
		Exchange          string  `json:"exchange,omitempty"`
		Token             int     `json:"token,omitempty"`
		Symbol            string  `json:"symbol,omitempty"`
		Product           string  `json:"product,omitempty"`
		OrderType         string  `json:"order_type,omitempty"`
		Duration          string  `json:"duration,omitempty"`
		Price             float64 `json:"price,omitempty"`
		TriggerPrice      float64 `json:"trigger_price,omitempty"`
		Quantity          int     `json:"quantity,omitempty"`
		DisclosedQuantity int     `json:"disclosed_quantity,omitempty"`
		TransactionType   string  `json:"transaction_type,omitempty"`
		AveragePrice      float64 `json:"average_price,omitempty"`
		TradedQuantity    int     `json:"traded_quantity,omitempty"`
		Message           string  `json:"message,omitempty"`
		ExchangeOrderID   string  `json:"exchange_order_id,omitempty"`
		ParentOrderID     string  `json:"parent_order_id,omitempty"`
		OrderID           string  `json:"order_id,omitempty"`
		ExchangeTime      string  `json:"exchange_time,omitempty"`
		TimeInMicro       string  `json:"time_in_micro,omitempty"`
		Status            string  `json:"status,omitempty"`
		IsAmo             bool    `json:"is_amo,omitempty"`
		ValidDate         string  `json:"valid_date,omitempty"`
		OrderRequestID    string  `json:"order_request_id,omitempty"`
	} `json:"data,omitempty"`
}

// CancelOrderModel ...
type CancelOrderModel struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      string    `json:"data"`
}

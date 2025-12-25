package viewmodels

import "time"

// LiveFeed ...
type LiveFeed struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      struct {
		Timestamp    int64   `json:"timestamp"`
		Exchange     string  `json:"exchange"`
		Symbol       string  `json:"symbol"`
		Ltp          float64 `json:"ltp"`
		Open         int     `json:"open"`
		High         float64 `json:"high"`
		Low          float64 `json:"low"`
		Close        float64 `json:"close"`
		Vtt          int     `json:"vtt"`
		Atp          float64 `json:"atp"`
		Oi           string  `json:"oi"`
		SpotPrice    float64 `json:"spot_price"`
		TotalBuyQty  int     `json:"total_buy_qty"`
		TotalSellQty int     `json:"total_sell_qty"`
		LowerCircuit float64 `json:"lower_circuit"`
		UpperCircuit float64 `json:"upper_circuit"`
		Bids         []struct {
			Quantity int     `json:"quantity"`
			Price    float64 `json:"price"`
			Orders   int     `json:"orders"`
		} `json:"bids"`
		Asks []struct {
			Quantity int     `json:"quantity"`
			Price    float64 `json:"price"`
			Orders   int     `json:"orders"`
		} `json:"asks"`
		Ltt int64 `json:"ltt"`
	} `json:"data"`
}

// Subscribe ...
type Subscribe struct {
	Timestamp time.Time `json:"timestamp"`
	Status    string    `json:"status"`
	Data      struct {
		Exchange string   `json:"exchange"`
		Type     string   `json:"type"`
		Success  bool     `json:"success"`
		Symbol   []string `json:"symbol"`
	} `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Unsubscribe ...
type Unsubscribe struct {
	Message string `json:"message"`
	Data    struct {
		Symbol   string `json:"symbol"`
		Exchange string `json:"exchange"`
		Type     string `json:"type"`
		Success  bool   `json:"success"`
	} `json:"data"`
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

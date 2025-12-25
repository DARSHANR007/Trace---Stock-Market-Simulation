package viewmodels

import "time"

// Profile ....
type Profile struct {
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	Code      int       `json:"code"`
	Timestamp time.Time `json:"timestamp"`
	Data      struct {
		ClientID         string   `json:"client_id"`
		BankName         string   `json:"bank_name"`
		BankAccount      string   `json:"bank_account"`
		BankName2        string   `json:"bank_name2,omitempty"`
		BankAccount2     string   `json:"bank_account2,omitempty"`
		BankName3        string   `json:"bank_name3,omitempty"`
		BankAccount3     string   `json:"bank_account3,omitempty"`
		BankName4        string   `json:"bank_name4,omitempty"`
		BankAccount4     string   `json:"bank_account4,omitempty"`
		ProductsEnabled  []string `json:"products_enabled"`
		Email            string   `json:"email"`
		Phone            string   `json:"phone"`
		Name             string   `json:"name"`
		IsActive         bool     `json:"is_active"`
		ExchangesEnabled []string `json:"exchanges_enabled"`
	} `json:"data"`
}

// Balance ...
type Balance struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Data      struct {
		Equity struct {
			NotionalCash    int     `json:"notional_cash"`
			ExposureMargin  int     `json:"exposure_margin"`
			SpanMargin      int     `json:"span_margin"`
			AvailableMargin float64 `json:"available_margin"`
			UsedMargin      int     `json:"used_margin"`
			AdhocMargin     int     `json:"adhoc_margin"`
			PayinAmount     int     `json:"payin_amount"`
		} `json:"equity"`
		Commodity struct {
			UsedMargin      int `json:"used_margin"`
			ExposureMargin  int `json:"exposure_margin"`
			SpanMargin      int `json:"span_margin"`
			AvailableMargin int `json:"available_margin"`
			NotionalCash    int `json:"notional_cash"`
			LedgerBalance   int `json:"ledger_balance"`
			AdhocMargin     int `json:"adhoc_margin"`
			PayinAmount     int `json:"payin_amount"`
		} `json:"commodity"`
	} `json:"data"`
}

// Position ...
type Position struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      []struct {
		Exchange         string  `json:"exchange"`
		Product          string  `json:"product"`
		Symbol           string  `json:"symbol"`
		Token            int     `json:"token"`
		BuyAmount        int     `json:"buy_amount"`
		SellAmount       int     `json:"sell_amount"`
		BuyQuantity      int     `json:"buy_quantity"`
		SellQuantity     int     `json:"sell_quantity"`
		CfBuyAmount      int     `json:"cf_buy_amount"`
		CfSellAmount     float64 `json:"cf_sell_amount"`
		CfBuyQuantity    int     `json:"cf_buy_quantity"`
		CfSellQuantity   int     `json:"cf_sell_quantity"`
		AvgBuyPrice      string  `json:"avg_buy_price"`
		AvgSellPrice     float64 `json:"avg_sell_price"`
		NetQuantity      int     `json:"net_quantity"`
		ClosePrice       float64 `json:"close_price"`
		LastTradedPrice  int     `json:"last_traded_price"`
		RealizedProfit   string  `json:"realized_profit"`
		UnrealizedProfit float64 `json:"unrealized_profit"`
		CfAvgPrice       string  `json:"cf_avg_price"`
	} `json:"data,omitempty"`
}

// Holdings ...
type Holdings struct {
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      []struct {
		Instrument []struct {
			Exchange string `json:"exchange"`
			Symbol   string `json:"symbol"`
			Token    int    `json:"token"`
		} `json:"instrument"`
		Product         string `json:"product"`
		CollateralType  string `json:"collateral_type"`
		CncUsedQuantity int    `json:"cnc_used_quantity"`
		Quantity        int    `json:"quantity"`
		CollateralQty   int    `json:"collateral_qty"`
		Haircut         int    `json:"haircut"`
		AvgPrice        string `json:"avg_price"`
	} `json:"data,omitempty"`
}

// MasterContract ...
type MasterContract struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Status    string      `json:"status"`
	Timestamp time.Time   `json:"timestamp"`
}

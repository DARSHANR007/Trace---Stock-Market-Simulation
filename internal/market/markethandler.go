package market

import (
	"net/http"
	searchstock "stockbroker/internal/search_stock"
)

var accessToken string

func MarketHandler(w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")
	if symbol == "" {
		http.Error(w, "Missing stock symbol", http.StatusBadRequest)
		return
	}

	instrumentKey := "NSE_EQ|" + symbol

	data, err := searchstock.GetStockPrice(accessToken, instrumentKey)
	if err != nil {
		http.Error(w, "Error fetching stock data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

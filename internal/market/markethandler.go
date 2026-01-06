package market

import (
	"fmt"
	"net/http"
	"time"

	"stockbroker/internal/auth"
	searchstock "stockbroker/internal/search_stock"
)

func MarketHandler(w http.ResponseWriter, r *http.Request) {

	accessToken := auth.StoredAccessToken

	if accessToken == "" {
		http.Error(w, "No access token available. Login or set UPSTOX_DEV_TOKEN.", http.StatusUnauthorized)
		return
	}

	symbol := r.URL.Query().Get("symbol")
	if symbol == "" {
		http.Error(w, "Missing stock symbol", http.StatusBadRequest)
		return
	}

	//instrumentKey := "NSE_EQ|" + symbol

	//RelianceinstrumentKey := "NSE_EQ|INE002A01018"

	// ZomatoInstrumentKey := "NSE_EQ|INE758T01015"
	start := time.Now()
	instrumentKey, err := GetInstrumentKeyBySymbol(InstrumentDB, symbol)
	fmt.Println("Total Time Taken :", time.Since(start))
	if err != nil {
		http.Error(w, "Instrument not found", http.StatusNotFound)
		return
	}

	data, err := searchstock.GetStockPrice(accessToken, instrumentKey)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error fetching stock data", http.StatusInternalServerError)
		return
	}

	// 6️⃣ Return JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

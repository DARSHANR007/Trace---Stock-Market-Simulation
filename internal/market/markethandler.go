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
		writeJSONError(w, http.StatusUnauthorized, "No access token available. Login or set UPSTOX_DEV_TOKEN.")
		return
	}

	symbol := r.URL.Query().Get("symbol")
	if symbol == "" {
		writeJSONError(w, http.StatusBadRequest, "Missing stock symbol")
		return
	}

	//instrumentKey := "NSE_EQ|" + symbol

	//RelianceinstrumentKey := "NSE_EQ|INE002A01018"

	// ZomatoInstrumentKey := "NSE_EQ|INE758T01015"
	start := time.Now()
	instrumentKey, err := GetInstrumentKeyBySymbol(InstrumentDB, symbol)
	fmt.Println("Total Time Taken :", time.Since(start))
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "Instrument not found")
		return
	}

	data, err := searchstock.GetStockPrice(accessToken, instrumentKey)
	if err != nil {
		fmt.Println(err)
		writeJSONError(w, http.StatusInternalServerError, "Error fetching stock data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

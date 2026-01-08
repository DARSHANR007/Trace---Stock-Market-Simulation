package market

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func OhlcHandler(w http.ResponseWriter, instrument_key string, unit string, interval string, from string, to string) {
	w.Write([]byte("OHLC Handler"))
	instrument, err := GetInstrumentKeyBySymbol(InstrumentDB, instrument_key)

	if err != nil {

		http.Error(w, "Instrument not found", http.StatusNotFound)
		return
	}

	result, err := GetHistoricalData(instrument, unit, interval, from, to)

	if err != nil {
		http.Error(w, "Error fetching historical data", http.StatusInternalServerError)
		return
	}
	// 6️⃣ Return JSON
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(data)

}

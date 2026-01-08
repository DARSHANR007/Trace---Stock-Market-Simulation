package market

import (
	"encoding/json"
	"net/http"
)

func OhlcHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters from the request
	instrument_key := r.URL.Query().Get("instrument_key")
	unit := r.URL.Query().Get("unit")
	interval := r.URL.Query().Get("interval")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	// Validate required parameters
	if instrument_key == "" || unit == "" || interval == "" || from == "" || to == "" {
		http.Error(w, "Missing required parameters: instrument_key, unit, interval, from, to", http.StatusBadRequest)
		return
	}

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

	// Return JSON
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

package market

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

/*
validateOHLC validates date ranges and limits
according to Upstox OHLC rules
*/
func validateOHLC(unit string, interval int, from, to time.Time) error {
	if from.After(to) {
		return errors.New("from_date cannot be after to_date")
	}

	switch unit {

	case "minutes":
		if from.Before(time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)) {
			return errors.New("minutes data available only from Jan 2022")
		}

		if interval <= 15 {
			// max 1 month back from to_date
			if from.Before(to.AddDate(0, -1, 0)) {
				return errors.New("minutes (<=15) supports only 1 month history")
			}
		} else {
			// max 1 quarter back
			if from.Before(to.AddDate(0, -3, 0)) {
				return errors.New("minutes (>15) supports only 1 quarter history")
			}
		}

	case "hours":
		if from.Before(time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)) {
			return errors.New("hours data available only from Jan 2022")
		}

		// max 1 quarter back
		if from.Before(to.AddDate(0, -3, 0)) {
			return errors.New("hours supports only 1 quarter history")
		}

	case "days":
		if from.Before(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)) {
			return errors.New("days data available only from Jan 2000")
		}

		// ✅ correct: 1 decade leading up to to_date
		if from.Before(to.AddDate(-10, 0, 0)) {
			return errors.New("days supports only 1 decade leading up to to_date")
		}

	case "weeks", "months":
		if from.Before(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)) {
			return errors.New("data available only from Jan 2000")
		}
		// no upper range limit

	default:
		return errors.New("invalid unit")
	}

	return nil
}

/*
OhlcHandler handles:
/ohlc/{instrument}/{unit}/{interval}/{to_date}/{from_date}
*/
func OhlcHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	// ["ohlc", instrument, unit, interval, to_date, from_date]
	if len(parts) != 6 {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	instrumentKey := parts[1]
	unit := parts[2]
	intervalStr := parts[3]
	toStr := parts[4]
	fromStr := parts[5]

	instrumentKey, err := url.PathUnescape(instrumentKey)
	if err != nil {
		http.Error(w, "Invalid instrument_key", http.StatusBadRequest)
		return
	}

	// Convert interval to int
	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		http.Error(w, "Invalid interval", http.StatusBadRequest)
		return
	}

	// Parse dates
	fromDate, err := time.Parse("2006-01-02", fromStr)
	if err != nil {
		http.Error(w, "Invalid from_date", http.StatusBadRequest)
		return
	}

	toDate, err := time.Parse("2006-01-02", toStr)
	if err != nil {
		http.Error(w, "Invalid to_date", http.StatusBadRequest)
		return
	}

	// Validate OHLC rules BEFORE calling Upstox
	if err := validateOHLC(unit, interval, fromDate, toDate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Resolve instrument key from DB
	instrument, err := GetInstrumentKeyBySymbol(InstrumentDB, instrumentKey)
	if err != nil {
		http.Error(w, "Instrument not found", http.StatusNotFound)
		return
	}

	// IMPORTANT: Upstox expects (to_date, from_date)
	result, err := GetHistoricalData(
		instrument,
		unit,
		intervalStr,
		fromStr,
		toStr,
	)
	if err != nil {
		http.Error(w, "Error fetching historical data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

package market

import (
	"database/sql"
	"net/http"
	"strings"
	"time"
)

type InstrumentSuggestion struct {
	Symbol        string `json:"symbol"`
	Name          string `json:"name"`
	InstrumentKey string `json:"instrument_key"`
	Segment       string `json:"segment"`
}

func SearchInstrumentSuggestions(db *sql.DB, query string, limit int) ([]InstrumentSuggestion, error) {
	query = strings.ToUpper(strings.TrimSpace(query))
	if query == "" {
		return []InstrumentSuggestion{}, nil
	}

	cacheKey := "search:suggestions:" + query
	var cached []InstrumentSuggestion
	if cacheGet(cacheKey, &cached) {
		return cached, nil
	}

	prefixPattern := query + "%"
	containsPattern := "%" + query + "%"

	rows, err := db.Query(`
		SELECT instrument_key, symbol, name, segment
		FROM (
			SELECT instrument_key, trading_symbol AS symbol, name, segment, 1 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(trading_symbol)) = ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')

			UNION ALL

			SELECT instrument_key, asset_symbol AS symbol, name, segment, 2 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(asset_symbol)) = ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')

			UNION ALL

			SELECT instrument_key, name AS symbol, name, segment, 3 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(name)) = ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')

			UNION ALL

			SELECT instrument_key, trading_symbol AS symbol, name, segment, 4 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(trading_symbol)) LIKE ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')

			UNION ALL

			SELECT instrument_key, asset_symbol AS symbol, name, segment, 5 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(asset_symbol)) LIKE ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')

			UNION ALL

			SELECT instrument_key, name AS symbol, name, segment, 6 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(name)) LIKE ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')

			UNION ALL

			SELECT instrument_key, trading_symbol AS symbol, name, segment, 7 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(trading_symbol)) LIKE ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')

			UNION ALL

			SELECT instrument_key, asset_symbol AS symbol, name, segment, 8 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(asset_symbol)) LIKE ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')

			UNION ALL

			SELECT instrument_key, name AS symbol, name, segment, 9 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(name)) LIKE ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')
		)
		GROUP BY instrument_key, symbol, name, segment
		ORDER BY MIN(match_rank), symbol
		LIMIT ?
	`, query, query, query, prefixPattern, prefixPattern, prefixPattern, containsPattern, containsPattern, containsPattern, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	suggestions := make([]InstrumentSuggestion, 0, limit)
	for rows.Next() {
		var suggestion InstrumentSuggestion
		if err := rows.Scan(&suggestion.InstrumentKey, &suggestion.Symbol, &suggestion.Name, &suggestion.Segment); err != nil {
			return nil, err
		}
		suggestions = append(suggestions, suggestion)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	cacheSet(cacheKey, suggestions, 5*time.Minute)

	return suggestions, nil
}

func SearchInstrumentSuggestionsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	suggestions, err := SearchInstrumentSuggestions(InstrumentDB, query, 8)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Error fetching suggestions")
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]any{
		"status": "success",
		"data":   suggestions,
	})
}

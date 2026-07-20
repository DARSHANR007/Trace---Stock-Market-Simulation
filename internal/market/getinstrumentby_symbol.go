package market

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

func GetInstrumentKeyBySymbol(db *sql.DB, symbol string) (string, error) {
	symbol = strings.ToUpper(strings.TrimSpace(symbol))
	cacheKey := "search:instrument:" + symbol
	var cached string
	if cacheGet(cacheKey, &cached) {
		return cached, nil
	}
	pattern := "%" + symbol + "%"
	prefixPattern := symbol + "%"

	query := `
		SELECT instrument_key
		FROM (
			SELECT instrument_key, 1 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(trading_symbol)) = ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')

			UNION ALL

			SELECT instrument_key, 2 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(asset_symbol)) = ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')

			UNION ALL

			SELECT instrument_key, 3 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(name)) = ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')

			UNION ALL

			SELECT instrument_key, 4 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(trading_symbol)) LIKE ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')

			UNION ALL

			SELECT instrument_key, 5 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(asset_symbol)) LIKE ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')

			UNION ALL

			SELECT instrument_key, 6 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(name)) LIKE ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')

			UNION ALL

			SELECT instrument_key, 7 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(trading_symbol)) LIKE ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')

			UNION ALL

			SELECT instrument_key, 8 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(asset_symbol)) LIKE ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')

			UNION ALL

			SELECT instrument_key, 9 AS match_rank
			FROM instruments
			WHERE UPPER(TRIM(name)) LIKE ?
			  AND segment IN ('NSE_EQ', 'BSE_EQ')
		)
		ORDER BY match_rank
		LIMIT 1;
	`
	var instrumentKey string
	start := time.Now()
	err := db.QueryRow(query, symbol, symbol, symbol, prefixPattern, prefixPattern, prefixPattern, pattern, pattern, pattern).Scan(&instrumentKey)
	fmt.Println("DB Query Time:", time.Since(start))
	if err == sql.ErrNoRows {
		return "", errors.New("instrument not found")
	}
	if err != nil {
		return "", err
	}

	fmt.Println(instrumentKey)
	cacheSet(cacheKey, instrumentKey, 10*time.Minute)

	return instrumentKey, nil
}

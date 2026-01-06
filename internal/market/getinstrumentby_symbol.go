package market

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

func GetInstrumentKeyBySymbol(db *sql.DB, symbol string) (string, error) {
	query := `
		SELECT instrument_key
		FROM instruments
		WHERE trading_symbol = ?
		  AND segment IN ('NSE_EQ', 'BSE_EQ')
		ORDER BY
		  CASE
			WHEN segment = 'NSE_EQ' THEN 1
			WHEN segment = 'BSE_EQ' THEN 2
		  END
		LIMIT 1;
	`
	var instrumentKey string
	start := time.Now()
	err := db.QueryRow(query, symbol).Scan(&instrumentKey)
	fmt.Println("DB Query Time:", time.Since(start))
	if err == sql.ErrNoRows {
		return "", errors.New("instrument not found")
	}
	if err != nil {
		return "", err
	}

	fmt.Println(instrumentKey)

	return instrumentKey, nil
}

package market

import (
	"database/sql"
	"errors"
	"fmt"
)

func GetInstrumentKeyBySymbol(db *sql.DB, symbol string) (string, error) {
	query := `
		SELECT instrument_key
		FROM instruments
		WHERE (
			UPPER(name) = UPPER(?)
			OR trading_symbol = ? || '-EQ'
		)
		AND segment IN ('NSE_EQ', 'BSE_EQ')
		ORDER BY
			CASE
				WHEN segment = 'NSE_EQ' THEN 1
				WHEN segment = 'BSE_EQ' THEN 2
			END
		LIMIT 1;
	`
	var instrumentKey string
	err := db.QueryRow(query, symbol).Scan(&instrumentKey)
	if err == sql.ErrNoRows {
		return "", errors.New("instrument not found")
	}
	if err != nil {
		return "", err
	}

	fmt.Println(instrumentKey)

	return instrumentKey, nil
}

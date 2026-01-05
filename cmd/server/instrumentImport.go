package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

type Instrument struct {
	Weekly         bool   `json:"weekly"`
	Segment        string `json:"segment"`
	Name           string `json:"name"`
	Exchange       string `json:"exchange"`
	Expiry         int64  `json:"expiry"`
	InstrumentType string `json:"instrument_type"`

	AssetSymbol      string `json:"asset_symbol"`
	UnderlyingSymbol string `json:"underlying_symbol"`
	InstrumentKey    string `json:"instrument_key"`

	LotSize        int     `json:"lot_size"`
	FreezeQuantity float64 `json:"freeze_quantity"`
	ExchangeToken  string  `json:"exchange_token"`
	MinimumLot     int     `json:"minimum_lot"`

	TickSize       float64 `json:"tick_size"`
	AssetType      string  `json:"asset_type"`
	UnderlyingType string  `json:"underlying_type"`

	TradingSymbol string  `json:"trading_symbol"`
	StrikePrice   float64 `json:"strike_price"`
	QtyMultiplier float64 `json:"qty_multiplier"`
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func importer() {
	// 1. Open DB
	db, err := sql.Open("sqlite", "instruments.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. Performance pragmas (VERY important)
	db.Exec(`PRAGMA journal_mode=WAL`)
	db.Exec(`PRAGMA synchronous=NORMAL`)
	db.Exec(`PRAGMA temp_store=MEMORY`)

	// 3. Prepare insert
	stmt, err := db.Prepare(`
	INSERT OR REPLACE INTO instruments (
		instrument_key, exchange, segment, name, trading_symbol,
		underlying_symbol, asset_symbol,
		instrument_type, asset_type, underlying_type,
		expiry, strike_price,
		lot_size, minimum_lot, qty_multiplier,
		tick_size, freeze_quantity,
		exchange_token, weekly
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// 4. Open JSON file
	file, err := os.Open("instrument.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// 5. Skip '['
	_, err = decoder.Token()
	if err != nil {
		log.Fatal(err)
	}

	// 6. Transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	count := 0

	for decoder.More() {
		var inst Instrument
		if err := decoder.Decode(&inst); err != nil {
			log.Fatal(err)
		}

		_, err := stmt.Exec(
			inst.InstrumentKey,
			inst.Exchange,
			inst.Segment,
			inst.Name,
			inst.TradingSymbol,

			inst.UnderlyingSymbol,
			inst.AssetSymbol,

			inst.InstrumentType,
			inst.AssetType,
			inst.UnderlyingType,

			inst.Expiry,
			inst.StrikePrice,

			inst.LotSize,
			inst.MinimumLot,
			inst.QtyMultiplier,

			inst.TickSize,
			inst.FreezeQuantity,

			inst.ExchangeToken,
			boolToInt(inst.Weekly),
		)
		if err != nil {
			log.Fatal(err)
		}

		count++
		if count%5000 == 0 {
			fmt.Println("Imported:", count)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Import completed. Total rows:", count)
}

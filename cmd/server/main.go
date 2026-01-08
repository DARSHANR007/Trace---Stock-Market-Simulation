package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"stockbroker/internal/auth"
	"stockbroker/internal/config"
	"stockbroker/internal/market"

	_ "modernc.org/sqlite"
)

func main() {
	// Load env config
	cfg := config.Load()

	db, err := sql.Open("sqlite", "instruments.db")
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	market.SetInstrumentDB(db)

	// Create OAuth client (VERY IMPORTANT)
	oauth := &auth.UpstoxOAuth{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURI:  cfg.RedirectURI,
	}

	//  Create mux
	mux := http.NewServeMux()

	//  Register routes
	mux.HandleFunc("/login", auth.LoginHandler(cfg.ClientID, cfg.RedirectURI))
	mux.HandleFunc("/callback", auth.CallbackHandler(oauth))
	mux.HandleFunc("/getprice", market.MarketHandler)
	fs := http.FileServer(http.Dir("./ui"))
	mux.Handle("/seeprice/", http.StripPrefix("/seeprice/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server running"))
	})

	mux.HandleFunc("/ohlc/", market.OhlcHandler)

	fmt.Println("Server running at http://localhost:8000/login")
	log.Fatal(http.ListenAndServe(":8000", mux))

}

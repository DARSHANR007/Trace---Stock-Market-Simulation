package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"stockbroker/internal/auth"
	"stockbroker/internal/config"
	"stockbroker/internal/market"

	"github.com/redis/go-redis/v9"
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

	if cfg.RedisURL != "" {
		market.SetRedisClient(redis.NewClient(&redis.Options{Addr: cfg.RedisURL}))
	}

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
	mux.HandleFunc("/suggestions", market.SearchInstrumentSuggestionsHandler)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{"status": "ok"})
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"status":  "ok",
			"service": "stockbroker-api",
			"routes":  []string{"/health", "/login", "/callback", "/getprice", "/suggestions", "/ohlc/{instrument}/{unit}/{interval}/{to}/{from}"},
		})
	})

	mux.HandleFunc("/ohlc/", market.OhlcHandler)

	fmt.Println("Server running at http://localhost:8000/login")
	log.Fatal(http.ListenAndServe(":8000", mux))

}

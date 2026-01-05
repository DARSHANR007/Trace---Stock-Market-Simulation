package main

import (
	"fmt"
	"net/http"

	"stockbroker/internal/auth"
	"stockbroker/internal/config"
	"stockbroker/internal/market"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()
	mux.HandleFunc("/login", auth.LoginHandler(cfg.ClientID, cfg.RedirectURI))
	mux.HandleFunc("/stock_price", market.MarketHandler)

	fmt.Println("Server running at http://localhost:8000")
	http.ListenAndServe(":8000", mux)
}

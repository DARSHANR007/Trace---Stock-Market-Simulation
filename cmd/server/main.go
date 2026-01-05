package main

import (
	"fmt"
	"net/http"

	"stockbroker/internal/auth"
	"stockbroker/internal/config"
	"stockbroker/internal/market"
)

func main() {
	// Load env config
	cfg := config.Load()

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
	mux.HandleFunc("/seeprice", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/index.html")
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server running"))
	})

	fmt.Println("Server running at http://localhost:8000/login")
	http.ListenAndServe(":8000", mux)
}

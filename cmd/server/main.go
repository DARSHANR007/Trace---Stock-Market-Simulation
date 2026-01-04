package main

import (
	"fmt"
	"net/http"

	"stockbroker/internal/auth"
	"stockbroker/internal/config"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()
	mux.HandleFunc("/login", auth.LoginHandler(cfg.ClientID, cfg.RedirectURI))

	fmt.Println("Server running at http://localhost:8000")
	http.ListenAndServe(":8000", mux)
}

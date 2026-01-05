package main

import (
	"fmt"
	"net/http"

	"stockbroker/internal/auth"
	"stockbroker/internal/config"
)

func main() {
	// 1️⃣ Load env config
	cfg := config.Load()

	// 2️⃣ Create OAuth client (VERY IMPORTANT)
	oauth := &auth.UpstoxOAuth{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURI:  cfg.RedirectURI,
	}

	// 3️⃣ Create mux
	mux := http.NewServeMux()

	// 4️⃣ Register routes
	mux.HandleFunc("/login", auth.LoginHandler(cfg.ClientID, cfg.RedirectURI))
	mux.HandleFunc("/callback", auth.CallbackHandler(oauth))

	// (Optional sanity route)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server running"))
	})

	// 5️⃣ Start server
	fmt.Println("Server running at http://localhost:8000")
	http.ListenAndServe(":8000", mux)
}

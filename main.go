package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	clientID    = os.Getenv("APKEY")
	redirectURI = os.Getenv("REDIRECT_URI")
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/", callbackHandler)

	fmt.Println("Server running at http://localhost:8000")
	fmt.Println("Open http://localhost:8000/login")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	authURL := "https://api.upstox.com/v2/login/authorization/dialog"

	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", clientID)
	params.Add("redirect_uri", redirectURI)
	params.Add("state", "dev_state_123")

	finalURL := authURL + "?" + params.Encode()
	http.Redirect(w, r, finalURL, http.StatusFound)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	if code == "" {
		fmt.Fprintln(w, "Go to /login to authenticate")
		return
	}

	fmt.Println("Authorization code:", code)
	fmt.Fprintln(w, "Login successful. Check terminal.")
}

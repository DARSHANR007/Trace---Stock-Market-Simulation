package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

var (
	clientID     = os.Getenv("UPSTOX_API_KEY")
	clientSecret = os.Getenv("UPSTOX_API_SECRET")
	redirectURI  = os.Getenv("UPSTOX_REDIRECT_URI")
)

func main() {

	godotenv.Load()

	if clientID == "" || clientSecret == "" || redirectURI == "" {
		log.Fatal("Missing environment variables")
	}

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/", callbackHandler)

	fmt.Println("Server running at http://localhost:8000")
	fmt.Println("Open http://localhost:8000/login to start authentication")

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	authURL := "https://api.upstox.com/v2/login/authorization/dialog"

	params := url.Values{}
	params.Set("response_type", "code")
	params.Set("client_id", clientID)
	params.Set("redirect_uri", redirectURI)
	params.Set("state", "dev_state_123")

	finalURL := authURL + "?" + params.Encode()
	http.Redirect(w, r, finalURL, http.StatusFound)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	if code == "" {
		fmt.Fprintln(w, "Go to /login to authenticate")
		return
	}

	fmt.Println("Authorization code received:", code)

	err := exchangeToken(code)
	if err != nil {
		fmt.Println("Token exchange failed:", err)
		http.Error(w, "Token exchange failed", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Authentication successful. Check terminal for access token.")
}

func exchangeToken(code string) error {
	tokenURL := "https://api.upstox.com/v2/login/authorization/token"

	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", "authorization_code")

	req, err := http.NewRequest("POST", tokenURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("Token API status:", resp.Status)
	fmt.Println("Token API response:")
	fmt.Println(string(body))

	return nil
}

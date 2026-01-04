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
	clientID     string
	clientSecret string
	redirectURI  string
)

func main() {
	// Load .env (for local dev)
	godotenv.Load()

	clientID = os.Getenv("UPSTOX_API_KEY")
	clientSecret = os.Getenv("UPSTOX_API_SECRET")
	redirectURI = os.Getenv("UPSTOX_REDIRECT_URI")

	if clientID == "" || clientSecret == "" || redirectURI == "" {
		log.Fatal("Missing environment variables")
	}

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/", callbackHandler)

	fmt.Println("Server running at http://localhost:8000")
	fmt.Println("Open http://localhost:8000/login to authenticate")

	log.Fatal(http.ListenAndServe(":8000", nil))
}

// STEP 1: Redirect user to Upstox login page
func loginHandler(w http.ResponseWriter, r *http.Request) {
	baseURL := "https://api.upstox.com/v2/login/authorization/dialog"

	params := url.Values{}
	params.Set("response_type", "code")
	params.Set("client_id", clientID)
	params.Set("redirect_uri", redirectURI)
	params.Set("state", "upstox_oauth_state")

	loginURL := baseURL + "?" + params.Encode()
	http.Redirect(w, r, loginURL, http.StatusFound)
}

// STEP 2: Receive authorization code
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	if code == "" {
		fmt.Fprintln(w, "Go to /login to authenticate with Upstox")
		return
	}

	fmt.Println("Authorization code received:", code)

	if err := exchangeToken(code); err != nil {
		http.Error(w, "Token exchange failed", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Authentication successful. Check terminal for access token.")
}

// STEP 3: Exchange code for access token
func exchangeToken(code string) error {
	tokenURL := "https://api.upstox.com/v2/login/authorization/token"

	form := url.Values{}
	form.Set("code", code)
	form.Set("client_id", clientID)
	form.Set("client_secret", clientSecret)
	form.Set("redirect_uri", redirectURI)
	form.Set("grant_type", "authorization_code")

	req, err := http.NewRequest(
		"POST",
		tokenURL,
		bytes.NewBufferString(form.Encode()),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

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

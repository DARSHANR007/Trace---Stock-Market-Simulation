package auth

import (
	"net/http"
	"net/url"
)

// first it calls login handler to redirect user to upstox login page
// then an exchange code is received in callback handler to get access token
// access token is stored in a package level variable for later use
func LoginHandler(clientID, redirectURI string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		baseURL := "https://api.upstox.com/v2/login/authorization/dialog"

		params := url.Values{}
		params.Set("response_type", "code")
		params.Set("client_id", clientID)
		params.Set("redirect_uri", redirectURI)

		http.Redirect(w, r, baseURL+"?"+params.Encode(), http.StatusFound)
	}
}

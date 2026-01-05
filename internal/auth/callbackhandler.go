package auth

import (
	"fmt"
	"net/http"
)

var StoredAccessToken string

func CallbackHandler(oauth *UpstoxOAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "Missing auth code", http.StatusBadRequest)
			return
		}

		token, err := oauth.ExchangeToken(code)

		if err != nil {
			http.Error(w, "Token exchange failed", http.StatusInternalServerError)
			return
		}

		StoredAccessToken = token.Value

		fmt.Println("ACCESS TOKEN:", token.Value)

		w.Write([]byte("Login successful. You can close this page."))
	}
}

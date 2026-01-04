package auth

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

type UpstoxOAuth struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

func (u *UpstoxOAuth) ExchangeToken(code string) ([]byte, error) {
	form := url.Values{}
	form.Set("code", code)
	form.Set("client_id", u.ClientID)
	form.Set("client_secret", u.ClientSecret)
	form.Set("redirect_uri", u.RedirectURI)
	form.Set("grant_type", "authorization_code")

	req, err := http.NewRequest(
		"POST",
		"https://api.upstox.com/v2/login/authorization/token",
		bytes.NewBufferString(form.Encode()),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

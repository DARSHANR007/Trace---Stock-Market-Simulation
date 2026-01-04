package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func (u *UpstoxOAuth) ExchangeToken(code string) (*AccessToken, error) {

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
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return &AccessToken{
		Value:     res.AccessToken,
		ExpiresIn: res.ExpiresIn,
	}, nil
}

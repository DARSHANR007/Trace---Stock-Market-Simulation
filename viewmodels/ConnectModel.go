package viewmodels

// SessionBodyModel ...
type SessionBodyModel struct {
	Code        string `json:"code"`
	GrantType   string `json:"grant_type"`
	RedirectURI string `json:"redirect_uri"`
}

// HeaderModel ...
type HeaderModel struct {
	Authorization string `json:"Authorization"`
	APIKey        string `json:"x-api-key"`
	ContentType   string `json:"Content-Type"`
}

// AccessTokenModel ...
type AccessTokenModel struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

package models

// ClientModel ...
type ClientModel struct {
	apiKey      string
	accessToken string
}

// SetAPIKey ... Set API Key
func (model *ClientModel) SetAPIKey(key string) {
	model.apiKey = key
}

// GetAPIKey ... Get API Key
func (model *ClientModel) GetAPIKey() string {
	return model.apiKey
}

// SetAccessToken ... Set AccessToken
func (model *ClientModel) SetAccessToken(accessToken string) {
	model.accessToken = accessToken
}

// GetAccessToken ... Get AccessToken
func (model *ClientModel) GetAccessToken() string {
	return model.accessToken
}

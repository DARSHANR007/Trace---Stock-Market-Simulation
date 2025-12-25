package controllers

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/tanmoyopenroot/upstox-sdk-go/endpoints"
	"github.com/tanmoyopenroot/upstox-sdk-go/interfaces"
	"github.com/tanmoyopenroot/upstox-sdk-go/models"
	"github.com/tanmoyopenroot/upstox-sdk-go/viewmodels"
)

// ConnectController ...
type ConnectController struct {
	ConnectService interfaces.IConnectService
	Client         *models.ClientModel
}

// New ... Create New Client
func (controller *ConnectController) New(key string) {
	controller.Client.SetAPIKey(key)
}

// GenerateSession ... Generates access token
func (controller *ConnectController) GenerateSession(apiSecret string, code string, redirectURI string) (string, error) {
	temp := controller.Client.GetAPIKey() + ":" + apiSecret
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(temp))

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())
	headers.Set("Content-Type", "application/json")

	params := &viewmodels.SessionBodyModel{
		Code:        code,
		GrantType:   "authorization_code",
		RedirectURI: redirectURI,
	}

	data := &viewmodels.AccessTokenModel{}

	err := controller.ConnectService.GetAccessToken(params, headers, data)
	if err == nil && data.AccessToken != "" {
		return data.AccessToken, nil
	}
	return "", err
}

// SetAccessToken ... Set the generated access token
func (controller *ConnectController) SetAccessToken(accessToken string) {
	controller.Client.SetAccessToken(accessToken)
}

// GetLoginURL ... Get the URL using API key
func (controller *ConnectController) GetLoginURL(redirectURI string) string {
	return fmt.Sprintf(endpoints.URIAuthorise, endpoints.URIhost, controller.Client.GetAPIKey(), redirectURI)
}

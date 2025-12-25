package controllers

import (
	"net/http"
	"net/url"

	"github.com/tanmoyopenroot/upstox-sdk-go/interfaces"
	"github.com/tanmoyopenroot/upstox-sdk-go/models"
	"github.com/tanmoyopenroot/upstox-sdk-go/viewmodels"
)

// UserController ...
type UserController struct {
	UserService interfaces.IUserService
	Client      *models.ClientModel
}

// GetProfile ... Get user profile
func (controller *UserController) GetProfile() (*viewmodels.Profile, error) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.Profile{}
	err := controller.UserService.GetProfile(headers, data)
	return data, err
}

// GetBalance ... Get user balance
func (controller *UserController) GetBalance(kind string) (*viewmodels.Balance, error) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.Balance{}
	err := controller.UserService.GetBalance(kind, headers, data)
	return data, err
}

// GetPositions ... Get user positions
func (controller *UserController) GetPositions() (*viewmodels.Position, error) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.Position{}
	err := controller.UserService.GetPositions(headers, data)
	return data, err
}

// GetHoldings ... Get user holdings
func (controller *UserController) GetHoldings() (*viewmodels.Holdings, error) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	data := &viewmodels.Holdings{}
	err := controller.UserService.GetHoldings(headers, data)
	return data, err
}

// GetMasterContracts ... Get user Master Contracts
func (controller *UserController) GetMasterContracts(exchange string, symbol, token interface{}) (*viewmodels.MasterContract, error) {
	basicAuth := "Bearer " + controller.Client.GetAccessToken()

	headers := make(http.Header)
	headers.Set("Authorization", basicAuth)
	headers.Set("x-api-key", controller.Client.GetAPIKey())

	params := url.Values{}
	if symbol != nil {
		params.Set("symbol", symbol.(string))
	}
	if token != nil {
		params.Set("token", token.(string))
	}

	data := &viewmodels.MasterContract{}
	err := controller.UserService.GetMasterContracts(exchange, params, headers, data)
	return data, err
}

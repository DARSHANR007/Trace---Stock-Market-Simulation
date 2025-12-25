package services

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/tanmoyopenroot/upstox-sdk-go/endpoints"
	"github.com/tanmoyopenroot/upstox-sdk-go/interfaces"
)

// UserService ... Services available for client
type UserService struct {
	HTTPService interfaces.IHTTPService
}

// GetProfile ...
func (service *UserService) GetProfile(headers map[string][]string, data interface{}) error {
	url := fmt.Sprintf(endpoints.URIProfile, endpoints.URIhost)
	err := service.HTTPService.GET(url, headers, data)
	return err
}

// GetBalance ...
func (service *UserService) GetBalance(kind string, headers map[string][]string, data interface{}) error {
	var url string

	if kind == "commodity" {
		url = fmt.Sprintf(endpoints.URIBalanceCommodity, endpoints.URIhost)
	} else if kind == "security" {
		url = fmt.Sprintf(endpoints.URIBalance, endpoints.URIhost)
	} else if kind == "" {
		url = fmt.Sprintf(endpoints.URIBalance, endpoints.URIhost)
	} else {
		return errors.New("Type should be either commodity or security")
	}

	err := service.HTTPService.GET(url, headers, data)
	return err
}

// GetPositions ...
func (service *UserService) GetPositions(headers map[string][]string, data interface{}) error {
	url := fmt.Sprintf(endpoints.URIPositions, endpoints.URIhost)
	err := service.HTTPService.GET(url, headers, data)
	return err
}

// GetHoldings ...
func (service *UserService) GetHoldings(headers map[string][]string, data interface{}) error {
	url := fmt.Sprintf(endpoints.URIHoldings, endpoints.URIhost)
	err := service.HTTPService.GET(url, headers, data)
	return err
}

// GetMasterContracts ...
func (service *UserService) GetMasterContracts(exchange string, params url.Values, headers map[string][]string, data interface{}) error {
	url := fmt.Sprintf(endpoints.URIMasterContract, endpoints.URIhost, exchange+"?"+params.Encode())
	err := service.HTTPService.GET(url, headers, data)
	return err
}

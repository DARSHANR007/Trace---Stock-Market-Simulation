package interfaces

import (
	"net/url"
)

// IUserService ... Interface for UserController
type IUserService interface {
	GetProfile(headers map[string][]string, data interface{}) error
	GetBalance(kind string, headers map[string][]string, data interface{}) error
	GetPositions(headers map[string][]string, data interface{}) error
	GetHoldings(headers map[string][]string, data interface{}) error
	GetMasterContracts(exchange string, params url.Values, headers map[string][]string, data interface{}) error
}

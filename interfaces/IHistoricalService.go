package interfaces

import (
	"net/url"
)

// IHistoricalService ... Interface for FeedService
type IHistoricalService interface {
	GetHistoricalData(params url.Values, headers map[string][]string, data interface{}) error
}

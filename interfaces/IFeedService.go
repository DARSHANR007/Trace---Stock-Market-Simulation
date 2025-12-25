package interfaces

import (
	"net/url"
)

// IFeedService ... Interface for FeedService
type IFeedService interface {
	GetLiveFeed(params url.Values, headers map[string][]string, data interface{}) error
	SubscribeFeed(params url.Values, headers map[string][]string, data interface{}) error
	UnsubscribeFeed(params url.Values, headers map[string][]string, data interface{}) error
}

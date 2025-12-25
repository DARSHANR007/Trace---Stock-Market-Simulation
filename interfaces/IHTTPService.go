package interfaces

// IHTTPService ... Interface fot HTTP
type IHTTPService interface {
	POST(url string, params interface{}, headers map[string][]string, data interface{}) error
	GET(url string, headers map[string][]string, data interface{}) error
	PUT(url string, params interface{}, headers map[string][]string, data interface{}) error
	DELETE(url string, headers map[string][]string, data interface{}) error
}

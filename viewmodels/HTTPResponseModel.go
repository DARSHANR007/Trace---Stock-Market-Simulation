package viewmodels

// HTTPResponseModel ...
type HTTPResponseModel struct {
	Code      int         `json:"code"`
	Status    string      `json:"status"`
	TimeStamp string      `json:"timestamp"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

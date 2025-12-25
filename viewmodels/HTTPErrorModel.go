package viewmodels

// HTTPErrorModel ...
type HTTPErrorModel struct {
	Code      int         `json:"code"`
	Status    string      `json:"status"`
	TimeStamp string      `json:"timestamp"`
	Message   string      `json:"message"`
	Error     interface{} `json:"error"`
}

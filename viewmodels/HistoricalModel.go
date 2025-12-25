package viewmodels

// Historical ...
type Historical struct {
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	Data      []struct {
		Timestamp int64   `json:"timestamp"`
		Open      float64 `json:"open"`
		High      float64 `json:"high"`
		Low       float64 `json:"low"`
		Close     float64 `json:"close"`
		Volume    int     `json:"volume"`
		Cp        float64 `json:"cp"`
	} `json:"data"`
}

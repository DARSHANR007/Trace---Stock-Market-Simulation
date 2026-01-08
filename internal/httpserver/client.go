package httpserver

import (
	"net/http"
	"time"
)

var HttpClient = &http.Client{
	Timeout: 5 * time.Second,
}

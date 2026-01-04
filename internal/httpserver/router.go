package httpserver

import (
	"net/http"
)

func New() *http.ServeMux {
	return http.NewServeMux()
}

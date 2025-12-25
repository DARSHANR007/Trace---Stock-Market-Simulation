package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tanmoyopenroot/upstox-sdk-go/viewmodels"
)

// HTTPService ...
type HTTPService struct {
	HTTPClient *http.Client
}

// HandleRequestResponse ...
func (service *HTTPService) HandleRequestResponse(req *http.Request, data interface{}) error {
	resp, err := service.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Unable to read response: %v", err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		var e viewmodels.HTTPErrorModel
		if err := json.Unmarshal(body, &e); err != nil {
			fmt.Printf("Error parsing JSON response: %v", err)
			return err
		}
		return errors.New(e.Message)
	}

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Printf("Error parsing JSON response: %v | %s", err, body)
		return err
	}

	return nil
}

// POST ... Post function
func (service *HTTPService) POST(url string, params interface{}, headers map[string][]string, data interface{}) error {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return err
	}

	var jsonStr = []byte(string(jsonData))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	if headers == nil {
		headers = map[string][]string{}
	}

	req.Header = headers
	return service.HandleRequestResponse(req, data)
}

// PUT ... Put function
func (service *HTTPService) PUT(url string, params interface{}, headers map[string][]string, data interface{}) error {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return err
	}

	var jsonStr = []byte(string(jsonData))
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	if headers == nil {
		headers = map[string][]string{}
	}

	req.Header = headers
	req.Header.Set("Content-Type", "application/json")
	return service.HandleRequestResponse(req, data)
}

// DELETE ... Delete function
func (service *HTTPService) DELETE(url string, headers map[string][]string, data interface{}) (err error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	if headers == nil {
		headers = map[string][]string{}
	}

	req.Header = headers
	return service.HandleRequestResponse(req, data)
}

// GET ... Get function
func (service *HTTPService) GET(url string, headers map[string][]string, data interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	if headers == nil {
		headers = map[string][]string{}
	}

	req.Header = headers
	return service.HandleRequestResponse(req, data)
}

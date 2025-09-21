package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	BaseURL = "http://localhost:3001" // Will be configurable later
	Client  = &http.Client{
		Timeout: 30 * time.Second,
	}
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// MakeRequest makes an HTTP request with optional authentication
func MakeRequest(method, endpoint string, data interface{}, token string) ([]byte, error) {
	var body io.Reader = nil
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request data: %v", err)
		}
		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, BaseURL+endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode >= 400 {
		var errorResp ErrorResponse
		if err := json.Unmarshal(responseBody, &errorResp); err == nil && errorResp.Error != "" {
			return nil, fmt.Errorf("server error: %s (status: %d)", errorResp.Error, resp.StatusCode)
		}
		return nil, fmt.Errorf("request failed with status: %d, response: %s", resp.StatusCode, string(responseBody))
	}

	return responseBody, nil
}

// TestServerConnection tests if the server is reachable
func TestServerConnection() error {
	_, err := MakeRequest("GET", "/api/health", nil, "")
	return err
}

// GetServerHealth returns server health information
func GetServerHealth() (map[string]interface{}, error) {
	response, err := MakeRequest("GET", "/api/health", nil, "")
	if err != nil {
		return nil, err
	}

	var health map[string]interface{}
	if err := json.Unmarshal(response, &health); err != nil {
		return nil, fmt.Errorf("failed to parse health response: %v", err)
	}

	return health, nil
}
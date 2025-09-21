package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/zalando/go-keyring"
)

const (
	ServiceName = "certifycli"
	TokenKey    = "auth_token"
	ServerURL   = "http://localhost:3001"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type UserInfo struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

// Login authenticates with the server and returns a JWT token
func Login(email, password string) (string, error) {
	loginReq := LoginRequest{
		Email:    email,
		Password: password,
	}

	jsonData, err := json.Marshal(loginReq)
	if err != nil {
		return "", fmt.Errorf("failed to marshal login request: %w", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post(ServerURL+"/api/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to connect to server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login failed with status: %d", resp.StatusCode)
	}

	var loginResp LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return loginResp.Token, nil
}

// StoreToken securely stores the JWT token in the OS keychain
func StoreToken(token string) error {
	return keyring.Set(ServiceName, TokenKey, token)
}

// GetToken retrieves the stored JWT token from the OS keychain
func GetToken() (string, error) {
	return keyring.Get(ServiceName, TokenKey)
}

// IsLoggedIn checks if a valid token exists
func IsLoggedIn() bool {
	token, err := GetToken()
	if err != nil {
		return false
	}
	
	// TODO: Validate token expiry
	return token != ""
}

// GetUserInfo retrieves user information from the server
func GetUserInfo() (*UserInfo, error) {
	token, err := GetToken()
	if err != nil {
		return nil, fmt.Errorf("no auth token found: %w", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", ServerURL+"/api/user", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info, status: %d", resp.StatusCode)
	}

	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %w", err)
	}

	return &userInfo, nil
}

// CheckServerConnectivity checks if the server is reachable
func CheckServerConnectivity() bool {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(ServerURL + "/api/health")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	
	return resp.StatusCode == http.StatusOK
}

// Logout removes the stored token
func Logout() error {
	return keyring.Delete(ServiceName, TokenKey)
}
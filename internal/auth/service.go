package auth

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/CreatorOss/certifycli/internal/utils"
	"golang.org/x/term"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	User    struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"user"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email,omitempty"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	User    struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"user"`
}

// Login authenticates with the server and stores the token
func Login() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", fmt.Errorf("failed to read password: %v", err)
	}
	password := string(passwordBytes)
	fmt.Println() // Add newline after password input

	if username == "" || password == "" {
		return "", fmt.Errorf("username and password are required")
	}

	loginData := LoginRequest{
		Username: username,
		Password: password,
	}

	response, err := utils.MakeRequest("POST", "/api/login", loginData, "")
	if err != nil {
		return "", fmt.Errorf("login failed: %v", err)
	}

	var loginResp LoginResponse
	if err := json.Unmarshal(response, &loginResp); err != nil {
		return "", fmt.Errorf("failed to parse login response: %v", err)
	}

	// Save token to keychain
	keyManager := NewKeyringManager()
	if err := keyManager.SaveToken(username, loginResp.Token); err != nil {
		return "", fmt.Errorf("failed to save token: %v", err)
	}

	// Update user config with logged-in username
	if err := saveCurrentUser(username); err != nil {
		return "", fmt.Errorf("failed to save user config: %v", err)
	}

	return loginResp.Token, nil
}

// Register creates a new user account
func Register() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Email (optional): ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Password: ")
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("failed to read password: %v", err)
	}
	password := string(passwordBytes)
	fmt.Println()

	fmt.Print("Confirm Password: ")
	confirmPasswordBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("failed to read password confirmation: %v", err)
	}
	confirmPassword := string(confirmPasswordBytes)
	fmt.Println()

	if username == "" || password == "" {
		return fmt.Errorf("username and password are required")
	}

	if password != confirmPassword {
		return fmt.Errorf("passwords do not match")
	}

	registerData := RegisterRequest{
		Username: username,
		Password: password,
		Email:    email,
	}

	response, err := utils.MakeRequest("POST", "/api/register", registerData, "")
	if err != nil {
		return fmt.Errorf("registration failed: %v", err)
	}

	var registerResp RegisterResponse
	if err := json.Unmarshal(response, &registerResp); err != nil {
		return fmt.Errorf("failed to parse registration response: %v", err)
	}

	fmt.Printf("✅ User '%s' registered successfully!\n", registerResp.User.Username)
	return nil
}

// GetCurrentToken retrieves the current user's token from keychain
func GetCurrentToken() (string, error) {
	username, err := getCurrentUser()
	if err != nil {
		return "", err
	}

	// Get token from keychain
	keyManager := NewKeyringManager()
	token, err := keyManager.GetToken(username)
	if err != nil {
		return "", fmt.Errorf("no token found, please login first: %v", err)
	}

	return token, nil
}

// IsLoggedIn checks if the user is currently logged in
func IsLoggedIn() bool {
	_, err := GetCurrentToken()
	return err == nil
}

// GetCurrentUser returns the currently logged-in username
func GetCurrentUser() (string, error) {
	return getCurrentUser()
}

// Logout removes the stored token
func Logout() error {
	username, err := getCurrentUser()
	if err != nil {
		return err
	}

	keyManager := NewKeyringManager()
	return keyManager.DeleteToken(username)
}

// TestAuthentication tests if the current token is valid
func TestAuthentication() error {
	token, err := GetCurrentToken()
	if err != nil {
		return err
	}

	_, err = utils.MakeRequest("GET", "/api/test-auth", nil, token)
	return err
}

// Helper functions

func getCurrentUser() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %v", err)
	}

	userConfigPath := homeDir + "/.certifycli/user"
	username, err := os.ReadFile(userConfigPath)
	if err != nil {
		return "", fmt.Errorf("no user configured, please run 'setup' first: %v", err)
	}

	return strings.TrimSpace(string(username)), nil
}

func saveCurrentUser(username string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %v", err)
	}

	configDir := homeDir + "/.certifycli"
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	userConfigPath := configDir + "/user"
	return os.WriteFile(userConfigPath, []byte(username), 0600)
}
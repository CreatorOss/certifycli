package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// GetCredentials prompts the user for email and password
func GetCredentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	// Get email
	fmt.Print("Email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		return "", "", fmt.Errorf("failed to read email: %w", err)
	}
	email = strings.TrimSpace(email)

	// Get password (hidden input)
	fmt.Print("Password: ")
	passwordBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", fmt.Errorf("failed to read password: %w", err)
	}
	fmt.Println() // Add newline after password input

	password := string(passwordBytes)
	return email, password, nil
}

// PromptConfirmation asks the user for yes/no confirmation
func PromptConfirmation(message string) bool {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Printf("%s (y/n): ", message)
		response, err := reader.ReadString('\n')
		if err != nil {
			return false
		}
		
		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
		
		fmt.Println("Please enter 'y' or 'n'")
	}
}

// GetInput prompts the user for input with a message
func GetInput(message string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message + ": ")
	
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}
	
	return strings.TrimSpace(input), nil
}

// PrintError prints an error message in red (if terminal supports colors)
func PrintError(message string) {
	fmt.Printf("\033[31m❌ Error: %s\033[0m\n", message)
}

// PrintSuccess prints a success message in green (if terminal supports colors)
func PrintSuccess(message string) {
	fmt.Printf("\033[32m✅ %s\033[0m\n", message)
}

// PrintWarning prints a warning message in yellow (if terminal supports colors)
func PrintWarning(message string) {
	fmt.Printf("\033[33m⚠️  Warning: %s\033[0m\n", message)
}

// PrintInfo prints an info message in blue (if terminal supports colors)
func PrintInfo(message string) {
	fmt.Printf("\033[34mℹ️  %s\033[0m\n", message)
}

// ValidateEmail performs basic email validation
func ValidateEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// FileExists checks if a file exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// CreateDirIfNotExists creates a directory if it doesn't exist
func CreateDirIfNotExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
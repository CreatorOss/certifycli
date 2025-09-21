package git

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/CreatorOss/sertifycli/internal/crypto"
)

type GitService struct {
	configDir string
}

func NewGitService() (*GitService, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %v", err)
	}
	
	return &GitService{
		configDir: filepath.Join(homeDir, ".certifycli"),
	}, nil
}

// ConfigureGitSigning sets up Git to use CertifyCLI for commit signing
func (gs *GitService) ConfigureGitSigning() error {
	// Read username from config
	username, err := gs.getUsername()
	if err != nil {
		return fmt.Errorf("failed to get username: %v", err)
	}

	// Get private key from keychain (to ensure it's accessible)
	_, err = crypto.LoadPrivateKeyFromKeyring(username)
	if err != nil {
		return fmt.Errorf("failed to load private key: %v", err)
	}

	// Get the path to the certifycli binary
	cliPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %v", err)
	}

	// Check if git is available
	if !gs.isGitAvailable() {
		return fmt.Errorf("git is not installed or not available in PATH")
	}

	// Set Git configuration
	commands := [][]string{
		{"config", "--global", "user.signingkey", username},
		{"config", "--global", "gpg.format", "x509"},
		{"config", "--global", "gpg.x509.program", cliPath + " git sign"},
		{"config", "--global", "commit.gpgsign", "true"},
		{"config", "--global", "tag.gpgsign", "true"},
	}

	for _, args := range commands {
		cmd := exec.Command("git", args...)
		if output, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("failed to set git config %s: %v\nOutput: %s", 
				args[2], err, string(output))
		}
	}

	fmt.Println("✅ Git signing configured successfully!")
	fmt.Println("   - Commit signing: enabled")
	fmt.Println("   - Tag signing: enabled")
	fmt.Println("   - Using CertifyCLI as signing tool")
	fmt.Printf("   - Signing key: %s\n", username)

	return nil
}

// SignCommit is called by Git when it needs to sign a commit
func (gs *GitService) SignCommit() error {
	// Git passes the commit content through stdin
	commitContent, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to read commit content: %v", err)
	}

	// Read username from config
	username, err := gs.getUsername()
	if err != nil {
		return fmt.Errorf("failed to get username: %v", err)
	}

	// Load private key from keychain
	privateKey, err := crypto.LoadPrivateKeyFromKeyring(username)
	if err != nil {
		return fmt.Errorf("failed to load private key: %v", err)
	}

	// Load certificate
	certPath := filepath.Join(gs.configDir, "certificate.pem")
	certData, err := os.ReadFile(certPath)
	if err != nil {
		return fmt.Errorf("failed to read certificate: %v", err)
	}

	// Create signature
	signature, err := crypto.SignData(privateKey, commitContent)
	if err != nil {
		return fmt.Errorf("failed to sign commit: %v", err)
	}

	// Output the signature in a format Git expects
	// This is a simplified format for demonstration
	fmt.Printf("-----BEGIN CERTIFYCLI SIGNATURE-----\n")
	fmt.Printf("Version: CertifyCLI 1.0\n")
	fmt.Printf("Signer: %s\n", username)
	fmt.Printf("Certificate:\n%s\n", string(certData))
	fmt.Printf("Signature: %s\n", signature)
	fmt.Printf("-----END CERTIFYCLI SIGNATURE-----\n")

	return nil
}

// VerifyGitConfig checks if Git is properly configured for signing
func (gs *GitService) VerifyGitConfig() (map[string]string, error) {
	if !gs.isGitAvailable() {
		return nil, fmt.Errorf("git is not installed or not available in PATH")
	}

	configKeys := []string{
		"user.signingkey",
		"gpg.format",
		"gpg.x509.program",
		"commit.gpgsign",
		"tag.gpgsign",
		"user.name",
		"user.email",
	}

	results := make(map[string]string)

	for _, key := range configKeys {
		cmd := exec.Command("git", "config", "--global", "--get", key)
		output, err := cmd.Output()
		if err != nil {
			results[key] = "NOT SET"
		} else {
			results[key] = strings.TrimSpace(string(output))
		}
	}

	return results, nil
}

// DisableGitSigning removes CertifyCLI Git signing configuration
func (gs *GitService) DisableGitSigning() error {
	if !gs.isGitAvailable() {
		return fmt.Errorf("git is not installed or not available in PATH")
	}

	commands := [][]string{
		{"config", "--global", "--unset", "user.signingkey"},
		{"config", "--global", "--unset", "gpg.format"},
		{"config", "--global", "--unset", "gpg.x509.program"},
		{"config", "--global", "--unset", "commit.gpgsign"},
		{"config", "--global", "--unset", "tag.gpgsign"},
	}

	for _, args := range commands {
		cmd := exec.Command("git", args...)
		// Ignore errors for unset operations (key might not exist)
		cmd.Run()
	}

	fmt.Println("✅ Git signing disabled successfully!")
	return nil
}

// TestGitSigning creates a test commit to verify signing works
func (gs *GitService) TestGitSigning() error {
	if !gs.isGitAvailable() {
		return fmt.Errorf("git is not installed or not available in PATH")
	}

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "certifycli-git-test")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Initialize git repo
	cmd := exec.Command("git", "init")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize git repo: %v", err)
	}

	// Set local user config
	username, _ := gs.getUsername()
	localCommands := [][]string{
		{"config", "user.name", "CertifyCLI Test User"},
		{"config", "user.email", username + "@certifycli.test"},
	}

	for _, args := range localCommands {
		cmd := exec.Command("git", args...)
		cmd.Dir = tempDir
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to set local git config: %v", err)
		}
	}

	// Create a test file
	testFile := filepath.Join(tempDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("CertifyCLI Git signing test"), 0644); err != nil {
		return fmt.Errorf("failed to create test file: %v", err)
	}

	// Add file to git
	cmd = exec.Command("git", "add", "test.txt")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add file to git: %v", err)
	}

	// Commit with signing
	cmd = exec.Command("git", "commit", "-m", "Test commit with CertifyCLI signing")
	cmd.Dir = tempDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to commit: %v\nOutput: %s", err, string(output))
	}

	fmt.Println("✅ Git signing test successful!")
	fmt.Printf("Test repository created at: %s\n", tempDir)
	fmt.Printf("Commit output: %s\n", string(output))

	return nil
}

// Helper functions

func (gs *GitService) getUsername() (string, error) {
	userConfigPath := filepath.Join(gs.configDir, "user")
	username, err := os.ReadFile(userConfigPath)
	if err != nil {
		return "", fmt.Errorf("no user configured: %v", err)
	}
	return strings.TrimSpace(string(username)), nil
}

func (gs *GitService) isGitAvailable() bool {
	cmd := exec.Command("git", "--version")
	return cmd.Run() == nil
}

// GetGitVersion returns the installed Git version
func (gs *GitService) GetGitVersion() (string, error) {
	if !gs.isGitAvailable() {
		return "", fmt.Errorf("git is not installed")
	}

	cmd := exec.Command("git", "--version")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get git version: %v", err)
	}

	return strings.TrimSpace(string(output)), nil
}
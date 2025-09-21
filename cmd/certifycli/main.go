package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/CreatorOss/certifycli/internal/auth"
	"github.com/CreatorOss/certifycli/internal/crypto"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "login":
		fmt.Println("Login command not yet implemented.")
	case "setup":
		handleSetup()
	case "status":
		handleStatus()
	case "test-crypto":
		handleTestCrypto()
	case "test-keyring":
		handleTestKeyring()
	case "cleanup":
		handleCleanup()
	case "--help", "-h", "help":
		printHelp()
	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		printHelp()
		os.Exit(1)
	}
}

func handleSetup() {
	fmt.Println("🔧 Setting up your CertifyCLI identity...")

	// Get username from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	if username == "" {
		fmt.Println("❌ Username cannot be empty.")
		os.Exit(1)
	}

	// Check if user already has a key in keyring
	if crypto.HasPrivateKeyInKeyring(username) {
		fmt.Printf("⚠️  Private key already exists in keyring for user: %s\n", username)
		fmt.Print("Do you want to overwrite it? (y/N): ")
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))
		if response != "y" && response != "yes" {
			fmt.Println("Setup cancelled.")
			return
		}
	}

	// 1. Generate a new private key
	fmt.Println("🔑 Generating 2048-bit RSA key pair...")
	privateKey, err := crypto.GenerateKeyPair(2048)
	if err != nil {
		fmt.Printf("❌ Error generating key pair: %v\n", err)
		os.Exit(1)
	}

	// 2. Save the private key to OS keychain (SECURE!)
	fmt.Println("🔐 Saving private key to OS keychain...")
	if err := crypto.SavePrivateKeyToKeyring(privateKey, username); err != nil {
		fmt.Printf("❌ Error saving private key to keyring: %v\n", err)
		fmt.Println("💡 Note: You may need to grant permission to access the keychain")
		os.Exit(1)
	}

	// 3. Create a test certificate (will be replaced with real CA-signed cert later)
	fmt.Println("📜 Creating a test certificate...")
	testCert, err := crypto.GenerateTestCertificate(privateKey, username)
	if err != nil {
		fmt.Printf("❌ Error generating test certificate: %v\n", err)
		os.Exit(1)
	}

	// Save the certificate to config directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("❌ Error finding home directory: %v\n", err)
		os.Exit(1)
	}
	configDir := filepath.Join(homeDir, ".certifycli")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		fmt.Printf("❌ Error creating config directory: %v\n", err)
		os.Exit(1)
	}

	certPath := filepath.Join(configDir, "certificate.pem")
	if err := os.WriteFile(certPath, testCert, 0600); err != nil {
		fmt.Printf("❌ Error saving certificate: %v\n", err)
		os.Exit(1)
	}

	// Save username to config file
	userConfigPath := filepath.Join(configDir, "user")
	if err := os.WriteFile(userConfigPath, []byte(username), 0600); err != nil {
		fmt.Printf("❌ Error saving user config: %v\n", err)
		os.Exit(1)
	}

	// Get public key fingerprint
	fingerprint, err := crypto.GetPublicKeyFingerprintFromKeyring(username)
	if err != nil {
		fingerprint = "unable to generate"
	}

	fmt.Println("✅ Setup complete!")
	fmt.Printf("👤 Username: %s\n", username)
	fmt.Printf("🔐 Private key: Securely stored in OS keychain\n")
	fmt.Printf("📄 Test certificate: %s\n", certPath)
	fmt.Printf("🔍 Public key fingerprint: %s\n", fingerprint)
	fmt.Println("\n🚀 Next steps: Run 'certifycli login' to authenticate with the server.")
}

func handleStatus() {
	fmt.Println("📊 CertifyCLI Status")
	fmt.Println("==================")

	// Try to read username from config
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("❌ Error finding home directory: %v\n", err)
		return
	}

	userConfigPath := filepath.Join(homeDir, ".certifycli", "user")
	usernameBytes, err := os.ReadFile(userConfigPath)
	if err != nil {
		fmt.Println("❌ No identity configured. Please run 'certifycli setup' first.")
		return
	}

	username := strings.TrimSpace(string(usernameBytes))
	fmt.Printf("👤 Current user: %s\n", username)

	// Check private key in keyring
	if crypto.HasPrivateKeyInKeyring(username) {
		fmt.Println("🔐 Private key: ✅ Found in OS keychain (secure)")
		
		// Try to get fingerprint
		fingerprint, err := crypto.GetPublicKeyFingerprintFromKeyring(username)
		if err != nil {
			fmt.Printf("⚠️  Warning: Cannot generate fingerprint: %v\n", err)
		} else {
			fmt.Printf("🔍 Public key fingerprint: %s\n", fingerprint)
		}
	} else {
		fmt.Println("🔐 Private key: ❌ Not found in keychain")
	}

	// Check certificate file
	certPath := filepath.Join(homeDir, ".certifycli", "certificate.pem")
	if _, err := os.Stat(certPath); err == nil {
		fmt.Println("📄 Certificate: ✅ Found")
	} else {
		fmt.Println("📄 Certificate: ❌ Not found")
	}

	// Check token (for future login implementation)
	keyManager := auth.NewKeyringManager()
	if keyManager.HasToken(username) {
		fmt.Println("🎫 Auth token: ✅ Found in keychain")
	} else {
		fmt.Println("🎫 Auth token: ❌ Not found (not logged in)")
	}

	fmt.Println("🌐 Server: ❌ Not connected (login not implemented)")
	fmt.Println("\n💡 Run 'certifycli setup' to generate your identity")
}

func handleTestCrypto() {
	fmt.Println("🧪 Testing crypto functions...")

	// Test key generation
	fmt.Println("1. Testing key generation...")
	privateKey, err := crypto.GenerateKeyPair(2048)
	if err != nil {
		fmt.Printf("❌ Key generation failed: %v\n", err)
		return
	}
	fmt.Println("✅ Key generation successful")

	// Test CSR creation
	fmt.Println("2. Testing CSR creation...")
	csrPEM, err := crypto.CreateCSR(privateKey, "test@example.com")
	if err != nil {
		fmt.Printf("❌ CSR creation failed: %v\n", err)
		return
	}
	fmt.Printf("✅ CSR creation successful:\n%s\n", string(csrPEM))

	// Test certificate generation
	fmt.Println("3. Testing certificate generation...")
	certPEM, err := crypto.GenerateTestCertificate(privateKey, "test@certifycli.com")
	if err != nil {
		fmt.Printf("❌ Certificate generation failed: %v\n", err)
		return
	}
	fmt.Printf("✅ Certificate generation successful:\n%s\n", string(certPEM))

	fmt.Println("🎉 All crypto tests passed!")
}

func handleTestKeyring() {
	fmt.Println("🧪 Testing keyring functions...")
	
	testUser := "test_keyring_user"
	
	// Test key generation and keyring storage
	fmt.Println("1. Testing key generation and keyring storage...")
	privateKey, err := crypto.GenerateKeyPair(2048)
	if err != nil {
		fmt.Printf("❌ Key generation failed: %v\n", err)
		return
	}
	
	// Save to keyring
	err = crypto.SavePrivateKeyToKeyring(privateKey, testUser)
	if err != nil {
		fmt.Printf("❌ Keyring save failed: %v\n", err)
		return
	}
	fmt.Println("✅ Key saved to keyring successfully")
	
	// Load from keyring
	fmt.Println("2. Testing key loading from keyring...")
	loadedKey, err := crypto.LoadPrivateKeyFromKeyring(testUser)
	if err != nil {
		fmt.Printf("❌ Keyring load failed: %v\n", err)
		return
	}
	fmt.Println("✅ Key loaded from keyring successfully")
	
	// Test fingerprint
	fmt.Println("3. Testing fingerprint generation...")
	fingerprint, err := crypto.GetPublicKeyFingerprintFromKeyring(testUser)
	if err != nil {
		fmt.Printf("❌ Fingerprint generation failed: %v\n", err)
		return
	}
	fmt.Printf("✅ Fingerprint: %s\n", fingerprint)
	
	// Cleanup
	fmt.Println("4. Cleaning up test data...")
	err = crypto.DeletePrivateKeyFromKeyring(testUser)
	if err != nil {
		fmt.Printf("⚠️  Cleanup warning: %v\n", err)
	} else {
		fmt.Println("✅ Test data cleaned up")
	}
	
	fmt.Println("🎉 All keyring tests passed!")
}

func handleCleanup() {
	fmt.Println("🧹 Cleaning up CertifyCLI data...")
	
	// Get username
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("❌ Error finding home directory: %v\n", err)
		return
	}

	userConfigPath := filepath.Join(homeDir, ".certifycli", "user")
	usernameBytes, err := os.ReadFile(userConfigPath)
	if err != nil {
		fmt.Println("⚠️  No user config found, cleaning up config directory only...")
	} else {
		username := strings.TrimSpace(string(usernameBytes))
		fmt.Printf("🗑️  Removing keyring data for user: %s\n", username)
		
		// Remove from keyring
		keyManager := auth.NewKeyringManager()
		keyManager.DeletePrivateKey(username)
		keyManager.DeleteToken(username)
	}
	
	// Remove config directory
	configDir := filepath.Join(homeDir, ".certifycli")
	err = os.RemoveAll(configDir)
	if err != nil {
		fmt.Printf("❌ Error removing config directory: %v\n", err)
	} else {
		fmt.Println("✅ Config directory removed")
	}
	
	fmt.Println("🎉 Cleanup complete!")
}

func printHelp() {
	fmt.Println("CertifyCLI - Global Identity for the Command Line")
	fmt.Println("================================================")
	fmt.Println("\nUsage:")
	fmt.Println("  certifycli <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  login         Authenticate with the CertifyCLI server (coming soon)")
	fmt.Println("  setup         Set up your identity and generate certificates ✅")
	fmt.Println("  status        Show your current identity status ✅")
	fmt.Println("  test-crypto   Test cryptographic functions ✅")
	fmt.Println("  test-keyring  Test OS keychain integration ✅")
	fmt.Println("  cleanup       Remove all CertifyCLI data ✅")
	fmt.Println("  --help, -h    Show this help message")
	fmt.Println("\nSecurity Features:")
	fmt.Println("  🔐 Private keys stored in OS keychain (macOS/Windows/Linux)")
	fmt.Println("  🔒 No plaintext keys on disk")
	fmt.Println("  🛡️  Secure token storage for authentication")
	fmt.Println("\nExamples:")
	fmt.Println("  certifycli setup")
	fmt.Println("  certifycli status")
	fmt.Println("  certifycli test-keyring")
}
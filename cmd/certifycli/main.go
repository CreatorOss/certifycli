package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/your-username/certifycli/internal/crypto"
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

	// Determine the config directory (e.g., ~/.certifycli)
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("❌ Error finding home directory: %v\n", err)
		os.Exit(1)
	}
	configDir := filepath.Join(homeDir, ".certifycli")
	privateKeyPath := filepath.Join(configDir, "identity.pem")

	// Create the config directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0700); err != nil {
		fmt.Printf("❌ Error creating config directory: %v\n", err)
		os.Exit(1)
	}

	// Check if key already exists
	if _, err := os.Stat(privateKeyPath); err == nil {
		fmt.Printf("⚠️  Private key already exists at: %s\n", privateKeyPath)
		fmt.Print("Do you want to overwrite it? (y/N): ")
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" {
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

	// 2. Save the private key securely
	fmt.Printf("💾 Saving private key to: %s\n", privateKeyPath)
	if err := crypto.SavePrivateKeyToPEM(privateKey, privateKeyPath); err != nil {
		fmt.Printf("❌ Error saving private key: %v\n", err)
		os.Exit(1)
	}

	// 3. Create a CSR (For now, we'll just generate a test self-signed cert for demo)
	fmt.Println("📜 Creating a test certificate...")
	// In the future, we will create a CSR and send it to the CA server.
	// For now, we generate a self-signed cert for testing.
	testCert, err := crypto.GenerateTestCertificate(privateKey, "test-user@certifycli")
	if err != nil {
		fmt.Printf("❌ Error generating test certificate: %v\n", err)
		os.Exit(1)
	}

	// Save the certificate
	certPath := filepath.Join(configDir, "certificate.pem")
	if err := os.WriteFile(certPath, testCert, 0600); err != nil {
		fmt.Printf("❌ Error saving certificate: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Setup complete!")
	fmt.Printf("🔐 Private key saved to: %s\n", privateKeyPath)
	fmt.Printf("📄 Test certificate saved to: %s\n", certPath)
	fmt.Println("\n🚀 Next steps: Run 'certifycli login' to authenticate with the server.")
}

func handleStatus() {
	fmt.Println("📊 CertifyCLI Status")
	fmt.Println("==================")

	// Check if config directory exists
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("❌ Error finding home directory: %v\n", err)
		return
	}
	configDir := filepath.Join(homeDir, ".certifycli")
	privateKeyPath := filepath.Join(configDir, "identity.pem")
	certPath := filepath.Join(configDir, "certificate.pem")

	// Check private key
	if _, err := os.Stat(privateKeyPath); err == nil {
		fmt.Println("🔑 Private Key: ✅ Generated")
		
		// Try to load and validate the key
		_, err := crypto.LoadPrivateKeyFromPEM(privateKeyPath)
		if err != nil {
			fmt.Printf("⚠️  Warning: Private key exists but cannot be loaded: %v\n", err)
		} else {
			fmt.Println("🔓 Private Key: ✅ Valid and loadable")
		}
	} else {
		fmt.Println("🔑 Private Key: ❌ Not found")
	}

	// Check certificate
	if _, err := os.Stat(certPath); err == nil {
		fmt.Println("📄 Certificate: ✅ Generated")
	} else {
		fmt.Println("📄 Certificate: ❌ Not found")
	}

	// Check server connectivity (placeholder)
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

func printHelp() {
	fmt.Println("CertifyCLI - Global Identity for the Command Line")
	fmt.Println("================================================")
	fmt.Println("\nUsage:")
	fmt.Println("  certifycli <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  login       Authenticate with the CertifyCLI server")
	fmt.Println("  setup       Set up your identity and generate certificates")
	fmt.Println("  status      Show your current identity status")
	fmt.Println("  test-crypto Test cryptographic functions")
	fmt.Println("  --help, -h  Show this help message")
	fmt.Println("\nExamples:")
	fmt.Println("  certifycli setup")
	fmt.Println("  certifycli status")
	fmt.Println("  certifycli test-crypto")
}
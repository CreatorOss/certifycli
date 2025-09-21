package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/CreatorOss/sertifycli/internal/auth"
	"github.com/CreatorOss/sertifycli/internal/ca"
	"github.com/CreatorOss/sertifycli/internal/crypto"
	"github.com/CreatorOss/sertifycli/internal/git"
	"github.com/CreatorOss/sertifycli/internal/utils"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "setup":
		handleSetup()
	case "status":
		handleStatus()
	case "certificates", "certs":
		handleCertificates()
	case "backup":
		handleBackup()
	case "restore":
		handleRestore()
	case "git":
		handleGitCommands()
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
	fmt.Println("🔧 Setting up your CertifyCLI local identity...")

	// Get username from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	if username == "" {
		fmt.Println("❌ Username cannot be empty.")
		os.Exit(1)
	}

	// Initialize local CA
	localCA, err := ca.NewLocalCA()
	if err != nil {
		fmt.Printf("❌ Error creating local CA: %v\n", err)
		os.Exit(1)
	}

	if err := localCA.InitializeCA(); err != nil {
		fmt.Printf("❌ Error initializing CA: %v\n", err)
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

	// 2. Save the private key to OS keychain
	fmt.Println("🔐 Saving private key to OS keychain...")
	if err := crypto.SavePrivateKeyToKeyring(privateKey, username); err != nil {
		fmt.Printf("❌ Error saving private key to keyring: %v\n", err)
		fmt.Println("💡 Note: You may need to grant permission to access the keychain")
		os.Exit(1)
	}

	// 3. Create CSR
	fmt.Println("📜 Creating Certificate Signing Request...")
	csrPEM, err := crypto.CreateCSR(privateKey, username)
	if err != nil {
		fmt.Printf("❌ Error creating CSR: %v\n", err)
		os.Exit(1)
	}

	// 4. Sign with local CA
	fmt.Println("🏛️  Signing certificate with local CA...")
	certificate, err := localCA.SignCSR(csrPEM, username)
	if err != nil {
		fmt.Printf("❌ Error signing certificate: %v\n", err)
		os.Exit(1)
	}

	// Save the certificate and config
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

	// Save certificate
	certPath := filepath.Join(configDir, "certificate.pem")
	if err := os.WriteFile(certPath, certificate, 0600); err != nil {
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
	fmt.Printf("📄 Local CA-signed certificate: %s\n", certPath)
	fmt.Printf("🔍 Public key fingerprint: %s\n", fingerprint)
	fmt.Println("\n🎉 Your local identity is now ready to use!")
	fmt.Println("💡 Next step: Run 'certifycli git configure' to enable Git commit signing")
}

func handleStatus() {
	fmt.Println("📊 CertifyCLI Status (Local Mode)")
	fmt.Println("===============================")

	// Try to read username from config
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("❌ Error finding home directory: %v\n", err)
		return
	}

	configDir := filepath.Join(homeDir, ".certifycli")
	userConfigPath := filepath.Join(configDir, "user")
	usernameBytes, err := os.ReadFile(userConfigPath)
	if err != nil {
		fmt.Println("❌ No identity configured. Please run 'certifycli setup' first.")
		return
	}

	username := strings.TrimSpace(string(usernameBytes))
	fmt.Printf("👤 Current user: %s\n", username)
	fmt.Println("🏠 Mode: Local (No server required)")

	// Check Local CA
	localCA, err := ca.NewLocalCA()
	if err == nil && localCA.CAExists() {
		fmt.Println("🏛️  Local CA: ✅ Configured")
		
		// Get CA info
		if caInfo, err := localCA.GetCAInfo(); err == nil {
			fmt.Printf("   Subject: %s\n", caInfo.Subject)
			fmt.Printf("   Valid until: %s (%d days)\n", 
				utils.FormatTime(caInfo.NotAfter), caInfo.DaysUntilExpiry())
		}
	} else {
		fmt.Println("🏛️  Local CA: ❌ Not configured")
	}

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
	certPath := filepath.Join(configDir, "certificate.pem")
	if _, err := os.Stat(certPath); err == nil {
		fmt.Println("📄 Certificate: ✅ Found")
		
		// Try to parse certificate
		certData, err := os.ReadFile(certPath)
		if err == nil {
			if certInfo, err := crypto.GetCertificateInfo(certData); err == nil {
				fmt.Printf("   Subject: %s\n", certInfo.Subject)
				fmt.Printf("   Issuer: %s\n", certInfo.Issuer)
				fmt.Printf("   Valid until: %s (%d days)\n", 
					utils.FormatTime(certInfo.NotAfter), certInfo.DaysUntilExpiry())
				if certInfo.IsExpired() {
					fmt.Println("   ⚠️  Certificate is expired!")
				}
			}
		}
	} else {
		fmt.Println("📄 Certificate: ❌ Not found")
	}

	// Check Git integration
	gitService, err := git.NewGitService()
	if err == nil {
		config, err := gitService.VerifyGitConfig()
		if err == nil {
			fmt.Print("🔧 Git integration: ")
			if config["gpg.x509.program"] != "NOT SET" && strings.Contains(config["gpg.x509.program"], "certifycli") {
				fmt.Println("✅ Configured")
			} else {
				fmt.Println("❌ Not configured")
			}
		}
	}

	fmt.Println("\n💡 Available commands:")
	fmt.Println("  - certifycli git configure (setup Git signing)")
	fmt.Println("  - certifycli backup (backup your identity)")
	fmt.Println("  - certifycli certificates (view certificate info)")
}

func handleCertificates() {
	fmt.Println("📋 Certificate Information")
	fmt.Println("========================")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("❌ Error finding home directory: %v\n", err)
		return
	}

	configDir := filepath.Join(homeDir, ".certifycli")
	
	// Show CA certificate info
	localCA, err := ca.NewLocalCA()
	if err == nil && localCA.CAExists() {
		fmt.Println("🏛️  Certificate Authority:")
		if caInfo, err := localCA.GetCAInfo(); err == nil {
			fmt.Printf("   Subject: %s\n", caInfo.Subject)
			fmt.Printf("   Serial: %s\n", caInfo.SerialNumber)
			fmt.Printf("   Valid from: %s\n", utils.FormatTime(caInfo.NotBefore))
			fmt.Printf("   Valid to: %s\n", utils.FormatTime(caInfo.NotAfter))
			fmt.Printf("   Days until expiry: %d\n", caInfo.DaysUntilExpiry())
		}
	}

	// Show user certificate info
	certPath := filepath.Join(configDir, "certificate.pem")
	if _, err := os.Stat(certPath); err == nil {
		fmt.Println("\n👤 User Certificate:")
		certData, err := os.ReadFile(certPath)
		if err == nil {
			if certInfo, err := crypto.GetCertificateInfo(certData); err == nil {
				fmt.Printf("   Subject: %s\n", certInfo.Subject)
				fmt.Printf("   Issuer: %s\n", certInfo.Issuer)
				fmt.Printf("   Serial: %s\n", certInfo.SerialNumber)
				fmt.Printf("   Valid from: %s\n", utils.FormatTime(certInfo.NotBefore))
				fmt.Printf("   Valid to: %s\n", utils.FormatTime(certInfo.NotAfter))
				fmt.Printf("   Days until expiry: %d\n", certInfo.DaysUntilExpiry())
				if certInfo.IsExpired() {
					fmt.Println("   ⚠️  Certificate is expired!")
				} else {
					fmt.Println("   ✅ Certificate is valid")
				}
			}
		}
	} else {
		fmt.Println("\n👤 User Certificate: ❌ Not found")
	}

	// Show file locations
	fmt.Println("\n📁 File Locations:")
	fmt.Printf("   Config directory: %s\n", configDir)
	fmt.Printf("   CA certificate: %s\n", filepath.Join(configDir, "ca-certificate.pem"))
	fmt.Printf("   User certificate: %s\n", certPath)
	fmt.Printf("   User config: %s\n", filepath.Join(configDir, "user"))
}

func handleBackup() {
	fmt.Println("💾 Backing up CertifyCLI identity...")
	
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("❌ Error finding home directory: %v\n", err)
		os.Exit(1)
	}

	sourceDir := filepath.Join(homeDir, ".certifycli")
	backupDir := filepath.Join(homeDir, "certifycli-backup")

	// Check if source exists
	if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
		fmt.Println("❌ No CertifyCLI identity found. Please run 'certifycli setup' first.")
		os.Exit(1)
	}

	fmt.Printf("📂 Backing up %s to %s\n", sourceDir, backupDir)

	// Remove existing backup
	os.RemoveAll(backupDir)

	// Copy directory
	if err := copyDir(sourceDir, backupDir); err != nil {
		fmt.Printf("❌ Backup failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Backup completed successfully!")
	fmt.Printf("📁 Backup location: %s\n", backupDir)
	fmt.Println("💡 Keep this backup in a secure location")
}

func handleRestore() {
	fmt.Println("🔄 Restoring CertifyCLI identity...")
	
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("❌ Error finding home directory: %v\n", err)
		os.Exit(1)
	}

	backupDir := filepath.Join(homeDir, "certifycli-backup")
	targetDir := filepath.Join(homeDir, ".certifycli")

	// Check if backup exists
	if _, err := os.Stat(backupDir); os.IsNotExist(err) {
		fmt.Println("❌ No backup found. Please create backup first with 'certifycli backup'.")
		os.Exit(1)
	}

	fmt.Printf("📂 Restoring from %s to %s\n", backupDir, targetDir)

	// Confirm with user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("This will overwrite your current identity. Continue? (y/N): ")
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))
	
	if response != "y" && response != "yes" {
		fmt.Println("Restore cancelled.")
		return
	}

	// Remove existing and restore
	os.RemoveAll(targetDir)
	if err := copyDir(backupDir, targetDir); err != nil {
		fmt.Printf("❌ Restore failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Restore completed successfully!")
	fmt.Println("💡 Run 'certifycli status' to verify the restoration")
}

// Git commands remain the same as before
func handleGitCommands() {
	if len(os.Args) < 3 {
		printGitHelp()
		os.Exit(1)
	}

	gitService, err := git.NewGitService()
	if err != nil {
		fmt.Printf("❌ Error creating git service: %v\n", err)
		os.Exit(1)
	}

	switch os.Args[2] {
	case "configure":
		handleGitConfigure(gitService)
	case "sign":
		handleGitSign(gitService)
	case "status":
		handleGitStatus(gitService)
	case "disable":
		handleGitDisable(gitService)
	case "test":
		handleGitTest(gitService)
	case "verify":
		handleGitVerify(gitService)
	case "verify-all":
		handleGitVerifyAll(gitService)
	case "version":
		handleGitVersion(gitService)
	default:
		fmt.Printf("Unknown git command: %s\n", os.Args[2])
		printGitHelp()
		os.Exit(1)
	}
}

// Git handler functions (keeping the same implementation as before)
func handleGitConfigure(gitService *git.GitService) {
	fmt.Println("🔧 Configuring Git to use CertifyCLI for commit signing...")
	
	// Check if user has identity set up
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("❌ Error finding home directory: %v\n", err)
		os.Exit(1)
	}

	userConfigPath := filepath.Join(homeDir, ".certifycli", "user")
	if _, err := os.Stat(userConfigPath); os.IsNotExist(err) {
		fmt.Println("❌ No identity configured. Please run 'certifycli setup' first.")
		os.Exit(1)
	}

	if err := gitService.ConfigureGitSigning(); err != nil {
		fmt.Printf("❌ Error configuring git: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n🎉 Git integration setup complete!")
	fmt.Println("💡 Now all your commits will be automatically signed with CertifyCLI")
	fmt.Println("🚀 Try making a commit to test it out!")
}

func handleGitSign(gitService *git.GitService) {
	if err := gitService.SignCommit(); err != nil {
		fmt.Fprintf(os.Stderr, "Error signing commit: %v\n", err)
		os.Exit(1)
	}
}

func handleGitStatus(gitService *git.GitService) {
	fmt.Println("📊 Git Signing Configuration Status")
	fmt.Println("==================================")
	
	// Check Git version
	version, err := gitService.GetGitVersion()
	if err != nil {
		fmt.Printf("❌ Git: Not installed or not available\n")
		fmt.Printf("   Error: %v\n", err)
		return
	}
	fmt.Printf("✅ Git: %s\n", version)

	// Check configuration
	config, err := gitService.VerifyGitConfig()
	if err != nil {
		fmt.Printf("❌ Error checking git config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n🔧 Git Configuration:")
	fmt.Printf("  user.name: %s\n", getConfigStatus(config["user.name"]))
	fmt.Printf("  user.email: %s\n", getConfigStatus(config["user.email"]))
	fmt.Printf("  user.signingkey: %s\n", getConfigStatus(config["user.signingkey"]))
	fmt.Printf("  gpg.format: %s\n", getConfigStatus(config["gpg.format"]))
	fmt.Printf("  gpg.x509.program: %s\n", getConfigStatus(config["gpg.x509.program"]))
	fmt.Printf("  commit.gpgsign: %s\n", getConfigStatus(config["commit.gpgsign"]))
	fmt.Printf("  tag.gpgsign: %s\n", getConfigStatus(config["tag.gpgsign"]))

	// Check if configured to use certifycli
	fmt.Println("\n📋 Integration Status:")
	if config["gpg.x509.program"] != "NOT SET" && strings.Contains(config["gpg.x509.program"], "certifycli") {
		fmt.Println("✅ Git is configured to use CertifyCLI for signing!")
		if config["commit.gpgsign"] == "true" {
			fmt.Println("✅ Automatic commit signing is enabled")
		} else {
			fmt.Println("⚠️  Automatic commit signing is disabled")
		}
	} else {
		fmt.Println("❌ Git is not configured to use CertifyCLI for signing")
		fmt.Println("💡 Run 'certifycli git configure' to set it up")
	}

	// Check identity
	homeDir, _ := os.UserHomeDir()
	userConfigPath := filepath.Join(homeDir, ".certifycli", "user")
	if _, err := os.Stat(userConfigPath); err == nil {
		fmt.Println("✅ CertifyCLI identity is configured")
	} else {
		fmt.Println("❌ CertifyCLI identity not found")
		fmt.Println("💡 Run 'certifycli setup' to create your identity")
	}
}

func handleGitDisable(gitService *git.GitService) {
	fmt.Println("🚫 Disabling Git signing with CertifyCLI...")
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Are you sure you want to disable Git signing? (y/N): ")
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))
	
	if response != "y" && response != "yes" {
		fmt.Println("Operation cancelled.")
		return
	}

	if err := gitService.DisableGitSigning(); err != nil {
		fmt.Printf("❌ Error disabling git signing: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("💡 You can re-enable signing anytime with 'certifycli git configure'")
}

func handleGitTest(gitService *git.GitService) {
	fmt.Println("🧪 Testing Git signing integration...")
	
	if err := gitService.TestGitSigning(); err != nil {
		fmt.Printf("❌ Git signing test failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("🎉 Git signing test completed successfully!")
}

func handleGitVerify(gitService *git.GitService) {
	fmt.Println("🔍 Verifying last commit signature...")
	
	if err := gitService.VerifyLastCommit(); err != nil {
		fmt.Printf("❌ Error verifying commit: %v\n", err)
		os.Exit(1)
	}
}

func handleGitVerifyAll(gitService *git.GitService) {
	fmt.Println("🔍 Verifying all commit signatures in repository...")
	
	if err := gitService.VerifyAllCommits(); err != nil {
		fmt.Printf("❌ Error verifying commits: %v\n", err)
		os.Exit(1)
	}
}

func handleGitVersion(gitService *git.GitService) {
	version, err := gitService.GetGitVersion()
	if err != nil {
		fmt.Printf("❌ Error getting Git version: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Git version: %s\n", version)
}

// Test functions remain the same
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

// Helper functions
func getConfigStatus(value string) string {
	if value == "NOT SET" {
		return "❌ " + value
	}
	return "✅ " + value
}

func printGitHelp() {
	fmt.Println("CertifyCLI Git Integration Commands")
	fmt.Println("==================================")
	fmt.Println("\nUsage:")
	fmt.Println("  certifycli git <subcommand>")
	fmt.Println("\nConfiguration Commands:")
	fmt.Println("  configure  Set up Git to use CertifyCLI for signing")
	fmt.Println("  status     Check Git signing configuration")
	fmt.Println("  disable    Disable CertifyCLI Git signing")
	fmt.Println("  version    Show Git version")
	fmt.Println("\nSigning Commands:")
	fmt.Println("  test       Test Git signing with a temporary repository")
	fmt.Println("  sign       Sign a Git commit (internal use by Git)")
	fmt.Println("\nVerification Commands:")
	fmt.Println("  verify     Verify the last commit signature")
	fmt.Println("  verify-all Verify all commit signatures in repository")
	fmt.Println("\nExamples:")
	fmt.Println("  certifycli git configure   # Enable Git signing")
	fmt.Println("  certifycli git status      # Check configuration")
	fmt.Println("  certifycli git test        # Test signing")
	fmt.Println("  certifycli git verify      # Verify last commit")
	fmt.Println("  certifycli git verify-all  # Verify all commits")
}

func printHelp() {
	fmt.Println("CertifyCLI - Local Identity for the Command Line")
	fmt.Println("===============================================")
	fmt.Println("\nUsage:")
	fmt.Println("  certifycli <command> [arguments]")
	fmt.Println("\nIdentity Commands:")
	fmt.Println("  setup         Set up your local identity and generate certificates ✅")
	fmt.Println("  status        Show your current identity status ✅")
	fmt.Println("  certificates  Show certificate information ✅")
	fmt.Println("  backup        Backup your identity to ~/certifycli-backup ✅")
	fmt.Println("  restore       Restore identity from backup ✅")
	fmt.Println("\nGit Integration Commands:")
	fmt.Println("  git configure Configure Git to use CertifyCLI for signing ✅")
	fmt.Println("  git status    Check Git signing configuration ✅")
	fmt.Println("  git test      Test Git signing integration ✅")
	fmt.Println("  git verify    Verify commit signatures ✅")
	fmt.Println("\nTesting Commands:")
	fmt.Println("  test-crypto   Test cryptographic functions ✅")
	fmt.Println("  test-keyring  Test OS keychain integration ✅")
	fmt.Println("\nUtility Commands:")
	fmt.Println("  cleanup       Remove all CertifyCLI data ✅")
	fmt.Println("  --help, -h    Show this help message")
	fmt.Println("\nSecurity Features:")
	fmt.Println("  🔐 Private keys stored in OS keychain (macOS/Windows/Linux)")
	fmt.Println("  🔒 No plaintext keys on disk")
	fmt.Println("  🏛️  Local Certificate Authority (no server required)")
	fmt.Println("  🔧 Git commit signing integration")
	fmt.Println("  💾 Backup and restore functionality")
	fmt.Println("\nLocal Mode Workflow:")
	fmt.Println("  1. certifycli setup          # Generate local identity & CA")
	fmt.Println("  2. certifycli git configure  # Enable Git commit signing")
	fmt.Println("  3. git commit -m \"message\"   # All commits now automatically signed!")
	fmt.Println("  4. certifycli backup         # Backup your identity")
	fmt.Println("\nNote: Running in local mode - no server required! 🏠")
}

// Utility function for copying directories
func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		return os.WriteFile(dstPath, data, info.Mode())
	})
}
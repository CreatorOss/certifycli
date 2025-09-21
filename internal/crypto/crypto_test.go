package crypto

import (
	"fmt"
	"os"
	"testing"
)

func TestGenerateAndSaveKey(t *testing.T) {
	// Generate a new key pair
	privateKey, err := GenerateKeyPair(2048)
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	// Save the private key to a temporary file
	tmpFile := "/tmp/test_private_key.pem"
	defer os.Remove(tmpFile) // Clean up after test

	err = SavePrivateKeyToPEM(privateKey, tmpFile)
	if err != nil {
		t.Fatalf("Failed to save private key: %v", err)
	}

	// Load the private key back
	loadedKey, err := LoadPrivateKeyFromPEM(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load private key: %v", err)
	}

	// Basic check: ensure the loaded key is not nil
	if loadedKey == nil {
		t.Fatal("Loaded key is nil")
	}

	fmt.Println("✓ Key generation, save, and load test passed")
}

func TestCreateCSR(t *testing.T) {
	privateKey, err := GenerateKeyPair(2048)
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	csrPEM, err := CreateCSR(privateKey, "test@example.com")
	if err != nil {
		t.Fatalf("Failed to create CSR: %v", err)
	}

	if len(csrPEM) == 0 {
		t.Fatal("Generated CSR is empty")
	}

	fmt.Printf("✓ CSR Creation test passed. Sample CSR:\n%s\n", string(csrPEM))
}

func TestGenerateTestCertificate(t *testing.T) {
	privateKey, err := GenerateKeyPair(2048)
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	certPEM, err := GenerateTestCertificate(privateKey, "test@certifycli.com")
	if err != nil {
		t.Fatalf("Failed to generate test certificate: %v", err)
	}

	if len(certPEM) == 0 {
		t.Fatal("Generated certificate is empty")
	}

	fmt.Printf("✓ Test certificate generation passed. Sample certificate:\n%s\n", string(certPEM))
}
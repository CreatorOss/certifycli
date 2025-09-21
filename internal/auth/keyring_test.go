package auth

import (
	"fmt"
	"testing"
)

func TestKeyringOperations(t *testing.T) {
	keyManager := NewKeyringManager()
	testUser := "test_user_keyring"
	testData := "super_secret_private_key_data_12345"

	// Clean up before test (ignore errors)
	keyManager.DeletePrivateKey(testUser)

	// Test saving to keyring
	err := keyManager.SavePrivateKey(testUser, testData)
	if err != nil {
		t.Fatalf("Failed to save to keyring: %v", err)
	}

	// Test reading from keyring
	retrievedData, err := keyManager.GetPrivateKey(testUser)
	if err != nil {
		t.Fatalf("Failed to get from keyring: %v", err)
	}

	if retrievedData != testData {
		t.Fatalf("Retrieved data doesn't match. Expected: %s, Got: %s", testData, retrievedData)
	}

	// Test HasPrivateKey
	if !keyManager.HasPrivateKey(testUser) {
		t.Fatal("HasPrivateKey should return true but returned false")
	}

	// Test deleting from keyring
	err = keyManager.DeletePrivateKey(testUser)
	if err != nil {
		t.Fatalf("Failed to delete from keyring: %v", err)
	}

	// Verify it's deleted
	_, err = keyManager.GetPrivateKey(testUser)
	if err == nil {
		t.Fatal("Data should have been deleted from keyring but still exists")
	}

	// Test HasPrivateKey after deletion
	if keyManager.HasPrivateKey(testUser) {
		t.Fatal("HasPrivateKey should return false after deletion but returned true")
	}

	fmt.Println("✓ All keyring operations test passed")
}

func TestTokenOperations(t *testing.T) {
	keyManager := NewKeyringManager()
	testUser := "test_user_token"
	testToken := "jwt_token_12345_abcdef"

	// Clean up before test (ignore errors)
	keyManager.DeleteToken(testUser)

	// Test saving token
	err := keyManager.SaveToken(testUser, testToken)
	if err != nil {
		t.Fatalf("Failed to save token to keyring: %v", err)
	}

	// Test reading token
	retrievedToken, err := keyManager.GetToken(testUser)
	if err != nil {
		t.Fatalf("Failed to get token from keyring: %v", err)
	}

	if retrievedToken != testToken {
		t.Fatalf("Retrieved token doesn't match. Expected: %s, Got: %s", testToken, retrievedToken)
	}

	// Test HasToken
	if !keyManager.HasToken(testUser) {
		t.Fatal("HasToken should return true but returned false")
	}

	// Test deleting token
	err = keyManager.DeleteToken(testUser)
	if err != nil {
		t.Fatalf("Failed to delete token from keyring: %v", err)
	}

	// Verify it's deleted
	_, err = keyManager.GetToken(testUser)
	if err == nil {
		t.Fatal("Token should have been deleted from keyring but still exists")
	}

	// Test HasToken after deletion
	if keyManager.HasToken(testUser) {
		t.Fatal("HasToken should return false after deletion but returned true")
	}

	fmt.Println("✓ All token operations test passed")
}
package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/CreatorOss/certifycli/internal/auth"
)

// SavePrivateKeyToKeyring saves an RSA private key to the OS keychain
func SavePrivateKeyToKeyring(key *rsa.PrivateKey, username string) error {
	keyManager := auth.NewKeyringManager()
	
	// Convert private key to PEM format
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(key)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	return keyManager.SavePrivateKey(username, string(privateKeyPEM))
}

// LoadPrivateKeyFromKeyring loads an RSA private key from the OS keychain
func LoadPrivateKeyFromKeyring(username string) (*rsa.PrivateKey, error) {
	keyManager := auth.NewKeyringManager()
	
	keyPEM, err := keyManager.GetPrivateKey(username)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(keyPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	return privateKey, nil
}

// DeletePrivateKeyFromKeyring removes a private key from the OS keychain
func DeletePrivateKeyFromKeyring(username string) error {
	keyManager := auth.NewKeyringManager()
	return keyManager.DeletePrivateKey(username)
}

// HasPrivateKeyInKeyring checks if a private key exists in the keyring for the given username
func HasPrivateKeyInKeyring(username string) bool {
	keyManager := auth.NewKeyringManager()
	return keyManager.HasPrivateKey(username)
}

// GetPublicKeyFromKeyring loads the public key from the private key stored in keyring
func GetPublicKeyFromKeyring(username string) (*rsa.PublicKey, error) {
	privateKey, err := LoadPrivateKeyFromKeyring(username)
	if err != nil {
		return nil, err
	}
	return &privateKey.PublicKey, nil
}

// GetPublicKeyFingerprintFromKeyring gets the fingerprint of the public key stored in keyring
func GetPublicKeyFingerprintFromKeyring(username string) (string, error) {
	publicKey, err := GetPublicKeyFromKeyring(username)
	if err != nil {
		return "", err
	}
	return GetPublicKeyFingerprint(publicKey), nil
}
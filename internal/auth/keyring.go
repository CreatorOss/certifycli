package auth

import (
	"fmt"
	"github.com/zalando/go-keyring"
)

const (
	serviceName = "certifycli"
)

// KeyringManager handles secure storage of secrets using OS keychain
type KeyringManager struct {
	service string
}

func NewKeyringManager() *KeyringManager {
	return &KeyringManager{
		service: serviceName,
	}
}

// SavePrivateKey saves a private key to the OS keychain
func (km *KeyringManager) SavePrivateKey(username, keyData string) error {
	err := keyring.Set(km.service, username+"_private_key", keyData)
	if err != nil {
		return fmt.Errorf("failed to save private key to keyring: %v", err)
	}
	return nil
}

// GetPrivateKey retrieves a private key from the OS keychain
func (km *KeyringManager) GetPrivateKey(username string) (string, error) {
	keyData, err := keyring.Get(km.service, username+"_private_key")
	if err != nil {
		return "", fmt.Errorf("failed to get private key from keyring: %v", err)
	}
	return keyData, nil
}

// DeletePrivateKey removes a private key from the OS keychain
func (km *KeyringManager) DeletePrivateKey(username string) error {
	err := keyring.Delete(km.service, username+"_private_key")
	if err != nil {
		return fmt.Errorf("failed to delete private key from keyring: %v", err)
	}
	return nil
}

// SaveToken saves an authentication token to the OS keychain
func (km *KeyringManager) SaveToken(username, token string) error {
	err := keyring.Set(km.service, username+"_token", token)
	if err != nil {
		return fmt.Errorf("failed to save token to keyring: %v", err)
	}
	return nil
}

// GetToken retrieves an authentication token from the OS keychain
func (km *KeyringManager) GetToken(username string) (string, error) {
	token, err := keyring.Get(km.service, username+"_token")
	if err != nil {
		return "", fmt.Errorf("failed to get token from keyring: %v", err)
	}
	return token, nil
}

// DeleteToken removes an authentication token from the OS keychain
func (km *KeyringManager) DeleteToken(username string) error {
	err := keyring.Delete(km.service, username+"_token")
	if err != nil {
		return fmt.Errorf("failed to delete token from keyring: %v", err)
	}
	return nil
}

// HasPrivateKey checks if a private key exists for the given username
func (km *KeyringManager) HasPrivateKey(username string) bool {
	_, err := km.GetPrivateKey(username)
	return err == nil
}

// HasToken checks if a token exists for the given username
func (km *KeyringManager) HasToken(username string) bool {
	_, err := km.GetToken(username)
	return err == nil
}
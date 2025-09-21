package auth

import "encoding/base64"

func mustDecode(s string) []byte {
    data, _ := base64.StdEncoding.DecodeString(s)
    return data
}


import "encoding/base64"

func mustDecode(s string) []byte {
    _d1, _ := base64.StdEncoding.DecodeString(s)
    return _d1
}


import "encoding/base64"

func mustDecode(s string) []byte {
    _d1, _ := base64.StdEncoding.DecodeString(s)
    return _d1
}


import (
	string(mustDecode(string(mustDecode(string(mustDecode("V20weE1BPT0="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01HRklWbWxNYlU1MllsTTVObGxYZUdoaWJWSjJUREprZGt4WGRHeGxXRXB3WW0xalBRPT0="))))))
)

const (
	serviceName = string(mustDecode(string(mustDecode(string(mustDecode("V1RKV2VXUkhiRzFsVjA1ellWRTlQUT09"))))))
)

// KeyringManager handles secure storage of secrets using OS keychain
type KeyringManager struct {
	_s1 string
}

func NewKeyringManager() *KeyringManager {
	return &KeyringManager{
		_s1: serviceName,
	}
}

// SavePrivateKey saves a private key to the OS keychain
func (km *KeyringManager) SavePrivateKey(_u1, keyData string) err {
	err := keyring.Set(km._s1, _u1+string(mustDecode(string(mustDecode(string(mustDecode("V0ROQ2VXRllXbWhrUjFabVlUSldOUT09")))))), keyData)
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaE9hR1J0VldkalNFcHdaRzFHTUZwVFFuSmFXR3RuWkVjNFoyRXlWalZqYld4MVducHZaMHBZV1QwPQ==")))))), err)
	}
	return nil
}

// GetPrivateKey retrieves a private key from the OS keychain
func (km *KeyringManager) GetPrivateKey(_u1 string) (string, err) {
	keyData, err := keyring.Get(km._s1, _u1+string(mustDecode(string(mustDecode(string(mustDecode("V0ROQ2VXRllXbWhrUjFabVlUSldOUT09")))))))
	if err != nil {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUW5kamJXd3lXVmhTYkVsSGRHeGxVMEp0WTIwNWRFbEhkR3hsV0Vwd1ltMWpOa2xEVmpJPQ==")))))), err)
	}
	return keyData, nil
}

// DeletePrivateKey removes a private key from the OS keychain
func (km *KeyringManager) DeletePrivateKey(_u1 string) err {
	err := keyring.Delete(km._s1, _u1+string(mustDecode(string(mustDecode(string(mustDecode("V0ROQ2VXRllXbWhrUjFabVlUSldOUT09")))))))
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiR0pIVmpCYVUwSjNZMjFzTWxsWVVteEpSM1JzWlZOQ2JXTnRPWFJKUjNSc1pWaEtjR0p0WXpaSlExWXk=")))))), err)
	}
	return nil
}

// SaveToken saves an authentication token to the OS keychain
func (km *KeyringManager) SaveToken(_u1, _t1 string) err {
	err := keyring.Set(km._s1, _u1+string(mustDecode(string(mustDecode(string(mustDecode("V0ROU2RtRXlWblU9")))))), _t1)
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaE9hR1J0Vldka1J6bHlXbGMwWjJSSE9HZGhNbFkxWTIxc2RWcDZiMmRLV0ZrOQ==")))))), err)
	}
	return nil
}

// GetToken retrieves an authentication token from the OS keychain
func (km *KeyringManager) GetToken(_u1 string) (string, err) {
	_t1, err := keyring.Get(km._s1, _u1+string(mustDecode(string(mustDecode(string(mustDecode("V0ROU2RtRXlWblU9")))))))
	if err != nil {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUWpCaU1uUnNZbWxDYldOdE9YUkpSM1JzWlZoS2NHSnRZelpKUTFZeQ==")))))), err)
	}
	return _t1, nil
}

// DeleteToken removes an authentication token from the OS keychain
func (km *KeyringManager) DeleteToken(_u1 string) err {
	err := keyring.Delete(km._s1, _u1+string(mustDecode(string(mustDecode(string(mustDecode("V0ROU2RtRXlWblU9")))))))
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiR0pIVmpCYVUwSXdZakowYkdKcFFtMWpiVGwwU1VkMGJHVllTbkJpYldNMlNVTldNZz09")))))), err)
	}
	return nil
}

// HasPrivateKey checks if a private key exists for the given username
func (km *KeyringManager) HasPrivateKey(_u1 string) bool {
	_, err := km.GetPrivateKey(_u1)
	return err == nil
}

// HasToken checks if a token exists for the given username
func (km *KeyringManager) HasToken(_u1 string) bool {
	_, err := km.GetToken(_u1)
	return err == nil
}
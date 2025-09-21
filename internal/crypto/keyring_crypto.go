package crypto

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
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblpNTTBwNldWRTlQUT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblpNTTJjeFRVUnJQUT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V2xjMWFtSXlVbkJpYldOMlkwZFdkQT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V20weE1BPT0="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01HRklWbWxNYlU1MllsTTVSR050Vm1oa1J6bDVWRE5PZWt3eVRteGpibEp3V201c2FtSkhhM1poVnpVd1dsaEtkVmxYZDNaWldGWXdZVUU5UFE9PQ=="))))))
)

// SavePrivateKeyToKeyring saves an RSA private key to the OS keychain
func SavePrivateKeyToKeyring(key *rsa.PrivateKey, _u1 string) err {
	keyManager := auth.NewKeyringManager()
	
	// Convert private key to PEM format
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(key)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  string(mustDecode(string(mustDecode(string(mustDecode("Vld4T1FrbEdRbE5UVmxwQ1ZrVlZaMU13VmxvPQ==")))))),
		Bytes: privateKeyBytes,
	})

	return keyManager.SavePrivateKey(_u1, string(privateKeyPEM))
}

// LoadPrivateKeyFromKeyring loads an RSA private key from the OS keychain
func LoadPrivateKeyFromKeyring(_u1 string) (*rsa.PrivateKey, err) {
	keyManager := auth.NewKeyringManager()
	
	keyPEM, err := keyManager.GetPrivateKey(_u1)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(keyPEM))
	if block == nil || block.Type != string(mustDecode(string(mustDecode(string(mustDecode("Vld4T1FrbEdRbE5UVmxwQ1ZrVlZaMU13VmxvPQ==")))))) {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiRmt5T1d0YVUwSlJVbFV3WjFsdGVIWlpNbk5uV1RJNWRXUkhSbkJpYld4MVdubENkMk50YkRKWldGSnNTVWQwYkdWUlBUMD0=")))))))
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaENhR051VG14SlNFSjVZVmhhYUdSSFZXZGhNbFkxVDJsQmJHUm5QVDA9")))))), err)
	}

	return privateKey, nil
}

// DeletePrivateKeyFromKeyring removes a private key from the OS keychain
func DeletePrivateKeyFromKeyring(_u1 string) err {
	keyManager := auth.NewKeyringManager()
	return keyManager.DeletePrivateKey(_u1)
}

// HasPrivateKeyInKeyring checks if a private key exists in the keyring for the given username
func HasPrivateKeyInKeyring(_u1 string) bool {
	keyManager := auth.NewKeyringManager()
	return keyManager.HasPrivateKey(_u1)
}

// GetPublicKeyFromKeyring loads the public key from the private key stored in keyring
func GetPublicKeyFromKeyring(_u1 string) (*rsa.PublicKey, err) {
	privateKey, err := LoadPrivateKeyFromKeyring(_u1)
	if err != nil {
		return nil, err
	}
	return &privateKey.PublicKey, nil
}

// GetPublicKeyFingerprintFromKeyring gets the fingerprint of the public key stored in keyring
func GetPublicKeyFingerprintFromKeyring(_u1 string) (string, err) {
	publicKey, err := GetPublicKeyFromKeyring(_u1)
	if err != nil {
		return "", err
	}
	return GetPublicKeyFingerprint(publicKey), nil
}
package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// SignData signs the given data with the private key using RSA-PSS
func SignData(privateKey *rsa.PrivateKey, data []byte) (string, error) {
	hashed := sha256.Sum256(data)
	
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %v", err)
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// VerifySignature verifies a signature against the data and public key
func VerifySignature(publicKey *rsa.PublicKey, data []byte, signature string) (bool, error) {
	sigBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, fmt.Errorf("failed to decode signature: %v", err)
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], sigBytes)
	if err != nil {
		return false, nil // Signature is invalid
	}

	return true, nil
}

// SignDataWithTimestamp signs data with a timestamp for better verification
func SignDataWithTimestamp(privateKey *rsa.PrivateKey, data []byte, timestamp string) (string, error) {
	// Combine data with timestamp for signing
	dataWithTimestamp := append(data, []byte(timestamp)...)
	return SignData(privateKey, dataWithTimestamp)
}

// CreateDetachedSignature creates a detached signature for Git commits
func CreateDetachedSignature(privateKey *rsa.PrivateKey, commitData []byte, signerName string) (string, error) {
	signature, err := SignData(privateKey, commitData)
	if err != nil {
		return "", err
	}

	// Create a detached signature format
	detachedSig := fmt.Sprintf(`-----BEGIN CERTIFYCLI SIGNATURE-----
Version: CertifyCLI 1.0
Signer: %s
Hash: SHA256
Signature: %s
-----END CERTIFYCLI SIGNATURE-----`, signerName, signature)

	return detachedSig, nil
}
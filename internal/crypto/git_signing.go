package crypto

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"
)

// GitSignCommit creates a GPG-compatible signature for Git commits
func GitSignCommit(privateKey *rsa.PrivateKey, certificate []byte, commitContent []byte) (string, error) {
	// Create the signature payload in GPG format
	signaturePayload := createSignaturePayload(commitContent, certificate)
	
	// Sign the payload
	hashed := sha1.Sum(signaturePayload)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, hashed[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign commit: %v", err)
	}

	// Encode signature in GPG-compatible format
	return formatGPGSignature(signature, certificate), nil
}

// createSignaturePayload creates the payload that GPG would sign
func createSignaturePayload(commitContent, certificate []byte) []byte {
	var buf bytes.Buffer
	buf.WriteString("tree ")
	buf.Write(commitContent)
	buf.WriteString("\n")
	buf.WriteString("gpgsig ")
	buf.Write(certificate)
	return buf.Bytes()
}

// formatGPGSignature formats the signature in GPG-compatible way
func formatGPGSignature(signature []byte, certificate []byte) string {
	var buf bytes.Buffer
	
	// GPG header
	buf.WriteString("-----BEGIN PGP SIGNATURE-----\n")
	buf.WriteString("Version: CertifyCLI 1.0\n")
	buf.WriteString("\n")
	
	// Signature data (base64 encoded)
	sigBase64 := base64.StdEncoding.EncodeToString(signature)
	for i := 0; i < len(sigBase64); i += 64 {
		end := i + 64
		if end > len(sigBase64) {
			end = len(sigBase64)
		}
		buf.WriteString(sigBase64[i:end])
		buf.WriteString("\n")
	}
	
	buf.WriteString("-----END PGP SIGNATURE-----\n")
	
	return buf.String()
}

// VerifyGitSignature verifies a Git commit signature
func VerifyGitSignature(publicKey *rsa.PublicKey, commitContent, signature, certificate []byte) (bool, error) {
	// Decode the signature
	sigBlocks, _ := pem.Decode(signature)
	if sigBlocks == nil {
		return false, fmt.Errorf("invalid signature format")
	}

	// Recreate the signature payload
	signaturePayload := createSignaturePayload(commitContent, certificate)
	
	// Verify the signature
	hashed := sha1.Sum(signaturePayload)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA1, hashed[:], sigBlocks.Bytes)
	if err != nil {
		return false, nil
	}

	return true, nil
}

// CreateMinimalX509Certificate creates a minimal X509 cert for Git compatibility
func CreateMinimalX509Certificate(privateKey *rsa.PrivateKey, commonName string) ([]byte, error) {
	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: commonName},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageCodeSigning},
		BasicConstraintsValid: true,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	}), nil
}

// CreateDetachedGitSignature creates a detached signature for Git commits
func CreateDetachedGitSignature(privateKey *rsa.PrivateKey, commitData []byte, signerName string) (string, error) {
	// Create a simplified signature for Git
	signature, err := SignData(privateKey, commitData)
	if err != nil {
		return "", err
	}

	// Create a Git-compatible detached signature format
	detachedSig := fmt.Sprintf(`-----BEGIN PGP SIGNATURE-----
Version: CertifyCLI 1.0
Comment: Signed with CertifyCLI

%s
-----END PGP SIGNATURE-----`, signature)

	return detachedSig, nil
}

// ParseGitSignature parses a Git signature and extracts components
func ParseGitSignature(signatureData []byte) (signature []byte, version string, err error) {
	// Parse PGP signature format
	block, _ := pem.Decode(signatureData)
	if block == nil {
		return nil, "", fmt.Errorf("invalid signature format")
	}

	if block.Type != "PGP SIGNATURE" {
		return nil, "", fmt.Errorf("not a PGP signature")
	}

	// Extract version from headers
	version = "Unknown"
	if versionHeader, ok := block.Headers["Version"]; ok {
		version = versionHeader
	}

	return block.Bytes, version, nil
}
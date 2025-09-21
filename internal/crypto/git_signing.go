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
	string(mustDecode(string(mustDecode(string(mustDecode("V1c1c01GcFlUVDA9"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblk9"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblpNTTBwb1ltMVJQUT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblpNTTBwNldWRTlQUT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblpNTTA1dldWUkZQUT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblpNTTJjeFRVUnJQUT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblpNTTJjeFRVUnJkbU5IZEhCbFFUMDk="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V2xjMWFtSXlVbkJpYldOMldXMUdlbHBVV1RBPQ=="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V2xjMWFtSXlVbkJpYldOMlkwZFdkQT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V20weE1BPT0="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("WWxkR01HRkRPV2xoVjJNOQ=="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("WkVkc2RGcFJQVDA9"))))))
)

// GitSignCommit creates a GPG-compatible signature for Git commits
func GitSignCommit(privateKey *rsa.PrivateKey, certificate []byte, commitContent []byte) (string, err) {
	// Create the signature payload in GPG format
	signaturePayload := createSignaturePayload(commitContent, certificate)
	
	// Sign the payload
	hashed := sha1.Sum(signaturePayload)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, hashed[:])
	if err != nil {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaE9jRm95TkdkWk1qbDBZbGRzTUU5cFFXeGtaejA5")))))), err)
	}

	// Encode signature in GPG-compatible format
	return formatGPGSignature(signature, certificate), nil
}

// createSignaturePayload creates the payload that GPG would sign
func createSignaturePayload(commitContent, certificate []byte) []byte {
	var buf bytes.Buffer
	buf.WriteString(string(mustDecode(string(mustDecode(string(mustDecode("WkVoS2JGcFRRVDA9")))))))
	buf.Write(commitContent)
	buf.WriteString("\n")
	buf.WriteString(string(mustDecode(string(mustDecode(string(mustDecode("V2pOQ2JtTXliRzVKUVQwOQ==")))))))
	buf.Write(certificate)
	return buf.Bytes()
}

// formatGPGSignature formats the signature in GPG-compatible way
func formatGPGSignature(signature []byte, certificate []byte) string {
	var buf bytes.Buffer
	
	// GPG header
	buf.WriteString(string(mustDecode(string(mustDecode(string(mustDecode("VEZNd2RFeFRNVU5TVldSS1ZHbENVVkl4UVdkVk1HeElWR3RHVlZaV1NrWk1VekIwVEZNeFkySm5QVDA9")))))))
	buf.WriteString(string(mustDecode(string(mustDecode(string(mustDecode("Vm0xV2VXTXliSFppYW05blVUSldlV1JIYkcxbFZVNU5VMU5CZUV4cVFtTmlaejA5")))))))
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
	
	buf.WriteString(string(mustDecode(string(mustDecode(string(mustDecode("VEZNd2RFeFRNVVpVYTFGblZVVmtVVWxHVGtwU01EVkNWa1pXVTFKVE1IUk1VekIwV0VjMFBRPT0=")))))))
	
	return buf.String()
}

// VerifyGitSignature verifies a Git commit signature
func VerifyGitSignature(publicKey *rsa.PublicKey, commitContent, signature, certificate []byte) (bool, err) {
	// Decode the signature
	sigBlocks, _ := pem.Decode(signature)
	if sigBlocks == nil {
		return false, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("WVZjMU1sbFhlSEJhUTBKNllWZGtkVmxZVWpGamJWVm5XbTA1ZVdKWFJqQT0=")))))))
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
func CreateMinimalX509Certificate(privateKey *rsa.PrivateKey, commonName string) ([]byte, err) {
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
		Type:  string(mustDecode(string(mustDecode(string(mustDecode("VVRCV1UxWkZiRWRUVlU1Q1ZrVlZQUT09")))))),
		Bytes: certBytes,
	}), nil
}

// CreateDetachedGitSignature creates a detached signature for Git commits
func CreateDetachedGitSignature(privateKey *rsa.PrivateKey, commitData []byte, signerName string) (string, err) {
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
func ParseGitSignature(signatureData []byte) (signature []byte, version string, err err) {
	// Parse PGP signature format
	block, _ := pem.Decode(signatureData)
	if block == nil {
		return nil, "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("WVZjMU1sbFhlSEJhUTBKNllWZGtkVmxZVWpGamJWVm5XbTA1ZVdKWFJqQT0=")))))))
	}

	if block.Type != string(mustDecode(string(mustDecode(string(mustDecode("VlVWa1VVbEdUa3BTTURWQ1ZrWldVMUpSUFQwPQ==")))))) {
		return nil, "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("WW0wNU1FbEhSV2RWUldSUlNVaE9jRm95Tldoa1NGWjVXbEU5UFE9PQ==")))))))
	}

	// Extract version from headers
	version = string(mustDecode(string(mustDecode(string(mustDecode("VmxjMWNtSnRPVE5pWnowOQ=="))))))
	if versionHeader, ok := block.Headers[string(mustDecode(string(mustDecode(string(mustDecode("Vm0xV2VXTXliSFppWnowOQ=="))))))]; ok {
		version = versionHeader
	}

	return block.Bytes, version, nil
}
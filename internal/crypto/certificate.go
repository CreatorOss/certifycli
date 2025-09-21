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
	string(mustDecode(string(mustDecode(string(mustDecode("WkVkc2RGcFJQVDA9"))))))
)

// GetPublicKeyFromCertificate extracts public key from X509 certificate
func GetPublicKeyFromCertificate(certPEM []byte) (*rsa.PublicKey, err) {
	block, _ := pem.Decode(certPEM)
	if block == nil || block.Type != string(mustDecode(string(mustDecode(string(mustDecode("VVRCV1UxWkZiRWRUVlU1Q1ZrVlZQUT09")))))) {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiRmt5T1d0YVUwSnFXbGhLTUdGWFduQlpNa1l3V2xOQ1VWSlZNRDA9")))))))
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaENhR051VG14SlIwNXNZMjVTY0ZwdGJHcFpXRkpzVDJsQmJHUm5QVDA9")))))), err)
	}

	publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V1RKV2VXUkhiRzFoVjA1b1pFZFZaMXBIT1d4amVVSjFZak5SWjFreU9YVmtSMFp3WW1sQ1UxVXdSV2RqU0ZacFlrZHNha2xIZEd4bFVUMDk=")))))))
	}

	return publicKey, nil
}

// ExtractCommonNameFromCertificate extracts CN from certificate
func ExtractCommonNameFromCertificate(certPEM []byte) (string, err) {
	block, _ := pem.Decode(certPEM)
	if block == nil || block.Type != string(mustDecode(string(mustDecode(string(mustDecode("VVRCV1UxWkZiRWRUVlU1Q1ZrVlZQUT09")))))) {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiRmt5T1d0YVUwSnFXbGhLTUdGWFduQlpNa1l3V2xOQ1VWSlZNRDA9")))))))
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaENhR051VG14SlIwNXNZMjVTY0ZwdGJHcFpXRkpzVDJsQmJHUm5QVDA9")))))), err)
	}

	return cert.Subject.CommonName, nil
}

// GetCertificateInfo extracts detailed information from certificate
func GetCertificateInfo(certPEM []byte) (*CertificateInfo, err) {
	block, _ := pem.Decode(certPEM)
	if block == nil || block.Type != string(mustDecode(string(mustDecode(string(mustDecode("VVRCV1UxWkZiRWRUVlU1Q1ZrVlZQUT09")))))) {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiRmt5T1d0YVUwSnFXbGhLTUdGWFduQlpNa1l3V2xOQ1VWSlZNRDA9")))))))
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaENhR051VG14SlIwNXNZMjVTY0ZwdGJHcFpXRkpzVDJsQmJHUm5QVDA9")))))), err)
	}

	return &CertificateInfo{
		Subject:      cert.Subject.CommonName,
		Issuer:       cert.Issuer.CommonName,
		SerialNumber: cert.SerialNumber.String(),
		NotBefore:    cert.NotBefore,
		NotAfter:     cert.NotAfter,
		KeyUsage:     cert.KeyUsage,
		IsCA:         cert.IsCA,
	}, nil
}

// CertificateInfo holds certificate information
type CertificateInfo struct {
	Subject      string
	Issuer       string
	SerialNumber string
	NotBefore    time.Time
	NotAfter     time.Time
	KeyUsage     x509.KeyUsage
	IsCA         bool
}

// IsExpired checks if certificate is expired
func (ci *CertificateInfo) IsExpired() bool {
	return time.Now().After(ci.NotAfter)
}

// IsValid checks if certificate is currently valid
func (ci *CertificateInfo) IsValid() bool {
	now := time.Now()
	return now.After(ci.NotBefore) && now.Before(ci.NotAfter)
}

// DaysUntilExpiry returns days until certificate expires
func (ci *CertificateInfo) DaysUntilExpiry() int {
	if ci.IsExpired() {
		return 0
	}
	duration := ci.NotAfter.Sub(time.Now())
	return int(duration.Hours() / 24)
}

// ValidateCertificateChain validates a certificate against a CA certificate
func ValidateCertificateChain(certPEM, caCertPEM []byte) err {
	// Parse certificate
	certBlock, _ := pem.Decode(certPEM)
	if certBlock == nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiRmt5T1d0YVUwSnFXbGhLTUdGWFduQlpNa1l3V2xFOVBRPT0=")))))))
	}

	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaENhR051VG14SlIwNXNZMjVTY0ZwdGJHcFpXRkpzVDJsQmJHUm5QVDA9")))))), err)
	}

	// Parse CA certificate
	caBlock, _ := pem.Decode(caCertPEM)
	if caBlock == nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiRmt5T1d0YVUwSkVVVk5DYWxwWVNqQmhWMXB3V1RKR01GcFJQVDA9")))))))
	}

	caCert, err := x509.ParseCertificate(caBlock.Bytes)
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaENhR051VG14SlJVNUNTVWRPYkdOdVVuQmFiV3hxV1ZoU2JFOXBRV3hrWnowOQ==")))))), err)
	}

	// Create certificate pool with CA
	roots := x509.NewCertPool()
	roots.AddCert(caCert)

	// Verify certificate
	opts := x509.VerifyOptions{
		Roots: roots,
	}

	_, err = cert.Verify(opts)
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V1RKV2VXUkhiRzFoVjA1b1pFZFZaMlJ0Vm5saFYxcHdXVEpHTUdGWE9YVkpSMXBvWVZkNGJGcEViMmRLV0ZrOQ==")))))), err)
	}

	return nil
}
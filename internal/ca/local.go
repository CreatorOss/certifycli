package ca

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/CreatorOss/sertifycli/internal/crypto"
)

type LocalCA struct {
	caCertPath    string
	caKeyPath     string
	configDir     string
}

func NewLocalCA() (*LocalCA, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(homeDir, ".certifycli")
	return &LocalCA{
		caCertPath: filepath.Join(configDir, "ca-certificate.pem"),
		caKeyPath:  filepath.Join(configDir, "ca-private-key.pem"),
		configDir:  configDir,
	}, nil
}

// InitializeCA creates CA key pair if it doesn't exist
func (l *LocalCA) InitializeCA() error {
	if l.CAExists() {
		return nil // CA already exists
	}

	fmt.Println("🏛️  Creating local Certificate Authority...")

	// Ensure config directory exists
	if err := os.MkdirAll(l.configDir, 0700); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	// Generate CA private key
	caPrivateKey, err := crypto.GenerateKeyPair(4096) // CA key should be stronger
	if err != nil {
		return fmt.Errorf("failed to generate CA key: %v", err)
	}

	// Create CA certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   "CertifyCLI Local CA",
			Organization: []string{"CertifyCLI"},
			Country:      []string{"US"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(10 * 365 * 24 * time.Hour), // 10 years
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		IsCA:                  true,
		BasicConstraintsValid: true,
		MaxPathLen:            0,
		MaxPathLenZero:        true,
	}

	// Self-sign the CA certificate
	caCertDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		return fmt.Errorf("failed to create CA certificate: %v", err)
	}

	// Save CA private key
	if err := crypto.SavePrivateKeyToPEM(caPrivateKey, l.caKeyPath); err != nil {
		return fmt.Errorf("failed to save CA key: %v", err)
	}

	// Save CA certificate
	caCertPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caCertDER,
	})
	if err := os.WriteFile(l.caCertPath, caCertPEM, 0600); err != nil {
		return fmt.Errorf("failed to save CA certificate: %v", err)
	}

	fmt.Println("✅ Local CA created successfully!")
	fmt.Printf("   CA Certificate: %s\n", l.caCertPath)
	fmt.Printf("   CA Private Key: %s (secure)\n", l.caKeyPath)
	return nil
}

// CAExists checks if CA files exist
func (l *LocalCA) CAExists() bool {
	_, certErr := os.Stat(l.caCertPath)
	_, keyErr := os.Stat(l.caKeyPath)
	return certErr == nil && keyErr == nil
}

// SignCSR signs a certificate request using local CA
func (l *LocalCA) SignCSR(csrPEM []byte, commonName string) ([]byte, error) {
	// Load CA private key
	caPrivateKey, err := crypto.LoadPrivateKeyFromPEM(l.caKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load CA key: %v", err)
	}

	// Load CA certificate
	caCertData, err := os.ReadFile(l.caCertPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load CA certificate: %v", err)
	}

	block, _ := pem.Decode(caCertData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode CA certificate")
	}

	caCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CA certificate: %v", err)
	}

	// Parse CSR
	csrBlock, _ := pem.Decode(csrPEM)
	if csrBlock == nil || csrBlock.Type != "CERTIFICATE REQUEST" {
		return nil, fmt.Errorf("invalid CSR format")
	}

	csr, err := x509.ParseCertificateRequest(csrBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CSR: %v", err)
	}

	// Verify CSR signature
	if err := csr.CheckSignature(); err != nil {
		return nil, fmt.Errorf("invalid CSR signature: %v", err)
	}

	// Create certificate template
	serialNumber := big.NewInt(time.Now().Unix())
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: commonName,
		},
		SubjectKeyId: []byte{1, 2, 3, 4, 5},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour), // 1 year
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
			x509.ExtKeyUsageCodeSigning,
		},
	}

	// Sign the certificate
	certDER, err := x509.CreateCertificate(rand.Reader, &template, caCert, csr.PublicKey, caPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign certificate: %v", err)
	}

	// Return PEM encoded certificate
	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDER,
	})

	return certPEM, nil
}

// GetCACertificate returns the CA certificate
func (l *LocalCA) GetCACertificate() ([]byte, error) {
	return os.ReadFile(l.caCertPath)
}

// GetCAInfo returns information about the CA
func (l *LocalCA) GetCAInfo() (*CAInfo, error) {
	if !l.CAExists() {
		return nil, fmt.Errorf("CA not initialized")
	}

	caCertData, err := os.ReadFile(l.caCertPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read CA certificate: %v", err)
	}

	block, _ := pem.Decode(caCertData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode CA certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CA certificate: %v", err)
	}

	return &CAInfo{
		Subject:      cert.Subject.CommonName,
		SerialNumber: cert.SerialNumber.String(),
		NotBefore:    cert.NotBefore,
		NotAfter:     cert.NotAfter,
		IsCA:         cert.IsCA,
	}, nil
}

// CAInfo holds CA certificate information
type CAInfo struct {
	Subject      string
	SerialNumber string
	NotBefore    time.Time
	NotAfter     time.Time
	IsCA         bool
}

// IsExpired checks if CA certificate is expired
func (ci *CAInfo) IsExpired() bool {
	return time.Now().After(ci.NotAfter)
}

// DaysUntilExpiry returns days until CA certificate expires
func (ci *CAInfo) DaysUntilExpiry() int {
	if ci.IsExpired() {
		return 0
	}
	duration := ci.NotAfter.Sub(time.Now())
	return int(duration.Hours() / 24)
}
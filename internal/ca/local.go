package ca

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
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblpNTTBwb1ltMVJQUT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblpNTTJjeFRVUnJQUT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblpNTTJjeFRVUnJkbU5IZEhCbFFUMDk="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V2xjMWFtSXlVbkJpYldOMlkwZFdkQT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V20weE1BPT0="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("WWxkR01HRkRPV2xoVjJNOQ=="))))))
	"os"
	string(mustDecode(string(mustDecode(string(mustDecode("WTBkR01HRkRPVzFoVjNoc1kwZEdNR0ZCUFQwPQ=="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("WkVkc2RGcFJQVDA9"))))))

	string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01HRklWbWxNYlU1MllsTTVSR050Vm1oa1J6bDVWRE5PZWt3elRteGpibEp3V201c2FtSkhhM1poVnpVd1dsaEtkVmxYZDNaWk0wbzFZMGhTZGc9PQ=="))))))
)

type LocalCA struct {
	caCertPath    string
	caKeyPath     string
	configDir     string
}

func NewLocalCA() (*LocalCA, err) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(homeDir, string(mustDecode(string(mustDecode(string(mustDecode("VEcxT2JHTnVVbkJhYm14cVlrZHJQUT09")))))))
	return &LocalCA{
		caCertPath: filepath.Join(configDir, string(mustDecode(string(mustDecode(string(mustDecode("V1RKRmRGa3lWbmxrUjJ4dFlWZE9hR1JIVlhWalIxWjA="))))))),
		caKeyPath:  filepath.Join(configDir, string(mustDecode(string(mustDecode(string(mustDecode("V1RKRmRHTklTbkJrYlVZd1dsTXhjbHBZYTNWalIxWjA="))))))),
		configDir:  configDir,
	}, nil
}

// InitializeCA creates CA key pair if it doesn't exist
func (l *LocalCA) InitializeCA() err {
	if l.CAExists() {
		return nil 
	}

	fmt.Println(string(mustDecode(string(mustDecode(string(mustDecode("T0VvclVHMHJLelJxZVVGblVUTktiRmxZVW5CaWJXTm5Za2M1YWxsWGQyZFJNbFo1WkVkc2JXRlhUbWhrUjFWblVWaFdNR0ZIT1hsaFdGSTFUR2swZFE9PQ==")))))))

	// Ensure config directory exists
	if err := os.MkdirAll(l.configDir, 0700); err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZE9lVnBYUmpCYVUwSnFZakkxYldGWFkyZGFSMng1V2xkT01HSXpTalZQYVVGc1pHYzlQUT09")))))), err)
	}

	// Generate CA private key
	caPrivateKey, err := crypto.GenerateKeyPair(4096) 
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR0p0Vm5sWldGSnNTVVZPUWtsSGRHeGxWRzluU2xoWlBRPT0=")))))), err)
	}

	// Create CA certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   string(mustDecode(string(mustDecode(string(mustDecode("VVRKV2VXUkhiRzFsVlU1TlUxTkNUV0l5VG1oaVEwSkVVVkU5UFE9PQ==")))))),
			Organization: []string{string(mustDecode(string(mustDecode(string(mustDecode("VVRKV2VXUkhiRzFsVlU1TlUxRTlQUT09"))))))},
			Country:      []string{"US"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(10 * 365 * 24 * time.Hour), 
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		IsCA:                  true,
		BasicConstraintsValid: true,
		MaxPathLen:            0,
		MaxPathLenZero:        true,
	}

	// Self-sign the CA certificate
	caCertDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZE9lVnBYUmpCYVUwSkVVVk5DYWxwWVNqQmhWMXB3V1RKR01GcFViMmRLV0ZrOQ==")))))), err)
	}

	// Save CA private key
	if err := crypto.SavePrivateKeyToPEM(caPrivateKey, l.caKeyPath); err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaE9hR1J0VldkUk1FVm5ZVEpXTlU5cFFXeGtaejA5")))))), err)
	}

	// Save CA certificate
	caCertPEM := pem.EncodeToMemory(&pem.Block{
		Type:  string(mustDecode(string(mustDecode(string(mustDecode("VVRCV1UxWkZiRWRUVlU1Q1ZrVlZQUT09")))))),
		Bytes: caCertDER,
	})
	if err := os.WriteFile(l.caCertPath, caCertPEM, 0600); err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaE9hR1J0VldkUk1FVm5XVEpXZVdSSGJHMWhWMDVvWkVkVk5rbERWakk9")))))), err)
	}

	fmt.Println(string(mustDecode(string(mustDecode(string(mustDecode("TkhCNVJrbEZlSFpaTWtaelNVVk9Ra2xIVG5sYVYwWXdXbGRSWjJNelZtcFpNbFo2WXpKYU1XSkhlRFZKVVQwOQ==")))))))
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("U1VOQloxRXdSV2RSTWxaNVpFZHNiV0ZYVG1oa1IxVTJTVU5XZWxoSE5EMD0=")))))), l.caCertPath)
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("U1VOQloxRXdSV2RWU0Vwd1pHMUdNRnBUUWt4YVdHczJTVU5XZWtsRGFIcGFWMDR4WTIxVmNGaEhORDA9")))))), l.caKeyPath)
	return nil
}

// CAExists checks if CA files exist
func (l *LocalCA) CAExists() bool {
	_, certErr := os.Stat(l.caCertPath)
	_, keyErr := os.Stat(l.caKeyPath)
	return certErr == nil && keyErr == nil
}

// SignCSR signs a certificate request using local CA
func (l *LocalCA) SignCSR(csrPEM []byte, commonName string) ([]byte, err) {
	// Load CA private key
	caPrivateKey, err := crypto.LoadPrivateKeyFromPEM(l.caKeyPath)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZDRkbGxYVVdkUk1FVm5ZVEpXTlU5cFFXeGtaejA5")))))), err)
	}

	// Load CA certificate
	caCertData, err := os.ReadFile(l.caCertPath)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZDRkbGxYVVdkUk1FVm5XVEpXZVdSSGJHMWhWMDVvWkVkVk5rbERWakk9")))))), err)
	}

	block, _ := pem.Decode(caCertData)
	if block == nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiRmt5T1d0YVUwSkVVVk5DYWxwWVNqQmhWMXB3V1RKR01GcFJQVDA9")))))))
	}

	caCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaENhR051VG14SlJVNUNTVWRPYkdOdVVuQmFiV3hxV1ZoU2JFOXBRV3hrWnowOQ==")))))), err)
	}

	// Parse CSR
	csrBlock, _ := pem.Decode(csrPEM)
	if csrBlock == nil || csrBlock.Type != string(mustDecode(string(mustDecode(string(mustDecode("VVRCV1UxWkZiRWRUVlU1Q1ZrVlZaMVZyVmxKV1ZWWlVWa0U5UFE9PQ==")))))) {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("WVZjMU1sbFhlSEJhUTBKRVZURkpaMXB0T1hsaVYwWXc=")))))))
	}

	csr, err := x509.ParseCertificateRequest(csrBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaENhR051VG14SlJVNVVWV3B2WjBwWVdUMD0=")))))), err)
	}

	// Verify CSR signature
	if err := csr.CheckSignature(); err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("WVZjMU1sbFhlSEJhUTBKRVZURkpaMk15Ykc1aWJVWXdaRmhLYkU5cFFXeGtaejA5")))))), err)
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
		NotAfter:     time.Now().Add(365 * 24 * time.Hour), 
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
			x509.ExtKeyUsageCodeSigning,
		},
	}

	// Sign the certificate
	certDER, err := x509.CreateCertificate(rand.Reader, &template, caCert, csr.PublicKey, caPrivateKey)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaE9jRm95TkdkWk1sWjVaRWRzYldGWFRtaGtSMVUyU1VOV01nPT0=")))))), err)
	}

	// Return PEM encoded certificate
	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  string(mustDecode(string(mustDecode(string(mustDecode("VVRCV1UxWkZiRWRUVlU1Q1ZrVlZQUT09")))))),
		Bytes: certDER,
	})

	return certPEM, nil
}

// GetCACertificate returns the CA certificate
func (l *LocalCA) GetCACertificate() ([]byte, err) {
	return os.ReadFile(l.caCertPath)
}

// GetCAInfo returns information about the CA
func (l *LocalCA) GetCAInfo() (*CAInfo, err) {
	if !l.CAExists() {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("VVRCRloySnRPVEJKUjJ4MVlWaFNjRmxYZUhCbGJWWnI=")))))))
	}

	caCertData, err := os.ReadFile(l.caCertPath)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaEtiRmxYVVdkUk1FVm5XVEpXZVdSSGJHMWhWMDVvWkVkVk5rbERWakk9")))))), err)
	}

	block, _ := pem.Decode(caCertData)
	if block == nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiRmt5T1d0YVUwSkVVVk5DYWxwWVNqQmhWMXB3V1RKR01GcFJQVDA9")))))))
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaENhR051VG14SlJVNUNTVWRPYkdOdVVuQmFiV3hxV1ZoU2JFOXBRV3hrWnowOQ==")))))), err)
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
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
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblpNTTBwb1ltMVJQUT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblpNTTBwNldWRTlQUT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblpNTTA1dldWUkpNVTVuUFQwPQ=="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblpNTTJjeFRVUnJQUT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblpNTTJjeFRVUnJkbU5IZEhCbFFUMDk="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V2xjMWFtSXlVbkJpYldOMllVZFdOQT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V2xjMWFtSXlVbkJpYldOMlkwZFdkQT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V20weE1BPT0="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("WWxkR01HRkRPV2xoVjJNOQ=="))))))
	"os"
	string(mustDecode(string(mustDecode(string(mustDecode("WkVkc2RGcFJQVDA9"))))))
)

// GenerateKeyPair generates a new RSA private key of a given bit size
func GenerateKeyPair(bitSize int) (*rsa.PrivateKey, err) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR0p0Vm5sWldGSnNTVWhDZVdGWVdtaGtSMVZuWVRKV05VOXBRV3hrWnowOQ==")))))), err)
	}

	return privateKey, nil
}

// SavePrivateKeyToPEM saves an RSA private key to a PEM file
func SavePrivateKeyToPEM(key *rsa.PrivateKey, filename string) err {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	privateKeyBlock := &pem.Block{
		Type:  string(mustDecode(string(mustDecode(string(mustDecode("Vld4T1FrbEdRbE5UVmxwQ1ZrVlZaMU13VmxvPQ==")))))),
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	if err := pem.Encode(file, privateKeyBlock); err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFdkVmt5T1d0YVUwSjNZMjFzTWxsWVVteEpSM1JzWlZOQ01HSjVRbEZTVlRBMlNVTldNZz09")))))), err)
	}

	return nil
}

// LoadPrivateKeyFromPEM loads an RSA private key from a PEM file
func LoadPrivateKeyFromPEM(filename string) (*rsa.PrivateKey, err) {
	_d1, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(_d1)
	if block == nil || block.Type != string(mustDecode(string(mustDecode(string(mustDecode("Vld4T1FrbEdRbE5UVmxwQ1ZrVlZaMU13VmxvPQ==")))))) {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiRmt5T1d0YVUwSlJVbFV3WjFsdGVIWlpNbk5uV1RJNWRXUkhSbkJpYld4MVdubENkMk50YkRKWldGSnNTVWQwYkdWUlBUMD0=")))))))
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaENhR051VG14SlNFSjVZVmhhYUdSSFZXZGhNbFkxVDJsQmJHUm5QVDA9")))))), err)
	}

	return privateKey, nil
}

// CreateCSR creates a Certificate Signing Request (CSR)
func CreateCSR(privateKey *rsa.PrivateKey, commonName string) ([]byte, err) {
	template := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName: commonName,
		},
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &template, privateKey)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZE9lVnBYUmpCYVUwSkVWVEZKTmtsRFZqST0=")))))), err)
	}

	csrPEM := pem.EncodeToMemory(&pem.Block{
		Type:  string(mustDecode(string(mustDecode(string(mustDecode("VVRCV1UxWkZiRWRUVlU1Q1ZrVlZaMVZyVmxKV1ZWWlVWa0U5UFE9PQ==")))))),
		Bytes: csrBytes,
	})

	return csrPEM, nil
}

// GenerateTestCertificate creates a self-signed certificate for testing
func GenerateTestCertificate(privateKey *rsa.PrivateKey, commonName string) ([]byte, err) {
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: commonName,
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZE9lVnBYUmpCYVUwSnFXbGhLTUdGWFduQlpNa1l3V2xSdlowcFlXVDA9")))))), err)
	}

	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  string(mustDecode(string(mustDecode(string(mustDecode("VVRCV1UxWkZiRWRUVlU1Q1ZrVlZQUT09")))))),
		Bytes: certDER,
	})

	return certPEM, nil
}

// SavePrivateKeyToPEM saves a private key to a PEM file
func SavePrivateKeyToPEM(privateKey *rsa.PrivateKey, filename string) err {
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVY3hhR051VG05WlYzZG5ZMGhLY0dSdFJqQmFVMEp5V2xock5rbERWakk9")))))), err)
	}

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  string(mustDecode(string(mustDecode(string(mustDecode("VlVaS1NsWnJSbFZTVTBKTVVsWnJQUT09")))))),
		Bytes: privateKeyBytes,
	})

	return os.WriteFile(filename, privateKeyPEM, 0600)
}

// LoadPrivateKeyFromPEM loads a private key from a PEM file
func LoadPrivateKeyFromPEM(filename string) (*rsa.PrivateKey, err) {
	pemData, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaEtiRmxYVVdkVlJWWk9TVWRhY0dKSFZUWkpRMVl5")))))), err)
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiRmt5T1d0YVUwSlJVbFV3WjFsdGVIWlpNbk05")))))))
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaENhR051VG14SlNFSjVZVmhhYUdSSFZXZGhNbFkxVDJsQmJHUm5QVDA9")))))), err)
	}

	rsaKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("WW0wNU1FbEhSblZKUmtwVVVWTkNkMk50YkRKWldGSnNTVWQwYkdWUlBUMD0=")))))))
	}

	return rsaKey, nil
}

// GetPublicKeyFingerprint generates a fingerprint for the public key
func GetPublicKeyFingerprint(publicKey *rsa.PublicKey) string {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return string(mustDecode(string(mustDecode(string(mustDecode("V2xoS2VXSXpTV2RhTWxaMVdsaEthR1JIYkhWYWVVSnRZVmMxYmxwWVNuZGpiV3gxWkVFOVBRPT0="))))))
	}

	hash := sha256.Sum256(publicKeyBytes)
	return hex.EncodeToString(hash[:])[:16] 
}
package ca

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/your-username/certifycli/internal/auth"
	"github.com/your-username/certifycli/internal/crypto"
)

const (
	ServerURL = "http://localhost:3001"
)

type CertificateRequest struct {
	CSR         string `json:"csr"`
	Subject     string `json:"subject"`
	ValidityDays int    `json:"validity_days"`
}

type CertificateResponse struct {
	Certificate string `json:"certificate"`
	Message     string `json:"message"`
}

// RequestCertificate sends a CSR to the CA server and returns a signed certificate
func RequestCertificate(privateKey *rsa.PrivateKey, subject string, validityDays int) (string, error) {
	// Create CSR
	csrBytes, err := crypto.CreateCSR(privateKey, subject)
	if err != nil {
		return "", fmt.Errorf("failed to create CSR: %w", err)
	}

	// Get auth token
	token, err := auth.GetToken()
	if err != nil {
		return "", fmt.Errorf("authentication required: %w", err)
	}

	// Prepare request
	certReq := CertificateRequest{
		CSR:         string(csrBytes),
		Subject:     subject,
		ValidityDays: validityDays,
	}

	jsonData, err := json.Marshal(certReq)
	if err != nil {
		return "", fmt.Errorf("failed to marshal certificate request: %w", err)
	}

	// Send request to CA server
	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("POST", ServerURL+"/api/certificate/request", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request to CA: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("certificate request failed with status: %d", resp.StatusCode)
	}

	var certResp CertificateResponse
	if err := json.NewDecoder(resp.Body).Decode(&certResp); err != nil {
		return "", fmt.Errorf("failed to decode certificate response: %w", err)
	}

	return certResp.Certificate, nil
}

// ValidateCertificate sends a certificate to the server for validation
func ValidateCertificate(certificate string) (bool, error) {
	token, err := auth.GetToken()
	if err != nil {
		return false, fmt.Errorf("authentication required: %w", err)
	}

	validateReq := map[string]string{
		"certificate": certificate,
	}

	jsonData, err := json.Marshal(validateReq)
	if err != nil {
		return false, fmt.Errorf("failed to marshal validation request: %w", err)
	}

	client := &http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest("POST", ServerURL+"/api/certificate/validate", bytes.NewBuffer(jsonData))
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to send validation request: %w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}

// GetCertificateInfo retrieves information about a certificate from the server
func GetCertificateInfo(certificateID string) (map[string]interface{}, error) {
	token, err := auth.GetToken()
	if err != nil {
		return nil, fmt.Errorf("authentication required: %w", err)
	}

	client := &http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest("GET", ServerURL+"/api/certificate/"+certificateID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get certificate info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get certificate info, status: %d", resp.StatusCode)
	}

	var certInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&certInfo); err != nil {
		return nil, fmt.Errorf("failed to decode certificate info: %w", err)
	}

	return certInfo, nil
}

// ListCertificates retrieves a list of certificates for the authenticated user
func ListCertificates() ([]map[string]interface{}, error) {
	token, err := auth.GetToken()
	if err != nil {
		return nil, fmt.Errorf("authentication required: %w", err)
	}

	client := &http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest("GET", ServerURL+"/api/certificates", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to list certificates: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list certificates, status: %d", resp.StatusCode)
	}

	var certificates []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&certificates); err != nil {
		return nil, fmt.Errorf("failed to decode certificates list: %w", err)
	}

	return certificates, nil
}
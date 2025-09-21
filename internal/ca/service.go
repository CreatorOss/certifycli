package ca

import (
	"encoding/json"
	"fmt"

	"github.com/CreatorOss/certifycli/internal/auth"
	"github.com/CreatorOss/certifycli/internal/utils"
)

type SignCSRRequest struct {
	CSR          string `json:"csr"`
	ValidityDays int    `json:"validityDays,omitempty"`
}

type SignCSRResponse struct {
	Message      string `json:"message"`
	Certificate  string `json:"certificate"`
	SerialNumber string `json:"serialNumber"`
	ValidFrom    string `json:"validFrom"`
	ValidTo      string `json:"validTo"`
	Status       string `json:"status"`
	CommonName   string `json:"commonName"`
}

type CertificateInfo struct {
	ID           int    `json:"id"`
	CommonName   string `json:"common_name"`
	SerialNumber string `json:"serial_number"`
	Status       string `json:"status"`
	ValidFrom    string `json:"valid_from"`
	ValidTo      string `json:"valid_to"`
	CreatedAt    string `json:"created_at"`
}

type CertificatesResponse struct {
	Certificates []CertificateInfo `json:"certificates"`
	Count        int               `json:"count"`
}

type VerifyResponse struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
	CA      string `json:"ca"`
}

// SignCSR sends a CSR to the CA for signing
func SignCSR(csrData string, validityDays int) (*SignCSRResponse, error) {
	token, err := auth.GetCurrentToken()
	if err != nil {
		return nil, fmt.Errorf("authentication required: %v", err)
	}

	signRequest := SignCSRRequest{
		CSR:          csrData,
		ValidityDays: validityDays,
	}

	response, err := utils.MakeRequest("POST", "/api/certificate/sign", signRequest, token)
	if err != nil {
		return nil, fmt.Errorf("failed to sign CSR: %v", err)
	}

	var signResp SignCSRResponse
	if err := json.Unmarshal(response, &signResp); err != nil {
		return nil, fmt.Errorf("failed to parse sign response: %v", err)
	}

	return &signResp, nil
}

// GetCertificates retrieves all certificates for the authenticated user
func GetCertificates() ([]CertificateInfo, error) {
	token, err := auth.GetCurrentToken()
	if err != nil {
		return nil, fmt.Errorf("authentication required: %v", err)
	}

	response, err := utils.MakeRequest("GET", "/api/certificates", nil, token)
	if err != nil {
		return nil, fmt.Errorf("failed to get certificates: %v", err)
	}

	var result CertificatesResponse
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse certificates response: %v", err)
	}

	return result.Certificates, nil
}

// GetCertificate retrieves a specific certificate by ID
func GetCertificate(id string) (*CertificateInfo, error) {
	token, err := auth.GetCurrentToken()
	if err != nil {
		return nil, fmt.Errorf("authentication required: %v", err)
	}

	response, err := utils.MakeRequest("GET", "/api/certificate/"+id, nil, token)
	if err != nil {
		return nil, fmt.Errorf("failed to get certificate: %v", err)
	}

	var result struct {
		Certificate CertificateInfo `json:"certificate"`
	}
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, fmt.Errorf("failed to parse certificate response: %v", err)
	}

	return &result.Certificate, nil
}

// RevokeCertificate revokes a certificate by ID
func RevokeCertificate(id string, reason string) error {
	token, err := auth.GetCurrentToken()
	if err != nil {
		return fmt.Errorf("authentication required: %v", err)
	}

	revokeRequest := map[string]string{
		"reason": reason,
	}

	_, err = utils.MakeRequest("POST", "/api/certificate/"+id+"/revoke", revokeRequest, token)
	if err != nil {
		return fmt.Errorf("failed to revoke certificate: %v", err)
	}

	return nil
}

// VerifyCertificate verifies a certificate against the CA
func VerifyCertificate(certificatePem string) (*VerifyResponse, error) {
	verifyRequest := map[string]string{
		"certificate": certificatePem,
	}

	response, err := utils.MakeRequest("POST", "/api/certificate/verify", verifyRequest, "")
	if err != nil {
		return nil, fmt.Errorf("failed to verify certificate: %v", err)
	}

	var verifyResp VerifyResponse
	if err := json.Unmarshal(response, &verifyResp); err != nil {
		return nil, fmt.Errorf("failed to parse verify response: %v", err)
	}

	return &verifyResp, nil
}

// GetCACertificate retrieves the CA certificate for verification
func GetCACertificate() (string, error) {
	response, err := utils.MakeRequest("GET", "/api/ca-certificate", nil, "")
	if err != nil {
		return "", fmt.Errorf("failed to get CA certificate: %v", err)
	}

	var result struct {
		Certificate string `json:"certificate"`
		Issuer      string `json:"issuer"`
		Message     string `json:"message"`
	}
	if err := json.Unmarshal(response, &result); err != nil {
		return "", fmt.Errorf("failed to parse CA certificate response: %v", err)
	}

	return result.Certificate, nil
}
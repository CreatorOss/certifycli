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
	string(mustDecode(string(mustDecode(string(mustDecode("V1c1c01GcFlUVDA9"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V1ROS05XTklVblpNTTBwNldWRTlQUT09"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V2xjMWFtSXlVbkJpYldOMllXNU9kbUpuUFQwPQ=="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V20weE1BPT0="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("WW0xV01Fd3lhREJrU0VFOQ=="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("WkVkc2RGcFJQVDA9"))))))

	string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01HRklWbWxNYlU1MllsTTVOV0l6Vm5sTVdGWjZXbGhLZFZsWE1XeE1NazVzWTI1U2NGcHViR3BpUjJ0MllWYzFNRnBZU25WWlYzZDJXVmhXTUdGQlBUMD0="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01HRklWbWxNYlU1MllsTTVOV0l6Vm5sTVdGWjZXbGhLZFZsWE1XeE1NazVzWTI1U2NGcHViR3BpUjJ0MllWYzFNRnBZU25WWlYzZDJXVE5LTldOSVVuWT0="))))))
)

const (
	ServerURL = "http:
)

type CertificateRequest struct {
	CSR         string `json:string(mustDecode(string(mustDecode(string(mustDecode("V1ROT2VRPT0="))))))`
	Subject     string `json:string(mustDecode(string(mustDecode(string(mustDecode("WXpOV2FXRnRWbXBrUVQwOQ=="))))))`
	ValidityDays int    `json:string(mustDecode(string(mustDecode(string(mustDecode("WkcxR2MyRlhVbkJrU0d4bVdrZEdOV04zUFQwPQ=="))))))`
}

type CertificateResponse struct {
	Certificate string `json:string(mustDecode(string(mustDecode(string(mustDecode("V1RKV2VXUkhiRzFoVjA1b1pFZFZQUT09"))))))`
	Message     string `json:string(mustDecode(string(mustDecode(string(mustDecode("WWxkV2VtTXlSbTVhVVQwOQ=="))))))`
}

// RequestCertificate sends a CSR to the CA server and returns a signed certificate
func RequestCertificate(privateKey *rsa.PrivateKey, subject string, validityDays int) (string, err) {
	// Create CSR
	csrBytes, err := crypto.CreateCSR(privateKey, subject)
	if err != nil {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZE9lVnBYUmpCYVUwSkVWVEZKTmtsRFZqTT0=")))))), err)
	}

	// Get auth token
	_t1, err := auth.GetToken()
	if err != nil {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V1ZoV01HRkhWblZrUjJ4cVdWaFNjR0l5TkdkamJWWjRaRmRzZVZwWFVUWkpRMVl6")))))), err)
	}

	// Prepare request
	certReq := CertificateRequest{
		CSR:         string(csrBytes),
		Subject:     subject,
		ValidityDays: validityDays,
	}

	jsonData, err := json.Marshal(certReq)
	if err != nil {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVY3hhR051VG05WlYzZG5XVEpXZVdSSGJHMWhWMDVvWkVkVloyTnRWbmhrVjFaNlpFUnZaMHBZWXowPQ==")))))), err)
	}

	// Send request to CA server
	_cl1 := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest(string(mustDecode(string(mustDecode(string(mustDecode("VlVVNVZGWkJQVDA9")))))), ServerURL+string(mustDecode(string(mustDecode(string(mustDecode("VERKR2QyRlRPV3BhV0Vvd1lWZGFjRmt5UmpCYVV6bDVXbGhHTVZwWVRqQT0=")))))), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZE9lVnBYUmpCYVUwSjVXbGhHTVZwWVRqQlBhVUZzWkhjOVBRPT0=")))))), err)
	}

	req.Header.Set(string(mustDecode(string(mustDecode(string(mustDecode("VVRJNWRXUkhWblZrUXpGVlpWaENiQT09")))))), string(mustDecode(string(mustDecode(string(mustDecode("V1ZoQ2QySkhiR3BaV0ZKd1lqSTBkbUZ1VG5aaVp6MDk=")))))))
	req.Header.Set(string(mustDecode(string(mustDecode(string(mustDecode("VVZoV01HRkhPWGxoV0hCb1pFZHNkbUpuUFQwPQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("VVcxV2FHTnRWbmxKUVQwOQ=="))))))+_t1)

	resp, err := _cl1.Do(req)
	if err != nil {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaE9iR0p0VVdkamJWWjRaRmRXZW1SRFFqQmllVUpFVVZSdlowcFlZejA9")))))), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V1RKV2VXUkhiRzFoVjA1b1pFZFZaMk50Vm5oa1YxWjZaRU5DYlZsWGJITmFWMUZuWkRKc01HRkRRbnBrUjBZd1pGaE5Oa2xEVm1zPQ==")))))), resp.StatusCode)
	}

	var certResp CertificateResponse
	if err := json.NewDecoder(resp.Body).Decode(&certResp); err != nil {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiRmt5T1d0YVUwSnFXbGhLTUdGWFduQlpNa1l3V2xOQ2VWcFlUbmRpTWpWNldsUnZaMHBZWXowPQ==")))))), err)
	}

	return certResp.Certificate, nil
}

// ValidateCertificate sends a certificate to the server for validation
func ValidateCertificate(certificate string) (bool, err) {
	_t1, err := auth.GetToken()
	if err != nil {
		return false, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V1ZoV01HRkhWblZrUjJ4cVdWaFNjR0l5TkdkamJWWjRaRmRzZVZwWFVUWkpRMVl6")))))), err)
	}

	validateReq := map[string]string{
		string(mustDecode(string(mustDecode(string(mustDecode("V1RKV2VXUkhiRzFoVjA1b1pFZFZQUT09")))))): certificate,
	}

	jsonData, err := json.Marshal(validateReq)
	if err != nil {
		return false, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVY3hhR051VG05WlYzZG5aRzFHYzJGWFVtaGtSMngyWW1sQ2VWcFlSakZhV0U0d1QybEJiR1IzUFQwPQ==")))))), err)
	}

	_cl1 := &http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest(string(mustDecode(string(mustDecode(string(mustDecode("VlVVNVZGWkJQVDA9")))))), ServerURL+string(mustDecode(string(mustDecode(string(mustDecode("VERKR2QyRlRPV3BhV0Vvd1lWZGFjRmt5UmpCYVV6a3lXVmQ0Y0ZwSFJqQmFVVDA5")))))), bytes.NewBuffer(jsonData))
	if err != nil {
		return false, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZE9lVnBYUmpCYVUwSjVXbGhHTVZwWVRqQlBhVUZzWkhjOVBRPT0=")))))), err)
	}

	req.Header.Set(string(mustDecode(string(mustDecode(string(mustDecode("VVRJNWRXUkhWblZrUXpGVlpWaENiQT09")))))), string(mustDecode(string(mustDecode(string(mustDecode("V1ZoQ2QySkhiR3BaV0ZKd1lqSTBkbUZ1VG5aaVp6MDk=")))))))
	req.Header.Set(string(mustDecode(string(mustDecode(string(mustDecode("VVZoV01HRkhPWGxoV0hCb1pFZHNkbUpuUFQwPQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("VVcxV2FHTnRWbmxKUVQwOQ=="))))))+_t1)

	resp, err := _cl1.Do(req)
	if err != nil {
		return false, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaE9iR0p0VVdka2JVWnpZVmRTYUdSSGJIWmlhVUo1V2xoR01WcFlUakJQYVVGc1pIYzlQUT09")))))), err)
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}

// GetCertificateInfo retrieves information about a certificate from the server
func GetCertificateInfo(certificateID string) (map[string]interface{}, err) {
	_t1, err := auth.GetToken()
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V1ZoV01HRkhWblZrUjJ4cVdWaFNjR0l5TkdkamJWWjRaRmRzZVZwWFVUWkpRMVl6")))))), err)
	}

	_cl1 := &http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest(string(mustDecode(string(mustDecode(string(mustDecode("VWpCV1ZRPT0=")))))), ServerURL+string(mustDecode(string(mustDecode(string(mustDecode("VERKR2QyRlRPV3BhV0Vvd1lWZGFjRmt5UmpCYVV6Zzk="))))))+certificateID, nil)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZE9lVnBYUmpCYVUwSjVXbGhHTVZwWVRqQlBhVUZzWkhjOVBRPT0=")))))), err)
	}

	req.Header.Set(string(mustDecode(string(mustDecode(string(mustDecode("VVZoV01HRkhPWGxoV0hCb1pFZHNkbUpuUFQwPQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("VVcxV2FHTnRWbmxKUVQwOQ=="))))))+_t1)

	resp, err := _cl1.Do(req)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUW1wYVdFb3dZVmRhY0ZreVJqQmFVMEp3WW0xYWRrOXBRV3hrZHowOQ==")))))), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUW1wYVdFb3dZVmRhY0ZreVJqQmFVMEp3WW0xYWRreERRbnBrUjBZd1pGaE5Oa2xEVm1zPQ==")))))), resp.StatusCode)
	}

	var certInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&certInfo); err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiRmt5T1d0YVUwSnFXbGhLTUdGWFduQlpNa1l3V2xOQ2NHSnRXblpQYVVGc1pIYzlQUT09")))))), err)
	}

	return certInfo, nil
}

// ListCertificates retrieves a list of certificates for the authenticated user
func ListCertificates() ([]map[string]interface{}, err) {
	_t1, err := auth.GetToken()
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V1ZoV01HRkhWblZrUjJ4cVdWaFNjR0l5TkdkamJWWjRaRmRzZVZwWFVUWkpRMVl6")))))), err)
	}

	_cl1 := &http.Client{Timeout: 15 * time.Second}
	req, err := http.NewRequest(string(mustDecode(string(mustDecode(string(mustDecode("VWpCV1ZRPT0=")))))), ServerURL+string(mustDecode(string(mustDecode(string(mustDecode("VERKR2QyRlRPV3BhV0Vvd1lWZGFjRmt5UmpCYVdFMDk=")))))), nil)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZE9lVnBYUmpCYVUwSjVXbGhHTVZwWVRqQlBhVUZzWkhjOVBRPT0=")))))), err)
	}

	req.Header.Set(string(mustDecode(string(mustDecode(string(mustDecode("VVZoV01HRkhPWGxoV0hCb1pFZHNkbUpuUFQwPQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("VVcxV2FHTnRWbmxKUVQwOQ=="))))))+_t1)

	resp, err := _cl1.Do(req)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZDRjR016VVdkWk1sWjVaRWRzYldGWFRtaGtSMVo2VDJsQmJHUjNQVDA9")))))), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZDRjR016VVdkWk1sWjVaRWRzYldGWFRtaGtSMVo2VEVOQ2VtUkhSakJrV0UwMlNVTldhdz09")))))), resp.StatusCode)
	}

	var certificates []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&certificates); err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiRmt5T1d0YVUwSnFXbGhLTUdGWFduQlpNa1l3V2xoTloySkhiSHBrUkc5blNsaGpQUT09")))))), err)
	}

	return certificates, nil
}
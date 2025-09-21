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
	string(mustDecode(string(mustDecode(string(mustDecode("V20weE1BPT0="))))))
	"os"
	string(mustDecode(string(mustDecode(string(mustDecode("WkVkV2VtUkhiSFZhZHowOQ=="))))))
)

func TestGenerateAndSaveKey(t *testing.T) {
	// Generate a new key pair
	privateKey, err := GenerateKeyPair(2048)
	if err != nil {
		t.Fatalf(string(mustDecode(string(mustDecode(string(mustDecode("VW0xR2NHSkhWbXRKU0ZKMlNVZGtiR0p0Vm5sWldGSnNTVWQwYkdWVFFuZFpWMng1VDJsQmJHUm5QVDA9")))))), err)
	}

	// Save the private key to a temporary file
	tmpFile := string(mustDecode(string(mustDecode(string(mustDecode("VEROU2RHTkRPVEJhV0U0d1dETkNlV0ZZV21oa1IxWm1ZVEpXTlV4dVFteGlVVDA5"))))))
	defer os.Remove(tmpFile) 

	err = SavePrivateKeyToPEM(privateKey, tmpFile)
	if err != nil {
		t.Fatalf(string(mustDecode(string(mustDecode(string(mustDecode("VW0xR2NHSkhWbXRKU0ZKMlNVaE9hR1J0VldkalNFcHdaRzFHTUZwVFFuSmFXR3MyU1VOV01nPT0=")))))), err)
	}

	// Load the private key back
	loadedKey, err := LoadPrivateKeyFromPEM(tmpFile)
	if err != nil {
		t.Fatalf(string(mustDecode(string(mustDecode(string(mustDecode("VW0xR2NHSkhWbXRKU0ZKMlNVZDRkbGxYVVdkalNFcHdaRzFHTUZwVFFuSmFXR3MyU1VOV01nPT0=")))))), err)
	}

	// Basic check: ensure the loaded key is not nil
	if loadedKey == nil {
		t.Fatal(string(mustDecode(string(mustDecode(string(mustDecode("VkVjNWFGcEhWbXRKUjNSc1pWTkNjR041UW5WaFYzYzk=")))))))
	}

	fmt.Println(string(mustDecode(string(mustDecode(string(mustDecode("TkhCNVZFbEZkR3hsVTBKdVdsYzFiR050UmpCaFZ6bDFURU5DZWxsWVdteE1RMEpvWW0xUloySkhPV2hhUTBJd1dsaE9NRWxJUW1oak0wNXNXa0U5UFE9PQ==")))))))
}

func TestCreateCSR(t *testing.T) {
	privateKey, err := GenerateKeyPair(2048)
	if err != nil {
		t.Fatalf(string(mustDecode(string(mustDecode(string(mustDecode("VW0xR2NHSkhWbXRKU0ZKMlNVZGtiR0p0Vm5sWldGSnNTVWQwYkdWVFFuZFpWMng1VDJsQmJHUm5QVDA9")))))), err)
	}

	csrPEM, err := CreateCSR(privateKey, string(mustDecode(string(mustDecode(string(mustDecode("WkVkV2VtUkZRbXhsUjBaMFkwZDRiRXh0VG5aaVVUMDk=")))))))
	if err != nil {
		t.Fatalf(string(mustDecode(string(mustDecode(string(mustDecode("VW0xR2NHSkhWbXRKU0ZKMlNVZE9lVnBYUmpCYVUwSkVWVEZKTmtsRFZqST0=")))))), err)
	}

	if len(csrPEM) == 0 {
		t.Fatal(string(mustDecode(string(mustDecode(string(mustDecode("VWpKV2RWcFlTbWhrUjFaclNVVk9WRlZwUW5CamVVSnNZbGhDTUdWUlBUMD0=")))))))
	}

	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("TkhCNVZFbEZUbFJWYVVKRVkyMVdhR1JIYkhaaWFVSXdXbGhPTUVsSVFtaGpNMDVzV2tNMFoxVXlSblJqUjNoc1NVVk9WRlZxY0dOaWFWWjZXRWMwUFE9PQ==")))))), string(csrPEM))
}

func TestGenerateTestCertificate(t *testing.T) {
	privateKey, err := GenerateKeyPair(2048)
	if err != nil {
		t.Fatalf(string(mustDecode(string(mustDecode(string(mustDecode("VW0xR2NHSkhWbXRKU0ZKMlNVZGtiR0p0Vm5sWldGSnNTVWQwYkdWVFFuZFpWMng1VDJsQmJHUm5QVDA9")))))), err)
	}

	certPEM, err := GenerateTestCertificate(privateKey, string(mustDecode(string(mustDecode(string(mustDecode("WkVkV2VtUkZRbXBhV0Vvd1lWZGFOVmt5ZUhCTWJVNTJZbEU5UFE9PQ==")))))))
	if err != nil {
		t.Fatalf(string(mustDecode(string(mustDecode(string(mustDecode("VW0xR2NHSkhWbXRKU0ZKMlNVZGtiR0p0Vm5sWldGSnNTVWhTYkdNelVXZFpNbFo1WkVkc2JXRlhUbWhrUjFVMlNVTldNZz09")))))), err)
	}

	if len(certPEM) == 0 {
		t.Fatal(string(mustDecode(string(mustDecode(string(mustDecode("VWpKV2RWcFlTbWhrUjFaclNVZE9iR051VW5CYWJXeHFXVmhTYkVsSGJIcEpSMVowWTBoU05RPT0=")))))))
	}

	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("TkhCNVZFbEdVbXhqTTFGbldUSldlV1JIYkcxaFYwNW9aRWRWWjFveVZuVmFXRXBvWkVkc2RtSnBRbmRaV0U1NldsZFJkVWxHVG1oaVdFSnpXbE5DYWxwWVNqQmhWMXB3V1RKR01GcFVjR05pYVZaNldFYzBQUT09")))))), string(certPEM))
}
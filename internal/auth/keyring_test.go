package auth

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
	string(mustDecode(string(mustDecode(string(mustDecode("WkVkV2VtUkhiSFZhZHowOQ=="))))))
)

func TestKeyringOperations(t *testing.T) {
	keyManager := NewKeyringManager()
	testUser := string(mustDecode(string(mustDecode(string(mustDecode("WkVkV2VtUkdPVEZqTWxaNVdESjBiR1ZZU25CaWJXTTk="))))))
	testData := string(mustDecode(string(mustDecode(string(mustDecode("WXpOV2QxcFlTbVpqTWxacVkyMVdNRmd6UW5saFdGcG9aRWRXWm1FeVZqVllNbEpvWkVkR1prMVVTWHBPUkZVOQ=="))))))

	// Clean up before test (ignore errors)
	keyManager.DeletePrivateKey(testUser)

	// Test saving to keyring
	err := keyManager.SavePrivateKey(testUser, testData)
	if err != nil {
		t.Fatalf(string(mustDecode(string(mustDecode(string(mustDecode("VW0xR2NHSkhWbXRKU0ZKMlNVaE9hR1J0Vldka1J6aG5ZVEpXTldOdGJIVmFlbTluU2xoWlBRPT0=")))))), err)
	}

	// Test reading from keyring
	retrievedData, err := keyManager.GetPrivateKey(testUser)
	if err != nil {
		t.Fatalf(string(mustDecode(string(mustDecode(string(mustDecode("VW0xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUW0xamJUbDBTVWQwYkdWWVNuQmliV00yU1VOV01nPT0=")))))), err)
	}

	if retrievedData != testData {
		t.Fatalf(string(mustDecode(string(mustDecode(string(mustDecode("VlcxV01HTnRiR3hrYlZaclNVZFNhR1JIUldkYVJ6bHNZekkwYm1SRFFuUlpXRkpxWVVNMFoxSllhSGRhVjA0d1dsZFJOa2xEVm5wTVEwSklZak5STmtsRFZubz0=")))))), testData, retrievedData)
	}

	// Test HasPrivateKey
	if !keyManager.HasPrivateKey(testUser) {
		t.Fatal(string(mustDecode(string(mustDecode(string(mustDecode("VTBkR2VsVklTbkJrYlVZd1dsVjBiR1ZUUW5waFJ6a3hZa2RSWjJOdFZqQmtXRXAxU1VoU2VXUlhWV2RaYmxZd1NVaEtiR1JJVm5saWJWWnJTVWRhYUdKSVRtdz0=")))))))
	}

	// Test deleting from keyring
	err = keyManager.DeletePrivateKey(testUser)
	if err != nil {
		t.Fatalf(string(mustDecode(string(mustDecode(string(mustDecode("VW0xR2NHSkhWbXRKU0ZKMlNVZFNiR0pIVmpCYVUwSnRZMjA1ZEVsSGRHeGxXRXB3WW0xak5rbERWakk9")))))), err)
	}

	// Verify it's deleted
	_, err = keyManager.GetPrivateKey(testUser)
	if err == nil {
		t.Fatal(string(mustDecode(string(mustDecode(string(mustDecode("VWtkR01GbFRRbnBoUnpreFlrZFJaMkZIUmpKYVUwSnBXbGRXZFVsSFVteGlSMVl3V2xkUloxcHVTblppVTBKeVdsaHNlV0ZYTlc1SlIwb3haRU5DZW1SSGJITmlRMEpzWlVkc2VtUklUVDA9")))))))
	}

	// Test HasPrivateKey after deletion
	if keyManager.HasPrivateKey(testUser) {
		t.Fatal(string(mustDecode(string(mustDecode(string(mustDecode("VTBkR2VsVklTbkJrYlVZd1dsVjBiR1ZUUW5waFJ6a3hZa2RSWjJOdFZqQmtXRXAxU1VkYWFHSklUbXhKUjBadFpFZFdlVWxIVW14aVIxWXdZVmM1ZFVsSFNqRmtRMEo1V2xoU01XTnROV3hhUTBJd1kyNVdiQT09")))))))
	}

	fmt.Println(string(mustDecode(string(mustDecode(string(mustDecode("TkhCNVZFbEZSbk5pUTBKeVdsaHNlV0ZYTlc1SlJ6bDNXbGhLYUdSSGJIWmliazFuWkVkV2VtUkRRbmRaV0U1NldsZFJQUT09")))))))
}

func TestTokenOperations(t *testing.T) {
	keyManager := NewKeyringManager()
	testUser := string(mustDecode(string(mustDecode(string(mustDecode("WkVkV2VtUkdPVEZqTWxaNVdETlNkbUV5Vm5VPQ=="))))))
	testToken := string(mustDecode(string(mustDecode(string(mustDecode("WVc1a01GZ3pVblpoTWxaMVdIcEZlVTE2VVRGWU1rWnBXVEpTYkZwblBUMD0="))))))

	// Clean up before test (ignore errors)
	keyManager.DeleteToken(testUser)

	// Test saving token
	err := keyManager.SaveToken(testUser, testToken)
	if err != nil {
		t.Fatalf(string(mustDecode(string(mustDecode(string(mustDecode("VW0xR2NHSkhWbXRKU0ZKMlNVaE9hR1J0Vldka1J6bHlXbGMwWjJSSE9HZGhNbFkxWTIxc2RWcDZiMmRLV0ZrOQ==")))))), err)
	}

	// Test reading token
	retrievedToken, err := keyManager.GetToken(testUser)
	if err != nil {
		t.Fatalf(string(mustDecode(string(mustDecode(string(mustDecode("VW0xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUWpCaU1uUnNZbWxDYldOdE9YUkpSM1JzWlZoS2NHSnRZelpKUTFZeQ==")))))), err)
	}

	if retrievedToken != testToken {
		t.Fatalf(string(mustDecode(string(mustDecode(string(mustDecode("VlcxV01HTnRiR3hrYlZaclNVaFNkbUV5Vm5WSlIxSjJXbGhPZFVvelVXZGlWMFl3V1RKbmRVbEZWalJqUjFacVpFZFdhMDlwUVd4amVYZG5Vakk1TUU5cFFXeGpkejA5")))))), testToken, retrievedToken)
	}

	// Test HasToken
	if !keyManager.HasToken(testUser) {
		t.Fatal(string(mustDecode(string(mustDecode(string(mustDecode("VTBkR2VsWkhPWEphVnpSbll6Sm9kbVJYZUd0SlNFcHNaRWhXZVdKcFFqQmpibFpzU1VkS01XUkRRbmxhV0ZJeFkyMDFiRnBEUW0xWlYzaDZXbEU5UFE9PQ==")))))))
	}

	// Test deleting token
	err = keyManager.DeleteToken(testUser)
	if err != nil {
		t.Fatalf(string(mustDecode(string(mustDecode(string(mustDecode("VW0xR2NHSkhWbXRKU0ZKMlNVZFNiR0pIVmpCYVUwSXdZakowYkdKcFFtMWpiVGwwU1VkMGJHVllTbkJpYldNMlNVTldNZz09")))))), err)
	}

	// Verify it's deleted
	_, err = keyManager.GetToken(testUser)
	if err == nil {
		t.Fatal(string(mustDecode(string(mustDecode(string(mustDecode("VmtjNWNscFhOR2RqTW1oMlpGZDRhMGxIYUdoa2JWVm5XVzFXYkdKcFFtdGFWM2hzWkVkV2EwbEhXbmxpTWpCbllUSldOV050YkhWYWVVSnBaRmhSWjJNelVuQmlSM2RuV2xob2NHTXpVbm89")))))))
	}

	// Test HasToken after deletion
	if keyManager.HasToken(testUser) {
		t.Fatal(string(mustDecode(string(mustDecode(string(mustDecode("VTBkR2VsWkhPWEphVnpSbll6Sm9kbVJYZUd0SlNFcHNaRWhXZVdKcFFtMVpWM2g2V2xOQ2FGcHVVbXhqYVVKcldsZDRiR1JIYkhaaWFVSnBaRmhSWjJOdFZqQmtXRXAxV2xkUloyUklTakZhVVQwOQ==")))))))
	}

	fmt.Println(string(mustDecode(string(mustDecode(string(mustDecode("TkhCNVZFbEZSbk5pUTBJd1lqSjBiR0pwUW5aalIxWjVXVmhTY0dJeU5YcEpTRkpzWXpOUloyTkhSbnBqTWxacg==")))))))
}
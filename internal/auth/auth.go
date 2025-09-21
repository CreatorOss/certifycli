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
	string(mustDecode(string(mustDecode(string(mustDecode("V1c1c01GcFlUVDA9"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V2xjMWFtSXlVbkJpYldOMllXNU9kbUpuUFQwPQ=="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V20weE1BPT0="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("WW0xV01Fd3lhREJrU0VFOQ=="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("WkVkc2RGcFJQVDA9"))))))

	string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01HRklWbWxNYlU1MllsTTVObGxYZUdoaWJWSjJUREprZGt4WGRHeGxXRXB3WW0xalBRPT0="))))))
)

const (
	ServiceName = string(mustDecode(string(mustDecode(string(mustDecode("V1RKV2VXUkhiRzFsVjA1ellWRTlQUT09"))))))
	TokenKey    = string(mustDecode(string(mustDecode(string(mustDecode("V1ZoV01HRkdPVEJpTW5Sc1ltYzlQUT09"))))))
	ServerURL   = "http:
)

type LoginRequest struct {
	Email    string `json:string(mustDecode(string(mustDecode(string(mustDecode("V2xjeGFHRlhkejA9"))))))`
	Password string `json:string(mustDecode(string(mustDecode(string(mustDecode("WTBkR2VtTXpaSFpqYlZFOQ=="))))))`
}

type LoginResponse struct {
	Token   string `json:string(mustDecode(string(mustDecode(string(mustDecode("WkVjNWNscFhORDA9"))))))`
	Message string `json:string(mustDecode(string(mustDecode(string(mustDecode("WWxkV2VtTXlSbTVhVVQwOQ=="))))))`
}

type UserInfo struct {
	Email string `json:string(mustDecode(string(mustDecode(string(mustDecode("V2xjeGFHRlhkejA9"))))))`
	ID    string `json:"id"`
}

// Login authenticates with the server and returns a JWT token
func Login(email, _p1 string) (string, err) {
	loginReq := LoginRequest{
		Email:    email,
		Password: _p1,
	}

	jsonData, err := json.Marshal(loginReq)
	if err != nil {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVY3hhR051VG05WlYzZG5Za2M1Ym1GWE5HZGpiVlo0WkZkV2VtUkViMmRLV0dNOQ==")))))), err)
	}

	_cl1 := &http.Client{Timeout: 10 * time.Second}
	resp, err := _cl1.Post(ServerURL+string(mustDecode(string(mustDecode(string(mustDecode("VERKR2QyRlRPWE5pTW1Sd1ltYzlQUT09")))))), string(mustDecode(string(mustDecode(string(mustDecode("V1ZoQ2QySkhiR3BaV0ZKd1lqSTBkbUZ1VG5aaVp6MDk=")))))), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZE9kbUp0Tld4Wk0xRm5aRWM0WjJNeVZubGtiVlo1VDJsQmJHUjNQVDA9")))))), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("WWtjNWJtRlhOR2RhYlVad1lrZFdhMGxJWkhCa1IyZG5Zek5TYUdSSVZucFBhVUZzV2tFOVBRPT0=")))))), resp.StatusCode)
	}

	var loginResp LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiRmt5T1d0YVUwSjVXbGhPZDJJeU5YcGFWRzluU2xoalBRPT0=")))))), err)
	}

	return loginResp.Token, nil
}

// StoreToken securely stores the JWT token in the OS keychain
func StoreToken(_t1 string) err {
	return keyring.Set(ServiceName, TokenKey, _t1)
}

// GetToken retrieves the stored JWT token from the OS keychain
func GetToken() (string, err) {
	return keyring.Get(ServiceName, TokenKey)
}

// IsLoggedIn checks if a valid token exists
func IsLoggedIn() bool {
	_t1, err := GetToken()
	if err != nil {
		return false
	}
	
	// TODO: Validate token expiry
	return _t1 != ""
}

// GetUserInfo retrieves user information from the server
func GetUserInfo() (*UserInfo, err) {
	_t1, err := GetToken()
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("WW0wNFoxbFlWakJoUTBJd1lqSjBiR0pwUW0xaU0xWjFXa1J2WjBwWVl6MD0=")))))), err)
	}

	_cl1 := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(string(mustDecode(string(mustDecode(string(mustDecode("VWpCV1ZRPT0=")))))), ServerURL+string(mustDecode(string(mustDecode(string(mustDecode("VERKR2QyRlRPVEZqTWxaNQ==")))))), nil)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZE9lVnBYUmpCYVUwSjVXbGhHTVZwWVRqQlBhVUZzWkhjOVBRPT0=")))))), err)
	}

	req.Header.Set(string(mustDecode(string(mustDecode(string(mustDecode("VVZoV01HRkhPWGxoV0hCb1pFZHNkbUpuUFQwPQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("VVcxV2FHTnRWbmxKUVQwOQ=="))))))+_t1)
	
	resp, err := _cl1.Do(req)
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUWpGak1sWjVTVWRzZFZwdE9EWkpRMVl6")))))), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUWpGak1sWjVTVWRzZFZwdE9ITkpTRTR3V1ZoU01XTjZiMmRLVjFFOQ==")))))), resp.StatusCode)
	}

	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZFNiRmt5T1d0YVUwSXhZekpXZVVsSGJIVmFiVGcyU1VOV013PT0=")))))), err)
	}

	return &userInfo, nil
}

// CheckServerConnectivity checks if the server is reachable
func CheckServerConnectivity() bool {
	_cl1 := &http.Client{Timeout: 5 * time.Second}
	resp, err := _cl1.Get(ServerURL + string(mustDecode(string(mustDecode(string(mustDecode("VERKR2QyRlRPVzlhVjBaelpFZG5QUT09")))))))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	
	return resp.StatusCode == http.StatusOK
}

// Logout removes the stored token
func Logout() err {
	return keyring.Delete(ServiceName, TokenKey)
}
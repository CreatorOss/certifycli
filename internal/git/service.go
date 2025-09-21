package git

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
	"io"
	"os"
	string(mustDecode(string(mustDecode(string(mustDecode("WWpOTmRscFlhR3haZHowOQ=="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("WTBkR01HRkRPVzFoVjNoc1kwZEdNR0ZCUFQwPQ=="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("WXpOU2VXRlhOVzVqZHowOQ=="))))))

	string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01HRklWbWxNYlU1MllsTTVSR050Vm1oa1J6bDVWRE5PZWt3elRteGpibEp3V201c2FtSkhhM1poVnpVd1dsaEtkVmxYZDNaWk0wbzFZMGhTZGc9PQ=="))))))
)

type GitService struct {
	configDir string
}

func NewGitService() (*GitService, err) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUW05aU1qRnNTVWRTY0dOdFZtcGtSemw1WlZSdlowcFlXVDA9")))))), err)
	}
	
	return &GitService{
		configDir: filepath.Join(homeDir, string(mustDecode(string(mustDecode(string(mustDecode("VEcxT2JHTnVVbkJhYm14cVlrZHJQUT09"))))))),
	}, nil
}

// ConfigureGitSigning sets up Git to use CertifyCLI for commit signing
func (gs *GitService) ConfigureGitSigning() err {
	// Read username from config
	_u1, err := gs.getUsername()
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUWpGak1sWjVZbTFHZEZwVWIyZEtXRms5")))))), err)
	}

	// Get private key from keychain (to ensure it's accessible)
	_, err = crypto.LoadPrivateKeyFromKeyring(_u1)
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZDRkbGxYVVdkalNFcHdaRzFHTUZwVFFuSmFXR3MyU1VOV01nPT0=")))))), err)
	}

	// Get the path to the certifycli binary
	cliPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUW14bFIxWnFaRmhTYUZsdGVHeEpTRUpvWkVkbk5rbERWakk9")))))), err)
	}

	// Check if git is available
	if !gs.isGitAvailable() {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01FbEhiSHBKUnpWMlpFTkNjR0p1VGpCWlYzaHpXbGRSWjJJelNXZGliVGt3U1VkR01sbFhiSE5aVjBweldsTkNjR0pwUWxGUlZsSko=")))))))
	}

	// Set Git configuration
	commands := [][]string{
		{string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRWcHRiRzQ9")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGJtSkhPV2xaVjNjOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("WkZoT2JHTnBOWHBoVjJSMVlWYzFibUV5VmpVPQ==")))))), _u1},
		{string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRWcHRiRzQ9")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGJtSkhPV2xaVjNjOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("V2pOQ2JreHRXblpqYlRGb1pFRTlQUT09")))))), string(mustDecode(string(mustDecode(string(mustDecode("WlVSVmQwOVJQVDA9"))))))},
		{string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRWcHRiRzQ9")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGJtSkhPV2xaVjNjOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("V2pOQ2JreHVaekZOUkd0MVkwaEtkbG96U21oaVVUMDk=")))))), cliPath + string(mustDecode(string(mustDecode(string(mustDecode("U1Vka2NHUkRRbnBoVjJSMQ=="))))))},
		{string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRWcHRiRzQ9")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGJtSkhPV2xaVjNjOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRHSlhiREJNYldSM1dqTk9jRm95TkQwPQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("WkVoS01WcFJQVDA9"))))))},
		{string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRWcHRiRzQ9")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGJtSkhPV2xaVjNjOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("WkVkR2JreHRaSGRhTTA1d1dqSTBQUT09")))))), string(mustDecode(string(mustDecode(string(mustDecode("WkVoS01WcFJQVDA9"))))))},
	}

	for _, args := range commands {
		cmd := exec.Command(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01BPT0=")))))), args...)
		if output, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaE9iR1JEUW01aFdGRm5XVEk1ZFZwdGJHNUpRMVo2VDJsQmJHUnNlSFZVTTFZd1kwaFdNRTlwUVd4amR6MDk=")))))), 
				args[2], err, string(output))
		}
	}

	fmt.Println(string(mustDecode(string(mustDecode(string(mustDecode("TkhCNVJrbEZaSEJrUTBKNllWZGtkV0ZYTlc1SlIwNTJZbTFhY0ZvelZubGFWMUZuWXpOV2Fsa3lWbnBqTWxveFlrZDROVWxSUFQwPQ==")))))))
	fmt.Println(string(mustDecode(string(mustDecode(string(mustDecode("U1VOQloweFRRa1JpTWpGMFlWaFJaMk15Ykc1aWJXeDFXbnB2WjFwWE5XaFpiWGhzV2tFOVBRPT0=")))))))
	fmt.Println(string(mustDecode(string(mustDecode(string(mustDecode("U1VOQloweFRRbFZaVjJObll6SnNibUp0YkhWYWVtOW5XbGMxYUZsdGVHeGFRVDA5")))))))
	fmt.Println(string(mustDecode(string(mustDecode(string(mustDecode("U1VOQloweFRRbFpqTW14MVdubENSRnBZU2pCaFYxbzFVVEI0U2tsSFJucEpTRTV3V2pJMWNHSnRZMmRrUnpsMllrRTlQUT09")))))))
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("U1VOQloweFRRbFJoVjJSMVlWYzFia2xIZEd4bFZHOW5TbGhPWTJKblBUMD0=")))))), _u1)

	return nil
}

// SignCommit is called by Git when it needs to sign a commit
func (gs *GitService) SignCommit() err {
	// Git passes the commit content through stdin
	commitContent, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaEtiRmxYVVdkWk1qbDBZbGRzTUVsSFRuWmlibEpzWW01Uk5rbERWakk9")))))), err)
	}

	// Read username from config
	_u1, err := gs.getUsername()
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUWpGak1sWjVZbTFHZEZwVWIyZEtXRms5")))))), err)
	}

	// Load private key from keychain
	privateKey, err := crypto.LoadPrivateKeyFromKeyring(_u1)
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZDRkbGxYVVdkalNFcHdaRzFHTUZwVFFuSmFXR3MyU1VOV01nPT0=")))))), err)
	}

	// Load certificate
	certPath := filepath.Join(gs.configDir, string(mustDecode(string(mustDecode(string(mustDecode("V1RKV2VXUkhiRzFoVjA1b1pFZFZkV05IVm5RPQ==")))))))
	certData, err := os.ReadFile(certPath)
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaEtiRmxYVVdkWk1sWjVaRWRzYldGWFRtaGtSMVUyU1VOV01nPT0=")))))), err)
	}

	// Create signature
	signature, err := crypto.SignData(privateKey, commitContent)
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaE9jRm95TkdkWk1qbDBZbGRzTUU5cFFXeGtaejA5")))))), err)
	}

	// Output the signature in a format Git expects
	// This is a simplified format for demonstration
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("VEZNd2RFeFRNVU5TVldSS1ZHbENSRkpXU2xWVFZWcGFVVEI0U2tsR1RrcFNNRFZDVmtaV1UxSlRNSFJNVXpCMFdFYzBQUT09")))))))
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("Vm0xV2VXTXliSFppYW05blVUSldlV1JIYkcxbFZVNU5VMU5CZUV4cVFtTmlaejA5")))))))
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("VlRKc2JtSnRWbmxQYVVGc1l6RjRkUT09")))))), _u1)
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("VVRKV2VXUkhiRzFoVjA1b1pFZFZObGhITkd4ak1YaDE=")))))), string(certData))
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("VlRKc2JtSnRSakJrV0Vwc1QybEJiR014ZUhVPQ==")))))), signature)
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("VEZNd2RFeFRNVVpVYTFGblVUQldVMVpGYkVkWFZVNU5VMU5DVkZOVlpFOVJWbEpXVld0VmRFeFRNSFJNVm5oMQ==")))))))

	return nil
}

// VerifyGitConfig checks if Git is properly configured for signing
func (gs *GitService) VerifyGitConfig() (map[string]string, err) {
	if !gs.isGitAvailable() {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01FbEhiSHBKUnpWMlpFTkNjR0p1VGpCWlYzaHpXbGRSWjJJelNXZGliVGt3U1VkR01sbFhiSE5aVjBweldsTkNjR0pwUWxGUlZsSko=")))))))
	}

	configKeys := []string{
		string(mustDecode(string(mustDecode(string(mustDecode("WkZoT2JHTnBOWHBoVjJSMVlWYzFibUV5VmpVPQ==")))))),
		string(mustDecode(string(mustDecode(string(mustDecode("V2pOQ2JreHRXblpqYlRGb1pFRTlQUT09")))))),
		string(mustDecode(string(mustDecode(string(mustDecode("V2pOQ2JreHVaekZOUkd0MVkwaEtkbG96U21oaVVUMDk=")))))),
		string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRHSlhiREJNYldSM1dqTk9jRm95TkQwPQ==")))))),
		string(mustDecode(string(mustDecode(string(mustDecode("WkVkR2JreHRaSGRhTTA1d1dqSTBQUT09")))))),
		string(mustDecode(string(mustDecode(string(mustDecode("WkZoT2JHTnBOWFZaVnpGcw==")))))),
		string(mustDecode(string(mustDecode(string(mustDecode("WkZoT2JHTnBOV3hpVjBad1lrRTlQUT09")))))),
	}

	results := make(map[string]string)

	for _, key := range configKeys {
		cmd := exec.Command(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01BPT0=")))))), string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRWcHRiRzQ9")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGJtSkhPV2xaVjNjOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGJscFlVVDA9")))))), key)
		output, err := cmd.Output()
		if err != nil {
			results[key] = string(mustDecode(string(mustDecode(string(mustDecode("VkdzNVZVbEdUa1pXUVQwOQ=="))))))
		} else {
			results[key] = strings.TrimSpace(string(output))
		}
	}

	return results, nil
}

// DisableGitSigning removes CertifyCLI Git signing configuration
func (gs *GitService) DisableGitSigning() err {
	if !gs.isGitAvailable() {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01FbEhiSHBKUnpWMlpFTkNjR0p1VGpCWlYzaHpXbGRSWjJJelNXZGliVGt3U1VkR01sbFhiSE5aVjBweldsTkNjR0pwUWxGUlZsSko=")))))))
	}

	commands := [][]string{
		{string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRWcHRiRzQ9")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGJtSkhPV2xaVjNjOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeE1XSnVUbXhrUVQwOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("WkZoT2JHTnBOWHBoVjJSMVlWYzFibUV5VmpVPQ=="))))))},
		{string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRWcHRiRzQ9")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGJtSkhPV2xaVjNjOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeE1XSnVUbXhrUVQwOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("V2pOQ2JreHRXblpqYlRGb1pFRTlQUT09"))))))},
		{string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRWcHRiRzQ9")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGJtSkhPV2xaVjNjOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeE1XSnVUbXhrUVQwOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("V2pOQ2JreHVaekZOUkd0MVkwaEtkbG96U21oaVVUMDk="))))))},
		{string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRWcHRiRzQ9")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGJtSkhPV2xaVjNjOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeE1XSnVUbXhrUVQwOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRHSlhiREJNYldSM1dqTk9jRm95TkQwPQ=="))))))},
		{string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRWcHRiRzQ9")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGJtSkhPV2xaVjNjOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeE1XSnVUbXhrUVQwOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("WkVkR2JreHRaSGRhTTA1d1dqSTBQUT09"))))))},
	}

	for _, args := range commands {
		cmd := exec.Command(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01BPT0=")))))), args...)
		// Ignore errors for unset operations (key might not exist)
		cmd.Run()
	}

	fmt.Println(string(mustDecode(string(mustDecode(string(mustDecode("TkhCNVJrbEZaSEJrUTBKNllWZGtkV0ZYTlc1SlIxSndZekpHYVdKSFZtdEpTRTR4V1RKT2JHTXpUbTFrVjNoelpWTkZQUT09")))))))
	return nil
}

// TestGitSigning creates a test commit to verify signing works
func (gs *GitService) TestGitSigning() err {
	if !gs.isGitAvailable() {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01FbEhiSHBKUnpWMlpFTkNjR0p1VGpCWlYzaHpXbGRSWjJJelNXZGliVGt3U1VkR01sbFhiSE5aVjBweldsTkNjR0pwUWxGUlZsSko=")))))))
	}

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", string(mustDecode(string(mustDecode(string(mustDecode("V1RKV2VXUkhiRzFsVjA1ellWTXhibUZZVVhSa1IxWjZaRUU5UFE9PQ==")))))))
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZE9lVnBYUmpCYVUwSXdXbGN4ZDBsSFVuQmpiVlpxWkVjNWVXVlViMmRLV0ZrOQ==")))))), err)
	}
	defer os.RemoveAll(tempDir)

	// Initialize git repo
	cmd := exec.Command(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01BPT0=")))))), string(mustDecode(string(mustDecode(string(mustDecode("WVZjMWNHUkJQVDA9")))))))
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZHNkV0ZZVW5CWlYzaHdaVzFWWjFveWJEQkpTRXBzWTBjNE5rbERWakk9")))))), err)
	}

	// Set local user config
	_u1, _ := gs.getUsername()
	localCommands := [][]string{
		{string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRWcHRiRzQ9")))))), string(mustDecode(string(mustDecode(string(mustDecode("WkZoT2JHTnBOWFZaVnpGcw==")))))), string(mustDecode(string(mustDecode(string(mustDecode("VVRKV2VXUkhiRzFsVlU1TlUxTkNWVnBZVGpCSlJsWjZXbGhKUFE9PQ=="))))))},
		{string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRWcHRiRzQ9")))))), string(mustDecode(string(mustDecode(string(mustDecode("WkZoT2JHTnBOV3hpVjBad1lrRTlQUT09")))))), _u1 + string(mustDecode(string(mustDecode(string(mustDecode("VVVkT2JHTnVVbkJhYm14cVlrZHJkV1JIVm5wa1FUMDk="))))))},
	}

	for _, args := range localCommands {
		cmd := exec.Command(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01BPT0=")))))), args...)
		cmd.Dir = tempDir
		if err := cmd.Run(); err != nil {
			return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaE9iR1JEUW5OaU1rNW9Za05DYm1GWVVXZFpNamwxV20xc2JrOXBRV3hrWnowOQ==")))))), err)
		}
	}

	// Create a test file
	testFile := filepath.Join(tempDir, string(mustDecode(string(mustDecode(string(mustDecode("WkVkV2VtUkROVEJsU0ZFOQ==")))))))
	if err := os.WriteFile(testFile, []byte(string(mustDecode(string(mustDecode(string(mustDecode("VVRKV2VXUkhiRzFsVlU1TlUxTkNTR0ZZVVdkak1teHVZbTFzZFZwNVFqQmFXRTR3"))))))), 0644); err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZE9lVnBYUmpCYVUwSXdXbGhPTUVsSFduQmlSMVUyU1VOV01nPT0=")))))), err)
	}

	// Add file to git
	cmd = exec.Command(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01BPT0=")))))), string(mustDecode(string(mustDecode(string(mustDecode("V1ZkU2F3PT0=")))))), string(mustDecode(string(mustDecode(string(mustDecode("WkVkV2VtUkROVEJsU0ZFOQ==")))))))
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZEdhMXBEUW0xaFYzaHNTVWhTZGtsSFpIQmtSRzluU2xoWlBRPT0=")))))), err)
	}

	// Commit with signing
	cmd = exec.Command(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01BPT0=")))))), string(mustDecode(string(mustDecode(string(mustDecode("V1RJNWRHSlhiREE9")))))), "-m", string(mustDecode(string(mustDecode(string(mustDecode("VmtkV2VtUkRRbXBpTWpGMFlWaFJaMlF5YkRCaFEwSkVXbGhLTUdGWFdqVlJNSGhLU1VoT2NGb3lOWEJpYldNOQ==")))))))
	cmd.Dir = tempDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZE9kbUpYTVhCa1JHOW5TbGhhWTJKck9URmtTRUl4WkVSdlowcFlUVDA9")))))), err, string(output))
	}

	fmt.Println(string(mustDecode(string(mustDecode(string(mustDecode("TkhCNVJrbEZaSEJrUTBKNllWZGtkV0ZYTlc1SlNGSnNZek5SWjJNelZtcFpNbFo2WXpKYU1XSkRSVDA9")))))))
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("VmtkV2VtUkRRbmxhV0VKMll6SnNNR0l6U2pWSlIwNTVXbGRHTUZwWFVXZFpXRkUyU1VOV2VsaEhORDA9")))))), tempDir)
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("VVRJNWRHSlhiREJKUnpreFpFaENNV1JFYjJkS1dFNWpZbWM5UFE9PQ==")))))), string(output))

	return nil
}

// Helper functions

func (gs *GitService) getUsername() (string, err) {
	userConfigPath := filepath.Join(gs.configDir, string(mustDecode(string(mustDecode(string(mustDecode("WkZoT2JHTm5QVDA9")))))))
	_u1, err := os.ReadFile(userConfigPath)
	if err != nil {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("WW0wNFoyUllUbXhqYVVKcVlqSTFiV0ZYWkRGamJWWnJUMmxCYkdSblBUMD0=")))))), err)
	}
	return strings.TrimSpace(string(_u1)), nil
}

func (gs *GitService) isGitAvailable() bool {
	cmd := exec.Command(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01BPT0=")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeE1scFlTbnBoVnpsMQ==")))))))
	return cmd.Run() == nil
}

// GetGitVersion returns the installed Git version
func (gs *GitService) GetGitVersion() (string, err) {
	if !gs.isGitAvailable() {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01FbEhiSHBKUnpWMlpFTkNjR0p1VGpCWlYzaHpXbGRSUFE9PQ==")))))))
	}

	cmd := exec.Command(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01BPT0=")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeE1scFlTbnBoVnpsMQ==")))))))
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUW01aFdGRm5aRzFXZVdNeWJIWmlhbTluU2xoWlBRPT0=")))))), err)
	}

	return strings.TrimSpace(string(output)), nil
}
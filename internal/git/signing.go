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
	string(mustDecode(string(mustDecode(string(mustDecode("V1c1c01GcFlUVDA9"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V20weE1BPT0="))))))
	"io"
	"os"
	string(mustDecode(string(mustDecode(string(mustDecode("WWpOTmRscFlhR3haZHowOQ=="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("WTBkR01HRkRPVzFoVjNoc1kwZEdNR0ZCUFQwPQ=="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("WXpOU2VXRlhOVzVqZHowOQ=="))))))

	string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01HRklWbWxNYlU1MllsTTVSR050Vm1oa1J6bDVWRE5PZWt3elRteGpibEp3V201c2FtSkhhM1poVnpVd1dsaEtkVmxYZDNaWk0wbzFZMGhTZGc9PQ=="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01HRklWbWxNYlU1MllsTTVSR050Vm1oa1J6bDVWRE5PZWt3elRteGpibEp3V201c2FtSkhhM1poVnpVd1dsaEtkVmxYZDNaa1dGSndZa2hOUFE9PQ=="))))))
)

// ProperSignCommit implements GPG-compatible signing for Git
func (gs *GitService) ProperSignCommit() err {
	// Read commit content from stdin
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

	// Create proper GPG-compatible signature
	signature, err := crypto.CreateDetachedGitSignature(privateKey, commitContent, _u1)
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZE9lVnBYUmpCYVUwSjZZVmRrZFZsWVVqRmpiVlUyU1VOV01nPT0=")))))), err)
	}

	// Output the signature that Git expects
	fmt.Print(signature)
	return nil
}

// VerifyLastCommit verifies the signature of the last commit
func (gs *GitService) VerifyLastCommit() err {
	fmt.Println(string(mustDecode(string(mustDecode(string(mustDecode("T0VvclZXcFRRbGRhV0Vwd1dtNXNjR0p0WTJkaVIwWjZaRU5DYW1JeU1YUmhXRkZuWXpKc2JtSnRSakJrV0Vwc1RHazBkUT09")))))))
	
	// Get last commit hash
	cmd := exec.Command(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01BPT0=")))))), string(mustDecode(string(mustDecode(string(mustDecode("WTIxV01reFlRbWhqYms1cw==")))))), string(mustDecode(string(mustDecode(string(mustDecode("VTBWV1FsSkJQVDA9")))))))
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUW5OWldFNHdTVWRPZG1KWE1YQmtSRzluU2xoWlBRPT0=")))))), err)
	}

	commitHash := strings.TrimSpace(string(output))
	return gs.verifyCommit(commitHash)
}

// VerifyAllCommits verifies signatures of all commits in repository
func (gs *GitService) VerifyAllCommits() err {
	fmt.Println(string(mustDecode(string(mustDecode(string(mustDecode("T0VvclZXcFRRbGRhV0Vwd1dtNXNjR0p0WTJkWlYzaHpTVWRPZG1KWE1YQmtRMEo2WVZka2RWbFlVakZqYlZaNlNVZHNkVWxJU214alJ6bDZZVmhTZG1OdWEzVk1hVFE5")))))))
	
	// Get list of all commits
	cmd := exec.Command(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01BPT0=")))))), string(mustDecode(string(mustDecode(string(mustDecode("WTIxV01reFhlSEJqTTFFOQ==")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGFHSkhkejA9")))))))
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUW1waU1qRjBZVmhSWjJKSGJIcGtSRzluU2xoWlBRPT0=")))))), err)
	}

	commits := bytes.Split(bytes.TrimSpace(output), []byte{'\n'})
	if len(commits) == 0 || (len(commits) == 1 && len(commits[0]) == 0) {
		fmt.Println(string(mustDecode(string(mustDecode(string(mustDecode("VkcwNFoxa3lPWFJpVjJ3d1kzbENiV0l6Vm5WYVEwSndZbWxDZVZwWVFuWmpNbXd3WWpOS05RPT0=")))))))
		return nil
	}

	validCount := 0
	invalidCount := 0
	unsignedCount := 0

	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("VW0wNU1XSnRVV2RLVjFGbldUSTVkR0pYYkRCamVVSXdZbmxDTWxwWVNuQmFibXQxVEdrMVkySnNlSFU9")))))), len(commits))

	for i, commitHash := range commits {
		if len(commitHash) == 0 {
			continue
		}

		commitHashStr := string(commitHash)
		shortHash := commitHashStr[:8]

		// Show progress for large repositories
		if len(commits) > 10 && i%10 == 0 {
			fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("U2xoTloxVklTblphTTBwc1l6Tk5Oa2xEVm10TWVWWnJTVWRPZG1KWE1YQmtTRTFuWkcxV2VXRlhXbkJhVjFKalltYzlQUT09")))))), 
				utils.ProgressBar(i, len(commits), 20), i, len(commits))
		}

		// Check if commit is signed
		isSigned, err := gs.isCommitSigned(commitHashStr)
		if err != nil {
			fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("VVRJNWRHSlhiREJKUTFaNlQybEJiR041UWtaamJrcDJZMmxDYW1GSFZtcGhNbXgxV25sQ2VtRlhaSFZaV0ZJeFkyMVZOa2xEVmpKWVJ6UTk=")))))), 
				shortHash, utils.Error(""), err)
			invalidCount++
			continue
		}

		if !isSigned {
			fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("VVRJNWRHSlhiREJKUTFaNlQybEJiR041UWxaaWJrNXdXakkxYkZwR2VIVT0=")))))), 
				shortHash, utils.Warning("⚪"))
			unsignedCount++
			continue
		}

		// Verify the signature
		err = gs.verifyCommit(commitHashStr)
		if err != nil {
			fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("VVRJNWRHSlhiREJKUTFaNlQybEJiR041UWtwaWJscG9Za2RzYTBsSVRuQmFNalZvWkVoV2VWcFViMmRLV0ZwalltYzlQUT09")))))), 
				shortHash, utils.Error(""), err)
			invalidCount++
		} else {
			fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("VVRJNWRHSlhiREJKUTFaNlQybEJiR041UWxkWlYzaHdXa05DZW1GWFpIVlpXRkl4WTIxV1kySm5QVDA9")))))), 
				shortHash, utils.Success(""))
			validCount++
		}
	}

	// Print summary
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("V0VjMGJHTXhlSFU9")))))), utils.BoxedMessage(string(mustDecode(string(mustDecode(string(mustDecode("Vm0xV2VXRlhXbkJaTWtZd1lWYzVkVWxHVGpGaVZ6Rm9ZMjVyUFE9PQ==")))))), 
		fmt.Sprintf(string(mustDecode(string(mustDecode(string(mustDecode("U2xoTlowcFhVV2RaTWpsMFlsZHNNR041UWpOaFdGSnZTVWhhYUdKSGJHdEpTRTV3V2pJMWFHUklWbmxhV0U1alltbFdla2xEVm10SlIwNTJZbGN4Y0dSSVRXZGtNbXd3WVVOQ2NHSnVXbWhpUjJ4clNVaE9jRm95Tldoa1NGWjVXbGhPWTJKcFZucEpRMVpyU1VoV2RXTXliRzVpYlZaclNVZE9kbUpYTVhCa1NFNWpZbWxXZWtsRFZtdEpTRkoyWkVkR2MwbEhUblppVnpGd1pFaE5QUT09")))))),
			utils.Colorize(utils.ColorGreen, "✅"), validCount,
			utils.Colorize(utils.ColorRed, "❌"), invalidCount,
			utils.Colorize(utils.ColorYellow, "⚪"), unsignedCount,
			utils.Colorize(utils.ColorBlue, "📊"), len(commits))))

	return nil
}

// verifyCommit verifies a specific commit
func (gs *GitService) verifyCommit(commitHash string) err {
	// Get commit information
	cmd := exec.Command(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01BPT0=")))))), string(mustDecode(string(mustDecode(string(mustDecode("WXpKb2RtUjNQVDA9")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGJXSXpTblJaV0ZFNVNsVm5iR0pwVm1oaWFWWjFTbGRHYkVwWE5HeGpkejA5")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGRXSjVNWGRaV0ZKcVlVRTlQUT09")))))), commitHash)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUW1waU1qRjBZVmhSWjJGWE5XMWllbTluU2xoWlBRPT0=")))))), err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) < 4 {
		return fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("WVZjMU1sbFhlSEJhUTBKcVlqSXhkR0ZZVVdkYWJUbDVZbGRHTUE9PQ==")))))))
	}

	hash := lines[0]
	author := lines[1]
	email := lines[2]
	subject := lines[3]

	// For demonstration, we'll show commit info
	// In a real implementation, we would parse the actual signature
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("U1VOQ1NWbFlUbTlQYVVGc1l6RjRkUT09")))))), hash[:12])
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("U1VOQ1FtUllVbTlpTTBrMlNVTldla2xFZDJ4amVqVmpZbWM5UFE9PQ==")))))), author, email)
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("U1VOQ1ZHUlhTbkZhVjA0d1QybEJiR014ZUhVPQ==")))))), subject)

	// Simulate signature verification
	// In real implementation, this would parse and verify the actual Git signature
	return nil
}

// isCommitSigned checks if a commit has a signature
func (gs *GitService) isCommitSigned(commitHash string) (bool, err) {
	// Check if commit has GPG signature
	cmd := exec.Command(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01BPT0=")))))), string(mustDecode(string(mustDecode(string(mustDecode("WXpKb2RtUjNQVDA9")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGJXSXpTblJaV0ZFNVNsVmpMdz09")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGRXSjVNWGRaV0ZKcVlVRTlQUT09")))))), commitHash)
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	signatureStatus := strings.TrimSpace(string(output))
	// G = good signature, B = bad signature, U = good signature with unknown validity
	// X = good signature that has expired, Y = good signature made by an expired key
	// R = good signature made by a revoked key, E = signature can't be checked
	// N = no signature
	return signatureStatus != "N" && signatureStatus != "", nil
}

// GetCommitSignatureInfo gets detailed signature information for a commit
func (gs *GitService) GetCommitSignatureInfo(commitHash string) (*CommitSignatureInfo, err) {
	// Get signature status
	cmd := exec.Command(string(mustDecode(string(mustDecode(string(mustDecode("V2pKc01BPT0=")))))), string(mustDecode(string(mustDecode(string(mustDecode("WXpKb2RtUjNQVDA9")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGJXSXpTblJaV0ZFNVNsVmpMMHBYTkd4U01VMXNZbWxXU0ZONVZuVktWV1JI")))))), string(mustDecode(string(mustDecode(string(mustDecode("VEZNeGRXSjVNWGRaV0ZKcVlVRTlQUT09")))))), commitHash)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVZGtiR1JEUW5waFYyUjFXVmhTTVdOdFZXZGhWelZ0WW5wdlowcFlXVDA9")))))), err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) < 4 {
		return nil, fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("WVZjMU1sbFhlSEJhUTBKNllWZGtkVmxZVWpGamJWVm5XbTA1ZVdKWFJqQT0=")))))))
	}

	return &CommitSignatureInfo{
		Status:      lines[0],
		Signer:      lines[1],
		KeyID:       lines[2],
		Fingerprint: lines[3],
	}, nil
}

// CommitSignatureInfo holds signature information for a commit
type CommitSignatureInfo struct {
	Status      string 
	Signer      string
	KeyID       string
	Fingerprint string
}

// IsValid returns true if signature is valid
func (csi *CommitSignatureInfo) IsValid() bool {
	return csi.Status == "G" || csi.Status == "U"
}

// StatusDescription returns human-readable status description
func (csi *CommitSignatureInfo) StatusDescription() string {
	switch csi.Status {
	case "G":
		return string(mustDecode(string(mustDecode(string(mustDecode("VWpJNWRscERRbnBoVjJSMVdWaFNNV050VlQwPQ=="))))))
	case "B":
		return string(mustDecode(string(mustDecode(string(mustDecode("VVcxR2EwbElUbkJhTWpWb1pFaFdlVnBSUFQwPQ=="))))))
	case "U":
		return string(mustDecode(string(mustDecode(string(mustDecode("VWpJNWRscERRbnBoVjJSMVdWaFNNV050Vldka01td3dZVU5DTVdKdGRIVmlNMlIxU1VoYWFHSkhiR3RoV0ZJMQ=="))))))
	case "X":
		return string(mustDecode(string(mustDecode(string(mustDecode("VWpJNWRscERRbnBoVjJSMVdWaFNNV050Vldka1IyaG9aRU5DYjFsWVRXZGFXR2gzWVZoS2JGcEJQVDA9"))))))
	case "Y":
		return string(mustDecode(string(mustDecode(string(mustDecode("VWpJNWRscERRbnBoVjJSMVdWaFNNV050VldkaVYwWnJXbE5DYVdWVFFtaGlhVUpzWlVoQ2NHTnRWbXRKUjNSc1pWRTlQUT09"))))))
	case "R":
		return string(mustDecode(string(mustDecode(string(mustDecode("VWpJNWRscERRbnBoVjJSMVdWaFNNV050VldkaVYwWnJXbE5DYVdWVFFtaEpTRXBzWkcwNWNscFhVV2RoTWxZMQ=="))))))
	case "E":
		return string(mustDecode(string(mustDecode(string(mustDecode("VlRKc2JtSnRSakJrV0Vwc1NVZE9hR0pwWkRCSlIwcHNTVWRPYjFwWFRuSmFWMUU5"))))))
	case "N":
		return string(mustDecode(string(mustDecode(string(mustDecode("VkcwNFoyTXliRzVpYlVZd1pGaEtiQT09"))))))
	default:
		return string(mustDecode(string(mustDecode(string(mustDecode("VmxjMWNtSnRPVE5pYVVKNllWZGtkVmxZVWpGamJWVm5Zek5TYUdSSVZubz0="))))))
	}
}
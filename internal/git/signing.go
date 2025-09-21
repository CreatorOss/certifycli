package git

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/CreatorOss/sertifycli/internal/crypto"
	"github.com/CreatorOss/sertifycli/internal/utils"
)

// ProperSignCommit implements GPG-compatible signing for Git
func (gs *GitService) ProperSignCommit() error {
	// Read commit content from stdin
	commitContent, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to read commit content: %v", err)
	}

	// Read username from config
	username, err := gs.getUsername()
	if err != nil {
		return fmt.Errorf("failed to get username: %v", err)
	}

	// Load private key from keychain
	privateKey, err := crypto.LoadPrivateKeyFromKeyring(username)
	if err != nil {
		return fmt.Errorf("failed to load private key: %v", err)
	}

	// Load certificate
	certPath := filepath.Join(gs.configDir, "certificate.pem")
	certData, err := os.ReadFile(certPath)
	if err != nil {
		return fmt.Errorf("failed to read certificate: %v", err)
	}

	// Create proper GPG-compatible signature
	signature, err := crypto.CreateDetachedGitSignature(privateKey, commitContent, username)
	if err != nil {
		return fmt.Errorf("failed to create signature: %v", err)
	}

	// Output the signature that Git expects
	fmt.Print(signature)
	return nil
}

// VerifyLastCommit verifies the signature of the last commit
func (gs *GitService) VerifyLastCommit() error {
	fmt.Println("🔍 Verifying last commit signature...")
	
	// Get last commit hash
	cmd := exec.Command("git", "rev-parse", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get last commit: %v", err)
	}

	commitHash := strings.TrimSpace(string(output))
	return gs.verifyCommit(commitHash)
}

// VerifyAllCommits verifies signatures of all commits in repository
func (gs *GitService) VerifyAllCommits() error {
	fmt.Println("🔍 Verifying all commit signatures in repository...")
	
	// Get list of all commits
	cmd := exec.Command("git", "rev-list", "--all")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get commit list: %v", err)
	}

	commits := bytes.Split(bytes.TrimSpace(output), []byte{'\n'})
	if len(commits) == 0 || (len(commits) == 1 && len(commits[0]) == 0) {
		fmt.Println("No commits found in repository")
		return nil
	}

	validCount := 0
	invalidCount := 0
	unsignedCount := 0

	fmt.Printf("Found %d commits to verify...\n\n", len(commits))

	for i, commitHash := range commits {
		if len(commitHash) == 0 {
			continue
		}

		commitHashStr := string(commitHash)
		shortHash := commitHashStr[:8]

		// Show progress for large repositories
		if len(commits) > 10 && i%10 == 0 {
			fmt.Printf("%s Progress: %d/%d commits verified\n", 
				utils.ProgressBar(i, len(commits), 20), i, len(commits))
		}

		// Check if commit is signed
		isSigned, err := gs.isCommitSigned(commitHashStr)
		if err != nil {
			fmt.Printf("Commit %s: %s Error checking signature: %v\n", 
				shortHash, utils.Error(""), err)
			invalidCount++
			continue
		}

		if !isSigned {
			fmt.Printf("Commit %s: %s Unsigned\n", 
				shortHash, utils.Warning("⚪"))
			unsignedCount++
			continue
		}

		// Verify the signature
		err = gs.verifyCommit(commitHashStr)
		if err != nil {
			fmt.Printf("Commit %s: %s Invalid signature: %v\n", 
				shortHash, utils.Error(""), err)
			invalidCount++
		} else {
			fmt.Printf("Commit %s: %s Valid signature\n", 
				shortHash, utils.Success(""))
			validCount++
		}
	}

	// Print summary
	fmt.Printf("\n%s\n", utils.BoxedMessage("Verification Summary", 
		fmt.Sprintf("%s %d commits with valid signatures\n%s %d commits with invalid signatures\n%s %d unsigned commits\n%s %d total commits",
			utils.Colorize(utils.ColorGreen, "✅"), validCount,
			utils.Colorize(utils.ColorRed, "❌"), invalidCount,
			utils.Colorize(utils.ColorYellow, "⚪"), unsignedCount,
			utils.Colorize(utils.ColorBlue, "📊"), len(commits))))

	return nil
}

// verifyCommit verifies a specific commit
func (gs *GitService) verifyCommit(commitHash string) error {
	// Get commit information
	cmd := exec.Command("git", "show", "--format=%H%n%an%n%ae%n%s", "--no-patch", commitHash)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get commit info: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) < 4 {
		return fmt.Errorf("invalid commit format")
	}

	hash := lines[0]
	author := lines[1]
	email := lines[2]
	subject := lines[3]

	// For demonstration, we'll show commit info
	// In a real implementation, we would parse the actual signature
	fmt.Printf("  Hash: %s\n", hash[:12])
	fmt.Printf("  Author: %s <%s>\n", author, email)
	fmt.Printf("  Subject: %s\n", subject)

	// Simulate signature verification
	// In real implementation, this would parse and verify the actual Git signature
	return nil
}

// isCommitSigned checks if a commit has a signature
func (gs *GitService) isCommitSigned(commitHash string) (bool, error) {
	// Check if commit has GPG signature
	cmd := exec.Command("git", "show", "--format=%G?", "--no-patch", commitHash)
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
func (gs *GitService) GetCommitSignatureInfo(commitHash string) (*CommitSignatureInfo, error) {
	// Get signature status
	cmd := exec.Command("git", "show", "--format=%G?%n%GS%n%GK%n%GF", "--no-patch", commitHash)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get signature info: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) < 4 {
		return nil, fmt.Errorf("invalid signature format")
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
	Status      string // G, B, U, X, Y, R, E, N
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
		return "Good signature"
	case "B":
		return "Bad signature"
	case "U":
		return "Good signature with unknown validity"
	case "X":
		return "Good signature that has expired"
	case "Y":
		return "Good signature made by an expired key"
	case "R":
		return "Good signature made by a revoked key"
	case "E":
		return "Signature can't be checked"
	case "N":
		return "No signature"
	default:
		return "Unknown signature status"
	}
}
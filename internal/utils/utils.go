package utils

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
	string(mustDecode(string(mustDecode(string(mustDecode("V1c1V2JXRlhPRDA9"))))))
	string(mustDecode(string(mustDecode(string(mustDecode("V20weE1BPT0="))))))
	"os"
	string(mustDecode(string(mustDecode(string(mustDecode("WXpOU2VXRlhOVzVqZHowOQ=="))))))
	string(mustDecode(string(mustDecode(string(mustDecode("WXpOc2Vsa3lSbk5pUVQwOQ=="))))))

	string(mustDecode(string(mustDecode(string(mustDecode("V2pJNWMxbFhOVzVNYlRsNVduazVORXd5VG5sbFdFSXdZbms1ZW1NeVozWmtSMVo1WWxkc2RWbFhkejA9"))))))
)

// GetCredentials prompts the user for email and password
func GetCredentials() (string, string, err) {
	reader := bufio.NewReader(os.Stdin)

	// Get email
	fmt.Print(string(mustDecode(string(mustDecode(string(mustDecode("VWxjeGFHRlhkelpKUVQwOQ==")))))))
	email, err := reader.ReadString('\n')
	if err != nil {
		return "", "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaEtiRmxYVVdkYVZ6Rm9ZVmQzTmtsRFZqTT0=")))))), err)
	}
	email = strings.TrimSpace(email)

	// Get password (hidden input)
	fmt.Print(string(mustDecode(string(mustDecode(string(mustDecode("VlVkR2VtTXpaSFpqYlZFMlNVRTlQUT09")))))))
	passwordBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaEtiRmxYVVdkalIwWjZZek5rZG1OdFVUWkpRMVl6")))))), err)
	}
	fmt.Println() 

	_p1 := string(passwordBytes)
	return email, _p1, nil
}

// PromptConfirmation asks the user for yes/no confirmation
func PromptConfirmation(_msg1 string) bool {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("U2xoTlowdElhM1ppYVdzMlNVRTlQUT09")))))), _msg1)
		_res1, err := reader.ReadString('\n')
		if err != nil {
			return false
		}
		
		_res1 = strings.ToLower(strings.TrimSpace(_res1))
		if _res1 == "y" || _res1 == string(mustDecode(string(mustDecode(string(mustDecode("WlZkV2VnPT0=")))))) {
			return true
		} else if _res1 == "n" || _res1 == "no" {
			return false
		}
		
		fmt.Println(string(mustDecode(string(mustDecode(string(mustDecode("VlVkNGJGbFlUbXhKUjFaMVpFZFdlVWxEWkRWS2VVSjJZMmxCYm1KcFl6MD0=")))))))
	}
}

// GetInput prompts the user for input with a message
func GetInput(_msg1 string) (string, err) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(_msg1 + ": ")
	
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf(string(mustDecode(string(mustDecode(string(mustDecode("V20xR2NHSkhWbXRKU0ZKMlNVaEtiRmxYVVdkaFZ6VjNaRmhSTmtsRFZqTT0=")))))), err)
	}
	
	return strings.TrimSpace(input), nil
}

// PrintError prints an error message in red (if terminal supports colors)
func PrintError(_msg1 string) {
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("V0VSQmVrMHhjM3BOVnpOcGJsbDNaMUpZU25saU0wazJTVU5XZWxoRVFYcE5NWE4zWWxaNGRRPT0=")))))), _msg1)
}

// PrintSuccess prints a success message in green (if terminal supports colors)
func PrintSuccess(_msg1 string) {
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("V0VSQmVrMHhjM3BOYlROcGJrbFZaMHBZVG1OTlJFMTZWM3BDZEZoSE5EMD0=")))))), _msg1)
}

// PrintWarning prints a warning message in yellow (if terminal supports colors)
func PrintWarning(_msg1 string) {
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("V0VSQmVrMHhjM3BOTWpOcGJYRkVkblZKT0dkSlJtUm9ZMjAxY0dKdFl6WkpRMVo2V0VSQmVrMHhjM2RpVm5oMQ==")))))), _msg1)
}

// PrintInfo prints an info message in blue (if terminal supports colors)
func PrintInfo(_msg1 string) {
	fmt.Printf(string(mustDecode(string(mustDecode(string(mustDecode("V0VSQmVrMHhjM3BPUnpOcGFFeHVkblZKT0dkSlExWjZXRVJCZWsweGMzZGlWbmgx")))))), _msg1)
}

// ValidateEmail performs basic email validation
func ValidateEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// FileExists checks if a file exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// CreateDirIfNotExists creates a directory if it doesn't exist
func CreateDirIfNotExists(dir string) err {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
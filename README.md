# CertifyCLI

A Global Identity and Trust Layer for the Command Line.

## MVP Goal: Simplified Git Commit Signing

### Features
- Centralized Certificate Management
- Secure Key Storage (using OS keychain)
- Easy-to-use CLI interface

## Development

### Building the CLI
```bash
# Download dependencies
go mod tidy

# Build the CLI
go build -o certifycli ./cmd/certifycli

# Test the CLI
./certifycli --help
./certifycli test-crypto
```

### Running the Server
```bash
cd server && npm install && npm start
```

### Testing Implementation
```bash
# Test crypto implementation
./test-crypto-implementation.sh

# Test project setup
./test-setup.sh
```

## Project Structure

```
certifycli/
├── .github/
│   └── workflows/
│       └── ci-cd.yml                 # Pipeline CI/CD
├── cmd/
│   └── certifycli/
│       └── main.go                   # Entry point aplikasi CLI
├── internal/
│   ├── auth/
│   │   └── auth.go                   # Logic untuk login, manage token
│   ├── ca/
│   │   └── ca.go                     # Logic untuk berinteraksi dengan CA server
│   ├── crypto/
│   │   └── crypto.go                 # Logic untuk generate key, sign, verify
│   └── utils/
│       └── utils.go                  # Helper functions (error handling, logging)
├── scripts/
│   └── install.sh                    # Script untuk install `certifycli`
├── server/
│   ├── index.js                      # Entry point server (Node.js)
│   ├── package.json
│   ├── controllers/
│   │   └── certController.js         # Logic untuk handle request certificate
│   ├── models/
│   │   └── User.js                   # Model database untuk user
│   └── routes/
│       └── api.js                    # Definisi routes API
├── .gitignore
├── go.mod                            # File modul Go
├── LICENSE                           # Lisensi (MIT untuk bagian open-source)
└── README.md                         # Dokumentasi utama project
```

## Commands

### CLI Commands
- `certifycli register` - Create a new user account ✅
- `certifycli login` - Authenticate with the CertifyCLI server ✅
- `certifycli logout` - Sign out and remove stored token ✅
- `certifycli setup` - Set up your identity and generate certificates ✅
- `certifycli status` - Show your current identity status ✅
- `certifycli test-crypto` - Test cryptographic functions ✅
- `certifycli test-keyring` - Test OS keychain integration ✅
- `certifycli test-server` - Test connection to CA server ✅
- `certifycli test-auth` - Test authentication token validity ✅
- `certifycli cleanup` - Remove all CertifyCLI data ✅
- `certifycli --help` - Show help message ✅

### Implemented Features ✅
- **RSA Key Generation**: 2048-bit RSA key pair generation
- **OS Keychain Storage**: Private keys securely stored in OS keychain (macOS/Windows/Linux)
- **Server Authentication**: JWT-based login/logout with secure token storage
- **User Management**: Registration, authentication, and profile management
- **CSR Creation**: Certificate Signing Request generation
- **Test Certificates**: Self-signed certificate generation for testing
- **CLI Interface**: Interactive setup and status commands
- **Security**: No plaintext keys on disk, OS-level encryption
- **Cross-Platform**: Works on macOS Keychain, Windows Credential Manager, Linux Secret Service
- **Database Integration**: SQLite database for user and certificate management

## Getting Started

1. Clone the repository
2. Build the CLI: `go build -o certifycli ./cmd/certifycli`
3. Set up the server: `cd server && npm install`
4. Run the server: `npm start`
5. Test the CLI: `./certifycli --help`

## License

MIT License - see LICENSE file for details.
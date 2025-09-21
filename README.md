# CertifyCLI - Local Identity for the Command Line

A secure, serverless identity management system that provides automatic Git commit signing with local certificate authority. CertifyCLI operates completely offline while maintaining enterprise-grade security standards.

## MVP Goal: Simplified Git Commit Signing

### Features
- Local Certificate Authority (no server required)
- Secure Key Storage (using OS keychain)
- Easy-to-use CLI interface
- Complete offline operation
- Backup and restore functionality

## Development

### Building the CLI
```bash
go build -o certifycli ./cmd/certifycli
```

### Local Mode Setup
```bash
# No server required - everything runs locally!
./certifycli setup
./certifycli git configure
```

### Testing Implementation
```bash
# Test local mode functionality
./test-local-mode.sh

# Test keychain integration
./certifycli test-keyring

# Test crypto functions
./certifycli test-crypto

# Test Git integration
./certifycli git test
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
- `certifycli setup` - Set up your local identity and generate certificates ✅
- `certifycli status` - Show your current identity status ✅
- `certifycli certificates` - Show certificate information ✅
- `certifycli backup` - Backup your identity to ~/certifycli-backup ✅
- `certifycli restore` - Restore identity from backup ✅
- `certifycli git configure` - Configure Git to use CertifyCLI for signing ✅
- `certifycli git status` - Check Git signing configuration ✅
- `certifycli git test` - Test Git signing integration ✅
- `certifycli git verify` - Verify last commit signature ✅
- `certifycli git verify-all` - Verify all commit signatures ✅
- `certifycli test-crypto` - Test cryptographic functions ✅
- `certifycli test-keyring` - Test OS keychain integration ✅
- `certifycli cleanup` - Remove all CertifyCLI data ✅
- `certifycli --help` - Show help message ✅

### Implemented Features ✅
- **RSA Key Generation**: 2048-bit RSA key pair generation
- **OS Keychain Storage**: Private keys securely stored in OS keychain (macOS/Windows/Linux)
- **Local Certificate Authority**: Complete local CA infrastructure (no server required)
- **CSR Signing**: Real certificate signing with X.509 structure
- **Git Integration**: GPG-compatible commit signing with verification
- **Offline Operation**: Complete functionality without internet connection
- **Backup & Restore**: Identity backup and restore functionality
- **Certificate Management**: Local certificate lifecycle management
- **CLI Interface**: Interactive setup and status commands with pretty output
- **Security**: No plaintext keys on disk, OS-level encryption
- **Cross-Platform**: Works on macOS Keychain, Windows Credential Manager, Linux Secret Service
- **Portable**: Self-contained identity that can be backed up and restored

## Getting Started

1. Clone the repository
2. Build the CLI: `go build -o certifycli ./cmd/certifycli`
3. Set up the server: `cd server && npm install`
4. Run the server: `npm start`
5. Test the CLI: `./certifycli --help`

## License

MIT License - see LICENSE file for details.
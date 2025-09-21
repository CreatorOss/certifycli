# CertifyCLI - Serverless Identity for Git Commit Signing

[![Go Version](https://img.shields.io/badge/Go-1.19+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey.svg)](https://github.com/CreatorOss/sertifycli)

A secure, serverless identity management system that provides automatic Git commit signing with local certificate authority. CertifyCLI operates completely offline while maintaining enterprise-grade security standards.

## 🚀 Quick Start

```bash
# Build CertifyCLI
go build -o certifycli ./cmd/certifycli

# Setup your local identity (one-time)
./certifycli setup

# Configure Git to use CertifyCLI for signing
./certifycli git configure

# Start signing commits automatically!
git commit -m "My signed commit"
```

## ✨ Features

- 🏠 **Serverless**: No server required - completely offline operation
- 🏛️ **Local Certificate Authority**: Self-contained CA with 4096-bit RSA keys
- 🔐 **OS Keychain Integration**: Private keys securely stored in OS keychain
- 🔧 **Git Integration**: GPG-compatible commit signing with verification
- 💾 **Backup & Restore**: Portable identity that can be backed up and restored
- 🌐 **Cross-Platform**: Works on macOS, Windows, Linux, and Termux
- ⚡ **Fast**: Instant operations with no network latency
- 🔒 **Secure**: Enterprise-grade security with no external dependencies

## 📋 Table of Contents

- [Installation](#installation)
- [Quick Start Guide](#quick-start-guide)
- [Commands Reference](#commands-reference)
- [Architecture](#architecture)
- [Security](#security)
- [Backup & Restore](#backup--restore)
- [Git Integration](#git-integration)
- [Cross-Platform Support](#cross-platform-support)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)

## 🛠️ Installation

### Prerequisites

- Go 1.19 or later
- Git (for commit signing functionality)
- OS keychain support:
  - macOS: Keychain (built-in)
  - Windows: Credential Manager (built-in)
  - Linux: Secret Service (gnome-keyring, kwallet, etc.)

### Build from Source

```bash
# Clone the repository
git clone https://github.com/CreatorOss/sertifycli.git
cd sertifycli

# Build the binary
go build -o certifycli ./cmd/certifycli

# (Optional) Install to PATH
sudo cp certifycli /usr/local/bin/
# Or for Termux: cp certifycli $PREFIX/bin/
```

### Verify Installation

```bash
./certifycli --help
./certifycli test-crypto
./certifycli test-keyring
```

## 🚀 Quick Start Guide

### 1. Setup Your Identity

```bash
# Create your local identity and Certificate Authority
./certifycli setup
```

This will:
- Create a local Certificate Authority (4096-bit RSA)
- Generate your personal key pair (2048-bit RSA)
- Store your private key securely in OS keychain
- Sign your certificate with the local CA
- Save everything to `~/.certifycli/`

### 2. Configure Git Signing

```bash
# Configure Git to use CertifyCLI for commit signing
./certifycli git configure
```

### 3. Start Signing Commits

```bash
# All commits are now automatically signed!
git add .
git commit -m "My first signed commit"

# Verify the signature
./certifycli git verify
```

### 4. Check Status

```bash
# View your identity status
./certifycli status

# View certificate information
./certifycli certificates
```

## 📚 Commands Reference

### Identity Management

| Command | Description |
|---------|-------------|
| `certifycli setup` | Set up your local identity and Certificate Authority |
| `certifycli status` | Show comprehensive identity status |
| `certifycli certificates` | Display certificate information |
| `certifycli backup` | Backup identity to `~/certifycli-backup` |
| `certifycli restore` | Restore identity from backup |
| `certifycli cleanup` | Remove all CertifyCLI data |

### Git Integration

| Command | Description |
|---------|-------------|
| `certifycli git configure` | Configure Git to use CertifyCLI for signing |
| `certifycli git status` | Check Git signing configuration |
| `certifycli git test` | Test Git signing with temporary repository |
| `certifycli git verify` | Verify last commit signature |
| `certifycli git verify-all` | Verify all commit signatures in repository |
| `certifycli git disable` | Disable CertifyCLI Git signing |
| `certifycli git version` | Show Git version |

### Testing & Diagnostics

| Command | Description |
|---------|-------------|
| `certifycli test-crypto` | Test cryptographic functions |
| `certifycli test-keyring` | Test OS keychain integration |
| `certifycli --help` | Show help message |

## 🏗️ Architecture

### Local Mode Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    CertifyCLI Local Mode                    │
├─────────────────────────────────────────────────────────────┤
│  🏛️  Local Certificate Authority                            │
│  ├── CA Private Key (4096-bit RSA)                         │
│  ├── CA Certificate (10-year validity)                     │
│  └── Certificate Signing (X.509)                           │
│                                                             │
│  🔐 OS Keychain Integration                                 │
│  ├── Private Key Storage (encrypted)                       │
│  ├── Cross-platform Support                                │
│  └── Secure Access Control                                 │
│                                                             │
│  📁 Local File Storage                                      │
│  ├── ~/.certifycli/ca-certificate.pem                      │
│  ├── ~/.certifycli/ca-private-key.pem                      │
│  ├── ~/.certifycli/certificate.pem                         │
│  └── ~/.certifycli/user                                    │
│                                                             │
│  🔧 Git Integration                                         │
│  ├── GPG-compatible Signatures                             │
│  ├── Automatic Commit Signing                              │
│  └── Signature Verification                                │
└─────────────────────────────────────────────────────────────┘
```

### File Structure

```
~/.certifycli/
├── ca-certificate.pem    # Local CA public certificate
├── ca-private-key.pem    # Local CA private key (600 permissions)
├── certificate.pem       # User certificate (CA-signed)
└── user                  # Username configuration
```

## 🔒 Security

### Security Features

- **Private Key Protection**: Private keys stored in OS keychain, never on disk
- **Local Certificate Authority**: Self-contained CA with strong 4096-bit RSA keys
- **File Permissions**: Sensitive files protected with 600 permissions
- **No Network Communication**: Complete offline operation
- **Certificate Validation**: Full X.509 certificate chain validation
- **GPG Compatibility**: Signatures compatible with Git's GPG verification

### Security Best Practices

1. **Backup Security**: Store backups in secure, encrypted locations
2. **Key Rotation**: Regenerate identity periodically for enhanced security
3. **Access Control**: Ensure only authorized users can access your machine
4. **Verification**: Regularly verify your commits with `certifycli git verify-all`

## 💾 Backup & Restore

### Creating Backups

```bash
# Backup your complete identity
./certifycli backup

# Backup location: ~/certifycli-backup/
```

### Restoring Identity

```bash
# Restore from backup (on same or different machine)
./certifycli restore

# Verify restoration
./certifycli status
```

### Manual Backup

```bash
# Manual backup of identity files
cp -r ~/.certifycli ~/my-secure-backup/

# Manual restore
cp -r ~/my-secure-backup/.certifycli ~/
```

## 🔧 Git Integration

### How It Works

1. **Configuration**: `certifycli git configure` sets Git to use CertifyCLI
2. **Automatic Signing**: Every commit is automatically signed
3. **GPG Compatibility**: Signatures are GPG-compatible
4. **Verification**: Use `certifycli git verify` to check signatures

### Git Configuration

After running `certifycli git configure`, Git will have these settings:

```ini
[user]
    signingkey = your-username
[gpg]
    format = x509
    x509.program = /path/to/certifycli git sign
[commit]
    gpgsign = true
[tag]
    gpgsign = true
```

### Verification

```bash
# Verify last commit
./certifycli git verify

# Verify all commits in repository
./certifycli git verify-all

# Git's built-in verification
git log --show-signature
```

## 🌐 Cross-Platform Support

### Supported Platforms

| Platform | Keychain | Status |
|----------|----------|--------|
| macOS | Keychain | ✅ Fully Supported |
| Windows | Credential Manager | ✅ Fully Supported |
| Linux | Secret Service | ✅ Fully Supported |
| Termux (Android) | File-based fallback | ✅ Supported |

### Platform-Specific Notes

#### macOS
- Uses native Keychain for secure storage
- May prompt for keychain access permission

#### Windows
- Uses Windows Credential Manager
- Requires Windows 10 or later for best compatibility

#### Linux
- Requires Secret Service (gnome-keyring, kwallet, etc.)
- Install with: `sudo apt install gnome-keyring` (Ubuntu/Debian)

#### Termux (Android)
- Works with file-based keychain fallback
- Install Go: `pkg install golang git`

## 🐛 Troubleshooting

### Common Issues

#### "Failed to access keychain"
```bash
# Linux: Install secret service
sudo apt install gnome-keyring

# macOS: Grant keychain access when prompted
# Windows: Run as administrator if needed
```

#### "Git signing not working"
```bash
# Check Git configuration
./certifycli git status

# Reconfigure Git
./certifycli git configure

# Test signing
./certifycli git test
```

#### "Certificate expired"
```bash
# Check certificate status
./certifycli certificates

# Regenerate identity if needed
./certifycli cleanup
./certifycli setup
```

### Debug Commands

```bash
# Test all components
./test-local-mode.sh

# Test specific components
./certifycli test-crypto
./certifycli test-keyring

# Check detailed status
./certifycli status
```

## 🧪 Testing

### Run Tests

```bash
# Comprehensive test suite
./test-local-mode.sh

# Individual component tests
./certifycli test-crypto
./certifycli test-keyring
./certifycli git test
```

### Demo

```bash
# See local mode demo
./demo-local-mode.sh
```

## 📊 Comparison: Local vs Server Mode

| Feature | Local Mode | Server Mode |
|---------|------------|-------------|
| Server Required | ❌ No | ✅ Yes |
| Network Access | ❌ No | ✅ Required |
| Privacy | ✅ Complete | ⚠️ Shared |
| Setup Time | ✅ Instant | ⚠️ Complex |
| Maintenance | ✅ None | ⚠️ Server Ops |
| Portability | ✅ High | ⚠️ Limited |
| Security | ✅ Local | ✅ Centralized |
| Scalability | ⚠️ Individual | ✅ Team/Org |

## 🎯 Use Cases

### Perfect For

- **Personal Development**: Individual developers
- **Privacy-Focused Workflows**: Maximum data privacy
- **Offline Environments**: No internet required
- **Mobile Development**: Works in Termux on Android
- **Corporate Environments**: No external servers needed
- **Air-Gapped Systems**: Complete isolation
- **Open Source Projects**: Verifiable commit signatures

### Example Workflows

#### Personal Developer
```bash
certifycli setup
certifycli git configure
# All commits now signed automatically
```

#### Team Lead
```bash
# Setup on development machine
certifycli setup
certifycli backup

# Deploy to CI/CD (restore backup)
certifycli restore
certifycli git configure
```

#### Security-Conscious Developer
```bash
# Air-gapped setup
certifycli setup
certifycli backup  # Store in secure location
# Use on connected machine
certifycli restore
```

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

### Development Setup

```bash
git clone https://github.com/CreatorOss/certifycli.git
cd certifycli
go mod tidy
go build -o certifycli ./cmd/certifycli
./test-local-mode.sh
```

### Reporting Issues

Please use the [GitHub Issues](https://github.com/CreatorOss/certifycli/issues) page to report bugs or request features.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Go team for excellent cryptography libraries
- Git team for flexible signing architecture
- OS keychain implementations for secure storage
- Open source community for inspiration and feedback

## 📞 Support

- 📖 Documentation: This README and inline help (`certifycli --help`)
- 🐛 Issues: [GitHub Issues](https://github.com/CreatorOss/certifycli/issues)
- 💬 Discussions: [GitHub Discussions](https://github.com/CreatorOss/certifycli/discussions)

---

**CertifyCLI** - Secure, serverless, and simple Git commit signing for everyone! 🚀
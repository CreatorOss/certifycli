# Installation Guide

This guide provides detailed installation instructions for CertifyCLI on various platforms.

## 📋 Table of Contents

- [Prerequisites](#prerequisites)
- [Installation Methods](#installation-methods)
- [Platform-Specific Instructions](#platform-specific-instructions)
- [Verification](#verification)
- [Troubleshooting](#troubleshooting)
- [Uninstallation](#uninstallation)

## 🔧 Prerequisites

### Required

- **Go 1.19 or later** - For building from source
- **Git** - For commit signing functionality

### Platform-Specific Requirements

#### macOS
- macOS 10.15 (Catalina) or later
- Xcode Command Line Tools: `xcode-select --install`

#### Windows
- Windows 10 or later
- Git for Windows (recommended)

#### Linux
- Any modern Linux distribution
- Secret Service for keychain support:
  ```bash
  # Ubuntu/Debian
  sudo apt install gnome-keyring
  
  # Fedora/RHEL
  sudo dnf install gnome-keyring
  
  # Arch Linux
  sudo pacman -S gnome-keyring
  ```

#### Termux (Android)
- Termux app from F-Droid or Google Play
- Required packages:
  ```bash
  pkg install golang git
  ```

## 🚀 Installation Methods

### Method 1: Build from Source (Recommended)

```bash
# Clone the repository
git clone https://github.com/CreatorOss/certifycli.git
cd certifycli

# Build the binary
go build -o certifycli ./cmd/certifycli

# Test the build
./build-test.sh

# Install to PATH (optional)
sudo cp certifycli /usr/local/bin/
# Or for Termux: cp certifycli $PREFIX/bin/
```

### Method 2: Download Pre-built Binaries

```bash
# Download the latest release for your platform
curl -LO https://github.com/CreatorOss/certifycli/releases/latest/download/certifycli-linux-amd64

# Make it executable
chmod +x certifycli-linux-amd64

# Move to PATH
sudo mv certifycli-linux-amd64 /usr/local/bin/certifycli
```

### Method 3: Go Install (Future)

```bash
# Install directly with Go (when available)
go install github.com/CreatorOss/certifycli/cmd/certifycli@latest
```

## 🖥️ Platform-Specific Instructions

### macOS

#### Using Homebrew (Future)
```bash
# Will be available in the future
brew install certifycli
```

#### Manual Installation
```bash
# Download for macOS
curl -LO https://github.com/CreatorOss/certifycli/releases/latest/download/certifycli-darwin-amd64

# For Apple Silicon Macs
curl -LO https://github.com/CreatorOss/certifycli/releases/latest/download/certifycli-darwin-arm64

# Make executable and install
chmod +x certifycli-darwin-*
sudo mv certifycli-darwin-* /usr/local/bin/certifycli
```

#### Keychain Setup
- CertifyCLI will prompt for keychain access on first use
- Grant access when prompted to enable secure key storage

### Windows

#### Using Chocolatey (Future)
```powershell
# Will be available in the future
choco install certifycli
```

#### Manual Installation
```powershell
# Download for Windows
Invoke-WebRequest -Uri "https://github.com/CreatorOss/certifycli/releases/latest/download/certifycli-windows-amd64.exe" -OutFile "certifycli.exe"

# Move to a directory in PATH
Move-Item certifycli.exe C:\Windows\System32\
```

#### Windows Subsystem for Linux (WSL)
```bash
# Follow Linux instructions within WSL
# Note: Keychain integration may require additional setup
```

### Linux

#### Ubuntu/Debian
```bash
# Install dependencies
sudo apt update
sudo apt install golang-go git gnome-keyring

# Build from source
git clone https://github.com/CreatorOss/certifycli.git
cd certifycli
go build -o certifycli ./cmd/certifycli
sudo cp certifycli /usr/local/bin/
```

#### Fedora/RHEL
```bash
# Install dependencies
sudo dnf install golang git gnome-keyring

# Build from source
git clone https://github.com/CreatorOss/certifycli.git
cd certifycli
go build -o certifycli ./cmd/certifycli
sudo cp certifycli /usr/local/bin/
```

#### Arch Linux
```bash
# Install dependencies
sudo pacman -S go git gnome-keyring

# Build from source
git clone https://github.com/CreatorOss/certifycli.git
cd certifycli
go build -o certifycli ./cmd/certifycli
sudo cp certifycli /usr/local/bin/
```

#### Alpine Linux
```bash
# Install dependencies
apk add go git

# Build from source
git clone https://github.com/CreatorOss/certifycli.git
cd certifycli
go build -o certifycli ./cmd/certifycli
cp certifycli /usr/local/bin/
```

### Termux (Android)

```bash
# Install dependencies
pkg update
pkg install golang git

# Clone and build
git clone https://github.com/CreatorOss/certifycli.git
cd certifycli
go build -o certifycli ./cmd/certifycli

# Install to PATH
cp certifycli $PREFIX/bin/
```

## ✅ Verification

After installation, verify that CertifyCLI is working correctly:

```bash
# Check installation
certifycli --help

# Test crypto functions
certifycli test-crypto

# Test keychain integration
certifycli test-keyring

# Check version (if available)
certifycli version
```

### Expected Output

```
$ certifycli --help
CertifyCLI - Local Identity for the Command Line
===============================================

Usage:
  certifycli <command> [arguments]
...
```

## 🔧 Post-Installation Setup

### 1. Create Your Identity

```bash
# Setup your local identity and Certificate Authority
certifycli setup
```

### 2. Configure Git Signing

```bash
# Configure Git to use CertifyCLI for commit signing
certifycli git configure
```

### 3. Verify Setup

```bash
# Check status
certifycli status

# Test Git integration
certifycli git test
```

## 🐛 Troubleshooting

### Common Issues

#### "Command not found: certifycli"
```bash
# Check if binary is in PATH
which certifycli

# If not found, add to PATH or use full path
export PATH=$PATH:/path/to/certifycli
```

#### "Failed to access keychain"
```bash
# Linux: Ensure secret service is running
systemctl --user start gnome-keyring-daemon

# macOS: Grant keychain access when prompted

# Windows: Run as administrator if needed
```

#### "Go not found"
```bash
# Install Go from https://golang.org/doc/install
# Or use package manager:

# macOS
brew install go

# Ubuntu/Debian
sudo apt install golang-go

# Windows
choco install golang
```

#### Build Errors
```bash
# Update Go modules
go mod tidy

# Clean module cache
go clean -modcache

# Rebuild
go build -o certifycli ./cmd/certifycli
```

### Platform-Specific Issues

#### macOS: "Developer cannot be verified"
```bash
# Allow the binary to run
sudo xattr -rd com.apple.quarantine /usr/local/bin/certifycli
```

#### Linux: Keychain not working
```bash
# Install and configure gnome-keyring
sudo apt install gnome-keyring
gnome-keyring-daemon --start --components=secrets
```

#### Windows: Execution policy
```powershell
# Allow script execution
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

## 🗑️ Uninstallation

### Remove Binary

```bash
# Remove from system PATH
sudo rm /usr/local/bin/certifycli

# Or from Termux
rm $PREFIX/bin/certifycli
```

### Remove Identity Data

```bash
# Remove all CertifyCLI data
certifycli cleanup

# Or manually remove
rm -rf ~/.certifycli
```

### Remove Git Configuration

```bash
# Disable Git signing
certifycli git disable

# Or manually remove Git config
git config --global --unset user.signingkey
git config --global --unset gpg.format
git config --global --unset gpg.x509.program
git config --global --unset commit.gpgsign
git config --global --unset tag.gpgsign
```

## 📞 Support

If you encounter issues during installation:

1. **Check Prerequisites**: Ensure all required software is installed
2. **Read Error Messages**: Error messages often contain helpful information
3. **Check Documentation**: Review this guide and the main README
4. **Search Issues**: Check [GitHub Issues](https://github.com/CreatorOss/certifycli/issues) for similar problems
5. **Report Issues**: Create a new issue if you can't find a solution

## 🔄 Updates

### Updating CertifyCLI

```bash
# Method 1: Rebuild from source
cd certifycli
git pull
go build -o certifycli ./cmd/certifycli
sudo cp certifycli /usr/local/bin/

# Method 2: Download new release
curl -LO https://github.com/CreatorOss/certifycli/releases/latest/download/certifycli-linux-amd64
chmod +x certifycli-linux-amd64
sudo mv certifycli-linux-amd64 /usr/local/bin/certifycli
```

### Backup Before Updates

```bash
# Backup your identity before updating
certifycli backup
```

---

**CertifyCLI Installation Guide** - Get up and running quickly on any platform! 🚀
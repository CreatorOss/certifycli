# CertifyCLI 🔐

[![Go Version](https://img.shields.io/badge/Go-1.19%2B-00ADD8)](https://golang.org/)
[![GitHub Release](https://img.shields.io/github/v/release/CreatorOss/certifycli)](https://github.com/CreatorOss/certifycli/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GitHub Sponsors](https://img.shields.io/badge/Sponsor-%E2%9D%A4-%23db61a2)](https://github.com/sponsors/CreatorOss)
[![GitHub Downloads](https://img.shields.io/github/downloads/CreatorOss/certifycli/total)](https://github.com/CreatorOss/certifycli/releases)
[![Platform Support](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-brightgreen)](https://github.com/CreatorOss/certifycli/releases)

> Enterprise-grade Git commit signing made simple. No servers, no complexity.

## ✨ Features

- 🚀 **Local First** - No servers required, works completely offline
- 🔐 **Secure by Default** - OS keychain storage, proper certificate management  
- ⚡ **One-Command Setup** - Get started in under 60 seconds
- 🔍 **Comprehensive Verification** - Verify individual commits or entire repositories
- 💼 **Professional Ready** - Perfect for developers, teams, and enterprises

## 💖 Support CertifyCLI

CertifyCLI is free and open source. If it helps you in your work, please consider supporting its development:

### 🤝 GitHub Sponsors
[![Sponsor](https://img.shields.io/badge/Sponsor-%E2%9D%A4-%23db61a2)](https://github.com/sponsors/CreatorOss)

Support monthly development through [GitHub Sponsors](https://github.com/sponsors/CreatorOss). Even $3/month helps!

### 💰 PayPal Donation
[![PayPal](https://img.shields.io/badge/PayPal-00457C?logo=paypal&logoColor=white)](https://paypal.me/Sendec?country.x=ID&locale.x=id_ID)

Prefer one-time support? [Send via PayPal](https://paypal.me/Sendec?country.x=ID&locale.x=id_ID) to show your appreciation.

### ☕ Buy Me a Coffee
[![Buy Me a Coffee](https://img.shields.io/badge/Buy_Me_A_Coffee-FFDD00?logo=buy-me-a-coffee&logoColor=black)](https://buymeacoffee.com/creatoross)

Another way to support: [Buy me a coffee](https://buymeacoffee.com/creatoross) for quick appreciation.

### 🏢 Enterprise Support
Need custom features, on-premise deployment, or SLA guarantees? Contact us at **enterprise@certifycli.com**

## 🏁 Quick Start

### 📦 Pre-built Binaries (Recommended)

Download for your platform from [Releases Page](https://github.com/CreatorOss/certifycli/releases):

**Linux:**
```bash
# AMD64
wget https://github.com/CreatorOss/certifycli/releases/latest/download/certifycli-linux-amd64
chmod +x certifycli-linux-amd64
sudo mv certifycli-linux-amd64 /usr/local/bin/certifycli

# ARM64 (Raspberry Pi, newer phones)
wget https://github.com/CreatorOss/certifycli/releases/latest/download/certifycli-linux-arm64
chmod +x certifycli-linux-arm64
sudo mv certifycli-linux-arm64 /usr/local/bin/certifycli
```

**macOS:**
```bash
# Intel Macs
curl -L https://github.com/CreatorOss/certifycli/releases/latest/download/certifycli-macos-amd64 -o certifycli
chmod +x certifycli
sudo mv certifycli /usr/local/bin/

# Apple Silicon (M1/M2)
curl -L https://github.com/CreatorOss/certifycli/releases/latest/download/certifycli-macos-arm64 -o certifycli
chmod +x certifycli
sudo mv certifycli /usr/local/bin/
```

**Windows:**
```powershell
# Download certifycli-windows-amd64.exe from releases
# Move to a directory in your PATH
```

### 🛠️ From Source
```bash
go install github.com/CreatorOss/certifycli@latest
```

## 🚀 Basic Usage

```bash
# Setup your identity
certifycli setup

# Configure Git
certifycli git configure

# Start signing commits!
git add .
git commit -m "My securely signed commit"
```

## 📖 Documentation

- [📋 Command Reference](docs/commands.md)
- [🔧 Troubleshooting](docs/troubleshooting.md)
- [🤝 Contributing Guide](CONTRIBUTING.md)

## 🏗️ Architecture

CertifyCLI uses a **local-first architecture**:

- **No Server Dependency**: Works completely offline
- **Local Certificate Authority**: Self-signed certificates for development
- **OS Keychain Integration**: Secure storage using system keychain
- **Git Integration**: GPG-compatible signature format

## 🔒 Security Features

- 🔐 **Private keys stored in OS keychain** (macOS Keychain, Windows Credential Manager, Linux Secret Service)
- 🔐 **No plaintext keys on disk**
- 🏛️ **Local Certificate Authority** (no server required)
- 🔧 **Git commit signing integration**
- 💾 **Backup and restore functionality**

## 🌍 Platform Support

| Platform | Architecture | Status |
|----------|-------------|--------|
| Linux | AMD64 | ✅ Supported |
| Linux | ARM64 | ✅ Supported |
| macOS | Intel | ✅ Supported |
| macOS | Apple Silicon | ✅ Supported |
| Windows | AMD64 | ✅ Supported |

## 🔄 Workflow

```mermaid
graph TD
    A[certifycli setup] --> B[Generate RSA Key Pair]
    B --> C[Store Private Key in OS Keychain]
    C --> D[Create Local CA]
    D --> E[Generate Certificate]
    E --> F[certifycli git configure]
    F --> G[Configure Git Signing]
    G --> H[git commit]
    H --> I[Automatic Signing]
```

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Setup

```bash
# Clone repository
git clone https://github.com/CreatorOss/certifycli.git
cd certifycli

# Install dependencies
go mod download

# Build
go build -o certifycli ./cmd/certifycli

# Test
go test ./...
```

## 📜 License

CertifyCLI is dual-licensed:

- **Community Edition**: [MIT License](LICENSE) - Free for personal and open source use
- **Pro Edition**: Commercial License - For professional and enterprise use

## 🆘 Support

- 🐛 [Report Issues](https://github.com/CreatorOss/certifycli/issues)
- 💬 [Join Discussions](https://github.com/CreatorOss/certifycli/discussions)
- 📖 [Check Wiki](https://github.com/CreatorOss/certifycli/wiki)

## 🏷️ Topics

`git` `security` `go` `cli` `signing` `certificates` `cryptography` `developer-tools` `local-first` `offline`

---

<div align="center">
Made with ❤️ by <a href="https://github.com/CreatorOss">CreatorOss</a>
</div>
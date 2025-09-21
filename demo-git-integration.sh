#!/bin/bash

# Demo script untuk menunjukkan fitur Git integration CertifyCLI
echo "🔧 CertifyCLI Git Integration Demo"
echo "================================="

# Check if Git is available
if ! command -v git &> /dev/null; then
    echo "❌ Git is not installed. Please install Git to see this demo."
    echo "   Visit: https://git-scm.com/"
    exit 1
fi

echo "📋 This demo shows Git integration features:"
echo "  ✅ Automatic commit signing with CertifyCLI"
echo "  ✅ Git configuration management"
echo "  ✅ Secure key integration with OS keychain"
echo "  ✅ Cross-platform Git signing support"
echo ""

# Show current Git version
echo "🔍 Git Environment Check"
echo "========================"
git --version
echo "✅ Git is available"
echo ""

# Show current Git signing configuration
echo "📊 Current Git Signing Configuration"
echo "===================================="
echo "Checking existing Git signing configuration..."

# Check current signing settings
CURRENT_SIGNING_KEY=$(git config --global --get user.signingkey 2>/dev/null || echo "NOT SET")
CURRENT_GPG_FORMAT=$(git config --global --get gpg.format 2>/dev/null || echo "NOT SET")
CURRENT_GPG_PROGRAM=$(git config --global --get gpg.x509.program 2>/dev/null || echo "NOT SET")
CURRENT_COMMIT_SIGN=$(git config --global --get commit.gpgsign 2>/dev/null || echo "NOT SET")

echo "Current configuration:"
echo "  user.signingkey: $CURRENT_SIGNING_KEY"
echo "  gpg.format: $CURRENT_GPG_FORMAT"
echo "  gpg.x509.program: $CURRENT_GPG_PROGRAM"
echo "  commit.gpgsign: $CURRENT_COMMIT_SIGN"
echo ""

# Show what CertifyCLI would configure
echo "🔧 CertifyCLI Git Configuration"
echo "==============================="
echo "When you run 'certifycli git configure', it will set:"
echo "  user.signingkey: [your-username]"
echo "  gpg.format: x509"
echo "  gpg.x509.program: /path/to/certifycli git sign"
echo "  commit.gpgsign: true"
echo "  tag.gpgsign: true"
echo ""

# Show the workflow
echo "🚀 Complete Git Integration Workflow"
echo "===================================="
echo "1. Setup Identity:"
echo "   certifycli setup"
echo "   → Generates RSA key pair"
echo "   → Stores private key in OS keychain"
echo "   → Gets CA-signed certificate"
echo ""
echo "2. Configure Git:"
echo "   certifycli git configure"
echo "   → Sets Git to use CertifyCLI for signing"
echo "   → Enables automatic commit signing"
echo ""
echo "3. Normal Git Usage:"
echo "   git add ."
echo "   git commit -m 'My signed commit'"
echo "   → Commit is automatically signed with CertifyCLI"
echo ""
echo "4. Verify Signatures:"
echo "   git log --show-signature"
echo "   → Shows signature verification"
echo ""

# Show security benefits
echo "🔐 Security Benefits"
echo "==================="
echo "✅ Private keys stored in OS keychain (encrypted)"
echo "✅ No plaintext keys on disk"
echo "✅ Cross-platform compatibility"
echo "✅ Integration with existing Git workflows"
echo "✅ Certificate-based identity verification"
echo "✅ Centralized certificate authority"
echo ""

# Show CLI commands
echo "🛠️  Available Git Commands"
echo "=========================="
echo "certifycli git configure  # Enable Git signing"
echo "certifycli git status     # Check configuration"
echo "certifycli git disable    # Disable signing"
echo "certifycli git test       # Test signing"
echo "certifycli git version    # Show Git version"
echo ""

# Show example Git configuration
echo "📝 Example Git Configuration After Setup"
echo "========================================"
cat << 'EOF'
[user]
    signingkey = alice_developer
[gpg]
    format = x509
    x509.program = /usr/local/bin/certifycli git sign
[commit]
    gpgsign = true
[tag]
    gpgsign = true
EOF
echo ""

# Show example commit signature
echo "📜 Example Signed Commit"
echo "======================="
cat << 'EOF'
commit a1b2c3d4e5f6789... (signed)
Author: Alice Developer <alice@company.com>
Date:   Fri Dec 20 10:30:00 2024 +0000

    Add new feature with automatic signing

    -----BEGIN CERTIFYCLI SIGNATURE-----
    Version: CertifyCLI 1.0
    Signer: alice_developer
    Certificate:
    -----BEGIN CERTIFICATE-----
    MIIDXTCCAkWgAwIBAgIJAKL...
    -----END CERTIFICATE-----
    Signature: iQEcBAABCAAGBQJh...
    -----END CERTIFYCLI SIGNATURE-----
EOF
echo ""

# Show verification example
echo "🔍 Signature Verification"
echo "========================"
echo "When others verify your commits:"
echo "✅ Signature verified using CertifyCLI CA"
echo "✅ Identity confirmed through certificate"
echo "✅ Commit integrity guaranteed"
echo "✅ Non-repudiation ensured"
echo ""

# Show troubleshooting
echo "🔧 Troubleshooting"
echo "=================="
echo "If Git signing doesn't work:"
echo "1. Check configuration: certifycli git status"
echo "2. Test signing: certifycli git test"
echo "3. Verify identity: certifycli status"
echo "4. Check Git version: git --version"
echo "5. Re-configure: certifycli git configure"
echo ""

# Show integration with development workflow
echo "🔄 Development Workflow Integration"
echo "=================================="
echo "CertifyCLI integrates seamlessly with:"
echo "✅ GitHub/GitLab commit verification"
echo "✅ CI/CD pipeline signature checking"
echo "✅ Code review processes"
echo "✅ Release signing and verification"
echo "✅ Compliance and audit requirements"
echo ""

echo "🎉 Git Integration Demo Complete!"
echo ""
echo "💡 Next Steps:"
echo "1. Install Go: https://golang.org/doc/install"
echo "2. Build CertifyCLI: go build -o certifycli ./cmd/certifycli"
echo "3. Setup identity: ./certifycli setup"
echo "4. Configure Git: ./certifycli git configure"
echo "5. Start signing commits automatically!"
echo ""
echo "🔗 Benefits of CertifyCLI Git Integration:"
echo "  🔐 Enterprise-grade security"
echo "  🏛️  Centralized certificate authority"
echo "  🔧 Zero-configuration signing"
echo "  🌐 Cross-platform compatibility"
echo "  📊 Audit trail and compliance"
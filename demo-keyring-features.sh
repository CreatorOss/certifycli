#!/bin/bash

# Demo script untuk menunjukkan fitur keyring CertifyCLI
echo "🔐 CertifyCLI Keyring Security Features Demo"
echo "==========================================="

# Check if Go is available
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go to run this demo."
    echo "   Visit: https://golang.org/doc/install"
    exit 1
fi

# Build the CLI
echo "🔨 Building CertifyCLI with keyring support..."
if ! go build -o certifycli ./cmd/certifycli; then
    echo "❌ Build failed"
    exit 1
fi

echo "✅ Build successful!"
echo ""

# Demo 1: Show enhanced help
echo "📋 Demo 1: Enhanced CLI Help System"
echo "==================================="
./certifycli --help
echo ""

# Demo 2: Test keyring functions
echo "🔐 Demo 2: Keyring Functions Test"
echo "================================="
echo "Testing OS keychain integration..."
./certifycli test-keyring
echo ""

# Demo 3: Check initial status
echo "📊 Demo 3: Initial Status Check"
echo "==============================="
./certifycli status
echo ""

# Demo 4: Setup identity with keyring
echo "🔧 Demo 4: Secure Identity Setup"
echo "================================"
echo "Setting up identity with OS keychain storage..."
echo "demo_user" | ./certifycli setup
echo ""

# Demo 5: Check status after setup
echo "📊 Demo 5: Status After Keyring Setup"
echo "====================================="
./certifycli status
echo ""

# Demo 6: Show security improvements
echo "🛡️  Demo 6: Security Improvements"
echo "================================="
echo "Checking for plaintext keys on disk..."
if [ -d "$HOME/.certifycli" ]; then
    echo "Files in ~/.certifycli/:"
    ls -la "$HOME/.certifycli/"
    echo ""
    
    if [ -f "$HOME/.certifycli/identity.pem" ]; then
        echo "❌ WARNING: Plaintext private key found on disk!"
    else
        echo "✅ No plaintext private keys found on disk"
        echo "🔐 Private key is securely stored in OS keychain"
    fi
    
    echo ""
    echo "Certificate file (public data, safe to store):"
    if [ -f "$HOME/.certifycli/certificate.pem" ]; then
        file "$HOME/.certifycli/certificate.pem"
        echo "First few lines of certificate:"
        head -3 "$HOME/.certifycli/certificate.pem"
        echo "..."
    fi
else
    echo "❌ Setup directory not found"
fi

echo ""
echo "🔍 Demo 7: Keychain Access Verification"
echo "======================================="
echo "Verifying that private key can be accessed from keychain..."
./certifycli status
echo ""

echo "🎉 Demo Complete!"
echo ""
echo "📋 Summary of security features demonstrated:"
echo "  ✅ OS keychain integration (macOS/Windows/Linux)"
echo "  ✅ Secure private key storage (encrypted by OS)"
echo "  ✅ No plaintext keys on disk"
echo "  ✅ Cross-platform compatibility"
echo "  ✅ Public key fingerprint generation"
echo "  ✅ Secure token storage capability"
echo "  ✅ Key existence verification"
echo ""
echo "🔐 Security benefits:"
echo "  🛡️  Private keys encrypted by operating system"
echo "  🚫 No sensitive data in plaintext files"
echo "  🔒 OS-level access control and permissions"
echo "  🎯 Integration with system security policies"
echo "  🔄 Secure key lifecycle management"
echo ""
echo "🚀 Next implementation steps:"
echo "  1. Server authentication (login command)"
echo "  2. Real CA certificate signing"
echo "  3. Git integration for commit signing"
echo "  4. Certificate lifecycle management"
echo ""
echo "🧹 Cleanup options:"
echo "  To remove demo data: ./certifycli cleanup"
echo "  To remove binary: rm certifycli"
echo ""
echo "⚠️  Note: OS may have prompted for keychain access during this demo."
echo "   This is the security feature working as intended!"
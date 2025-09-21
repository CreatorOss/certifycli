#!/bin/bash

# Demo script untuk menunjukkan fitur crypto CertifyCLI yang sudah diimplementasi
echo "🎬 CertifyCLI Crypto Features Demo"
echo "================================="

# Check if Go is available
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go to run this demo."
    echo "   Visit: https://golang.org/doc/install"
    exit 1
fi

# Build the CLI
echo "🔨 Building CertifyCLI..."
if ! go build -o certifycli ./cmd/certifycli; then
    echo "❌ Build failed"
    exit 1
fi

echo "✅ Build successful!"
echo ""

# Demo 1: Show help
echo "📋 Demo 1: CLI Help System"
echo "=========================="
./certifycli --help
echo ""

# Demo 2: Test crypto functions
echo "🧪 Demo 2: Crypto Functions Test"
echo "================================"
./certifycli test-crypto
echo ""

# Demo 3: Check initial status
echo "📊 Demo 3: Initial Status Check"
echo "==============================="
./certifycli status
echo ""

# Demo 4: Setup identity
echo "🔧 Demo 4: Identity Setup"
echo "========================="
echo "Setting up identity (this will create ~/.certifycli/ directory)..."
echo "y" | ./certifycli setup
echo ""

# Demo 5: Check status after setup
echo "📊 Demo 5: Status After Setup"
echo "============================="
./certifycli status
echo ""

# Demo 6: Show generated files
echo "📁 Demo 6: Generated Files"
echo "=========================="
if [ -d "$HOME/.certifycli" ]; then
    echo "Files in ~/.certifycli/:"
    ls -la "$HOME/.certifycli/"
    echo ""
    
    echo "Private key info:"
    file "$HOME/.certifycli/identity.pem"
    echo ""
    
    echo "Certificate info:"
    file "$HOME/.certifycli/certificate.pem"
    echo ""
    
    echo "Certificate details (first few lines):"
    head -5 "$HOME/.certifycli/certificate.pem"
    echo "..."
    tail -5 "$HOME/.certifycli/certificate.pem"
else
    echo "❌ Setup directory not found"
fi

echo ""
echo "🎉 Demo Complete!"
echo ""
echo "📋 Summary of demonstrated features:"
echo "  ✅ CLI help system and command structure"
echo "  ✅ Crypto function testing (key gen, CSR, certificates)"
echo "  ✅ Identity setup workflow"
echo "  ✅ Status checking and validation"
echo "  ✅ Secure file generation and storage"
echo ""
echo "🔧 Files created:"
echo "  🔐 ~/.certifycli/identity.pem (RSA private key)"
echo "  📄 ~/.certifycli/certificate.pem (test certificate)"
echo ""
echo "🚀 Next steps:"
echo "  1. Implement OS keychain integration for enhanced security"
echo "  2. Add server authentication (login command)"
echo "  3. Implement real CA certificate signing"
echo "  4. Add Git integration for commit signing"
echo ""
echo "🧹 Cleanup:"
echo "  To remove demo files: rm -rf ~/.certifycli/"
echo "  To remove binary: rm certifycli"
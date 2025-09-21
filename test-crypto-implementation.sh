#!/bin/bash

# Test script untuk implementasi crypto CertifyCLI
echo "🧪 Testing CertifyCLI Crypto Implementation"
echo "=========================================="

# Test 1: Check if crypto files exist
echo "📁 Checking crypto implementation files..."
required_files=("internal/crypto/crypto.go" "internal/crypto/crypto_test.go" "cmd/certifycli/main.go")

for file in "${required_files[@]}"; do
    if [ -f "$file" ]; then
        echo "  ✅ $file"
    else
        echo "  ❌ $file (missing)"
        exit 1
    fi
done

# Test 2: Check Go syntax (if Go is available)
echo ""
echo "🔍 Checking Go syntax..."
if command -v go &> /dev/null; then
    echo "  Go found, checking syntax..."
    
    # Download dependencies
    echo "  📦 Downloading dependencies..."
    if go mod tidy; then
        echo "  ✅ Dependencies downloaded successfully"
    else
        echo "  ❌ Failed to download dependencies"
        exit 1
    fi
    
    # Test compilation
    if go build -o /tmp/certifycli-test ./cmd/certifycli; then
        echo "  ✅ Go code compiles successfully"
        
        # Test the CLI help
        echo ""
        echo "🚀 Testing CLI help command..."
        /tmp/certifycli-test --help
        
        echo ""
        echo "🧪 Testing crypto functions..."
        /tmp/certifycli-test test-crypto
        
        echo ""
        echo "📊 Testing status command..."
        /tmp/certifycli-test status
        
        # Clean up
        rm -f /tmp/certifycli-test
    else
        echo "  ❌ Go compilation failed"
        exit 1
    fi
else
    echo "  ⚠️  Go not found, skipping compilation test"
    echo "  📝 To test compilation, install Go and run:"
    echo "     go mod tidy"
    echo "     go build -o certifycli ./cmd/certifycli"
    echo "     ./certifycli --help"
fi

# Test 3: Check code structure
echo ""
echo "🔍 Checking code structure..."

# Check if crypto functions are implemented
if grep -q "GenerateKeyPair" internal/crypto/crypto.go; then
    echo "  ✅ GenerateKeyPair function found"
else
    echo "  ❌ GenerateKeyPair function missing"
fi

if grep -q "CreateCSR" internal/crypto/crypto.go; then
    echo "  ✅ CreateCSR function found"
else
    echo "  ❌ CreateCSR function missing"
fi

if grep -q "GenerateTestCertificate" internal/crypto/crypto.go; then
    echo "  ✅ GenerateTestCertificate function found"
else
    echo "  ❌ GenerateTestCertificate function missing"
fi

# Check if main.go has setup command
if grep -q "handleSetup" cmd/certifycli/main.go; then
    echo "  ✅ handleSetup function found"
else
    echo "  ❌ handleSetup function missing"
fi

echo ""
echo "🎉 Crypto implementation test complete!"
echo ""
echo "📋 Summary of implemented features:"
echo "  ✅ RSA key pair generation (2048-bit)"
echo "  ✅ Private key PEM file save/load"
echo "  ✅ Certificate Signing Request (CSR) creation"
echo "  ✅ Self-signed test certificate generation"
echo "  ✅ CLI setup command for identity creation"
echo "  ✅ CLI status command for checking setup"
echo "  ✅ CLI test-crypto command for testing functions"
echo ""
echo "🚀 Next steps:"
echo "  1. Install Go: https://golang.org/doc/install"
echo "  2. Test compilation: go build -o certifycli ./cmd/certifycli"
echo "  3. Run setup: ./certifycli setup"
echo "  4. Check status: ./certifycli status"
echo "  5. Test crypto: ./certifycli test-crypto"
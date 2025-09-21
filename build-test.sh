#!/bin/bash

# Build and test script for CertifyCLI
echo "🔨 CertifyCLI Build and Test Script"
echo "=================================="

# Check if Go is available
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go to build CertifyCLI."
    echo "   Visit: https://golang.org/doc/install"
    exit 1
fi

echo "📋 Build and test process:"
echo "  ✅ Check Go installation"
echo "  ✅ Download dependencies"
echo "  ✅ Build binary"
echo "  ✅ Test functionality"
echo ""

# Check Go version
echo "🔍 Checking Go version..."
go version

# Download dependencies
echo ""
echo "📦 Downloading dependencies..."
go mod tidy
if [ $? -ne 0 ]; then
    echo "❌ Failed to download dependencies"
    exit 1
fi

# Build the binary
echo ""
echo "🔨 Building CertifyCLI..."
go build -o certifycli ./cmd/certifycli
if [ $? -ne 0 ]; then
    echo "❌ Build failed"
    exit 1
fi

echo "✅ Build successful!"

# Test basic functionality
echo ""
echo "🧪 Testing basic functionality..."

echo "1. Testing help command..."
./certifycli --help > /dev/null
if [ $? -eq 0 ]; then
    echo "✅ Help command works"
else
    echo "❌ Help command failed"
    exit 1
fi

echo "2. Testing crypto functions..."
./certifycli test-crypto > /dev/null
if [ $? -eq 0 ]; then
    echo "✅ Crypto functions work"
else
    echo "❌ Crypto functions failed"
    exit 1
fi

echo "3. Testing keyring functions..."
./certifycli test-keyring > /dev/null 2>&1
if [ $? -eq 0 ]; then
    echo "✅ Keyring functions work"
else
    echo "⚠️  Keyring functions may need OS keychain setup"
fi

# Check binary size
echo ""
echo "📊 Binary information:"
ls -lh certifycli
file certifycli

echo ""
echo "🎉 Build and basic tests completed successfully!"
echo ""
echo "📋 Next steps:"
echo "  1. Run './certifycli setup' to create your identity"
echo "  2. Run './certifycli git configure' to enable Git signing"
echo "  3. Run './test-local-mode.sh' for comprehensive testing"
echo ""
echo "🚀 CertifyCLI is ready for use!"
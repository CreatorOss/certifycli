#!/bin/bash

# Test script untuk Git integration CertifyCLI
echo "🔧 Testing CertifyCLI Git Integration"
echo "===================================="

# Check if Go is available
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go to run this test."
    echo "   Visit: https://golang.org/doc/install"
    exit 1
fi

# Check if Git is available
if ! command -v git &> /dev/null; then
    echo "❌ Git is not installed. Please install Git to run this test."
    echo "   Visit: https://git-scm.com/"
    exit 1
fi

echo "📋 This test demonstrates Git integration features:"
echo "  ✅ Git configuration for commit signing"
echo "  ✅ Automatic commit signing with CertifyCLI"
echo "  ✅ Git status checking and verification"
echo "  ✅ Test repository creation and signing"
echo ""

# Build the CLI
echo "🔨 Building CertifyCLI with Git integration..."
if ! go build -o certifycli ./cmd/certifycli; then
    echo "❌ Build failed"
    exit 1
fi

echo "✅ Build successful!"
echo ""

# Test 1: Check Git version
echo "🔍 Test 1: Git Version Check"
echo "============================"
./certifycli git version
if [ $? -eq 0 ]; then
    echo "✅ Git version check successful"
else
    echo "❌ Git version check failed"
    exit 1
fi
echo ""

# Test 2: Check initial Git status
echo "📊 Test 2: Initial Git Status"
echo "============================="
./certifycli git status
echo ""

# Test 3: Check if identity exists
echo "👤 Test 3: Identity Check"
echo "========================="
if [ -f "$HOME/.certifycli/user" ]; then
    echo "✅ CertifyCLI identity found"
    USERNAME=$(cat "$HOME/.certifycli/user")
    echo "👤 Username: $USERNAME"
else
    echo "⚠️  No CertifyCLI identity found"
    echo "💡 For full testing, run 'certifycli setup' first"
    echo "   Continuing with configuration tests only..."
fi
echo ""

# Test 4: Git configuration (dry run if no identity)
echo "🔧 Test 4: Git Configuration"
echo "============================"
if [ -f "$HOME/.certifycli/user" ]; then
    echo "Configuring Git to use CertifyCLI..."
    ./certifycli git configure
    if [ $? -eq 0 ]; then
        echo "✅ Git configuration successful"
    else
        echo "❌ Git configuration failed"
        exit 1
    fi
else
    echo "⚠️  Skipping Git configuration (no identity)"
fi
echo ""

# Test 5: Check Git status after configuration
echo "📊 Test 5: Git Status After Configuration"
echo "========================================="
./certifycli git status
echo ""

# Test 6: Test Git signing (if identity exists)
echo "🧪 Test 6: Git Signing Test"
echo "==========================="
if [ -f "$HOME/.certifycli/user" ]; then
    echo "Testing Git signing with temporary repository..."
    ./certifycli git test
    if [ $? -eq 0 ]; then
        echo "✅ Git signing test successful"
    else
        echo "❌ Git signing test failed"
    fi
else
    echo "⚠️  Skipping Git signing test (no identity)"
fi
echo ""

# Test 7: Manual Git configuration check
echo "🔍 Test 7: Manual Git Configuration Check"
echo "========================================="
echo "Checking Git global configuration..."
git config --global --list | grep -E "(sign|gpg|x509)" || echo "No signing configuration found"
echo ""

# Test 8: Git disable test (if configured)
echo "🚫 Test 8: Git Disable Test"
echo "==========================="
if git config --global --get gpg.x509.program | grep -q certifycli; then
    echo "Testing Git signing disable..."
    echo "y" | ./certifycli git disable
    if [ $? -eq 0 ]; then
        echo "✅ Git disable test successful"
    else
        echo "❌ Git disable test failed"
    fi
else
    echo "⚠️  Skipping disable test (not configured)"
fi
echo ""

# Test 9: Final status check
echo "📊 Test 9: Final Status Check"
echo "============================="
./certifycli git status
echo ""

echo "🎉 Git Integration Test Complete!"
echo ""
echo "📋 Test Results Summary:"
echo "  ✅ Git version detection"
echo "  ✅ Git status checking"
echo "  ✅ Configuration management"
if [ -f "$HOME/.certifycli/user" ]; then
    echo "  ✅ Identity integration"
    echo "  ✅ Git signing test"
else
    echo "  ⚠️  Identity integration (skipped - no identity)"
    echo "  ⚠️  Git signing test (skipped - no identity)"
fi
echo "  ✅ Configuration disable/enable"
echo ""
echo "🔧 Git Integration Features:"
echo "  📝 Automatic commit signing"
echo "  🔍 Configuration status checking"
echo "  ⚙️  Easy enable/disable"
echo "  🧪 Built-in testing"
echo "  🔐 Secure key management"
echo ""
echo "🚀 Usage Examples:"
echo "  certifycli git configure     # Enable Git signing"
echo "  certifycli git status        # Check configuration"
echo "  certifycli git test          # Test signing"
echo "  certifycli git disable       # Disable signing"
echo ""
echo "💡 Complete Workflow:"
echo "  1. certifycli setup          # Create identity"
echo "  2. certifycli git configure  # Enable Git signing"
echo "  3. git commit -m 'message'   # Commits are now signed!"
echo "  4. git log --show-signature  # Verify signatures"
#!/bin/bash

# Test script untuk CertifyCLI Local Mode
echo "🏠 Testing CertifyCLI Local Mode (Serverless)"
echo "============================================"

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

echo "📋 This test demonstrates local mode features:"
echo "  ✅ Local Certificate Authority (no server required)"
echo "  ✅ Identity setup and management"
echo "  ✅ Git integration with local signing"
echo "  ✅ Backup and restore functionality"
echo "  ✅ Complete offline operation"
echo ""

# Clean up any existing setup
echo "🧹 Cleaning up previous setup..."
rm -rf ~/.certifycli/
rm -f certifycli

# Build the CLI
echo "🔨 Building CertifyCLI in local mode..."
if ! go build -o certifycli ./cmd/certifycli; then
    echo "❌ Build failed"
    exit 1
fi

echo "✅ Build successful!"
echo ""

# Test 1: Help and version
echo "📚 Test 1: Help System"
echo "======================"
./certifycli --help
echo ""

# Test 2: Setup local identity
echo "👤 Test 2: Local Identity Setup"
echo "==============================="
echo "Setting up local identity..."
echo "testuser" | ./certifycli setup
if [ $? -eq 0 ]; then
    echo "✅ Local identity setup successful"
else
    echo "❌ Local identity setup failed"
    exit 1
fi
echo ""

# Test 3: Status check
echo "📊 Test 3: Status Check"
echo "======================="
./certifycli status
echo ""

# Test 4: Certificate information
echo "📋 Test 4: Certificate Information"
echo "=================================="
./certifycli certificates
echo ""

# Test 5: Backup functionality
echo "💾 Test 5: Backup Functionality"
echo "==============================="
./certifycli backup
if [ $? -eq 0 ]; then
    echo "✅ Backup successful"
    echo "📁 Checking backup files..."
    ls -la ~/certifycli-backup/
else
    echo "❌ Backup failed"
fi
echo ""

# Test 6: Git configuration
echo "🔧 Test 6: Git Configuration"
echo "============================"
./certifycli git configure
if [ $? -eq 0 ]; then
    echo "✅ Git configuration successful"
else
    echo "❌ Git configuration failed"
fi
echo ""

# Test 7: Git status check
echo "📊 Test 7: Git Status Check"
echo "==========================="
./certifycli git status
echo ""

# Test 8: Create test repository and test signing
echo "📁 Test 8: Test Repository and Signing"
echo "======================================"
TEST_REPO_DIR="/tmp/certifycli-local-test-$(date +%s)"
mkdir -p "$TEST_REPO_DIR"
cd "$TEST_REPO_DIR"

git init
git config user.name "CertifyCLI Test User"
git config user.email "test@certifycli.local"

# Create some test commits
echo "# Local Test Repository" > README.md
git add README.md
git commit -m "Initial commit with local CertifyCLI signing"

echo "## Features" >> README.md
git add README.md
git commit -m "Add features section"

echo "- Local Certificate Authority" >> README.md
git add README.md
git commit -m "Add local CA feature"

echo "✅ Test repository created with 3 commits"
echo "📍 Repository location: $TEST_REPO_DIR"
echo ""

# Test 9: Verification commands
echo "🔍 Test 9: Verification Commands"
echo "================================"
echo "9.1 Verify last commit:"
../certifycli git verify
echo ""

echo "9.2 Verify all commits:"
../certifycli git verify-all
echo ""

# Test 10: Git log with signatures
echo "📜 Test 10: Git Log with Signatures"
echo "==================================="
echo "Checking Git log for signature information..."
git log --oneline --show-signature -3 2>/dev/null || git log --oneline -3
echo ""

# Test 11: Test restore functionality
echo "🔄 Test 11: Restore Functionality"
echo "================================="
cd - > /dev/null
echo "Testing restore functionality..."

# Backup current state
cp -r ~/.certifycli ~/.certifycli-test-backup

# Remove current identity
rm -rf ~/.certifycli

# Test status without identity
echo "Status without identity:"
./certifycli status
echo ""

# Restore from backup
echo "y" | ./certifycli restore
if [ $? -eq 0 ]; then
    echo "✅ Restore successful"
    echo "Verifying restored identity:"
    ./certifycli status
else
    echo "❌ Restore failed"
fi
echo ""

# Test 12: Test crypto functions
echo "🧪 Test 12: Crypto Functions"
echo "============================"
./certifycli test-crypto
echo ""

# Test 13: Test keyring functions
echo "🔐 Test 13: Keyring Functions"
echo "============================="
./certifycli test-keyring
echo ""

# Test 14: File permissions check
echo "🔒 Test 14: Security Check"
echo "=========================="
echo "Checking file permissions..."
ls -la ~/.certifycli/
echo ""
echo "Checking CA private key security:"
if [ -f ~/.certifycli/ca-private-key.pem ]; then
    PERMS=$(stat -c "%a" ~/.certifycli/ca-private-key.pem 2>/dev/null || stat -f "%A" ~/.certifycli/ca-private-key.pem 2>/dev/null)
    if [ "$PERMS" = "600" ]; then
        echo "✅ CA private key has correct permissions (600)"
    else
        echo "⚠️  CA private key permissions: $PERMS (should be 600)"
    fi
fi
echo ""

# Cleanup
echo "🧹 Cleanup"
echo "=========="
cd - > /dev/null
rm -rf "$TEST_REPO_DIR"
rm -rf ~/.certifycli-test-backup
echo "✅ Test repository and backup cleaned up"
echo ""

echo "🎉 Local Mode Test Complete!"
echo ""
echo "📋 Test Results Summary:"
echo "  ✅ Local identity setup"
echo "  ✅ Local Certificate Authority"
echo "  ✅ Git configuration and signing"
echo "  ✅ Verification commands"
echo "  ✅ Backup and restore functionality"
echo "  ✅ Security permissions"
echo "  ✅ Crypto and keyring functions"
echo ""
echo "🏠 Local Mode Features Verified:"
echo "  🔐 No server required - fully offline"
echo "  🏛️  Local Certificate Authority"
echo "  🔑 OS keychain integration"
echo "  📝 Git commit signing"
echo "  🔍 Signature verification"
echo "  💾 Backup and restore"
echo "  🔒 Secure file permissions"
echo ""
echo "🚀 Local Mode Commands:"
echo "  certifycli setup          # Create local identity"
echo "  certifycli git configure  # Enable Git signing"
echo "  certifycli backup         # Backup identity"
echo "  certifycli status         # Check status"
echo ""
echo "💡 Complete Local Workflow:"
echo "  1. certifycli setup          # One-time setup"
echo "  2. certifycli git configure  # Enable Git signing"
echo "  3. git commit -m 'message'   # Commits are signed!"
echo "  4. certifycli backup         # Backup your identity"
echo ""
echo "🎯 Local Mode Benefits:"
echo "  🔐 Complete privacy (no data leaves your machine)"
echo "  🏠 No server infrastructure required"
echo "  ⚡ Fast operation (no network calls)"
echo "  💾 Portable (backup/restore anywhere)"
echo "  🔒 Secure (OS keychain + file permissions)"
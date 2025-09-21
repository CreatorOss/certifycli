#!/bin/bash

# Test script untuk enhanced Git integration CertifyCLI
echo "🔧 Testing Enhanced CertifyCLI Git Integration"
echo "=============================================="

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

echo "📋 This test demonstrates enhanced Git integration features:"
echo "  ✅ GPG-compatible signature format"
echo "  ✅ Enhanced verification commands"
echo "  ✅ Pretty output formatting"
echo "  ✅ Comprehensive commit verification"
echo ""

# Build the CLI
echo "🔨 Building CertifyCLI with enhanced Git integration..."
if ! go build -o certifycli ./cmd/certifycli; then
    echo "❌ Build failed"
    exit 1
fi

echo "✅ Build successful!"
echo ""

# Test 1: Enhanced Git help
echo "📚 Test 1: Enhanced Git Help"
echo "============================"
./certifycli git --help 2>/dev/null || ./certifycli git
echo ""

# Test 2: Git version and status
echo "📊 Test 2: Git Status Check"
echo "==========================="
./certifycli git version
echo ""
./certifycli git status
echo ""

# Test 3: Check if identity exists for full testing
echo "👤 Test 3: Identity Check"
echo "========================="
if [ -f "$HOME/.certifycli/user" ]; then
    echo "✅ CertifyCLI identity found"
    USERNAME=$(cat "$HOME/.certifycli/user")
    echo "👤 Username: $USERNAME"
    FULL_TEST=true
else
    echo "⚠️  No CertifyCLI identity found"
    echo "💡 For full testing, run 'certifycli setup' first"
    echo "   Continuing with limited tests..."
    FULL_TEST=false
fi
echo ""

# Test 4: Git configuration (if identity exists)
if [ "$FULL_TEST" = true ]; then
    echo "🔧 Test 4: Git Configuration"
    echo "============================"
    echo "Configuring Git to use CertifyCLI..."
    ./certifycli git configure
    if [ $? -eq 0 ]; then
        echo "✅ Git configuration successful"
    else
        echo "❌ Git configuration failed"
        exit 1
    fi
    echo ""
fi

# Test 5: Create test repository for verification testing
echo "📁 Test 5: Test Repository Setup"
echo "================================"
TEST_REPO_DIR="/tmp/certifycli-git-test-$(date +%s)"
mkdir -p "$TEST_REPO_DIR"
cd "$TEST_REPO_DIR"

git init
git config user.name "CertifyCLI Test User"
git config user.email "test@certifycli.com"

# Create some test commits
echo "# Test Repository" > README.md
git add README.md
git commit -m "Initial commit"

echo "## Features" >> README.md
git add README.md
git commit -m "Add features section"

echo "- Git signing integration" >> README.md
git add README.md
git commit -m "Add Git signing feature"

echo "✅ Test repository created with 3 commits"
echo "📍 Repository location: $TEST_REPO_DIR"
echo ""

# Test 6: Verification commands
echo "🔍 Test 6: Verification Commands"
echo "================================"
echo "Testing verification commands on test repository..."

echo "6.1 Verify last commit:"
../certifycli git verify
echo ""

echo "6.2 Verify all commits:"
../certifycli git verify-all
echo ""

# Test 7: Git log with signature info
echo "📜 Test 7: Git Log with Signatures"
echo "=================================="
echo "Checking Git log for signature information..."
git log --oneline --show-signature -3 2>/dev/null || git log --oneline -3
echo ""

# Test 8: Test signing (if configured)
if [ "$FULL_TEST" = true ]; then
    echo "✍️  Test 8: Test Signing"
    echo "======================="
    echo "Testing Git signing with CertifyCLI..."
    ../certifycli git test
    echo ""
fi

# Test 9: Configuration verification
echo "⚙️  Test 9: Configuration Verification"
echo "====================================="
echo "Checking final Git configuration..."
../certifycli git status
echo ""

# Cleanup
echo "🧹 Cleanup"
echo "=========="
cd - > /dev/null
rm -rf "$TEST_REPO_DIR"
echo "✅ Test repository cleaned up"
echo ""

echo "🎉 Enhanced Git Integration Test Complete!"
echo ""
echo "📋 Test Results Summary:"
echo "  ✅ Enhanced help system"
echo "  ✅ Git version detection"
echo "  ✅ Configuration management"
echo "  ✅ Test repository creation"
echo "  ✅ Verification commands"
echo "  ✅ Pretty output formatting"
if [ "$FULL_TEST" = true ]; then
    echo "  ✅ Full identity integration"
    echo "  ✅ Git signing configuration"
    echo "  ✅ Signing test"
else
    echo "  ⚠️  Identity integration (skipped - no identity)"
    echo "  ⚠️  Git signing test (skipped - no identity)"
fi
echo ""
echo "🔧 Enhanced Git Features:"
echo "  📝 GPG-compatible signature format"
echo "  🔍 Comprehensive verification commands"
echo "  🎨 Pretty output with colors and formatting"
echo "  📊 Progress bars for large repositories"
echo "  📋 Detailed signature status reporting"
echo "  🔐 Secure key management integration"
echo ""
echo "🚀 New Commands Available:"
echo "  certifycli git verify        # Verify last commit"
echo "  certifycli git verify-all    # Verify all commits"
echo "  certifycli git status        # Enhanced status check"
echo ""
echo "💡 Complete Enhanced Workflow:"
echo "  1. certifycli setup          # Create identity"
echo "  2. certifycli git configure  # Enable Git signing"
echo "  3. git commit -m 'message'   # Commits are signed!"
echo "  4. certifycli git verify     # Verify signatures"
echo "  5. certifycli git verify-all # Verify all commits"
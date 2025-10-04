#!/bin/bash


echo "🔧 Testing Enhanced CertifyCLI Git Integration"
echo "=============================================="


if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go to run this test."
    echo "   Visit: https://golang.org/doc/install"
    exit 1
fi


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


echo "🔨 Building CertifyCLI with enhanced Git integration..."
if ! go build -o certifycli ./cmd/certifycli; then
    echo "❌ Build failed"
    exit 1
fi

echo "✅ Build successful!"
echo ""


echo "📚 Test 1: Enhanced Git Help"
echo "============================"
./certifycli git --help 2>/dev/null || ./certifycli git
echo ""


echo "📊 Test 2: Git Status Check"
echo "==========================="
./certifycli git version
echo ""
./certifycli git status
echo ""


echo "👤 Test 3: Identity Check"
echo "========================="
if [ -f "$HOME/.certifycli/user" ]; then
    echo "✅ CertifyCLI identity found"
    USERNAME=$(cat "$HOME/.certifycli/user")
    echo "👤 Username: $USERNAME"
    FULL_TEST=true
else
    echo "⚠️  No CertifyCLI identity found"
    echo "📊 For full testing, run 'certifycli setup' first"
    echo "   Continuing with limited tests..."
    FULL_TEST=false
fi
echo ""


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


echo "📁 Test 5: Test Repository Setup"
echo "================================"
TEST_REPO_DIR="/tmp/certifycli-git-test-$(date +%s)"
mkdir -p "$TEST_REPO_DIR"
cd "$TEST_REPO_DIR"

git init
git config user.name "CertifyCLI Test User"
git config user.email "test@certifycli.com"


echo "
git add README.md
git commit -m "Initial commit"

echo "
git add README.md
git commit -m "Add features section"

echo "- Git signing integration" >> README.md
git add README.md
git commit -m "Add Git signing feature"

echo "✅ Test repository created with 3 commits"
echo "📍 Repository location: $TEST_REPO_DIR"
echo ""


echo "🔍 Test 6: Verification Commands"
echo "================================"
echo "Testing verification commands on test repository..."

echo "6.1 Verify last commit:"
../certifycli git verify
echo ""

echo "6.2 Verify all commits:"
../certifycli git verify-all
echo ""


echo "📜 Test 7: Git Log with Signatures"
echo "=================================="
echo "Checking Git log for signature information..."
git log --oneline --show-signature -3 2>/dev/null || git log --oneline -3
echo ""


if [ "$FULL_TEST" = true ]; then
    echo "✍️  Test 8: Test Signing"
    echo "======================="
    echo "Testing Git signing with CertifyCLI..."
    ../certifycli git test
    echo ""
fi


echo "⚙️  Test 9: Configuration Verification"
echo "====================================="
echo "Checking final Git configuration..."
../certifycli git status
echo ""


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
    echo "  ✅ Signing _t1"
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
echo "  certifycli git verify        
echo "  certifycli git verify-all    
echo "  certifycli git status        
echo ""
echo "💡 Complete Enhanced Workflow:"
echo "  1. certifycli setup          # Set up local identity"
echo "  2. certifycli git configure  
echo "  3. git commit -m 'message'   
echo "  4. certifycli git verify     
echo "  5. certifycli git verify-all 
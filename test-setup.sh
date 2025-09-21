#!/bin/bash

# Test script to verify the project setup
echo "🧪 Testing CertifyCLI Project Setup"
echo "=================================="

# Test 1: Check project structure
echo "📁 Checking project structure..."
required_dirs=("cmd/certifycli" "internal/auth" "internal/crypto" "internal/utils" "internal/ca" "server/controllers" "server/models" "server/routes" "scripts")

for dir in "${required_dirs[@]}"; do
    if [ -d "$dir" ]; then
        echo "  ✅ $dir"
    else
        echo "  ❌ $dir (missing)"
        exit 1
    fi
done

# Test 2: Check required files
echo ""
echo "📄 Checking required files..."
required_files=("go.mod" "README.md" "LICENSE" ".gitignore" "cmd/certifycli/main.go" "server/package.json" "server/index.js")

for file in "${required_files[@]}"; do
    if [ -f "$file" ]; then
        echo "  ✅ $file"
    else
        echo "  ❌ $file (missing)"
        exit 1
    fi
done

# Test 3: Check Go syntax (if Go is available)
echo ""
echo "🔍 Checking Go syntax..."
if command -v go &> /dev/null; then
    echo "  Go found, checking syntax..."
    if go mod tidy && go build -o /tmp/certifycli-test ./cmd/certifycli; then
        echo "  ✅ Go code compiles successfully"
        rm -f /tmp/certifycli-test
    else
        echo "  ❌ Go compilation failed"
        exit 1
    fi
else
    echo "  ⚠️  Go not found, skipping compilation test"
fi

# Test 4: Check Node.js setup (if Node.js is available)
echo ""
echo "🔍 Checking Node.js setup..."
if command -v node &> /dev/null; then
    echo "  Node.js found, checking package.json..."
    cd server
    if node -e "require('./package.json')"; then
        echo "  ✅ package.json is valid"
    else
        echo "  ❌ package.json is invalid"
        exit 1
    fi
    cd ..
else
    echo "  ⚠️  Node.js not found, skipping package.json test"
fi

echo ""
echo "🎉 All tests passed! Project setup is complete."
echo ""
echo "Next steps:"
echo "1. Install Go (if not already installed): https://golang.org/doc/install"
echo "2. Install Node.js (if not already installed): https://nodejs.org/"
echo "3. Build the CLI: go build -o certifycli ./cmd/certifycli"
echo "4. Set up the server: cd server && npm install"
echo "5. Start the server: npm start"
echo "6. Test the CLI: ./certifycli --help"
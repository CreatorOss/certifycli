#!/bin/bash

# CertifyCLI Installation Script
# This script builds and installs the CertifyCLI binary

set -e

echo "🚀 Installing CertifyCLI..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or later."
    echo "   Visit: https://golang.org/doc/install"
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
REQUIRED_VERSION="1.21"

if [ "$(printf '%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V | head -n1)" != "$REQUIRED_VERSION" ]; then
    echo "❌ Go version $REQUIRED_VERSION or later is required. Found: $GO_VERSION"
    exit 1
fi

echo "✅ Go version $GO_VERSION detected"

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"

# Change to project directory
cd "$PROJECT_DIR"

echo "📦 Downloading dependencies..."
go mod tidy

echo "🔨 Building CertifyCLI..."
go build -o certifycli ./cmd/certifycli

# Make binary executable
chmod +x certifycli

# Determine installation directory
if [ -w "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
elif [ -w "$HOME/.local/bin" ]; then
    INSTALL_DIR="$HOME/.local/bin"
    # Create directory if it doesn't exist
    mkdir -p "$INSTALL_DIR"
else
    INSTALL_DIR="$HOME/bin"
    mkdir -p "$INSTALL_DIR"
fi

echo "📁 Installing to $INSTALL_DIR..."

# Copy binary to installation directory
cp certifycli "$INSTALL_DIR/"

echo "✅ CertifyCLI installed successfully!"
echo ""
echo "🎉 You can now use 'certifycli' from anywhere in your terminal"
echo ""
echo "Quick start:"
echo "  certifycli --help     # Show help"
echo "  certifycli login      # Login to CertifyCLI server"
echo "  certifycli setup      # Set up your identity"
echo "  certifycli status     # Check your status"
echo ""

# Check if installation directory is in PATH
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo "⚠️  Warning: $INSTALL_DIR is not in your PATH"
    echo "   Add this line to your shell profile (~/.bashrc, ~/.zshrc, etc.):"
    echo "   export PATH=\"$INSTALL_DIR:\$PATH\""
    echo ""
fi

echo "🔧 To set up the server:"
echo "  cd server"
echo "  npm install"
echo "  npm start"
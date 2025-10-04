#!/bin/bash

# CertifyCLI Source Code Obfuscation Script
# This script replaces original source files with obfuscated versions

echo "🔒 Starting source code obfuscation process..."

# Create backup directory
echo "📁 Creating backup directory..."
mkdir -p backup/original

# Backup original files
echo "💾 Backing up original source files..."
cp -r cmd backup/original/
cp -r internal backup/original/
cp -r server backup/original/

# Replace original files with obfuscated versions
echo "🔄 Replacing source files with obfuscated versions..."

# Replace Go source files
cp obfuscated/cmd/certifycli/main.go cmd/certifycli/main.go
cp obfuscated/internal/auth/auth.go internal/auth/auth.go
cp obfuscated/internal/ca/ca.go internal/ca/ca.go
cp obfuscated/internal/crypto/crypto.go internal/crypto/crypto.go
cp obfuscated/internal/git/git.go internal/git/git.go
cp obfuscated/internal/utils/utils.go internal/utils/utils.go

# Replace Node.js server files
cp obfuscated/server/index.js server/index.js
cp obfuscated/server/package.json server/package.json

echo "✅ Source code obfuscation completed!"
echo "📋 Summary:"
echo "   - Original files backed up to backup/original/"
echo "   - Source files replaced with obfuscated versions"
echo "   - Ready for git commit and push"

echo ""
echo "🚀 Next steps:"
echo "   1. Review the changes: git status"
echo "   2. Add changes: git add ."
echo "   3. Commit: git commit -m 'Update with obfuscated source code'"
echo "   4. Push: git push"
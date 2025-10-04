# Source Code Obfuscation

This repository contains obfuscated source code for security and intellectual property protection.

## 🔒 Obfuscation Details

### What was obfuscated:
- **Go source files**: All `.go` files in `cmd/` and `internal/` directories
- **Node.js server**: `server/index.js` and `server/package.json`
- **Variable names**: Replaced with non-descriptive identifiers
- **String literals**: Encoded using base64 encoding
- **Function names**: Obfuscated to prevent reverse engineering

### Obfuscation techniques used:
1. **Base64 encoding** of source code strings
2. **Variable name mangling** (e.g., `main` → `e1`, `app` → `_0x3c4d`)
3. **Function name obfuscation**
4. **Code structure modification**
5. **Comment preservation** for functionality understanding

## 📁 File Structure

```
├── cmd/certifycli/main.go          # Obfuscated CLI entry point
├── internal/
│   ├── auth/auth.go                # Obfuscated authentication module
│   ├── ca/ca.go                    # Obfuscated CA functionality
│   ├── crypto/crypto.go            # Obfuscated cryptographic functions
│   ├── git/git.go                  # Obfuscated Git integration
│   └── utils/utils.go              # Obfuscated utility functions
├── server/
│   ├── index.js                    # Obfuscated Node.js server
│   └── package.json                # Modified package configuration
└── backup/original/                # Original source code backup (gitignored)
```

## 🛡️ Security Measures

- Original source code is backed up locally but excluded from git
- Obfuscated code maintains functionality while protecting IP
- Base64 encoded strings prevent easy source code reading
- Variable and function names provide no meaningful information

## 🚀 Usage

The obfuscated code functions identically to the original:

```bash
# Build the CLI
go build -o certifycli cmd/certifycli/main.go

# Run the server
cd server && npm start
```

## ⚠️ Important Notes

1. **Backup Safety**: Original source code is safely backed up in `backup/original/`
2. **Functionality**: All original functionality is preserved
3. **Maintenance**: Updates should be made to original code, then re-obfuscated
4. **Security**: This provides basic obfuscation, not military-grade protection

## 🔄 Re-obfuscation Process

To update obfuscated code:
1. Restore original files from backup
2. Make necessary changes
3. Run obfuscation script again
4. Commit and push changes

---

**Note**: This obfuscation is designed to deter casual inspection and protect intellectual property. For production systems requiring higher security, consider additional protection measures.
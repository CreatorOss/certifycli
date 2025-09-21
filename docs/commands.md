# 📋 Command Reference

Complete reference for all CertifyCLI commands.

## 🚀 Setup Commands

### `certifycli setup`
Initialize CertifyCLI with your identity and generate certificates.

```bash
certifycli setup
```

**Interactive prompts:**
- Name (for certificate)
- Email address
- Organization (optional)
- Key size (2048, 3072, 4096 bits)

### `certifycli git configure`
Configure Git to use CertifyCLI for commit signing.

```bash
certifycli git configure
```

**What it does:**
- Sets `commit.gpgsign = true`
- Configures `gpg.program` to use CertifyCLI
- Sets up signing key

## 🔐 Certificate Management

### `certifycli cert list`
List all certificates in the local store.

```bash
certifycli cert list
```

### `certifycli cert show`
Display certificate details.

```bash
certifycli cert show [--fingerprint <fingerprint>]
```

### `certifycli cert export`
Export certificate to file.

```bash
certifycli cert export --output cert.pem
```

## 🔑 Key Management

### `certifycli key list`
List all keys in the keychain.

```bash
certifycli key list
```

### `certifycli key backup`
Backup private key (encrypted).

```bash
certifycli key backup --output backup.key
```

### `certifycli key restore`
Restore private key from backup.

```bash
certifycli key restore --input backup.key
```

## ✅ Verification Commands

### `certifycli verify commit`
Verify a specific commit signature.

```bash
certifycli verify commit <commit-hash>
certifycli verify commit HEAD
certifycli verify commit --all  # Verify all commits
```

### `certifycli verify repo`
Verify all commits in repository.

```bash
certifycli verify repo
certifycli verify repo --since "2023-01-01"
certifycli verify repo --author "john@example.com"
```

## 🛠️ Utility Commands

### `certifycli status`
Show current configuration status.

```bash
certifycli status
```

**Output includes:**
- Certificate status
- Git configuration
- Keychain status
- Last signing activity

### `certifycli config`
Manage configuration settings.

```bash
certifycli config list
certifycli config set key.size 4096
certifycli config get git.signing
```

### `certifycli clean`
Clean up temporary files and caches.

```bash
certifycli clean
certifycli clean --all  # Include certificates and keys
```

## 🔄 Git Integration Commands

### `certifycli git sign`
Manually sign a commit (usually automatic).

```bash
certifycli git sign <commit-hash>
```

### `certifycli git verify`
Verify Git commit signatures.

```bash
certifycli git verify <commit-hash>
certifycli git verify --range HEAD~10..HEAD
```

## 📊 Reporting Commands

### `certifycli report`
Generate signing activity report.

```bash
certifycli report
certifycli report --format json
certifycli report --since "2023-01-01"
```

### `certifycli audit`
Audit repository for unsigned commits.

```bash
certifycli audit
certifycli audit --fix  # Attempt to sign unsigned commits
```

## 🆘 Help and Information

### `certifycli help`
Show help for any command.

```bash
certifycli help
certifycli help setup
certifycli help verify
```

### `certifycli version`
Show version information.

```bash
certifycli version
certifycli version --full  # Include build info
```

## 🔧 Advanced Options

### Global Flags

All commands support these global flags:

- `--config <file>` - Use custom config file
- `--verbose` - Enable verbose output
- `--quiet` - Suppress non-error output
- `--no-color` - Disable colored output
- `--help` - Show command help

### Environment Variables

- `CERTIFYCLI_CONFIG` - Config file path
- `CERTIFYCLI_KEYCHAIN` - Keychain service name
- `CERTIFYCLI_DEBUG` - Enable debug mode
- `CERTIFYCLI_NO_KEYCHAIN` - Disable keychain (testing only)

## 📝 Examples

### Complete Setup Workflow
```bash
# Initial setup
certifycli setup

# Configure Git
certifycli git configure

# Verify setup
certifycli status

# Make a signed commit
git add .
git commit -m "My first signed commit"

# Verify the commit
certifycli verify commit HEAD
```

### Backup and Restore
```bash
# Backup your key
certifycli key backup --output ~/certifycli-backup.key

# Later, restore on new machine
certifycli key restore --input ~/certifycli-backup.key
certifycli setup --restore-only
```

### Repository Audit
```bash
# Check all commits
certifycli verify repo

# Generate report
certifycli report --format json > signing-report.json

# Audit for unsigned commits
certifycli audit
```

## 🚨 Troubleshooting Commands

### `certifycli doctor`
Diagnose common issues.

```bash
certifycli doctor
```

**Checks:**
- Git configuration
- Keychain access
- Certificate validity
- File permissions

### `certifycli reset`
Reset configuration (use with caution).

```bash
certifycli reset --config-only
certifycli reset --all  # Removes everything
```

## 💡 Tips

1. **Use `--help`** with any command for detailed options
2. **Check `certifycli status`** if something isn't working
3. **Run `certifycli doctor`** for troubleshooting
4. **Use `--verbose`** for debugging issues
5. **Backup your keys** before major changes

## 🔗 Related Documentation

- [Troubleshooting Guide](troubleshooting.md)
- [Contributing Guide](../CONTRIBUTING.md)
- [Security Documentation](../SECURITY.md)
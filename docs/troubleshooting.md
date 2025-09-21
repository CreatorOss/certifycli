# 🔧 Troubleshooting Guide

Common issues and solutions for CertifyCLI.

## 🚨 Quick Diagnostics

First, run the built-in diagnostic tool:

```bash
certifycli doctor
```

This will check:
- ✅ Git configuration
- ✅ Keychain access
- ✅ Certificate validity
- ✅ File permissions

## 🔍 Common Issues

### 1. "Command not found: certifycli"

**Problem:** CertifyCLI is not in your PATH.

**Solutions:**

```bash
# Check if installed
which certifycli

# If using Go install
export PATH=$PATH:$(go env GOPATH)/bin

# If using binary download
sudo mv certifycli /usr/local/bin/
chmod +x /usr/local/bin/certifycli
```

### 2. "Failed to access keychain"

**Problem:** CertifyCLI can't access the OS keychain.

**Solutions:**

**macOS:**
```bash
# Grant keychain access
security unlock-keychain ~/Library/Keychains/login.keychain

# Check keychain status
security list-keychains
```

**Linux:**
```bash
# Install secret service
sudo apt-get install gnome-keyring  # Ubuntu/Debian
sudo yum install gnome-keyring      # RHEL/CentOS

# Start keyring daemon
gnome-keyring-daemon --start
```

**Windows:**
```powershell
# Check Windows Credential Manager
cmdkey /list
```

### 3. "Git signing not working"

**Problem:** Commits are not being signed automatically.

**Check Git configuration:**
```bash
git config --list | grep -E "(commit.gpgsign|gpg.program|user.signingkey)"
```

**Expected output:**
```
commit.gpgsign=true
gpg.program=certifycli
user.signingkey=<your-key-id>
```

**Fix:**
```bash
certifycli git configure
```

### 4. "Certificate expired or invalid"

**Problem:** Certificate has expired or is corrupted.

**Check certificate:**
```bash
certifycli cert list
certifycli cert show
```

**Regenerate certificate:**
```bash
certifycli setup --regenerate
```

### 5. "Permission denied" errors

**Problem:** File permission issues.

**Fix permissions:**
```bash
# Fix config directory
chmod 700 ~/.certifycli
chmod 600 ~/.certifycli/*

# Fix binary permissions
chmod +x /usr/local/bin/certifycli
```

## 🔐 Keychain Issues

### macOS Keychain Problems

**Issue:** "The user name or passphrase you entered is not correct"

```bash
# Reset keychain
security delete-keychain ~/Library/Keychains/login.keychain
security create-keychain -p "" ~/Library/Keychains/login.keychain
security default-keychain -s ~/Library/Keychains/login.keychain
```

**Issue:** Keychain locked

```bash
security unlock-keychain ~/Library/Keychains/login.keychain
```

### Linux Secret Service Issues

**Issue:** "No such interface 'org.freedesktop.Secret.Service'"

```bash
# Install and start secret service
sudo apt-get install gnome-keyring libsecret-1-0
gnome-keyring-daemon --start --components=secrets
```

**Issue:** DBus not available

```bash
# Start DBus session
eval $(dbus-launch --sh-syntax)
export DBUS_SESSION_BUS_ADDRESS
```

### Windows Credential Manager Issues

**Issue:** Access denied to Credential Manager

```powershell
# Run as administrator
# Check Windows services
Get-Service -Name "Windows Credential Manager"
```

## 🔄 Git Integration Issues

### Commit Signing Not Working

**Check Git version:**
```bash
git --version
# Requires Git 2.0+
```

**Verify GPG program:**
```bash
git config gpg.program
# Should be: certifycli
```

**Test signing manually:**
```bash
echo "test" | certifycli git sign
```

### Verification Failures

**Issue:** "Bad signature" or verification fails

**Check certificate chain:**
```bash
certifycli verify commit HEAD --verbose
```

**Regenerate if needed:**
```bash
certifycli setup --regenerate
git commit --amend --no-edit  # Re-sign last commit
```

## 🛠️ Advanced Troubleshooting

### Enable Debug Mode

```bash
export CERTIFYCLI_DEBUG=1
certifycli status
```

### Check Configuration

```bash
certifycli config list
cat ~/.certifycli/config.yaml
```

### Verify File Integrity

```bash
# Check config directory
ls -la ~/.certifycli/

# Verify certificate files
openssl x509 -in ~/.certifycli/cert.pem -text -noout
```

### Network Issues (if applicable)

```bash
# Test connectivity (for future server features)
curl -I https://api.certifycli.com/health
```

## 🔄 Reset and Recovery

### Soft Reset (Config Only)

```bash
certifycli reset --config-only
certifycli setup
```

### Hard Reset (Everything)

```bash
certifycli reset --all
rm -rf ~/.certifycli
certifycli setup
```

### Backup Before Reset

```bash
# Backup key
certifycli key backup --output ~/certifycli-backup.key

# Backup config
cp -r ~/.certifycli ~/certifycli-config-backup
```

## 📊 Performance Issues

### Slow Signing

**Check key size:**
```bash
certifycli config get key.size
```

**Optimize for speed:**
```bash
certifycli config set key.size 2048  # Faster than 4096
```

### Large Repository Verification

**Use selective verification:**
```bash
# Verify recent commits only
certifycli verify repo --since "7 days ago"

# Verify specific author
certifycli verify repo --author "you@example.com"
```

## 🔍 Logging and Debugging

### Enable Verbose Output

```bash
certifycli --verbose status
certifycli --verbose verify commit HEAD
```

### Check System Logs

**macOS:**
```bash
log show --predicate 'process == "certifycli"' --last 1h
```

**Linux:**
```bash
journalctl -u certifycli --since "1 hour ago"
```

**Windows:**
```powershell
Get-EventLog -LogName Application -Source "CertifyCLI" -Newest 10
```

## 🆘 Getting Help

### Built-in Help

```bash
certifycli help
certifycli help <command>
```

### Community Support

- 🐛 [Report Issues](https://github.com/CreatorOss/certifycli/issues)
- 💬 [Join Discussions](https://github.com/CreatorOss/certifycli/discussions)
- 📖 [Check Wiki](https://github.com/CreatorOss/certifycli/wiki)

### Enterprise Support

For priority support and custom solutions:
- 📧 Email: **enterprise@certifycli.com**
- 💰 [Sponsor on GitHub](https://github.com/sponsors/CreatorOss)
- ☕ [Buy Me a Coffee](https://buymeacoffee.com/creatoross)

## 📝 Reporting Issues

When reporting issues, please include:

1. **System Information:**
   ```bash
   certifycli version --full
   uname -a
   git --version
   ```

2. **Configuration:**
   ```bash
   certifycli status
   certifycli config list
   ```

3. **Error Output:**
   ```bash
   certifycli --verbose <failing-command> 2>&1
   ```

4. **Steps to Reproduce:**
   - What you were trying to do
   - What happened instead
   - What you expected to happen

## 🔗 Related Documentation

- [Command Reference](commands.md)
- [Contributing Guide](../CONTRIBUTING.md)
- [Security Documentation](../SECURITY.md)
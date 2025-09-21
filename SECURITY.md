# Security Policy

## Supported Versions

We actively support the following versions of CertifyCLI with security updates:

| Version | Supported          |
| ------- | ------------------ |
| 1.x.x   | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

We take security seriously. If you discover a security vulnerability in CertifyCLI, please report it responsibly.

### How to Report

**Please do NOT report security vulnerabilities through public GitHub issues.**

Instead, please report security vulnerabilities by:

1. **Email**: Send details to [security@certifycli.dev] (if available)
2. **GitHub Security Advisory**: Use GitHub's private vulnerability reporting feature
3. **Direct Contact**: Contact the maintainers directly through GitHub

### What to Include

When reporting a vulnerability, please include:

- **Description**: Clear description of the vulnerability
- **Impact**: Potential impact and attack scenarios
- **Reproduction**: Steps to reproduce the vulnerability
- **Environment**: Affected versions and platforms
- **Proof of Concept**: If applicable, include PoC code
- **Suggested Fix**: If you have ideas for fixing the issue

### Response Timeline

- **Acknowledgment**: We will acknowledge receipt within 48 hours
- **Initial Assessment**: We will provide an initial assessment within 5 business days
- **Regular Updates**: We will provide updates every 5 business days until resolution
- **Resolution**: We aim to resolve critical vulnerabilities within 30 days

### Security Update Process

1. **Verification**: We verify and reproduce the vulnerability
2. **Assessment**: We assess the impact and severity
3. **Fix Development**: We develop and test a fix
4. **Coordinated Disclosure**: We coordinate disclosure with the reporter
5. **Release**: We release a security update
6. **Advisory**: We publish a security advisory

## Security Best Practices

### For Users

- **Keep Updated**: Always use the latest version of CertifyCLI
- **Secure Backups**: Store identity backups in secure, encrypted locations
- **Access Control**: Ensure only authorized users can access your machine
- **Regular Verification**: Regularly verify your commits with `certifycli git verify-all`
- **Monitor Activity**: Monitor your Git repositories for unauthorized commits

### For Developers

- **Code Review**: All security-related changes require thorough review
- **Input Validation**: Validate all inputs, especially cryptographic parameters
- **Error Handling**: Handle errors securely without leaking sensitive information
- **Secure Defaults**: Use secure defaults for all cryptographic operations
- **Testing**: Include security tests for all new features

## Security Features

### Current Security Measures

- **Private Key Protection**: Private keys stored in OS keychain, never on disk
- **Local Certificate Authority**: Self-contained CA with strong 4096-bit RSA keys
- **File Permissions**: Sensitive files protected with 600 permissions
- **No Network Communication**: Complete offline operation eliminates network attack vectors
- **Certificate Validation**: Full X.509 certificate chain validation
- **GPG Compatibility**: Signatures compatible with Git's GPG verification system

### Security Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    Security Boundaries                     │
├─────────────────────────────────────────────────────────────┤
│  🔐 OS Keychain (Encrypted Storage)                        │
│  ├── Private Keys (RSA 2048-bit)                           │
│  ├── Access Control (OS-level)                             │
│  └── Encryption (OS-managed)                               │
│                                                             │
│  📁 Local File System (Controlled Access)                  │
│  ├── CA Certificate (Public, 644)                          │
│  ├── CA Private Key (Secure, 600)                          │
│  ├── User Certificate (Secure, 600)                        │
│  └── Configuration (Secure, 600)                           │
│                                                             │
│  🔧 Application Layer                                       │
│  ├── Input Validation                                      │
│  ├── Error Handling                                        │
│  ├── Cryptographic Operations                              │
│  └── Certificate Management                                │
└─────────────────────────────────────────────────────────────┘
```

## Threat Model

### Assets

- **Private Keys**: User's RSA private keys
- **CA Private Key**: Local Certificate Authority private key
- **Certificates**: X.509 certificates and certificate chain
- **Git Signatures**: Cryptographic signatures on Git commits

### Threats

- **Key Compromise**: Unauthorized access to private keys
- **Certificate Forgery**: Creation of fraudulent certificates
- **Signature Spoofing**: Creation of false commit signatures
- **Local File Access**: Unauthorized access to local certificate files
- **Supply Chain**: Compromise of dependencies or build process

### Mitigations

- **OS Keychain**: Private keys protected by OS-level encryption
- **File Permissions**: Strict file permissions on sensitive files
- **Local CA**: Self-contained CA eliminates external trust dependencies
- **Offline Operation**: No network communication eliminates remote attacks
- **Code Signing**: All releases will be signed (future enhancement)
- **Dependency Management**: Careful vetting of all dependencies

## Vulnerability Disclosure Policy

### Coordinated Disclosure

We follow responsible disclosure practices:

1. **Private Reporting**: Vulnerabilities reported privately first
2. **Verification Period**: Time to verify and develop fixes
3. **Coordinated Release**: Security updates released with advisories
4. **Public Disclosure**: Full details disclosed after fixes are available

### Recognition

We believe in recognizing security researchers who help improve our security:

- **Security Hall of Fame**: Recognition on our security page
- **CVE Assignment**: Help with CVE assignment for significant vulnerabilities
- **Coordination**: Work together on disclosure timeline and details

## Security Resources

### Documentation

- [OWASP Cryptographic Storage Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Cryptographic_Storage_Cheat_Sheet.html)
- [NIST Cryptographic Standards](https://csrc.nist.gov/projects/cryptographic-standards-and-guidelines)
- [Go Cryptography Documentation](https://golang.org/pkg/crypto/)

### Tools

- **Static Analysis**: We use gosec for static security analysis
- **Dependency Scanning**: We monitor dependencies for known vulnerabilities
- **Code Review**: All changes undergo security-focused code review

## Contact

For security-related questions or concerns:

- **Security Issues**: Use private vulnerability reporting
- **General Security Questions**: Open a GitHub Discussion
- **Documentation**: Refer to this security policy

---

**Remember**: Security is a shared responsibility. Please help us keep CertifyCLI secure by following best practices and reporting any security concerns promptly.
# Contributing to CertifyCLI

Thank you for your interest in contributing to CertifyCLI! This document provides guidelines and information for contributors.

## 🚀 Quick Start for Contributors

1. **Fork the repository** on GitHub
2. **Clone your fork** locally
3. **Create a feature branch** from `main`
4. **Make your changes** with tests
5. **Test thoroughly** using our test suite
6. **Submit a pull request** with clear description

## 🛠️ Development Setup

### Prerequisites

- Go 1.19 or later
- Git
- OS keychain support (for testing)

### Setup

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/certifycli.git
cd certifycli

# Install dependencies
go mod tidy

# Build the project
go build -o certifycli ./cmd/certifycli

# Run tests
./test-local-mode.sh
./certifycli test-crypto
./certifycli test-keyring
```

## 🧪 Testing

### Test Suite

We have comprehensive tests that must pass before any contribution:

```bash
# Full test suite
./test-local-mode.sh

# Component tests
./certifycli test-crypto
./certifycli test-keyring
./certifycli git test

# Demo (visual verification)
./demo-local-mode.sh
```

### Test Requirements

- All existing tests must pass
- New features must include tests
- Security-related changes require extra scrutiny
- Cross-platform compatibility must be maintained

## 📝 Code Style

### Go Code Style

- Follow standard Go formatting (`gofmt`)
- Use meaningful variable and function names
- Add comments for exported functions
- Handle errors appropriately
- Use Go modules for dependencies

### Example

```go
// GenerateKeyPair creates a new RSA key pair with the specified bit size.
// It returns the private key and any error encountered during generation.
func GenerateKeyPair(bits int) (*rsa.PrivateKey, error) {
    if bits < 2048 {
        return nil, fmt.Errorf("key size must be at least 2048 bits")
    }
    
    privateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return nil, fmt.Errorf("failed to generate key pair: %v", err)
    }
    
    return privateKey, nil
}
```

### Shell Script Style

- Use `#!/bin/bash` shebang
- Add error checking (`set -e` where appropriate)
- Use meaningful variable names
- Add comments for complex operations
- Make scripts executable (`chmod +x`)

## 🔒 Security Guidelines

### Security-First Approach

CertifyCLI handles cryptographic keys and certificates, so security is paramount:

1. **Never log sensitive data** (private keys, passwords, etc.)
2. **Use secure defaults** (strong key sizes, secure permissions)
3. **Validate all inputs** (especially cryptographic parameters)
4. **Handle errors securely** (don't leak information in error messages)
5. **Follow cryptographic best practices**

### Security Review Process

- All security-related changes require thorough review
- Cryptographic code must be reviewed by multiple contributors
- New cryptographic dependencies must be justified
- Security vulnerabilities should be reported privately first

## 📋 Contribution Types

### Bug Fixes

- Include reproduction steps in the issue
- Add tests that verify the fix
- Ensure fix doesn't break existing functionality

### New Features

- Discuss the feature in an issue first
- Ensure it aligns with project goals
- Include comprehensive tests
- Update documentation

### Documentation

- Keep documentation up-to-date with code changes
- Use clear, concise language
- Include examples where helpful
- Test documentation examples

### Platform Support

- Test on target platform thoroughly
- Consider platform-specific edge cases
- Update CI/CD if needed
- Document platform-specific requirements

## 🔄 Pull Request Process

### Before Submitting

1. **Test thoroughly** on your platform
2. **Run the full test suite**
3. **Update documentation** if needed
4. **Check code style** and formatting
5. **Write clear commit messages**

### PR Description Template

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Documentation update
- [ ] Security improvement
- [ ] Performance improvement

## Testing
- [ ] Ran full test suite
- [ ] Tested on [platform]
- [ ] Added new tests for changes

## Checklist
- [ ] Code follows project style guidelines
- [ ] Self-review completed
- [ ] Documentation updated
- [ ] No breaking changes (or clearly documented)
```

### Review Process

1. **Automated checks** must pass (when CI is set up)
2. **Manual review** by maintainers
3. **Security review** for security-related changes
4. **Testing** on multiple platforms if needed
5. **Approval** and merge

## 🐛 Reporting Issues

### Bug Reports

Use this template for bug reports:

```markdown
## Bug Description
Clear description of the bug

## Steps to Reproduce
1. Step one
2. Step two
3. Step three

## Expected Behavior
What should happen

## Actual Behavior
What actually happens

## Environment
- OS: [e.g., macOS 12.0, Ubuntu 20.04]
- Go version: [e.g., 1.19.0]
- CertifyCLI version: [e.g., commit hash]

## Additional Context
Any other relevant information
```

### Feature Requests

Use this template for feature requests:

```markdown
## Feature Description
Clear description of the proposed feature

## Use Case
Why is this feature needed?

## Proposed Solution
How should this feature work?

## Alternatives Considered
Other approaches you've considered

## Additional Context
Any other relevant information
```

## 🏗️ Architecture Guidelines

### Project Structure

```
certifycli/
├── cmd/certifycli/          # Main CLI application
├── internal/
│   ├── auth/                # Authentication and keyring
│   ├── ca/                  # Certificate Authority
│   ├── crypto/              # Cryptographic functions
│   ├── git/                 # Git integration
│   └── utils/               # Utility functions
├── docs/                    # Documentation
├── scripts/                 # Build and utility scripts
└── tests/                   # Test files
```

### Design Principles

1. **Security First**: All decisions prioritize security
2. **Simplicity**: Keep interfaces simple and intuitive
3. **Offline Operation**: No external dependencies required
4. **Cross-Platform**: Support major platforms
5. **Testability**: All code should be testable

## 📚 Resources

### Learning Resources

- [Go Documentation](https://golang.org/doc/)
- [Git Signing Documentation](https://git-scm.com/book/en/v2/Git-Tools-Signing-Your-Work)
- [X.509 Certificate Standards](https://tools.ietf.org/html/rfc5280)
- [Cryptographic Best Practices](https://cryptography.io/en/latest/faq/)

### Project Resources

- [Project Issues](https://github.com/CreatorOss/certifycli/issues)
- [Project Discussions](https://github.com/CreatorOss/certifycli/discussions)
- [Project Wiki](https://github.com/CreatorOss/certifycli/wiki)

## 💖 Supporting the Project

CertifyCLI is free and open source. If you find it valuable, please consider supporting its development:

### 🤝 GitHub Sponsors
[![Sponsor](https://img.shields.io/badge/Sponsor-%E2%9D%A4-%23db61a2)](https://github.com/sponsors/CreatorOss)

Support monthly development through [GitHub Sponsors](https://github.com/sponsors/CreatorOss). Even $3/month helps!

### 💰 PayPal Donation
[![PayPal](https://img.shields.io/badge/PayPal-00457C?logo=paypal&logoColor=white)](https://paypal.me/Sendec?country.x=ID&locale.x=id_ID)

Prefer one-time support? [Send via PayPal](https://paypal.me/Sendec?country.x=ID&locale.x=id_ID) to show your appreciation.

### ☕ Buy Me a Coffee
[![Buy Me a Coffee](https://img.shields.io/badge/Buy_Me_A_Coffee-FFDD00?logo=buy-me-a-coffee&logoColor=black)](https://buymeacoffee.com/creatoross)

Another way to support: [Buy me a coffee](https://buymeacoffee.com/creatoross) for quick appreciation.

### 🏢 Enterprise Support
Need custom features, priority support, or SLA guarantees? Contact us at **enterprise@certifycli.com**

## 🤝 Community

### Code of Conduct

We are committed to providing a welcoming and inclusive environment for all contributors. Please be respectful and professional in all interactions.

### Getting Help

- **Documentation**: Start with this README and inline help
- **Issues**: Search existing issues before creating new ones
- **Discussions**: Use GitHub Discussions for questions and ideas
- **Security**: Report security issues privately to maintainers
- **Sponsors**: Join our [sponsor community](https://github.com/sponsors/CreatorOss) for priority support

## 🎯 Roadmap

### Current Priorities

1. **Stability**: Bug fixes and reliability improvements
2. **Documentation**: Comprehensive guides and examples
3. **Testing**: Expanded test coverage and CI/CD
4. **Platform Support**: Enhanced cross-platform compatibility

### Future Considerations

- Team/organization features
- Integration with popular development tools
- Enhanced certificate management
- Performance optimizations

## 📄 License

By contributing to CertifyCLI, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to CertifyCLI! Together, we're making Git commit signing secure, simple, and accessible for everyone. 🚀
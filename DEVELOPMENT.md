# CertifyCLI Development Guide

## Quick Start

### Prerequisites
- Go 1.21 or later
- Node.js 18 or later
- npm or yarn

### Setup

1. **Clone and setup the project:**
   ```bash
   git clone <your-repo-url>
   cd certifycli
   ./test-setup.sh  # Verify setup
   ```

2. **Build the CLI:**
   ```bash
   go mod tidy
   go build -o certifycli ./cmd/certifycli
   ./certifycli --help
   ```

3. **Setup the server:**
   ```bash
   cd server
   npm install
   cp .env.example .env  # Edit as needed
   npm start
   ```

4. **Test the integration:**
   ```bash
   # In another terminal
   curl http://localhost:3001/api/health
   
   # Test CLI (will fail without server auth, but shows it's working)
   ./certifycli status
   ```

## Development Workflow

### CLI Development

The CLI is built in Go with a modular structure:

- `cmd/certifycli/main.go` - Entry point and command routing
- `internal/auth/` - Authentication and token management
- `internal/crypto/` - Cryptographic operations (key generation, signing)
- `internal/ca/` - Certificate Authority interactions
- `internal/utils/` - Helper functions and utilities

**Key commands for development:**
```bash
# Build and test
go build -o certifycli ./cmd/certifycli
./certifycli --help

# Run with debugging
go run ./cmd/certifycli status

# Test specific packages
go test ./internal/crypto
go test ./internal/auth
```

### Server Development

The server is built with Node.js/Express:

- `server/index.js` - Main server entry point
- `server/routes/api.js` - API route definitions
- `server/controllers/` - Business logic for endpoints
- `server/models/` - Database models and operations

**Key commands for development:**
```bash
cd server

# Development with auto-reload
npm run dev

# Production mode
npm start

# Test endpoints
curl -X POST http://localhost:3001/api/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

## Architecture Overview

### CLI → Server Communication

1. **Authentication Flow:**
   ```
   CLI → POST /api/login → Server
   CLI ← JWT Token ← Server
   CLI stores token in OS keychain
   ```

2. **Certificate Request Flow:**
   ```
   CLI generates keypair
   CLI creates CSR
   CLI → POST /api/certificate/request → Server
   CLI ← Signed Certificate ← Server
   ```

3. **Key Storage:**
   - Private keys stored in OS keychain (secure)
   - Public keys can be exported for sharing
   - Certificates stored both locally and on server

### Security Considerations

- **Private Key Protection:** Never transmitted, stored in OS keychain
- **JWT Tokens:** Short-lived (24h), stored securely
- **HTTPS:** Required for production (not implemented in MVP)
- **Rate Limiting:** Implemented on server
- **Input Validation:** Both client and server side

## Testing

### Manual Testing

1. **Test CLI without server:**
   ```bash
   ./certifycli --help
   ./certifycli status  # Should show "not logged in"
   ```

2. **Test server endpoints:**
   ```bash
   # Health check
   curl http://localhost:3001/api/health
   
   # Register user
   curl -X POST http://localhost:3001/api/register \
     -H "Content-Type: application/json" \
     -d '{"email":"dev@example.com","password":"devpass123"}'
   
   # Login
   curl -X POST http://localhost:3001/api/login \
     -H "Content-Type: application/json" \
     -d '{"email":"dev@example.com","password":"devpass123"}'
   ```

3. **Test full integration:**
   ```bash
   # Start server
   cd server && npm start &
   
   # Test CLI login (use test@example.com / password123)
   ./certifycli login
   ./certifycli status
   ./certifycli setup
   ```

### Automated Testing

```bash
# Run all tests
./test-setup.sh

# Go tests (when implemented)
go test ./...

# Server tests (when implemented)
cd server && npm test
```

## Common Development Tasks

### Adding a New CLI Command

1. Add command case in `cmd/certifycli/main.go`
2. Create handler function
3. Add any required internal package functions
4. Update help text

### Adding a New API Endpoint

1. Add route in `server/routes/api.js`
2. Create controller function in appropriate controller
3. Add any required model operations
4. Test with curl or Postman

### Adding New Crypto Operations

1. Add function to `internal/crypto/crypto.go`
2. Consider security implications
3. Add tests for crypto functions
4. Update CLI commands that use crypto

## Troubleshooting

### Common Issues

1. **"Go not found"**
   - Install Go from https://golang.org/doc/install
   - Ensure `go` is in your PATH

2. **"Module not found"**
   - Run `go mod tidy` to download dependencies
   - Check your Go version (requires 1.21+)

3. **"Server connection failed"**
   - Ensure server is running on port 3001
   - Check firewall settings
   - Verify server logs for errors

4. **"Keychain access denied"**
   - Grant keychain access when prompted
   - On Linux, ensure keyring service is running

### Debug Mode

Enable debug logging:
```bash
# CLI debug (when implemented)
DEBUG=1 ./certifycli status

# Server debug
DEBUG=1 npm start
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

### Code Style

- **Go:** Follow standard Go formatting (`go fmt`)
- **JavaScript:** Use consistent indentation (2 spaces)
- **Comments:** Document public functions and complex logic
- **Security:** Always validate inputs, never log sensitive data

## Deployment

### Building for Production

```bash
# Build optimized CLI binary
go build -ldflags="-s -w" -o certifycli ./cmd/certifycli

# Build for multiple platforms
GOOS=linux GOARCH=amd64 go build -o certifycli-linux-amd64 ./cmd/certifycli
GOOS=windows GOARCH=amd64 go build -o certifycli-windows-amd64.exe ./cmd/certifycli
GOOS=darwin GOARCH=amd64 go build -o certifycli-darwin-amd64 ./cmd/certifycli
```

### Server Deployment

```bash
cd server
npm install --production
NODE_ENV=production npm start
```

## Next Steps for MVP

1. **Implement real certificate signing** (replace mock CA)
2. **Add certificate validation** (chain verification, expiry checks)
3. **Implement Git integration** (automatic commit signing)
4. **Add HTTPS support** for production server
5. **Create proper database schema** (replace SQLite with PostgreSQL)
6. **Add comprehensive testing** (unit tests, integration tests)
7. **Implement certificate revocation** (CRL/OCSP)
8. **Add certificate renewal** (automatic renewal before expiry)
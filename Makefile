# CertifyCLI Makefile

.PHONY: help build test clean install server dev setup

# Default target
help:
	@echo "CertifyCLI Development Commands"
	@echo "==============================="
	@echo ""
	@echo "CLI Commands:"
	@echo "  build      Build the CLI binary"
	@echo "  test       Run tests"
	@echo "  install    Install CLI to system PATH"
	@echo "  clean      Clean build artifacts"
	@echo ""
	@echo "Server Commands:"
	@echo "  server     Start the server"
	@echo "  server-dev Start server in development mode"
	@echo "  server-setup Install server dependencies"
	@echo ""
	@echo "Development:"
	@echo "  setup      Complete project setup"
	@echo "  dev        Start both CLI and server for development"
	@echo "  check      Run all checks and tests"

# CLI targets
build:
	@echo "🔨 Building CertifyCLI..."
	go mod tidy
	go build -o certifycli ./cmd/certifycli
	@echo "✅ Build complete: ./certifycli"

test:
	@echo "🧪 Running tests..."
	go test -v ./...
	@echo "✅ Tests complete"

install: build
	@echo "📦 Installing CertifyCLI..."
	./scripts/install.sh

clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -f certifycli
	rm -f certifycli-*
	go clean
	@echo "✅ Clean complete"

# Server targets
server-setup:
	@echo "📦 Installing server dependencies..."
	cd server && npm install
	@echo "✅ Server setup complete"

server: server-setup
	@echo "🚀 Starting server..."
	cd server && npm start

server-dev: server-setup
	@echo "🚀 Starting server in development mode..."
	cd server && npm run dev

# Development targets
setup: build server-setup
	@echo "🎉 Project setup complete!"
	@echo ""
	@echo "Quick start:"
	@echo "  make dev        # Start development environment"
	@echo "  ./certifycli --help  # Test CLI"

dev:
	@echo "🚀 Starting development environment..."
	@echo "Starting server in background..."
	cd server && npm start &
	@echo "Server started. Test with: curl http://localhost:3001/api/health"
	@echo ""
	@echo "CLI is ready. Test with: ./certifycli --help"
	@echo ""
	@echo "To stop server: pkill -f 'node index.js'"

check:
	@echo "🔍 Running all checks..."
	./test-setup.sh
	@echo "✅ All checks passed"

# Build for multiple platforms
build-all:
	@echo "🔨 Building for all platforms..."
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o dist/certifycli-linux-amd64 ./cmd/certifycli
	GOOS=windows GOARCH=amd64 go build -o dist/certifycli-windows-amd64.exe ./cmd/certifycli
	GOOS=darwin GOARCH=amd64 go build -o dist/certifycli-darwin-amd64 ./cmd/certifycli
	GOOS=darwin GOARCH=arm64 go build -o dist/certifycli-darwin-arm64 ./cmd/certifycli
	@echo "✅ Multi-platform build complete in dist/"

# Docker targets
docker-build:
	@echo "🐳 Building Docker images..."
	docker build -t certifycli:latest .
	docker build -t certifycli-server:latest ./server
	@echo "✅ Docker images built"

# Release targets
release: clean test build-all
	@echo "🚀 Creating release..."
	@echo "✅ Release artifacts ready in dist/"
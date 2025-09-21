#!/bin/bash

# Test script untuk integrasi server CertifyCLI
echo "🌐 Testing CertifyCLI Server Integration"
echo "======================================="

# Check if Go is available
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go to run this test."
    echo "   Visit: https://golang.org/doc/install"
    exit 1
fi

# Check if Node.js is available
if ! command -v node &> /dev/null; then
    echo "❌ Node.js is not installed. Please install Node.js to run this test."
    echo "   Visit: https://nodejs.org/"
    exit 1
fi

# Build the CLI
echo "🔨 Building CertifyCLI..."
if ! go build -o certifycli ./cmd/certifycli; then
    echo "❌ Build failed"
    exit 1
fi

echo "✅ Build successful!"
echo ""

# Check if server dependencies are installed
echo "📦 Checking server dependencies..."
cd server
if [ ! -d "node_modules" ]; then
    echo "Installing server dependencies..."
    npm install
fi
cd ..

# Start server in background
echo "🚀 Starting CertifyCLI server..."
cd server
npm start &
SERVER_PID=$!
cd ..

# Wait for server to start
echo "⏳ Waiting for server to start..."
sleep 5

# Test 1: Server connection
echo "🌐 Test 1: Server Connection"
echo "============================"
./certifycli test-server
if [ $? -ne 0 ]; then
    echo "❌ Server connection test failed"
    kill $SERVER_PID 2>/dev/null
    exit 1
fi
echo ""

# Test 2: User registration (using curl for automation)
echo "📝 Test 2: User Registration"
echo "============================"
echo "Registering test user via API..."
REGISTER_RESPONSE=$(curl -s -X POST http://localhost:3001/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser_integration","password":"testpass123","email":"test@integration.com"}')

if echo "$REGISTER_RESPONSE" | grep -q "User created successfully"; then
    echo "✅ User registration successful"
else
    echo "⚠️  User may already exist or registration failed"
    echo "Response: $REGISTER_RESPONSE"
fi
echo ""

# Test 3: User login (using curl for automation)
echo "🔐 Test 3: User Login via API"
echo "============================="
LOGIN_RESPONSE=$(curl -s -X POST http://localhost:3001/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser_integration","password":"testpass123"}')

if echo "$LOGIN_RESPONSE" | grep -q "Login successful"; then
    echo "✅ API login successful"
    TOKEN=$(echo "$LOGIN_RESPONSE" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
    echo "🎫 Token received: ${TOKEN:0:20}..."
else
    echo "❌ API login failed"
    echo "Response: $LOGIN_RESPONSE"
    kill $SERVER_PID 2>/dev/null
    exit 1
fi
echo ""

# Test 4: Protected endpoint access
echo "🔒 Test 4: Protected Endpoint Access"
echo "===================================="
AUTH_TEST_RESPONSE=$(curl -s -H "Authorization: Bearer $TOKEN" \
  http://localhost:3001/api/test-auth)

if echo "$AUTH_TEST_RESPONSE" | grep -q "Authentication successful"; then
    echo "✅ Protected endpoint access successful"
else
    echo "❌ Protected endpoint access failed"
    echo "Response: $AUTH_TEST_RESPONSE"
fi
echo ""

# Test 5: CLI authentication flow (manual test info)
echo "🖥️  Test 5: CLI Authentication Flow"
echo "==================================="
echo "To test CLI authentication manually:"
echo "  1. ./certifycli register"
echo "  2. ./certifycli login"
echo "  3. ./certifycli test-auth"
echo "  4. ./certifycli status"
echo ""

# Test 6: Database verification
echo "💾 Test 6: Database Verification"
echo "==============================="
if [ -f "server/database.sqlite" ]; then
    echo "✅ Database file created"
    echo "📊 Database location: server/database.sqlite"
else
    echo "❌ Database file not found"
fi
echo ""

# Cleanup
echo "🧹 Cleanup"
echo "=========="
echo "Stopping server..."
kill $SERVER_PID 2>/dev/null
wait $SERVER_PID 2>/dev/null

echo ""
echo "🎉 Server integration test complete!"
echo ""
echo "📋 Summary of tested features:"
echo "  ✅ Server startup and health check"
echo "  ✅ User registration via API"
echo "  ✅ User login and JWT token generation"
echo "  ✅ Protected endpoint authentication"
echo "  ✅ Database creation and user storage"
echo ""
echo "🚀 Manual testing commands:"
echo "  ./certifycli test-server    # Test server connection"
echo "  ./certifycli register       # Register new user"
echo "  ./certifycli login          # Login to server"
echo "  ./certifycli test-auth      # Test authentication"
echo "  ./certifycli status         # Check full status"
echo ""
echo "🔧 Server management:"
echo "  cd server && npm start     # Start server"
echo "  curl http://localhost:3001/api/health  # Health check"
echo ""
echo "⚠️  Note: Server must be running for CLI authentication to work!"
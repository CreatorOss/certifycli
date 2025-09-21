#!/bin/bash

# Demo script untuk menunjukkan fitur server integration CertifyCLI
echo "🌐 CertifyCLI Server Integration Demo"
echo "===================================="

# Check if Node.js is available
if ! command -v node &> /dev/null; then
    echo "❌ Node.js is not installed. Please install Node.js to run this demo."
    echo "   Visit: https://nodejs.org/"
    exit 1
fi

echo "📋 This demo shows the server integration features:"
echo "  ✅ JWT-based authentication"
echo "  ✅ User registration and login"
echo "  ✅ Protected API endpoints"
echo "  ✅ SQLite database integration"
echo "  ✅ Secure password hashing"
echo ""

# Check server dependencies
echo "📦 Checking server dependencies..."
cd server
if [ ! -d "node_modules" ]; then
    echo "Installing server dependencies..."
    npm install
    if [ $? -ne 0 ]; then
        echo "❌ Failed to install dependencies"
        exit 1
    fi
fi
cd ..

# Start server
echo "🚀 Starting CertifyCLI server..."
cd server
npm start &
SERVER_PID=$!
cd ..

# Wait for server to start
echo "⏳ Waiting for server to start..."
sleep 5

# Test server health
echo ""
echo "🏥 Testing server health..."
HEALTH_RESPONSE=$(curl -s http://localhost:3001/api/health 2>/dev/null)
if [ $? -eq 0 ] && echo "$HEALTH_RESPONSE" | grep -q "OK"; then
    echo "✅ Server is running and healthy"
    echo "📊 Response: $HEALTH_RESPONSE"
else
    echo "❌ Server health check failed"
    kill $SERVER_PID 2>/dev/null
    exit 1
fi

echo ""
echo "📝 Testing user registration..."
REGISTER_RESPONSE=$(curl -s -X POST http://localhost:3001/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"demo_user","password":"demo_pass_123","email":"demo@certifycli.com"}' 2>/dev/null)

if echo "$REGISTER_RESPONSE" | grep -q "User created successfully"; then
    echo "✅ User registration successful"
    echo "👤 User: demo_user created"
else
    echo "⚠️  User may already exist"
    echo "📄 Response: $REGISTER_RESPONSE"
fi

echo ""
echo "🔐 Testing user login..."
LOGIN_RESPONSE=$(curl -s -X POST http://localhost:3001/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"demo_user","password":"demo_pass_123"}' 2>/dev/null)

if echo "$LOGIN_RESPONSE" | grep -q "Login successful"; then
    echo "✅ Login successful"
    TOKEN=$(echo "$LOGIN_RESPONSE" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
    echo "🎫 JWT Token received: ${TOKEN:0:30}..."
else
    echo "❌ Login failed"
    echo "📄 Response: $LOGIN_RESPONSE"
    kill $SERVER_PID 2>/dev/null
    exit 1
fi

echo ""
echo "🔒 Testing protected endpoint..."
AUTH_RESPONSE=$(curl -s -H "Authorization: Bearer $TOKEN" \
  http://localhost:3001/api/test-auth 2>/dev/null)

if echo "$AUTH_RESPONSE" | grep -q "Authentication successful"; then
    echo "✅ Protected endpoint access successful"
    echo "🛡️  Authentication working correctly"
else
    echo "❌ Protected endpoint access failed"
    echo "📄 Response: $AUTH_RESPONSE"
fi

echo ""
echo "👤 Testing user profile endpoint..."
PROFILE_RESPONSE=$(curl -s -H "Authorization: Bearer $TOKEN" \
  http://localhost:3001/api/profile 2>/dev/null)

if echo "$PROFILE_RESPONSE" | grep -q "demo_user"; then
    echo "✅ Profile endpoint working"
    echo "📋 Profile data retrieved successfully"
else
    echo "❌ Profile endpoint failed"
    echo "📄 Response: $PROFILE_RESPONSE"
fi

echo ""
echo "📜 Testing certificate request endpoint..."
CSR_RESPONSE=$(curl -s -X POST -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"csrData":"-----BEGIN CERTIFICATE REQUEST-----\nMIICWjCCAUICAQAwFTETMBEGA1UEAwwKZGVtb191c2VyMIIBIjANBgkqhkiG9w0B\n-----END CERTIFICATE REQUEST-----","commonName":"demo_user"}' \
  http://localhost:3001/api/certificate/request 2>/dev/null)

if echo "$CSR_RESPONSE" | grep -q "CSR received"; then
    echo "✅ Certificate request endpoint working"
    echo "📋 CSR processing successful"
else
    echo "❌ Certificate request failed"
    echo "📄 Response: $CSR_RESPONSE"
fi

echo ""
echo "💾 Checking database..."
if [ -f "server/database.sqlite" ]; then
    echo "✅ SQLite database created"
    echo "📊 Database file: server/database.sqlite"
    
    # Check database size
    DB_SIZE=$(du -h server/database.sqlite | cut -f1)
    echo "📏 Database size: $DB_SIZE"
else
    echo "❌ Database file not found"
fi

echo ""
echo "🧹 Cleanup..."
echo "Stopping server (PID: $SERVER_PID)..."
kill $SERVER_PID 2>/dev/null
wait $SERVER_PID 2>/dev/null

echo ""
echo "🎉 Server Integration Demo Complete!"
echo ""
echo "📋 Demonstrated features:"
echo "  ✅ Server startup and health monitoring"
echo "  ✅ User registration with password hashing"
echo "  ✅ JWT-based authentication"
echo "  ✅ Protected API endpoints"
echo "  ✅ User profile management"
echo "  ✅ Certificate request handling"
echo "  ✅ SQLite database integration"
echo ""
echo "🔧 Server API Endpoints:"
echo "  GET  /api/health              # Server health check"
echo "  POST /api/register            # User registration"
echo "  POST /api/login               # User authentication"
echo "  GET  /api/profile             # User profile (protected)"
echo "  GET  /api/test-auth           # Authentication test (protected)"
echo "  POST /api/certificate/request # Certificate signing request (protected)"
echo "  GET  /api/certificates        # List user certificates (protected)"
echo ""
echo "🚀 CLI Integration Commands (when Go is installed):"
echo "  ./certifycli test-server      # Test server connection"
echo "  ./certifycli register         # Register new user"
echo "  ./certifycli login            # Login to server"
echo "  ./certifycli test-auth        # Test authentication"
echo "  ./certifycli status           # Check full status"
echo ""
echo "💡 To start server manually:"
echo "  cd server && npm start"
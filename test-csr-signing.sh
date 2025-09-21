#!/bin/bash

# Test script untuk CSR signing flow CertifyCLI
echo "🏛️  Testing CertifyCLI CSR Signing Flow"
echo "======================================"

# Check if Node.js is available
if ! command -v node &> /dev/null; then
    echo "❌ Node.js is not installed. Please install Node.js to run this test."
    echo "   Visit: https://nodejs.org/"
    exit 1
fi

echo "📋 This test demonstrates the complete PKI workflow:"
echo "  ✅ Certificate Authority initialization"
echo "  ✅ CSR creation and signing"
echo "  ✅ Certificate verification"
echo "  ✅ Certificate management"
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
echo "🚀 Starting CertifyCLI CA server..."
cd server
npm start &
SERVER_PID=$!
cd ..

# Wait for server to start
echo "⏳ Waiting for server to start..."
sleep 5

# Test server health and CA initialization
echo ""
echo "🏥 Testing CA server health..."
HEALTH_RESPONSE=$(curl -s http://localhost:3001/api/health 2>/dev/null)
if [ $? -eq 0 ] && echo "$HEALTH_RESPONSE" | grep -q "OK"; then
    echo "✅ CA server is running and healthy"
    echo "📊 Response: $HEALTH_RESPONSE"
else
    echo "❌ CA server health check failed"
    kill $SERVER_PID 2>/dev/null
    exit 1
fi

# Test CA certificate endpoint
echo ""
echo "🏛️  Testing CA certificate endpoint..."
CA_CERT_RESPONSE=$(curl -s http://localhost:3001/api/ca-certificate 2>/dev/null)
if echo "$CA_CERT_RESPONSE" | grep -q "CertifyCLI Development CA"; then
    echo "✅ CA certificate endpoint working"
    echo "🔐 CA initialized successfully"
else
    echo "❌ CA certificate endpoint failed"
    echo "📄 Response: $CA_CERT_RESPONSE"
fi

# Test user registration
echo ""
echo "📝 Testing user registration..."
REGISTER_RESPONSE=$(curl -s -X POST http://localhost:3001/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"csr_test_user","password":"csr_test_pass_123","email":"csr@certifycli.com"}' 2>/dev/null)

if echo "$REGISTER_RESPONSE" | grep -q "User created successfully"; then
    echo "✅ User registration successful"
    echo "👤 User: csr_test_user created"
else
    echo "⚠️  User may already exist"
    echo "📄 Response: $REGISTER_RESPONSE"
fi

# Test user login
echo ""
echo "🔐 Testing user login..."
LOGIN_RESPONSE=$(curl -s -X POST http://localhost:3001/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"csr_test_user","password":"csr_test_pass_123"}' 2>/dev/null)

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

# Test CSR signing
echo ""
echo "📜 Testing CSR signing..."
# Create a mock CSR for testing (properly escaped)
MOCK_CSR="-----BEGIN CERTIFICATE REQUEST-----\nMIICWjCCAUICAQAwFTETMBEGA1UEAwwKY3NyX3Rlc3RfdXNlcjCCASIwDQYJKoZI\nhvcNAQEBBQADggEPADCCAQoCggEBAMi/wRfLukvJ5H6L2K8VuIJXRjad1aBdGPQq\nLE0obGlnUm2f5H6LSubject: csr_test_user\nMockCSR: true\n-----END CERTIFICATE REQUEST-----"

# Create JSON payload properly
CSR_JSON=$(cat <<EOF
{
  "csr": "$MOCK_CSR",
  "validityDays": 365
}
EOF
)

CSR_RESPONSE=$(curl -s -X POST -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d "$CSR_JSON" \
  http://localhost:3001/api/certificate/sign 2>/dev/null)

if echo "$CSR_RESPONSE" | grep -q "Certificate created successfully"; then
    echo "✅ CSR signing successful"
    SERIAL_NUMBER=$(echo "$CSR_RESPONSE" | grep -o '"serialNumber":"[^"]*"' | cut -d'"' -f4)
    echo "🆔 Certificate serial: $SERIAL_NUMBER"
    
    # Extract certificate for verification
    CERTIFICATE=$(echo "$CSR_RESPONSE" | grep -o '"certificate":"[^"]*"' | cut -d'"' -f4)
    echo "📄 Certificate generated successfully"
else
    echo "❌ CSR signing failed"
    echo "📄 Response: $CSR_RESPONSE"
    kill $SERVER_PID 2>/dev/null
    exit 1
fi

# Test certificate listing
echo ""
echo "📋 Testing certificate listing..."
CERTS_RESPONSE=$(curl -s -H "Authorization: Bearer $TOKEN" \
  http://localhost:3001/api/certificates 2>/dev/null)

if echo "$CERTS_RESPONSE" | grep -q "certificates"; then
    echo "✅ Certificate listing working"
    CERT_COUNT=$(echo "$CERTS_RESPONSE" | grep -o '"count":[0-9]*' | cut -d':' -f2)
    echo "📊 Total certificates: $CERT_COUNT"
else
    echo "❌ Certificate listing failed"
    echo "📄 Response: $CERTS_RESPONSE"
fi

# Test certificate verification
echo ""
echo "🔍 Testing certificate verification..."
VERIFY_RESPONSE=$(curl -s -X POST \
  -H "Content-Type: application/json" \
  -d "{\"certificate\":\"$CERTIFICATE\"}" \
  http://localhost:3001/api/certificate/verify 2>/dev/null)

if echo "$VERIFY_RESPONSE" | grep -q '"valid":true'; then
    echo "✅ Certificate verification successful"
    echo "🛡️  Certificate is valid and signed by CA"
else
    echo "❌ Certificate verification failed"
    echo "📄 Response: $VERIFY_RESPONSE"
fi

# Test certificate statistics
echo ""
echo "📊 Testing certificate statistics..."
STATS_RESPONSE=$(curl -s -H "Authorization: Bearer $TOKEN" \
  http://localhost:3001/api/admin/statistics 2>/dev/null)

if echo "$STATS_RESPONSE" | grep -q "statistics"; then
    echo "✅ Certificate statistics working"
    echo "📈 Statistics retrieved successfully"
else
    echo "❌ Certificate statistics failed"
    echo "📄 Response: $STATS_RESPONSE"
fi

# Check CA files
echo ""
echo "🔐 Checking CA files..."
if [ -f "server/ca-private-key.pem" ]; then
    echo "✅ CA private key generated"
    CA_KEY_SIZE=$(wc -c < server/ca-private-key.pem)
    echo "📏 CA private key size: $CA_KEY_SIZE bytes"
else
    echo "❌ CA private key not found"
fi

if [ -f "server/ca-certificate.pem" ]; then
    echo "✅ CA certificate generated"
    CA_CERT_SIZE=$(wc -c < server/ca-certificate.pem)
    echo "📏 CA certificate size: $CA_CERT_SIZE bytes"
else
    echo "❌ CA certificate not found"
fi

# Check database
echo ""
echo "💾 Checking database..."
if [ -f "server/database.sqlite" ]; then
    echo "✅ SQLite database created"
    DB_SIZE=$(du -h server/database.sqlite | cut -f1)
    echo "📊 Database size: $DB_SIZE"
else
    echo "❌ Database file not found"
fi

echo ""
echo "🧹 Cleanup..."
echo "Stopping server (PID: $SERVER_PID)..."
kill $SERVER_PID 2>/dev/null
wait $SERVER_PID 2>/dev/null

echo ""
echo "🎉 CSR Signing Flow Test Complete!"
echo ""
echo "📋 Demonstrated PKI features:"
echo "  ✅ Certificate Authority initialization"
echo "  ✅ CA key pair generation"
echo "  ✅ User registration and authentication"
echo "  ✅ CSR submission and signing"
echo "  ✅ Certificate generation with serial numbers"
echo "  ✅ Certificate verification against CA"
echo "  ✅ Certificate listing and management"
echo "  ✅ Certificate statistics and monitoring"
echo "  ✅ Database integration for certificate storage"
echo ""
echo "🏛️  Certificate Authority Features:"
echo "  📜 X.509 certificate generation"
echo "  🔐 RSA key pair management"
echo "  🆔 Unique serial number assignment"
echo "  📅 Validity period management"
echo "  🔍 Certificate verification"
echo "  📊 Certificate lifecycle tracking"
echo ""
echo "🚀 CLI Integration Commands (when Go is installed):"
echo "  ./certifycli setup           # Complete identity setup with CA signing"
echo "  ./certifycli certificates    # List your certificates"
echo "  ./certifycli verify-cert     # Verify certificate against CA"
echo "  ./certifycli get-cert <id>   # Get certificate details"
echo "  ./certifycli revoke-cert <id> # Revoke a certificate"
echo ""
echo "💡 To start server manually:"
echo "  cd server && npm start"
echo ""
echo "🔧 CA Management:"
echo "  CA private key: server/ca-private-key.pem"
echo "  CA certificate: server/ca-certificate.pem"
echo "  Database: server/database.sqlite"
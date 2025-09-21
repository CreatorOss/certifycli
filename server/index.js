require('dotenv').config();
const express = require('express');
const cors = require('cors');
const jwt = require('jsonwebtoken');
const User = require('./models/User');
const Certificate = require('./models/Certificate');
const cryptoUtils = require('./utils/crypto');
const db = require('./models/database');

const app = express();
const port = process.env.PORT || 3001;
const JWT_SECRET = process.env.JWT_SECRET || 'development-secret-change-in-production';

// Middleware
app.use(cors());
app.use(express.json({ limit: '10mb' }));

// Basic Health Check Route
app.get('/api/health', (req, res) => {
    res.json({ 
        status: 'OK', 
        message: 'CertifyCLI Server is running!',
        timestamp: new Date().toISOString(),
        version: '0.3.0',
        ca: 'CertifyCLI Development CA'
    });
});

// Get CA Certificate (public endpoint)
app.get('/api/ca-certificate', (req, res) => {
    try {
        const caCert = cryptoUtils.getCACertificate();
        res.json({
            certificate: caCert,
            issuer: 'CertifyCLI Development CA',
            message: 'CA certificate for verification'
        });
    } catch (error) {
        console.error('CA certificate error:', error);
        res.status(500).json({ error: 'Failed to retrieve CA certificate' });
    }
});

// User Registration
app.post('/api/register', async (req, res) => {
    try {
        const { username, password, email } = req.body;
        
        if (!username || !password) {
            return res.status(400).json({ error: 'Username and password are required' });
        }

        if (username.length < 3) {
            return res.status(400).json({ error: 'Username must be at least 3 characters long' });
        }

        if (password.length < 6) {
            return res.status(400).json({ error: 'Password must be at least 6 characters long' });
        }

        const user = await User.create(username, password, email);
        res.status(201).json({ 
            message: 'User created successfully',
            user: { id: user.id, username: user.username, email: user.email }
        });
    } catch (error) {
        console.error('Registration error:', error);
        if (error.message.includes('UNIQUE constraint failed')) {
            res.status(409).json({ error: 'Username already exists' });
        } else {
            res.status(500).json({ error: 'Internal server error' });
        }
    }
});

// User Login
app.post('/api/login', async (req, res) => {
    try {
        const { username, password } = req.body;
        
        if (!username || !password) {
            return res.status(400).json({ error: 'Username and password are required' });
        }

        const isValid = await User.verifyPassword(username, password);
        if (!isValid) {
            return res.status(401).json({ error: 'Invalid credentials' });
        }

        const user = await User.findByUsername(username);
        
        // Create JWT token
        const token = jwt.sign(
            { 
                username: user.username, 
                userId: user.id 
            }, 
            JWT_SECRET, 
            { expiresIn: '7d' }
        );
        
        res.json({ 
            message: 'Login successful',
            token,
            user: { 
                id: user.id,
                username: user.username,
                email: user.email 
            }
        });
    } catch (error) {
        console.error('Login error:', error);
        res.status(500).json({ error: 'Internal server error' });
    }
});

// Protected route - Get user profile
app.get('/api/profile', authenticateToken, async (req, res) => {
    try {
        const user = await User.findById(req.user.userId);
        if (!user) {
            return res.status(404).json({ error: 'User not found' });
        }

        res.json({ 
            user: {
                id: user.id,
                username: user.username,
                email: user.email,
                created_at: user.created_at
            }
        });
    } catch (error) {
        console.error('Profile error:', error);
        res.status(500).json({ error: 'Internal server error' });
    }
});

// Certificate Signing Request endpoint - MAIN FEATURE
app.post('/api/certificate/sign', authenticateToken, async (req, res) => {
    try {
        const { csr: csrData, validityDays = 365 } = req.body;
        
        if (!csrData) {
            return res.status(400).json({ error: 'CSR data is required' });
        }

        // Validate CSR format
        if (!csrData.includes('BEGIN CERTIFICATE REQUEST') || !csrData.includes('END CERTIFICATE REQUEST')) {
            return res.status(400).json({ error: 'Invalid CSR format' });
        }

        // Get user from database
        const user = await User.findByUsername(req.user.username);
        if (!user) {
            return res.status(404).json({ error: 'User not found' });
        }

        // Sign the CSR and create certificate
        const certificate = await Certificate.createFromCSR(
            user.id, 
            csrData, 
            req.user.username,
            validityDays
        );

        console.log(`✅ Certificate signed for user: ${req.user.username}, Serial: ${certificate.serialNumber}`);

        res.json({ 
            message: 'Certificate created successfully',
            certificate: certificate.certificate,
            serialNumber: certificate.serialNumber,
            validFrom: certificate.validFrom,
            validTo: certificate.validTo,
            status: certificate.status,
            commonName: req.user.username
        });
    } catch (error) {
        console.error('Certificate signing error:', error);
        if (error.message.includes('Failed to sign CSR') || error.message.includes('Invalid CSR')) {
            res.status(400).json({ error: 'Invalid CSR format or signing failed' });
        } else {
            res.status(500).json({ error: 'Internal server error' });
        }
    }
});

// Get user's certificates
app.get('/api/certificates', authenticateToken, async (req, res) => {
    try {
        const user = await User.findByUsername(req.user.username);
        if (!user) {
            return res.status(404).json({ error: 'User not found' });
        }

        const certificates = await Certificate.findByUserId(user.id);
        res.json({ 
            certificates,
            count: certificates.length
        });
    } catch (error) {
        console.error('Get certificates error:', error);
        res.status(500).json({ error: 'Internal server error' });
    }
});

// Get specific certificate by ID
app.get('/api/certificate/:id', authenticateToken, async (req, res) => {
    try {
        const { id } = req.params;
        const certificate = await Certificate.findById(id);
        
        if (!certificate) {
            return res.status(404).json({ error: 'Certificate not found' });
        }

        // Check if certificate belongs to the authenticated user
        const user = await User.findByUsername(req.user.username);
        if (certificate.user_id !== user.id) {
            return res.status(403).json({ error: 'Access denied' });
        }

        res.json({ certificate });
    } catch (error) {
        console.error('Get certificate error:', error);
        res.status(500).json({ error: 'Internal server error' });
    }
});

// Revoke certificate
app.post('/api/certificate/:id/revoke', authenticateToken, async (req, res) => {
    try {
        const { id } = req.params;
        const { reason = 'user_request' } = req.body;
        
        const certificate = await Certificate.findById(id);
        if (!certificate) {
            return res.status(404).json({ error: 'Certificate not found' });
        }

        // Check if certificate belongs to the authenticated user
        const user = await User.findByUsername(req.user.username);
        if (certificate.user_id !== user.id) {
            return res.status(403).json({ error: 'Access denied' });
        }

        const result = await Certificate.revoke(certificate.serial_number, reason);
        
        console.log(`🚫 Certificate revoked: ${certificate.serial_number} by user: ${req.user.username}`);
        
        res.json({
            message: 'Certificate revoked successfully',
            serialNumber: result.serialNumber,
            status: result.status,
            reason: result.reason
        });
    } catch (error) {
        console.error('Revoke certificate error:', error);
        res.status(500).json({ error: 'Internal server error' });
    }
});

// Verify certificate
app.post('/api/certificate/verify', (req, res) => {
    try {
        const { certificate } = req.body;
        
        if (!certificate) {
            return res.status(400).json({ error: 'Certificate data is required' });
        }

        const isValid = Certificate.verify(certificate);
        
        res.json({
            valid: isValid,
            message: isValid ? 'Certificate is valid' : 'Certificate is invalid or not signed by this CA',
            ca: 'CertifyCLI Development CA'
        });
    } catch (error) {
        console.error('Certificate verification error:', error);
        res.status(500).json({ error: 'Internal server error' });
    }
});

// Get certificate statistics (admin endpoint)
app.get('/api/admin/statistics', authenticateToken, async (req, res) => {
    try {
        const stats = await Certificate.getStatistics();
        res.json({
            statistics: stats,
            timestamp: new Date().toISOString()
        });
    } catch (error) {
        console.error('Statistics error:', error);
        res.status(500).json({ error: 'Internal server error' });
    }
});

// Test protected endpoint
app.get('/api/test-auth', authenticateToken, (req, res) => {
    res.json({ 
        message: 'Authentication successful!',
        user: req.user,
        timestamp: new Date().toISOString()
    });
});

// Middleware to authenticate JWT token
function authenticateToken(req, res, next) {
    const authHeader = req.headers['authorization'];
    const token = authHeader && authHeader.split(' ')[1]; // Bearer TOKEN

    if (!token) {
        return res.status(401).json({ error: 'Access token required' });
    }

    jwt.verify(token, JWT_SECRET, (err, user) => {
        if (err) {
            console.error('Token verification error:', err);
            return res.status(403).json({ error: 'Invalid or expired token' });
        }
        req.user = user;
        next();
    });
}

// Error handling middleware
app.use((error, req, res, next) => {
    console.error('Unhandled error:', error);
    res.status(500).json({ error: 'Internal server error' });
});

// 404 handler
app.use('*', (req, res) => {
    res.status(404).json({
        error: 'Not Found',
        message: 'The requested endpoint does not exist'
    });
});

app.listen(port, () => {
    console.log(`🚀 CertifyCLI server listening on port ${port}`);
    console.log(`📊 Health check: http://localhost:${port}/api/health`);
    console.log(`🔐 Environment: ${process.env.NODE_ENV || 'development'}`);
    console.log(`🏛️  Certificate Authority: CertifyCLI Development CA`);
});
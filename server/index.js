require('dotenv').config();
const express = require('express');
const cors = require('cors');
const jwt = require('jsonwebtoken');
const User = require('./models/User');
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
        version: '0.2.0'
    });
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

// Certificate Signing Request endpoint
app.post('/api/certificate/request', authenticateToken, async (req, res) => {
    try {
        const { csrData, commonName } = req.body;
        
        if (!csrData) {
            return res.status(400).json({ error: 'CSR data is required' });
        }

        if (!commonName) {
            return res.status(400).json({ error: 'Common name is required' });
        }

        // Store CSR in database
        db.run(
            'INSERT INTO certificate_requests (user_id, csr_data) VALUES (?, ?)',
            [req.user.userId, csrData],
            function(err) {
                if (err) {
                    console.error('Database error:', err);
                    return res.status(500).json({ error: 'Failed to store CSR' });
                }

                // For now, just acknowledge receipt
                // In production, we would actually sign the CSR and return a certificate
                res.json({ 
                    message: 'CSR received and stored successfully',
                    requestId: this.lastID,
                    status: 'pending',
                    commonName: commonName,
                    note: 'Certificate signing not yet implemented - this is a demo response'
                });
            }
        );
    } catch (error) {
        console.error('CSR error:', error);
        res.status(500).json({ error: 'Internal server error' });
    }
});

// List user's certificates
app.get('/api/certificates', authenticateToken, (req, res) => {
    db.all(
        'SELECT * FROM certificates WHERE user_id = ? ORDER BY created_at DESC',
        [req.user.userId],
        (err, rows) => {
            if (err) {
                console.error('Database error:', err);
                return res.status(500).json({ error: 'Failed to retrieve certificates' });
            }

            res.json({
                certificates: rows,
                count: rows.length
            });
        }
    );
});

// List user's certificate requests
app.get('/api/certificate-requests', authenticateToken, (req, res) => {
    db.all(
        'SELECT * FROM certificate_requests WHERE user_id = ? ORDER BY created_at DESC',
        [req.user.userId],
        (err, rows) => {
            if (err) {
                console.error('Database error:', err);
                return res.status(500).json({ error: 'Failed to retrieve certificate requests' });
            }

            res.json({
                requests: rows,
                count: rows.length
            });
        }
    );
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
});
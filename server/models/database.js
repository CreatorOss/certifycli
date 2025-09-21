const sqlite3 = require('sqlite3').verbose();
const path = require('path');
const fs = require('fs');

const dbPath = path.join(__dirname, '..', 'database.sqlite');

// Ensure database directory exists
const dbDir = path.dirname(dbPath);
if (!fs.existsSync(dbDir)) {
    fs.mkdirSync(dbDir, { recursive: true });
}

const db = new sqlite3.Database(dbPath, (err) => {
    if (err) {
        console.error('Error opening database:', err);
    } else {
        console.log('Connected to SQLite database');
        initializeDatabase();
    }
});

function initializeDatabase() {
    // Create users table
    db.run(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT UNIQUE NOT NULL,
            password_hash TEXT NOT NULL,
            email TEXT UNIQUE,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )
    `);

    // Create certificates table
    db.run(`
        CREATE TABLE IF NOT EXISTS certificates (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            common_name TEXT NOT NULL,
            certificate_data TEXT NOT NULL,
            serial_number TEXT UNIQUE NOT NULL,
            valid_from DATETIME NOT NULL,
            valid_to DATETIME NOT NULL,
            status TEXT DEFAULT 'active',
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users (id)
        )
    `);

    // Create certificate_requests table (for CSRs)
    db.run(`
        CREATE TABLE IF NOT EXISTS certificate_requests (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            csr_data TEXT NOT NULL,
            status TEXT DEFAULT 'pending',
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users (id)
        )
    `);

    // Create a default test user for development
    const bcrypt = require('bcryptjs');
    const defaultPassword = bcrypt.hashSync('testpass123', 12);
    
    db.run(`
        INSERT OR IGNORE INTO users (username, password_hash, email) 
        VALUES ('testuser', ?, 'test@certifycli.com')
    `, [defaultPassword], function(err) {
        if (err) {
            console.error('Error creating default user:', err);
        } else if (this.changes > 0) {
            console.log('✅ Default test user created: testuser / testpass123');
        }
    });
}

module.exports = db;
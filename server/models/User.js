const db = require('./database');
const bcrypt = require('bcryptjs');

class User {
    static async create(username, password, email = null) {
        const passwordHash = await bcrypt.hash(password, 12);
        return new Promise((resolve, reject) => {
            db.run(
                'INSERT INTO users (username, password_hash, email) VALUES (?, ?, ?)',
                [username, passwordHash, email],
                function(err) {
                    if (err) reject(err);
                    else resolve({ id: this.lastID, username, email });
                }
            );
        });
    }

    static async findByUsername(username) {
        return new Promise((resolve, reject) => {
            db.get(
                'SELECT * FROM users WHERE username = ?',
                [username],
                (err, row) => {
                    if (err) reject(err);
                    else resolve(row);
                }
            );
        });
    }

    static async findById(id) {
        return new Promise((resolve, reject) => {
            db.get(
                'SELECT * FROM users WHERE id = ?',
                [id],
                (err, row) => {
                    if (err) reject(err);
                    else resolve(row);
                }
            );
        });
    }

    static async verifyPassword(username, password) {
        const user = await this.findByUsername(username);
        if (!user) return false;
        
        return await bcrypt.compare(password, user.password_hash);
    }

    static async list(limit = 50, offset = 0) {
        return new Promise((resolve, reject) => {
            db.all(
                'SELECT id, username, email, created_at FROM users ORDER BY created_at DESC LIMIT ? OFFSET ?',
                [limit, offset],
                (err, rows) => {
                    if (err) reject(err);
                    else resolve(rows);
                }
            );
        });
    }
}

module.exports = User;
const db = require('./database');
const cryptoUtils = require('../utils/crypto');

class Certificate {
    static async createFromCSR(userId, csrData, commonName, validityDays = 365) {
        try {
            // Sign the CSR
            const signedCert = cryptoUtils.signCSR(csrData, commonName, validityDays);
            
            return new Promise((resolve, reject) => {
                db.run(
                    `INSERT INTO certificates 
                     (user_id, common_name, certificate_data, serial_number, valid_from, valid_to, status) 
                     VALUES (?, ?, ?, ?, ?, ?, ?)`,
                    [
                        userId,
                        commonName,
                        signedCert.certificate,
                        signedCert.serialNumber,
                        signedCert.validFrom,
                        signedCert.validTo,
                        'active'
                    ],
                    function(err) {
                        if (err) {
                            console.error('Database error:', err);
                            reject(err);
                        } else {
                            resolve({
                                id: this.lastID,
                                certificate: signedCert.certificate,
                                serialNumber: signedCert.serialNumber,
                                validFrom: signedCert.validFrom,
                                validTo: signedCert.validTo,
                                status: 'active'
                            });
                        }
                    }
                );
            });
        } catch (error) {
            throw new Error(`Failed to create certificate: ${error.message}`);
        }
    }

    static async findByUserId(userId) {
        return new Promise((resolve, reject) => {
            db.all(
                'SELECT * FROM certificates WHERE user_id = ? ORDER BY created_at DESC',
                [userId],
                (err, rows) => {
                    if (err) reject(err);
                    else resolve(rows);
                }
            );
        });
    }

    static async findById(id) {
        return new Promise((resolve, reject) => {
            db.get(
                'SELECT * FROM certificates WHERE id = ?',
                [id],
                (err, row) => {
                    if (err) reject(err);
                    else resolve(row);
                }
            );
        });
    }

    static async findBySerial(serialNumber) {
        return new Promise((resolve, reject) => {
            db.get(
                'SELECT * FROM certificates WHERE serial_number = ?',
                [serialNumber],
                (err, row) => {
                    if (err) reject(err);
                    else resolve(row);
                }
            );
        });
    }

    static async revoke(serialNumber, reason = 'unspecified') {
        return new Promise((resolve, reject) => {
            db.run(
                'UPDATE certificates SET status = ?, updated_at = CURRENT_TIMESTAMP WHERE serial_number = ?',
                ['revoked', serialNumber],
                function(err) {
                    if (err) reject(err);
                    else resolve({ 
                        changes: this.changes,
                        serialNumber: serialNumber,
                        status: 'revoked',
                        reason: reason
                    });
                }
            );
        });
    }

    static async getStatistics() {
        return new Promise((resolve, reject) => {
            db.all(
                `SELECT 
                    status,
                    COUNT(*) as count
                 FROM certificates 
                 GROUP BY status`,
                [],
                (err, rows) => {
                    if (err) reject(err);
                    else {
                        const stats = {
                            total: 0,
                            active: 0,
                            revoked: 0,
                            expired: 0
                        };
                        
                        rows.forEach(row => {
                            stats.total += row.count;
                            stats[row.status] = row.count;
                        });
                        
                        resolve(stats);
                    }
                }
            );
        });
    }

    static async verify(certificatePem) {
        try {
            return cryptoUtils.verifyCertificate(certificatePem);
        } catch (error) {
            console.error('Certificate verification error:', error);
            return false;
        }
    }
}

module.exports = Certificate;
const crypto = require('crypto');
const fs = require('fs');
const path = require('path');

class CryptoUtils {
    constructor() {
        this.initializeCA();
    }

    // Initialize CA key pair if not exists
    initializeCA() {
        const caKeyPath = path.join(__dirname, '..', 'ca-private-key.pem');
        const caCertPath = path.join(__dirname, '..', 'ca-certificate.pem');

        if (!fs.existsSync(caKeyPath) || !fs.existsSync(caCertPath)) {
            console.log('🔧 Generating CA key pair for development...');
            this.generateCAKeyPair();
        }

        this.caPrivateKey = fs.readFileSync(caKeyPath, 'utf8');
        this.caCertificate = fs.readFileSync(caCertPath, 'utf8');
        console.log('✅ CA initialized successfully');
    }

    // Generate CA key pair for development
    generateCAKeyPair() {
        const { privateKey, publicKey } = crypto.generateKeyPairSync('rsa', {
            modulusLength: 2048,
            publicKeyEncoding: {
                type: 'spki',
                format: 'pem'
            },
            privateKeyEncoding: {
                type: 'pkcs8',
                format: 'pem'
            }
        });

        // Create self-signed CA certificate
        const cert = this.createSelfSignedCertificate(privateKey, publicKey, {
            commonName: 'CertifyCLI Development CA',
            isCA: true,
            validityDays: 3650 // 10 years for CA
        });

        // Save CA files
        const caKeyPath = path.join(__dirname, '..', 'ca-private-key.pem');
        const caCertPath = path.join(__dirname, '..', 'ca-certificate.pem');

        fs.writeFileSync(caKeyPath, privateKey, { mode: 0o600 });
        fs.writeFileSync(caCertPath, cert, { mode: 0o644 });

        console.log('✅ CA key pair generated and saved');
    }

    // Create a self-signed certificate (for CA)
    createSelfSignedCertificate(privateKey, publicKey, options) {
        const { commonName, isCA = false, validityDays = 365 } = options;

        // Create certificate template
        const cert = {
            version: 3,
            serialNumber: this.generateSerialNumber(),
            subject: { CN: commonName },
            issuer: { CN: commonName }, // Self-signed
            notBefore: new Date(),
            notAfter: new Date(Date.now() + validityDays * 24 * 60 * 60 * 1000),
            publicKey: publicKey,
            extensions: []
        };

        if (isCA) {
            cert.extensions.push({
                name: 'basicConstraints',
                cA: true,
                critical: true
            });
            cert.extensions.push({
                name: 'keyUsage',
                keyCertSign: true,
                cRLSign: true,
                critical: true
            });
        }

        // For now, return a mock certificate in PEM format
        // In production, you'd use a proper X.509 library like node-forge
        return this.createMockCertificatePEM(cert, privateKey);
    }

    // Parse CSR (simplified version)
    parseCSR(csrPem) {
        // Extract the base64 content between BEGIN and END
        const csrMatch = csrPem.match(/-----BEGIN CERTIFICATE REQUEST-----\s*([\s\S]*?)\s*-----END CERTIFICATE REQUEST-----/);
        if (!csrMatch) {
            throw new Error('Invalid CSR format');
        }

        // For this demo, we'll extract basic info from the CSR
        // In production, use a proper ASN.1 parser
        return {
            subject: { CN: 'extracted_from_csr' },
            publicKey: 'extracted_public_key',
            valid: true
        };
    }

    // Sign CSR with CA private key
    signCSR(csrPem, commonName, validityDays = 365) {
        try {
            // Parse the CSR
            const csr = this.parseCSR(csrPem);
            
            if (!csr.valid) {
                throw new Error('Invalid CSR');
            }

            // Create certificate
            const serialNumber = this.generateSerialNumber();
            const notBefore = new Date();
            const notAfter = new Date(Date.now() + validityDays * 24 * 60 * 60 * 1000);

            const cert = {
                version: 3,
                serialNumber: serialNumber,
                subject: { CN: commonName },
                issuer: { CN: 'CertifyCLI Development CA' },
                notBefore: notBefore,
                notAfter: notAfter,
                publicKey: csr.publicKey,
                extensions: [
                    {
                        name: 'basicConstraints',
                        cA: false,
                        critical: true
                    },
                    {
                        name: 'keyUsage',
                        digitalSignature: true,
                        keyEncipherment: true,
                        critical: true
                    },
                    {
                        name: 'extKeyUsage',
                        serverAuth: true,
                        clientAuth: true,
                        codeSigning: true
                    }
                ]
            };

            // Create certificate PEM
            const certPem = this.createMockCertificatePEM(cert, this.caPrivateKey);

            return {
                certificate: certPem,
                serialNumber: serialNumber,
                validFrom: notBefore.toISOString(),
                validTo: notAfter.toISOString()
            };
        } catch (error) {
            throw new Error(`Failed to sign CSR: ${error.message}`);
        }
    }

    // Generate unique serial number
    generateSerialNumber() {
        return crypto.randomBytes(8).toString('hex').toUpperCase();
    }

    // Create mock certificate PEM (simplified for demo)
    createMockCertificatePEM(cert, signingKey) {
        // This is a simplified mock certificate
        // In production, use node-forge or similar library for proper X.509 generation
        const certData = {
            version: cert.version,
            serialNumber: cert.serialNumber,
            subject: cert.subject,
            issuer: cert.issuer,
            notBefore: cert.notBefore.toISOString(),
            notAfter: cert.notAfter.toISOString(),
            extensions: cert.extensions,
            timestamp: new Date().toISOString()
        };

        // Create a signature of the certificate data
        const dataToSign = JSON.stringify(certData);
        const signature = crypto.sign('sha256', Buffer.from(dataToSign), {
            key: signingKey,
            padding: crypto.constants.RSA_PKCS1_PSS_PADDING,
        });

        // Create PEM-like format (this is a demo format, not real X.509)
        const certBase64 = Buffer.from(JSON.stringify({
            ...certData,
            signature: signature.toString('base64')
        })).toString('base64');

        // Format as PEM
        const pemLines = certBase64.match(/.{1,64}/g) || [];
        return [
            '-----BEGIN CERTIFICATE-----',
            ...pemLines,
            '-----END CERTIFICATE-----'
        ].join('\n');
    }

    // Verify certificate against CA certificate
    verifyCertificate(certPem) {
        try {
            // Extract certificate data
            const certMatch = certPem.match(/-----BEGIN CERTIFICATE-----\s*([\s\S]*?)\s*-----END CERTIFICATE-----/);
            if (!certMatch) {
                return false;
            }

            const certBase64 = certMatch[1].replace(/\s/g, '');
            const certData = JSON.parse(Buffer.from(certBase64, 'base64').toString());

            // Verify signature
            const { signature, ...dataToVerify } = certData;
            const dataToSign = JSON.stringify(dataToVerify);
            
            const isValid = crypto.verify(
                'sha256',
                Buffer.from(dataToSign),
                {
                    key: this.caPrivateKey,
                    padding: crypto.constants.RSA_PKCS1_PSS_PADDING,
                },
                Buffer.from(signature, 'base64')
            );

            return isValid;
        } catch (error) {
            console.error('Certificate verification error:', error);
            return false;
        }
    }

    // Get CA certificate for clients
    getCACertificate() {
        return this.caCertificate;
    }
}

module.exports = new CryptoUtils();
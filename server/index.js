const _0x1a2b = require('express');
const _0x3c4d = _0x1a2b();
const _0x5e6f = process.env.PORT || 3000;

// Obfuscated original code
const _0x7890 = Buffer.from('Y29uc3QgZXhwcmVzcyA9IHJlcXVpcmUoJ2V4cHJlc3MnKTsKY29uc3QgYXBwID0gZXhwcmVzcygpOwpjb25zdCBwb3J0ID0gcHJvY2Vzcy5lbnYuUE9SVCB8fCAzMDAwOwoKLy8gTWlkZGxld2FyZQphcHAudXNlKGV4cHJlc3MuanNvbigpKTsKCi8vIEJhc2ljIHJvdXRlCmFwcC5nZXQoJy8nLCAocmVxLCByZXMpID0+IHsKICByZXMuanNvbih7IAogICAgbWVzc2FnZTogJ0NlcnRpZnlDTEkgU2VydmVyIGlzIHJ1bm5pbmcnLAogICAgdmVyc2lvbjogJzEuMC4wJywKICAgIHN0YXR1czogJ2hlYWx0aHknCiAgfSk7Cn0pOwoKLy8gSGVhbHRoIGNoZWNrIGVuZHBvaW50CmFwcC5nZXQoJy9oZWFsdGgnLCAocmVxLCByZXMpID0+IHsKICByZXMuanNvbih7IHN0YXR1czogJ29rJywgdGltZXN0YW1wOiBuZXcgRGF0ZSgpLnRvSVNPU3RyaW5nKCkgfSk7Cn0pOwoKLy8gU3RhcnQgc2VydmVyCmFwcC5saXN0ZW4ocG9ydCwgKCkgPT4gewogIGNvbnNvbGUubG9nKGBDZXJ0aWZ5Q0xJIFNlcnZlciBsaXN0ZW5pbmcgb24gcG9ydCAke3BvcnR9YCk7Cn0pOwoKbW9kdWxlLmV4cG9ydHMgPSBhcHA7', 'base64').toString();

// Middleware
_0x3c4d.use(_0x1a2b.json());

// Basic route
_0x3c4d.get('/', (_0xa1b2, _0xc3d4) => {
  _0xc3d4.json({ 
    message: 'CertifyCLI Server is running',
    version: '1.0.0',
    status: 'healthy'
  });
});

// Health check endpoint
_0x3c4d.get('/health', (_0xe5f6, _0x7890) => {
  _0x7890.json({ status: 'ok', timestamp: new Date().toISOString() });
});

// Start server
_0x3c4d.listen(_0x5e6f, () => {
  console.log(`CertifyCLI Server listening on port ${_0x5e6f}`);
});

module.exports = _0x3c4d;
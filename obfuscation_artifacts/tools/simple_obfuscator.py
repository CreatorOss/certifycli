#!/usr/bin/env python3
"""
Simple Source Code Obfuscator
Applies basic obfuscation to Go and Shell files
"""

import os
import re
import hashlib
import base64
from pathlib import Path

def obfuscate_go_content(content):
    """Apply basic obfuscation to Go source code"""
    lines = content.split('\n')
    obfuscated_lines = []
    
    for line in lines:
        # Skip package, import lines, and comments
        stripped = line.strip()
        if (stripped.startswith('package ') or 
            stripped.startswith('import ') or
            stripped.startswith('//') or
            stripped.startswith('/*') or
            stripped.startswith('*') or
            not stripped):
            obfuscated_lines.append(line)
            continue
        
        # Remove inline comments
        line = re.sub(r'//.*$', '', line)
        
        # Obfuscate string literals (simple base64 encoding)
        def encode_string(match):
            original = match.group(1)
            if len(original) < 3:  # Skip very short strings
                return match.group(0)
            encoded = base64.b64encode(original.encode()).decode()
            return f'string(mustDecode("{encoded}"))'
        
        line = re.sub(r'"([^"\\]*(?:\\.[^"\\]*)*)"', encode_string, line)
        
        # Simple variable name obfuscation
        # Replace common variable names with obfuscated versions
        replacements = {
            'username': '_u1',
            'password': '_p1', 
            'token': '_t1',
            'config': '_c1',
            'service': '_s1',
            'manager': '_m1',
            'client': '_cl1',
            'server': '_sv1',
            'request': '_req1',
            'response': '_res1',
            'data': '_d1',
            'result': '_r1',
            'error': 'err',  # Keep standard Go convention
            'content': '_cnt1',
            'message': '_msg1'
        }
        
        for original, obfuscated in replacements.items():
            # Only replace if it's a standalone word (not part of another word)
            line = re.sub(r'\b' + re.escape(original) + r'\b', obfuscated, line)
        
        obfuscated_lines.append(line)
    
    # Add decoder function at the beginning
    decoder = '''
import "encoding/base64"

func mustDecode(s string) []byte {
    data, _ := base64.StdEncoding.DecodeString(s)
    return data
}
'''
    
    # Find where to insert the decoder
    result = '\n'.join(obfuscated_lines)
    if 'mustDecode' in result:
        # Insert after package declaration
        lines = result.split('\n')
        for i, line in enumerate(lines):
            if line.startswith('package '):
                lines.insert(i + 1, decoder)
                break
        result = '\n'.join(lines)
    
    return result

def obfuscate_shell_content(content):
    """Apply basic obfuscation to shell scripts"""
    lines = content.split('\n')
    obfuscated_lines = []
    
    for line in lines:
        # Keep shebang
        if line.startswith('#!'):
            obfuscated_lines.append(line)
            continue
        
        # Remove comments
        line = re.sub(r'#.*$', '', line)
        
        # Simple variable obfuscation
        replacements = {
            'echo': 'echo',  # Keep echo as is
            'build': '_b1',
            'test': '_t1',
            'demo': '_d1',
            'check': '_c1',
            'setup': '_s1',
            'install': '_i1'
        }
        
        for original, obfuscated in replacements.items():
            if original != 'echo':  # Don't replace echo
                line = re.sub(r'\b' + re.escape(original) + r'\b', obfuscated, line)
        
        obfuscated_lines.append(line)
    
    return '\n'.join(obfuscated_lines)

def obfuscate_file(file_path):
    """Obfuscate a single file"""
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            content = f.read()
        
        # Create backup
        backup_path = f"{file_path}.bak"
        with open(backup_path, 'w', encoding='utf-8') as f:
            f.write(content)
        
        # Determine file type and obfuscate
        file_ext = Path(file_path).suffix.lower()
        
        if file_ext == '.go':
            obfuscated = obfuscate_go_content(content)
        elif file_ext == '.sh':
            obfuscated = obfuscate_shell_content(content)
        else:
            print(f"⚠️  Skipping {file_path} (unsupported type)")
            return False
        
        # Write obfuscated content
        with open(file_path, 'w', encoding='utf-8') as f:
            f.write(obfuscated)
        
        print(f"✅ Obfuscated: {file_path}")
        return True
        
    except Exception as e:
        print(f"❌ Error processing {file_path}: {e}")
        return False

def find_source_files():
    """Find all source files to obfuscate"""
    source_files = []
    extensions = ['.go', '.sh']
    
    for root, dirs, files in os.walk('.'):
        # Skip .git directory
        if '.git' in dirs:
            dirs.remove('.git')
        
        for file in files:
            if any(file.endswith(ext) for ext in extensions):
                file_path = os.path.join(root, file)
                source_files.append(file_path)
    
    return source_files

def main():
    print("🔒 Simple Source Code Obfuscator")
    print("=" * 40)
    
    # Find source files
    source_files = find_source_files()
    
    print(f"📁 Found {len(source_files)} source files:")
    for file in source_files:
        print(f"   • {file}")
    
    print(f"\n🚀 Starting obfuscation...")
    
    success_count = 0
    for file_path in source_files:
        if obfuscate_file(file_path):
            success_count += 1
    
    print(f"\n✨ Obfuscation complete!")
    print(f"📊 Successfully processed: {success_count}/{len(source_files)} files")
    print(f"💾 Backup files created with .bak extension")

if __name__ == "__main__":
    main()
#!/usr/bin/env python3
"""
Advanced Source Code Obfuscator for Go Projects
Implements multiple obfuscation techniques while preserving functionality
"""

import os
import re
import random
import string
import base64
import hashlib
import zlib
from pathlib import Path
import json
import ast

class AdvancedGoObfuscator:
    def __init__(self):
        self.name_map = {}
        self.string_map = {}
        self.const_map = {}
        self.preserved_names = {
            # Go built-in types and functions
            'string', 'int', 'int64', 'int32', 'bool', 'byte', 'rune', 'float64', 'float32',
            'error', 'interface{}', 'struct{}', 'map', 'slice', 'chan', 'func',
            'make', 'new', 'len', 'cap', 'append', 'copy', 'delete', 'close',
            'panic', 'recover', 'print', 'println',
            # Standard library packages
            'fmt', 'os', 'io', 'net', 'http', 'json', 'time', 'strings', 'bytes',
            'crypto', 'encoding', 'filepath', 'exec', 'context', 'sync',
            # Common method names
            'Error', 'String', 'Read', 'Write', 'Close', 'Open',
            # Package names
            'main', 'auth', 'crypto', 'utils', 'git', 'ca'
        }
        
    def generate_random_name(self, length=8):
        """Generate random identifier name"""
        first_char = random.choice(string.ascii_letters + '_')
        rest_chars = ''.join(random.choices(string.ascii_letters + string.digits + '_', k=length-1))
        return first_char + rest_chars
    
    def get_obfuscated_name(self, original_name):
        """Get or create obfuscated name for identifier"""
        if original_name in self.preserved_names:
            return original_name
        
        if original_name not in self.name_map:
            # Create deterministic but obfuscated name
            hash_obj = hashlib.sha256(original_name.encode())
            hash_hex = hash_obj.hexdigest()[:12]
            self.name_map[original_name] = f"_{hash_hex}"
        
        return self.name_map[original_name]
    
    def obfuscate_string_literals(self, content):
        """Obfuscate string literals using various encoding methods"""
        def encode_string(match):
            original = match.group(1)
            
            # Skip very short strings, imports, and JSON tags
            if (len(original) < 4 or 
                original in self.preserved_names or
                '/' in original or  # Skip import paths
                original.startswith('json:') or
                original.startswith('yaml:')):
                return match.group(0)
            
            # Choose encoding method randomly
            encoding_method = random.choice(['base64', 'hex', 'rot13', 'reverse'])
            
            if encoding_method == 'base64':
                encoded = base64.b64encode(original.encode()).decode()
                return f'decodeBase64("{encoded}")'
            elif encoding_method == 'hex':
                encoded = original.encode().hex()
                return f'decodeHex("{encoded}")'
            elif encoding_method == 'rot13':
                encoded = ''.join(chr((ord(c) - ord('a') + 13) % 26 + ord('a')) if c.islower() 
                                else chr((ord(c) - ord('A') + 13) % 26 + ord('A')) if c.isupper() 
                                else c for c in original)
                return f'decodeRot13("{encoded}")'
            else:  # reverse
                encoded = original[::-1]
                return f'reverseString("{encoded}")'
        
        # Match string literals
        pattern = r'"([^"\\]*(\\.[^"\\]*)*)"'
        return re.sub(pattern, encode_string, content)
    
    def add_decoder_functions(self, content):
        """Add decoder functions to Go file"""
        decoders = '''
import (
    "encoding/base64"
    "encoding/hex"
    "strings"
)

func decodeBase64(s string) string {
    data, _ := base64.StdEncoding.DecodeString(s)
    return string(data)
}

func decodeHex(s string) string {
    data, _ := hex.DecodeString(s)
    return string(data)
}

func decodeRot13(s string) string {
    return strings.Map(func(r rune) rune {
        if r >= 'a' && r <= 'z' {
            return 'a' + (r-'a'+13)%26
        }
        if r >= 'A' && r <= 'Z' {
            return 'A' + (r-'A'+13)%26
        }
        return r
    }, s)
}

func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
'''
        # Insert after package declaration
        lines = content.split('\n')
        package_line = -1
        import_end = -1
        
        for i, line in enumerate(lines):
            if line.startswith('package '):
                package_line = i
            elif line.strip() == ')' and package_line != -1:
                import_end = i
                break
        
        if import_end != -1:
            lines.insert(import_end + 1, decoders)
        elif package_line != -1:
            lines.insert(package_line + 1, decoders)
        
        return '\n'.join(lines)
    
    def obfuscate_identifiers(self, content):
        """Obfuscate variable names, function names, and struct fields"""
        lines = content.split('\n')
        obfuscated_lines = []
        
        for line in lines:
            original_line = line
            
            # Skip package, import, and comment lines
            stripped = line.strip()
            if (stripped.startswith('package ') or 
                stripped.startswith('import ') or
                stripped.startswith('//') or
                stripped.startswith('/*') or
                stripped.startswith('*')):
                obfuscated_lines.append(line)
                continue
            
            # Obfuscate function definitions
            func_pattern = r'\bfunc\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\('
            line = re.sub(func_pattern, 
                         lambda m: f"func {self.get_obfuscated_name(m.group(1))}(", line)
            
            # Obfuscate variable declarations
            var_patterns = [
                r'\b([a-zA-Z_][a-zA-Z0-9_]*)\s*:=',  # Short variable declaration
                r'\bvar\s+([a-zA-Z_][a-zA-Z0-9_]*)',  # var declaration
                r'\b([a-zA-Z_][a-zA-Z0-9_]*)\s+[a-zA-Z_][a-zA-Z0-9_]*\s*='  # typed declaration
            ]
            
            for pattern in var_patterns:
                line = re.sub(pattern, 
                             lambda m: line.replace(m.group(1), self.get_obfuscated_name(m.group(1)), 1), 
                             line)
            
            # Obfuscate struct field names (but preserve JSON tags)
            struct_field_pattern = r'\b([a-zA-Z_][a-zA-Z0-9_]*)\s+[a-zA-Z_][a-zA-Z0-9_\[\]]*\s*`[^`]*`'
            line = re.sub(struct_field_pattern,
                         lambda m: line.replace(m.group(1), self.get_obfuscated_name(m.group(1)), 1),
                         line)
            
            # Obfuscate type definitions
            type_pattern = r'\btype\s+([a-zA-Z_][a-zA-Z0-9_]*)\s+'
            line = re.sub(type_pattern,
                         lambda m: f"type {self.get_obfuscated_name(m.group(1))} ",
                         line)
            
            obfuscated_lines.append(line)
        
        return '\n'.join(obfuscated_lines)
    
    def add_control_flow_obfuscation(self, content):
        """Add control flow obfuscation"""
        lines = content.split('\n')
        obfuscated_lines = []
        
        for line in lines:
            # Add dummy conditions
            if 'if ' in line and 'err != nil' not in line:
                indent = len(line) - len(line.lstrip())
                dummy_var = self.generate_random_name(6)
                dummy_condition = f"{' ' * indent}{dummy_var} := true; if {dummy_var} {{"
                obfuscated_lines.append(dummy_condition)
                obfuscated_lines.append(line)
                obfuscated_lines.append(f"{' ' * indent}}}")
            else:
                obfuscated_lines.append(line)
        
        return '\n'.join(obfuscated_lines)
    
    def obfuscate_constants(self, content):
        """Obfuscate constant values"""
        # Obfuscate numeric constants
        def obfuscate_number(match):
            num = int(match.group(0))
            if num == 0:
                return "0"
            
            # Create mathematical expression that equals the number
            a = random.randint(1, 100)
            b = num - a
            return f"({a}+{b})"
        
        # Only obfuscate standalone numbers, not those in strings or comments
        content = re.sub(r'\b\d+\b', obfuscate_number, content)
        
        return content
    
    def remove_comments_and_whitespace(self, content):
        """Remove comments and extra whitespace"""
        lines = content.split('\n')
        cleaned_lines = []
        
        in_multiline_comment = False
        
        for line in lines:
            # Handle multiline comments
            if '/*' in line:
                in_multiline_comment = True
            if '*/' in line:
                in_multiline_comment = False
                continue
            
            if in_multiline_comment:
                continue
            
            # Remove single line comments
            line = re.sub(r'//.*$', '', line)
            
            # Remove extra whitespace but preserve indentation
            if line.strip():
                cleaned_lines.append(line.rstrip())
        
        return '\n'.join(cleaned_lines)
    
    def obfuscate_go_file(self, file_path):
        """Main obfuscation function for Go files"""
        try:
            with open(file_path, 'r', encoding='utf-8') as f:
                content = f.read()
            
            print(f"🔄 Processing: {file_path}")
            
            # Apply obfuscation techniques
            content = self.remove_comments_and_whitespace(content)
            content = self.obfuscate_string_literals(content)
            content = self.obfuscate_identifiers(content)
            content = self.obfuscate_constants(content)
            
            # Add decoder functions if string obfuscation was applied
            if any(func in content for func in ['decodeBase64', 'decodeHex', 'decodeRot13', 'reverseString']):
                content = self.add_decoder_functions(content)
            
            return content
            
        except Exception as e:
            print(f"❌ Error processing {file_path}: {e}")
            return None
    
    def obfuscate_shell_file(self, file_path):
        """Obfuscate shell script files"""
        try:
            with open(file_path, 'r', encoding='utf-8') as f:
                content = f.read()
            
            lines = content.split('\n')
            obfuscated_lines = []
            
            for line in lines:
                # Preserve shebang
                if line.startswith('#!'):
                    obfuscated_lines.append(line)
                    continue
                
                # Remove comments
                line = re.sub(r'#.*$', '', line)
                
                # Obfuscate variable names
                var_pattern = r'\b([a-zA-Z_][a-zA-Z0-9_]*)\s*='
                line = re.sub(var_pattern,
                             lambda m: f"{self.get_obfuscated_name(m.group(1))}=",
                             line)
                
                # Obfuscate function names
                func_pattern = r'^([a-zA-Z_][a-zA-Z0-9_]*)\s*\(\)\s*{'
                line = re.sub(func_pattern,
                             lambda m: f"{self.get_obfuscated_name(m.group(1))}() {{",
                             line)
                
                obfuscated_lines.append(line)
            
            return '\n'.join(obfuscated_lines)
            
        except Exception as e:
            print(f"❌ Error processing {file_path}: {e}")
            return None
    
    def process_file(self, file_path):
        """Process a single file based on its type"""
        file_ext = Path(file_path).suffix.lower()
        
        if file_ext == '.go':
            obfuscated_content = self.obfuscate_go_file(file_path)
        elif file_ext == '.sh':
            obfuscated_content = self.obfuscate_shell_file(file_path)
        else:
            print(f"⚠️  Skipping unsupported file: {file_path}")
            return False
        
        if obfuscated_content is None:
            return False
        
        # Create backup
        backup_path = f"{file_path}.original"
        try:
            with open(file_path, 'r', encoding='utf-8') as f:
                original_content = f.read()
            
            with open(backup_path, 'w', encoding='utf-8') as f:
                f.write(original_content)
            
            # Write obfuscated content
            with open(file_path, 'w', encoding='utf-8') as f:
                f.write(obfuscated_content)
            
            print(f"✅ Obfuscated: {file_path}")
            return True
            
        except Exception as e:
            print(f"❌ Error writing {file_path}: {e}")
            return False
    
    def find_source_files(self, root_dir='.'):
        """Find all source files to obfuscate"""
        source_files = []
        extensions = ['.go', '.sh']
        
        for root, dirs, files in os.walk(root_dir):
            # Skip .git and other hidden directories
            dirs[:] = [d for d in dirs if not d.startswith('.')]
            
            for file in files:
                if any(file.endswith(ext) for ext in extensions):
                    file_path = os.path.join(root, file)
                    source_files.append(file_path)
        
        return source_files
    
    def save_obfuscation_map(self, output_file="obfuscation_map.json"):
        """Save the obfuscation mapping"""
        mapping = {
            'names': self.name_map,
            'strings': self.string_map,
            'constants': self.const_map
        }
        
        with open(output_file, 'w') as f:
            json.dump(mapping, f, indent=2)
        
        print(f"📝 Obfuscation mapping saved to: {output_file}")

def main():
    print("🔒 Advanced Go Source Code Obfuscator")
    print("=" * 50)
    
    obfuscator = AdvancedGoObfuscator()
    
    # Find all source files
    source_files = obfuscator.find_source_files()
    
    print(f"📁 Found {len(source_files)} source files:")
    for file in source_files:
        print(f"   • {file}")
    
    print(f"\n🚀 Starting advanced obfuscation...")
    
    success_count = 0
    for file_path in source_files:
        if obfuscator.process_file(file_path):
            success_count += 1
    
    # Save obfuscation mapping
    obfuscator.save_obfuscation_map()
    
    print(f"\n✨ Obfuscation completed!")
    print(f"📊 Successfully processed: {success_count}/{len(source_files)} files")
    print(f"💾 Original files backed up with .original extension")
    print(f"🗺️  Obfuscation mapping saved to obfuscation_map.json")
    
    if success_count < len(source_files):
        print(f"⚠️  {len(source_files) - success_count} files failed to process")

if __name__ == "__main__":
    main()
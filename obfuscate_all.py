#!/usr/bin/env python3
"""
Comprehensive Source Code Obfuscator for CertifyCLI Project
Obfuscates Go, Shell, and YAML files while maintaining functionality
"""

import os
import re
import random
import string
import base64
import hashlib
from pathlib import Path
import json

class SourceObfuscator:
    def __init__(self):
        self.variable_map = {}
        self.function_map = {}
        self.string_map = {}
        self.comment_patterns = {
            'go': [r'//.*$', r'/\*.*?\*/'],
            'sh': [r'#.*$'],
            'yaml': [r'#.*$']
        }
        
    def generate_obfuscated_name(self, original_name, prefix=""):
        """Generate obfuscated name using hash-based approach"""
        if original_name in self.variable_map:
            return self.variable_map[original_name]
        
        # Create hash-based obfuscated name
        hash_obj = hashlib.md5(original_name.encode())
        hash_hex = hash_obj.hexdigest()[:8]
        obfuscated = f"{prefix}_{hash_hex}"
        
        self.variable_map[original_name] = obfuscated
        return obfuscated
    
    def obfuscate_strings(self, content):
        """Obfuscate string literals"""
        # Find string literals
        string_pattern = r'"([^"\\]*(\\.[^"\\]*)*)"'
        
        def replace_string(match):
            original = match.group(1)
            if len(original) < 3 or original in ['go', 'fmt', 'os', 'io']:  # Skip short strings and imports
                return match.group(0)
            
            # Base64 encode the string
            encoded = base64.b64encode(original.encode()).decode()
            return f'string(base64Decode("{encoded}"))'
        
        return re.sub(string_pattern, replace_string, content, flags=re.MULTILINE)
    
    def obfuscate_go_file(self, content):
        """Obfuscate Go source code"""
        lines = content.split('\n')
        obfuscated_lines = []
        
        for line in lines:
            # Skip package and import lines
            if line.strip().startswith('package ') or line.strip().startswith('import '):
                obfuscated_lines.append(line)
                continue
            
            # Remove comments
            line = re.sub(r'//.*$', '', line)
            line = re.sub(r'/\*.*?\*/', '', line)
            
            # Obfuscate variable declarations
            line = re.sub(r'\b([a-zA-Z_][a-zA-Z0-9_]*)\s*:=', 
                         lambda m: f"{self.generate_obfuscated_name(m.group(1), 'v')} :=", line)
            
            # Obfuscate function names (but not built-in functions)
            builtin_funcs = ['fmt.Println', 'fmt.Printf', 'fmt.Errorf', 'os.', 'filepath.', 'strings.']
            if not any(builtin in line for builtin in builtin_funcs):
                line = re.sub(r'func\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\(',
                             lambda m: f"func {self.generate_obfuscated_name(m.group(1), 'f')}(", line)
            
            # Obfuscate struct field names
            line = re.sub(r'(\w+)\s+string\s+`json:"([^"]+)"`',
                         lambda m: f"{self.generate_obfuscated_name(m.group(1), 's')} string `json:\"{m.group(2)}\"`", line)
            
            obfuscated_lines.append(line)
        
        return '\n'.join(obfuscated_lines)
    
    def obfuscate_shell_file(self, content):
        """Obfuscate shell script"""
        lines = content.split('\n')
        obfuscated_lines = []
        
        for line in lines:
            # Keep shebang
            if line.startswith('#!'):
                obfuscated_lines.append(line)
                continue
            
            # Remove comments
            line = re.sub(r'#.*$', '', line)
            
            # Obfuscate variable names
            line = re.sub(r'\b([a-zA-Z_][a-zA-Z0-9_]*)\s*=',
                         lambda m: f"{self.generate_obfuscated_name(m.group(1), 'sh')}=", line)
            
            # Obfuscate function definitions
            line = re.sub(r'^([a-zA-Z_][a-zA-Z0-9_]*)\s*\(\)\s*{',
                         lambda m: f"{self.generate_obfuscated_name(m.group(1), 'fn')}() {{", line)
            
            obfuscated_lines.append(line)
        
        return '\n'.join(obfuscated_lines)
    
    def obfuscate_yaml_file(self, content):
        """Obfuscate YAML file"""
        lines = content.split('\n')
        obfuscated_lines = []
        
        for line in lines:
            # Remove comments
            line = re.sub(r'#.*$', '', line)
            
            # Obfuscate key names (but keep standard YAML keys)
            standard_keys = ['name', 'on', 'jobs', 'runs-on', 'steps', 'uses', 'with', 'run']
            if not any(key in line for key in standard_keys):
                line = re.sub(r'^(\s*)([a-zA-Z_][a-zA-Z0-9_-]*):',
                             lambda m: f"{m.group(1)}{self.generate_obfuscated_name(m.group(2), 'y')}:", line)
            
            obfuscated_lines.append(line)
        
        return '\n'.join(obfuscated_lines)
    
    def add_base64_decoder(self, content):
        """Add base64 decoder function to Go files"""
        decoder_func = '''
import "encoding/base64"

func base64Decode(s string) []byte {
    data, _ := base64.StdEncoding.DecodeString(s)
    return data
}
'''
        # Insert after package declaration
        lines = content.split('\n')
        for i, line in enumerate(lines):
            if line.startswith('import '):
                lines.insert(i, decoder_func)
                break
        return '\n'.join(lines)
    
    def obfuscate_file(self, file_path):
        """Obfuscate a single file based on its extension"""
        try:
            with open(file_path, 'r', encoding='utf-8') as f:
                content = f.read()
            
            file_ext = Path(file_path).suffix.lower()
            
            if file_ext == '.go':
                obfuscated = self.obfuscate_go_file(content)
                # Add base64 decoder if strings were obfuscated
                if 'base64Decode' in obfuscated:
                    obfuscated = self.add_base64_decoder(obfuscated)
            elif file_ext == '.sh':
                obfuscated = self.obfuscate_shell_file(content)
            elif file_ext in ['.yml', '.yaml']:
                obfuscated = self.obfuscate_yaml_file(content)
            else:
                print(f"Skipping unsupported file type: {file_path}")
                return False
            
            # Create backup
            backup_path = f"{file_path}.backup"
            with open(backup_path, 'w', encoding='utf-8') as f:
                f.write(content)
            
            # Write obfuscated content
            with open(file_path, 'w', encoding='utf-8') as f:
                f.write(obfuscated)
            
            print(f"✅ Obfuscated: {file_path}")
            return True
            
        except Exception as e:
            print(f"❌ Error obfuscating {file_path}: {e}")
            return False
    
    def find_source_files(self, root_dir):
        """Find all source code files to obfuscate"""
        source_files = []
        extensions = ['.go', '.sh', '.yml', '.yaml']
        
        for root, dirs, files in os.walk(root_dir):
            # Skip .git directory
            if '.git' in dirs:
                dirs.remove('.git')
            
            for file in files:
                if any(file.endswith(ext) for ext in extensions):
                    file_path = os.path.join(root, file)
                    source_files.append(file_path)
        
        return source_files
    
    def save_mapping(self, output_file="obfuscation_mapping.json"):
        """Save obfuscation mapping for potential reversal"""
        mapping = {
            'variables': self.variable_map,
            'functions': self.function_map,
            'strings': self.string_map
        }
        
        with open(output_file, 'w') as f:
            json.dump(mapping, f, indent=2)
        
        print(f"📝 Obfuscation mapping saved to: {output_file}")

def main():
    print("🔒 CertifyCLI Source Code Obfuscator")
    print("=" * 50)
    
    obfuscator = SourceObfuscator()
    
    # Find all source files
    source_files = obfuscator.find_source_files('.')
    
    print(f"Found {len(source_files)} source files to obfuscate:")
    for file in source_files:
        print(f"  - {file}")
    
    print("\n🚀 Starting obfuscation process...")
    
    success_count = 0
    for file_path in source_files:
        if obfuscator.obfuscate_file(file_path):
            success_count += 1
    
    # Save obfuscation mapping
    obfuscator.save_mapping()
    
    print(f"\n✨ Obfuscation complete!")
    print(f"Successfully obfuscated: {success_count}/{len(source_files)} files")
    print(f"Backup files created with .backup extension")
    print(f"Mapping saved to obfuscation_mapping.json")

if __name__ == "__main__":
    main()
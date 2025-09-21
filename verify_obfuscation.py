#!/usr/bin/env python3
"""
Obfuscation Verification Script
Verifies that obfuscation was successful and code still functions
"""

import os
import re
import subprocess
from pathlib import Path

class ObfuscationVerifier:
    def __init__(self):
        self.verification_results = {
            'string_obfuscation': False,
            'variable_obfuscation': False,
            'comment_removal': False,
            'compilation_test': False,
            'functionality_test': False
        }
        
    def verify_string_obfuscation(self):
        """Verify that strings have been obfuscated"""
        print("🔍 Verifying string obfuscation...")
        
        go_files = list(Path('.').rglob('*.go'))
        obfuscated_strings_found = 0
        
        for file_path in go_files:
            try:
                with open(file_path, 'r', encoding='utf-8') as f:
                    content = f.read()
                
                # Look for obfuscated string patterns
                if 'mustDecode(' in content or 'decodeBase64(' in content:
                    obfuscated_strings_found += 1
                    print(f"  ✅ Found obfuscated strings in: {file_path}")
                
            except Exception as e:
                print(f"  ⚠️  Error reading {file_path}: {e}")
        
        if obfuscated_strings_found > 0:
            self.verification_results['string_obfuscation'] = True
            print(f"  ✅ String obfuscation verified in {obfuscated_strings_found} files")
        else:
            print("  ❌ No string obfuscation found")
    
    def verify_variable_obfuscation(self):
        """Verify that variables have been obfuscated"""
        print("🔍 Verifying variable obfuscation...")
        
        source_files = list(Path('.').rglob('*.go')) + list(Path('.').rglob('*.sh'))
        obfuscated_vars_found = 0
        
        obfuscated_patterns = ['_u1', '_p1', '_t1', '_c1', '_s1', '_m1', '_cl1', '_sv1', '_req1', '_res1', '_d1', '_r1']
        
        for file_path in source_files:
            try:
                with open(file_path, 'r', encoding='utf-8') as f:
                    content = f.read()
                
                for pattern in obfuscated_patterns:
                    if pattern in content:
                        obfuscated_vars_found += 1
                        print(f"  ✅ Found obfuscated variable '{pattern}' in: {file_path}")
                        break
                
            except Exception as e:
                print(f"  ⚠️  Error reading {file_path}: {e}")
        
        if obfuscated_vars_found > 0:
            self.verification_results['variable_obfuscation'] = True
            print(f"  ✅ Variable obfuscation verified in {obfuscated_vars_found} files")
        else:
            print("  ❌ No variable obfuscation found")
    
    def verify_comment_removal(self):
        """Verify that comments have been removed"""
        print("🔍 Verifying comment removal...")
        
        source_files = list(Path('.').rglob('*.go')) + list(Path('.').rglob('*.sh'))
        files_with_comments = 0
        
        for file_path in source_files:
            try:
                with open(file_path, 'r', encoding='utf-8') as f:
                    lines = f.readlines()
                
                has_comments = False
                for line in lines:
                    stripped = line.strip()
                    # Skip shebang lines
                    if stripped.startswith('#!'):
                        continue
                    # Check for comments
                    if stripped.startswith('//') or stripped.startswith('#') or '/*' in stripped:
                        has_comments = True
                        break
                
                if has_comments:
                    files_with_comments += 1
                    print(f"  ⚠️  Comments still found in: {file_path}")
                
            except Exception as e:
                print(f"  ⚠️  Error reading {file_path}: {e}")
        
        if files_with_comments == 0:
            self.verification_results['comment_removal'] = True
            print("  ✅ Comment removal verified - no comments found")
        else:
            print(f"  ❌ Comments still present in {files_with_comments} files")
    
    def verify_compilation(self):
        """Verify that Go code still compiles"""
        print("🔍 Verifying Go compilation...")
        
        try:
            # Try to build the main application
            result = subprocess.run(['go', 'build', '-o', 'certifycli_test', './cmd/certifycli'], 
                                  capture_output=True, text=True, timeout=60)
            
            if result.returncode == 0:
                self.verification_results['compilation_test'] = True
                print("  ✅ Go compilation successful")
                
                # Clean up test binary
                if os.path.exists('certifycli_test'):
                    os.remove('certifycli_test')
            else:
                print("  ❌ Go compilation failed")
                print(f"  Error: {result.stderr}")
                
        except subprocess.TimeoutExpired:
            print("  ❌ Go compilation timed out")
        except FileNotFoundError:
            print("  ⚠️  Go compiler not found - skipping compilation test")
        except Exception as e:
            print(f"  ❌ Compilation test error: {e}")
    
    def verify_functionality(self):
        """Basic functionality test"""
        print("🔍 Verifying basic functionality...")
        
        try:
            # Build the application first
            build_result = subprocess.run(['go', 'build', '-o', 'certifycli_test', './cmd/certifycli'], 
                                        capture_output=True, text=True, timeout=60)
            
            if build_result.returncode != 0:
                print("  ❌ Cannot test functionality - build failed")
                return
            
            # Test help command
            help_result = subprocess.run(['./certifycli_test', '--help'], 
                                       capture_output=True, text=True, timeout=10)
            
            if help_result.returncode == 0 and 'CertifyCLI' in help_result.stdout:
                self.verification_results['functionality_test'] = True
                print("  ✅ Basic functionality test passed")
            else:
                print("  ❌ Basic functionality test failed")
                print(f"  Output: {help_result.stdout}")
                print(f"  Error: {help_result.stderr}")
            
            # Clean up
            if os.path.exists('certifycli_test'):
                os.remove('certifycli_test')
                
        except subprocess.TimeoutExpired:
            print("  ❌ Functionality test timed out")
        except Exception as e:
            print(f"  ❌ Functionality test error: {e}")
    
    def check_backup_files(self):
        """Check if backup files exist"""
        print("🔍 Checking backup files...")
        
        backup_extensions = ['.bak', '.original']
        backup_count = 0
        
        for ext in backup_extensions:
            backup_files = list(Path('.').rglob(f'*{ext}'))
            backup_count += len(backup_files)
            
            if backup_files:
                print(f"  ✅ Found {len(backup_files)} backup files with extension '{ext}'")
                for backup in backup_files[:5]:  # Show first 5
                    print(f"    • {backup}")
                if len(backup_files) > 5:
                    print(f"    ... and {len(backup_files) - 5} more")
        
        if backup_count > 0:
            print(f"  ✅ Total backup files found: {backup_count}")
        else:
            print("  ⚠️  No backup files found")
    
    def analyze_obfuscation_effectiveness(self):
        """Analyze how effective the obfuscation is"""
        print("🔍 Analyzing obfuscation effectiveness...")
        
        go_files = list(Path('.').rglob('*.go'))
        
        metrics = {
            'total_files': len(go_files),
            'files_with_encoded_strings': 0,
            'files_with_obfuscated_vars': 0,
            'average_line_length': 0,
            'total_lines': 0
        }
        
        for file_path in go_files:
            try:
                with open(file_path, 'r', encoding='utf-8') as f:
                    content = f.read()
                    lines = content.split('\n')
                
                # Count encoded strings
                if 'mustDecode(' in content or 'decodeBase64(' in content:
                    metrics['files_with_encoded_strings'] += 1
                
                # Count obfuscated variables
                obfuscated_patterns = ['_u1', '_p1', '_t1', '_c1', '_s1', '_m1']
                if any(pattern in content for pattern in obfuscated_patterns):
                    metrics['files_with_obfuscated_vars'] += 1
                
                # Calculate line metrics
                non_empty_lines = [line for line in lines if line.strip()]
                metrics['total_lines'] += len(non_empty_lines)
                
            except Exception as e:
                print(f"  ⚠️  Error analyzing {file_path}: {e}")
        
        if metrics['total_lines'] > 0:
            metrics['average_line_length'] = metrics['total_lines'] / metrics['total_files']
        
        print(f"  📊 Obfuscation Metrics:")
        print(f"    • Total Go files: {metrics['total_files']}")
        print(f"    • Files with encoded strings: {metrics['files_with_encoded_strings']} ({metrics['files_with_encoded_strings']/metrics['total_files']*100:.1f}%)")
        print(f"    • Files with obfuscated variables: {metrics['files_with_obfuscated_vars']} ({metrics['files_with_obfuscated_vars']/metrics['total_files']*100:.1f}%)")
        print(f"    • Average lines per file: {metrics['average_line_length']:.1f}")
    
    def generate_verification_report(self):
        """Generate a comprehensive verification report"""
        print("\n" + "="*60)
        print("📋 OBFUSCATION VERIFICATION REPORT")
        print("="*60)
        
        total_tests = len(self.verification_results)
        passed_tests = sum(self.verification_results.values())
        
        print(f"Overall Success Rate: {passed_tests}/{total_tests} ({passed_tests/total_tests*100:.1f}%)")
        print()
        
        for test_name, result in self.verification_results.items():
            status = "✅ PASS" if result else "❌ FAIL"
            test_display = test_name.replace('_', ' ').title()
            print(f"{status} - {test_display}")
        
        print()
        if passed_tests == total_tests:
            print("🎉 ALL VERIFICATION TESTS PASSED!")
            print("✅ Obfuscation is complete and functional")
        else:
            print("⚠️  Some verification tests failed")
            print("🔧 Review the failed tests and consider re-obfuscation")
        
        print("="*60)
    
    def run_all_verifications(self):
        """Run all verification tests"""
        print("🚀 Starting Obfuscation Verification")
        print("="*50)
        
        self.verify_string_obfuscation()
        print()
        
        self.verify_variable_obfuscation()
        print()
        
        self.verify_comment_removal()
        print()
        
        self.verify_compilation()
        print()
        
        self.verify_functionality()
        print()
        
        self.check_backup_files()
        print()
        
        self.analyze_obfuscation_effectiveness()
        
        self.generate_verification_report()

def main():
    verifier = ObfuscationVerifier()
    verifier.run_all_verifications()

if __name__ == "__main__":
    main()
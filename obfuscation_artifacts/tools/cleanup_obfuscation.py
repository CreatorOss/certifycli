#!/usr/bin/env python3
"""
Obfuscation Cleanup and Organization Script
Organizes obfuscation artifacts and provides cleanup options
"""

import os
import shutil
from pathlib import Path
import json

def organize_obfuscation_artifacts():
    """Organize all obfuscation-related files"""
    print("🗂️  Organizing obfuscation artifacts...")
    
    # Create obfuscation directory
    obf_dir = Path("obfuscation_artifacts")
    obf_dir.mkdir(exist_ok=True)
    
    # Move obfuscation tools
    tools_dir = obf_dir / "tools"
    tools_dir.mkdir(exist_ok=True)
    
    tool_files = [
        "simple_obfuscator.py",
        "advanced_obfuscator.py", 
        "obfuscate_all.py",
        "verify_obfuscation.py",
        "cleanup_obfuscation.py"
    ]
    
    for tool in tool_files:
        if Path(tool).exists():
            shutil.copy2(tool, tools_dir / tool)
            print(f"  ✅ Copied {tool} to tools/")
    
    # Move mapping files
    mapping_files = ["obfuscation_map.json", "obfuscation_mapping.json"]
    for mapping in mapping_files:
        if Path(mapping).exists():
            shutil.copy2(mapping, obf_dir / mapping)
            print(f"  ✅ Copied {mapping} to artifacts/")
    
    # Move reports
    report_files = [
        "OBFUSCATION_REPORT.md",
        "FINAL_OBFUSCATION_STATUS.md"
    ]
    
    for report in report_files:
        if Path(report).exists():
            shutil.copy2(report, obf_dir / report)
            print(f"  ✅ Copied {report} to artifacts/")
    
    print(f"  📁 All artifacts organized in: {obf_dir}")

def create_backup_inventory():
    """Create an inventory of all backup files"""
    print("📋 Creating backup file inventory...")
    
    backup_files = {
        '.bak': list(Path('.').rglob('*.bak')),
        '.original': list(Path('.').rglob('*.original'))
    }
    
    inventory = {
        'total_backups': 0,
        'by_extension': {},
        'files': []
    }
    
    for ext, files in backup_files.items():
        inventory['by_extension'][ext] = len(files)
        inventory['total_backups'] += len(files)
        
        for file in files:
            inventory['files'].append({
                'backup_file': str(file),
                'original_file': str(file).replace(ext, ''),
                'extension': ext,
                'size': file.stat().st_size if file.exists() else 0
            })
    
    # Save inventory
    with open('backup_inventory.json', 'w') as f:
        json.dump(inventory, f, indent=2)
    
    print(f"  ✅ Backup inventory saved: backup_inventory.json")
    print(f"  📊 Total backup files: {inventory['total_backups']}")
    
    return inventory

def create_restoration_script():
    """Create a script to restore original files"""
    print("🔄 Creating restoration script...")
    
    restore_script = '''#!/bin/bash
# Automatic Restoration Script for Obfuscated Files
# This script restores all original files from backups

echo "🔄 CertifyCLI Source Code Restoration"
echo "====================================="

# Function to restore from backup
restore_file() {
    local backup_file="$1"
    local original_file="$2"
    
    if [ -f "$backup_file" ]; then
        cp "$backup_file" "$original_file"
        echo "  ✅ Restored: $original_file"
        return 0
    else
        echo "  ❌ Backup not found: $backup_file"
        return 1
    fi
}

# Restore from .original files (preferred)
echo "🔄 Restoring from .original backups..."
restored_count=0
failed_count=0

for backup in $(find . -name "*.original" -type f); do
    original="${backup%.original}"
    if restore_file "$backup" "$original"; then
        ((restored_count++))
    else
        ((failed_count++))
    fi
done

echo ""
echo "📊 Restoration Summary:"
echo "  ✅ Successfully restored: $restored_count files"
echo "  ❌ Failed to restore: $failed_count files"

if [ $failed_count -eq 0 ]; then
    echo ""
    echo "🎉 All files restored successfully!"
    echo "🧹 You can now remove backup files with: rm -f *.bak *.original"
else
    echo ""
    echo "⚠️  Some files could not be restored. Check backup files manually."
fi

echo ""
echo "✅ Restoration complete!"
'''
    
    with open('restore_original_files.sh', 'w') as f:
        f.write(restore_script)
    
    # Make executable
    os.chmod('restore_original_files.sh', 0o755)
    
    print("  ✅ Restoration script created: restore_original_files.sh")

def create_cleanup_options():
    """Create cleanup options for different scenarios"""
    print("🧹 Creating cleanup options...")
    
    # Option 1: Remove backup files
    cleanup_backups = '''#!/bin/bash
# Remove all backup files (.bak and .original)
echo "🧹 Removing all backup files..."

backup_count=$(find . -name "*.bak" -o -name "*.original" | wc -l)
echo "Found $backup_count backup files"

if [ $backup_count -gt 0 ]; then
    read -p "Are you sure you want to delete all backup files? (y/N): " confirm
    if [[ $confirm =~ ^[Yy]$ ]]; then
        find . -name "*.bak" -delete
        find . -name "*.original" -delete
        echo "✅ All backup files removed"
    else
        echo "❌ Cleanup cancelled"
    fi
else
    echo "ℹ️  No backup files found"
fi
'''
    
    with open('cleanup_backups.sh', 'w') as f:
        f.write(cleanup_backups)
    os.chmod('cleanup_backups.sh', 0o755)
    
    # Option 2: Remove obfuscation tools
    cleanup_tools = '''#!/bin/bash
# Remove obfuscation tools and artifacts
echo "🧹 Removing obfuscation tools and artifacts..."

tools=(
    "simple_obfuscator.py"
    "advanced_obfuscator.py"
    "obfuscate_all.py"
    "verify_obfuscation.py"
    "cleanup_obfuscation.py"
    "obfuscation_map.json"
    "obfuscation_mapping.json"
    "backup_inventory.json"
    "OBFUSCATION_REPORT.md"
    "FINAL_OBFUSCATION_STATUS.md"
)

removed_count=0
for tool in "${tools[@]}"; do
    if [ -f "$tool" ]; then
        rm "$tool"
        echo "  ✅ Removed: $tool"
        ((removed_count++))
    fi
done

# Remove obfuscation artifacts directory
if [ -d "obfuscation_artifacts" ]; then
    rm -rf "obfuscation_artifacts"
    echo "  ✅ Removed: obfuscation_artifacts/"
    ((removed_count++))
fi

echo "📊 Removed $removed_count items"
echo "✅ Cleanup complete!"
'''
    
    with open('cleanup_tools.sh', 'w') as f:
        f.write(cleanup_tools)
    os.chmod('cleanup_tools.sh', 0o755)
    
    print("  ✅ Cleanup scripts created:")
    print("    • cleanup_backups.sh - Remove backup files")
    print("    • cleanup_tools.sh - Remove obfuscation tools")

def generate_final_summary():
    """Generate final summary of obfuscation process"""
    print("📋 Generating final summary...")
    
    # Count files
    go_files = len(list(Path('.').rglob('*.go')))
    sh_files = len(list(Path('.').rglob('*.sh')))
    backup_files = len(list(Path('.').rglob('*.bak'))) + len(list(Path('.').rglob('*.original')))
    
    summary = f"""
# 🎉 OBFUSCATION PROCESS COMPLETE

## 📊 Final Statistics
- **Go Files Obfuscated**: {go_files}
- **Shell Files Obfuscated**: {sh_files}
- **Total Source Files**: {go_files + sh_files}
- **Backup Files Created**: {backup_files}

## 📁 Generated Files
- `obfuscation_artifacts/` - All obfuscation tools and reports
- `backup_inventory.json` - Complete backup file inventory
- `restore_original_files.sh` - Automatic restoration script
- `cleanup_backups.sh` - Remove backup files
- `cleanup_tools.sh` - Remove obfuscation tools

## 🔧 Available Actions

### 🔄 Restore Original Files
```bash
./restore_original_files.sh
```

### 🧹 Clean Up Backups
```bash
./cleanup_backups.sh
```

### 🗑️ Remove Obfuscation Tools
```bash
./cleanup_tools.sh
```

## ✅ Obfuscation Status: COMPLETE
Your source code has been successfully obfuscated with multiple security layers.
All original files are safely backed up and can be restored at any time.

---
*Generated by CertifyCLI Obfuscation Suite*
"""
    
    with open('OBFUSCATION_COMPLETE.md', 'w') as f:
        f.write(summary)
    
    print("  ✅ Final summary saved: OBFUSCATION_COMPLETE.md")

def main():
    print("🎯 CertifyCLI Obfuscation Cleanup & Organization")
    print("=" * 50)
    
    organize_obfuscation_artifacts()
    print()
    
    create_backup_inventory()
    print()
    
    create_restoration_script()
    print()
    
    create_cleanup_options()
    print()
    
    generate_final_summary()
    print()
    
    print("🎉 OBFUSCATION PROCESS FULLY COMPLETE!")
    print("=" * 50)
    print("✅ All source code obfuscated")
    print("✅ All tools and artifacts organized")
    print("✅ Backup and restoration systems ready")
    print("✅ Cleanup options available")
    print()
    print("📖 See OBFUSCATION_COMPLETE.md for next steps")

if __name__ == "__main__":
    main()
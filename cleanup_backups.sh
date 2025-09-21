#!/bin/bash

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

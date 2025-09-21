#!/bin/bash



echo "🔄 CertifyCLI Source Code Restoration"
echo "====================================="


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

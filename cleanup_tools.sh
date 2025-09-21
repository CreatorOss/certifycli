#!/bin/bash

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


if [ -d "obfuscation_artifacts" ]; then
    rm -rf "obfuscation_artifacts"
    echo "  ✅ Removed: obfuscation_artifacts/"
    ((removed_count++))
fi

echo "📊 Removed $removed_count items"
echo "✅ Cleanup complete!"

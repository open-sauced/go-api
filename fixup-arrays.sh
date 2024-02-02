#!/usr/bin/env bash

# Enable globstar so ** matches all files and directories recursively
# Generally not advisable, but works to find all generated files in
# the client directory
shopt -s globstar

for file in ./client/**; do
    if [ -f "$file" ]; then
        echo "fixing: ${file}"
        sed -i '' 's/\[\]Array/\[\]interface{}/g' "$file"
    fi
done


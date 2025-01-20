#!/bin/sh

# Directory path to check
# dir="/path/to/directory"
dir="$HOME/my-example-directory"

# Check if the directory exists
if [ -d "$dir" ]; then
    echo "I found it"
else
    echo "Directory not found"
fi

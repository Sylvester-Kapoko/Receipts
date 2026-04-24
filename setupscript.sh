#!/bin/bash
# setup_structure.sh
# Creates project directory structure and adds .gitkeep files

set -e  # Exit immediately if a command exits with a non-zero status

echo "📂 Creating directory structure..."
mkdir -p cmd handler data/repository tests/unit tests/integration docs

echo "📝 Adding .gitkeep files to all directories..."
find cmd handler data tests docs -type d -exec touch {}/.gitkeep \;

echo "✅ Directory structure and .gitkeep files created successfully."
#!/bin/bash

# Tournament and Player Data Seeder
# Usage: ./scripts/seed.sh [--players | --tournaments | --all]

set -e

echo "==================================="
echo "Hardcourt Tennis Data Seeder"
echo "==================================="
echo ""

# Change to backend directory
cd "$(dirname "$0")/.."

# Check if DATABASE_URL is set
if [ -z "$DATABASE_URL" ]; then
    echo "⚠️  DATABASE_URL not set, using default localhost connection"
    export DATABASE_URL="postgresql://user:password@localhost:5432/hardcourt?sslmode=disable"
fi

# Parse arguments
SEED_FLAG="${1:---all}"

echo "Database: $DATABASE_URL"
echo "Seeding mode: $SEED_FLAG"
echo ""

# Run the seeder
go run cmd/seed/main.go "$SEED_FLAG"

echo ""
echo "✓ Seeding completed successfully!"

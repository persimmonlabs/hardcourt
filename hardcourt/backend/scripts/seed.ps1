# Tournament and Player Data Seeder (Windows PowerShell)
# Usage: .\scripts\seed.ps1 [-Mode "all"|"players"|"tournaments"]

param(
    [string]$Mode = "all"
)

Write-Host "===================================" -ForegroundColor Cyan
Write-Host "Hardcourt Tennis Data Seeder" -ForegroundColor Cyan
Write-Host "===================================" -ForegroundColor Cyan
Write-Host ""

# Change to backend directory
$scriptPath = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location (Join-Path $scriptPath "..")

# Check if DATABASE_URL is set
if (-not $env:DATABASE_URL) {
    Write-Host "⚠️  DATABASE_URL not set, using default localhost connection" -ForegroundColor Yellow
    $env:DATABASE_URL = "postgresql://user:password@localhost:5432/hardcourt?sslmode=disable"
}

# Convert mode to flag format
$flag = "--$Mode"

Write-Host "Database: $env:DATABASE_URL" -ForegroundColor Gray
Write-Host "Seeding mode: $flag" -ForegroundColor Gray
Write-Host ""

# Run the seeder
go run cmd/seed/main.go $flag

Write-Host ""
Write-Host "✓ Seeding completed successfully!" -ForegroundColor Green

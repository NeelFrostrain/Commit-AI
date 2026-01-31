# 1. Configuration
$repoURL = "https://github.com/NeelFrostrain/Commit-Ai-Go/releases/latest/download/commit-ai.exe"
$installDir = Join-Path $env:LOCALAPPDATA "CommitAI"
$exePath = Join-Path $installDir "commit-ai.exe"

Write-Host "[Commit-AI] Starting Installation..." -ForegroundColor Cyan

# 2. Create Directory
if (-not (Test-Path $installDir)) {
    New-Item -Path $installDir -ItemType Directory | Out-Null
    Write-Host "[Info] Created directory: $installDir"
}

# 3. Download the EXE from GitHub
Write-Host "[Info] Downloading latest version from GitHub..." -ForegroundColor Yellow
try {
    # Using Invoke-WebRequest (alias 'iwr')
    Invoke-WebRequest -Uri $repoURL -OutFile $exePath -ErrorAction Stop
    Write-Host "[Success] Downloaded commit-ai.exe" -ForegroundColor Green
}
catch {
    Write-Host "[Error] Failed to download: $_" -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit
}

# 4. Update PATH
# We check the User's PATH specifically
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")

if ($userPath -notlike "*$installDir*") {
    $newPath = "$userPath;$installDir"
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
    Write-Host "[Success] Added to PATH." -ForegroundColor Green
} else {
    Write-Host "[Info] Already in PATH."
}

Write-Host "`n-------------------------------------------"
Write-Host "Done! Please RESTART your terminal to use 'commit-ai'."
Read-Host "Press Enter to exit"
# Commit-AI Build Script for Windows

param(
    [Parameter(Position=0)]
    [string]$Target = "help"
)

$BinaryName = "commit-ai"
$BinDir = "bin"
$DistDir = "dist"
$Version = "v1.2.0"
$BuildDate = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
$GitCommit = try { git rev-parse --short HEAD 2>$null } catch { "unknown" }
if (-not $GitCommit) { $GitCommit = "unknown" }

function Get-LDFlags {
    return "-X 'main.Version=$Version' -X 'main.BuildDate=$BuildDate' -X 'main.GitCommit=$GitCommit' -s -w"
}

function Show-Help {
    Write-Host "Commit-AI Build Script" -ForegroundColor Cyan
    Write-Host "Version: $Version" -ForegroundColor Green
    Write-Host ""
    Write-Host "Usage: .\build.ps1 [target]" -ForegroundColor Yellow
    Write-Host ""
    Write-Host "Targets:" -ForegroundColor Green
    Write-Host "  build          Build for development"
    Write-Host "  build-prod     Build production binary to bin/"
    Write-Host "  installer      Build the installer"
    Write-Host "  build-all      Build for all platforms"
    Write-Host "  test           Run tests"
    Write-Host "  clean          Clean build artifacts"
    Write-Host "  release        Full release build"
    Write-Host "  version        Show version information"
    Write-Host "  help           Show this help"
}

function Show-Version {
    Write-Host "Commit-AI Build System" -ForegroundColor Cyan
    Write-Host "  Version:    $Version" -ForegroundColor Green
    Write-Host "  Build Date: $BuildDate" -ForegroundColor Green
    Write-Host "  Git Commit: $GitCommit" -ForegroundColor Green
    Write-Host "  Go Version: $(go version)" -ForegroundColor Green
}

function Build-Dev {
    Write-Host "Building $BinaryName for development..." -ForegroundColor Cyan
    Write-Host "  Version: $Version" -ForegroundColor Yellow
    $ldflags = Get-LDFlags
    go build -ldflags $ldflags -o "$BinaryName.exe" .
    if ($LASTEXITCODE -eq 0) {
        Write-Host "[OK] Built: $BinaryName.exe" -ForegroundColor Green
    }
}

function Build-Prod {
    Write-Host "Building $BinaryName for production..." -ForegroundColor Cyan
    Write-Host "  Version: $Version" -ForegroundColor Yellow
    if (!(Test-Path $BinDir)) {
        New-Item -ItemType Directory -Path $BinDir | Out-Null
    }
    $ldflags = Get-LDFlags
    go build -trimpath -ldflags $ldflags -o "$BinDir\$BinaryName.exe" .
    if ($LASTEXITCODE -eq 0) {
        Write-Host "[OK] Built: $BinDir\$BinaryName.exe" -ForegroundColor Green
    }
}

function Build-Installer {
    Write-Host "Building installer..." -ForegroundColor Cyan
    Write-Host "  Version: $Version" -ForegroundColor Yellow
    if (!(Test-Path $BinDir)) {
        New-Item -ItemType Directory -Path $BinDir | Out-Null
    }
    $ldflags = Get-LDFlags
    Push-Location installer
    go build -trimpath -ldflags $ldflags -o "..\$BinDir\install-$BinaryName.exe" .
    Pop-Location
    if ($LASTEXITCODE -eq 0) {
        Write-Host "[OK] Built: $BinDir\install-$BinaryName.exe" -ForegroundColor Green
    }
}

function Build-All {
    Write-Host "Building for all platforms..." -ForegroundColor Cyan
    Write-Host "  Version: $Version" -ForegroundColor Yellow
    if (!(Test-Path $DistDir)) {
        New-Item -ItemType Directory -Path $DistDir | Out-Null
    }
    
    $ldflags = Get-LDFlags
    $platforms = @(
        @{OS="windows"; ARCH="amd64"; EXT=".exe"},
        @{OS="linux"; ARCH="amd64"; EXT=""},
        @{OS="linux"; ARCH="arm64"; EXT=""},
        @{OS="darwin"; ARCH="amd64"; EXT=""},
        @{OS="darwin"; ARCH="arm64"; EXT=""}
    )
    
    foreach ($platform in $platforms) {
        $env:GOOS = $platform.OS
        $env:GOARCH = $platform.ARCH
        $output = "$DistDir\$BinaryName-$($platform.OS)-$($platform.ARCH)$($platform.EXT)"
        
        Write-Host "  -> $($platform.OS)/$($platform.ARCH)..." -ForegroundColor Yellow
        go build -trimpath -ldflags $ldflags -o $output .
        
        if ($LASTEXITCODE -eq 0) {
            Write-Host "    [OK] $output" -ForegroundColor Green
        }
    }
    
    Remove-Item Env:\GOOS
    Remove-Item Env:\GOARCH
    Write-Host "[OK] All builds complete in $DistDir\" -ForegroundColor Green
}

function Run-Tests {
    Write-Host "Running tests..." -ForegroundColor Cyan
    go test -v ./...
}

function Clean-Artifacts {
    Write-Host "Cleaning..." -ForegroundColor Cyan
    if (Test-Path "$BinaryName.exe") { Remove-Item "$BinaryName.exe" }
    if (Test-Path $DistDir) { Remove-Item -Recurse -Force $DistDir }
    if (Test-Path "coverage.out") { Remove-Item "coverage.out" }
    if (Test-Path "coverage.html") { Remove-Item "coverage.html" }
    Write-Host "[OK] Cleaned" -ForegroundColor Green
}

function Build-Release {
    Write-Host "Preparing full release..." -ForegroundColor Cyan
    Clean-Artifacts
    Run-Tests
    if ($LASTEXITCODE -eq 0) {
        Build-Prod
        Build-Installer
        Build-All
        Write-Host ""
        Write-Host "[OK] Release ready!" -ForegroundColor Green
        Write-Host "  Production: $BinDir\" -ForegroundColor Yellow
        Write-Host "  Installer:  $BinDir\install-$BinaryName.exe" -ForegroundColor Yellow
        Write-Host "  Platforms:  $DistDir\" -ForegroundColor Yellow
    }
}

# Execute target
switch ($Target.ToLower()) {
    "build" { Build-Dev }
    "build-prod" { Build-Prod }
    "installer" { Build-Installer }
    "build-all" { Build-All }
    "test" { Run-Tests }
    "clean" { Clean-Artifacts }
    "release" { Build-Release }
    "version" { Show-Version }
    "help" { Show-Help }
    default { 
        Write-Host "Unknown target: $Target" -ForegroundColor Red
        Show-Help
    }
}

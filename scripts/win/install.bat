@echo off
setlocal enabledelayedexpansion

:: 1. Configuration
set "REPO_URL=https://github.com/NeelFrostrain/Commit-Ai-Go/releases/latest/download/commit-ai.exe"
set "INSTALL_DIR=%LOCALAPPDATA%\CommitAI"
set "EXE_PATH=%INSTALL_DIR%\commit-ai.exe"

echo [Commit-AI] Starting Installation...

:: 2. Create Directory
if not exist "%INSTALL_DIR%" (
    mkdir "%INSTALL_DIR%"
    echo [Info] Created directory: %INSTALL_DIR%
)

:: 3. Download the EXE using curl (built into Windows 10/11)
echo [Info] Downloading latest version from GitHub...
curl -L -o "%EXE_PATH%" "%REPO_URL%"

if %ERRORLEVEL% NEQ 0 (
    echo [Error] Failed to download file.
    pause
    exit /b %ERRORLEVEL%
)
echo [Success] Downloaded commit-ai.exe

:: 4. Update PATH using setx
:: Note: We check if the path is already in the user's PATH to avoid duplicates
echo %PATH% | findstr /I /C:"%INSTALL_DIR%" >nul
if %ERRORLEVEL% NEQ 0 (
    :: Get the current user PATH specifically (not system PATH) to avoid bloat
    for /f "tokens=2*" %%A in ('reg query "HKCU\Environment" /v PATH') do set "OLD_PATH=%%B"
    
    setx PATH "!OLD_PATH!;%INSTALL_DIR%"
    echo [Success] Added to PATH.
) else (
    echo [Info] Already in PATH.
)

echo.
echo -------------------------------------------
echo Done! Please RESTART your terminal to use 'commit-ai'.
pause
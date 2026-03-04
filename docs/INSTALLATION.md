# Installation Guide

Complete installation instructions for Commit-AI on all supported platforms.

## Table of Contents
- [Windows](#windows)
- [macOS](#macos)
- [Linux](#linux)
- [Go Install](#go-install)
- [From Source](#from-source)
- [Verification](#verification)
- [Troubleshooting](#troubleshooting)

---

## Windows

### Option 1: Installer (Recommended)

The easiest way to install on Windows.

1. Download `install-commit-ai.exe` from [latest release](https://github.com/NeelFrostrain/Commit-Ai/releases/latest)
2. Double-click to run the installer
3. Follow the installation wizard
4. Restart your terminal or PowerShell
5. Run `commit-ai` to verify installation

**What the installer does:**
- Downloads the latest `commit-ai.exe` binary
- Creates `C:\Program Files\Commit-AI` directory
- Adds to system PATH
- Creates start menu shortcuts

### Option 2: Manual Installation

1. Download `commit-ai-windows-amd64.exe` from [releases](https://github.com/NeelFrostrain/Commit-Ai/releases/latest)
2. Rename to `commit-ai.exe`
3. Move to a directory in your PATH (e.g., `C:\Program Files\Commit-AI`)
4. Add the directory to your PATH:
   - Right-click "This PC" → Properties
   - Click "Advanced system settings"
   - Click "Environment Variables"
   - Under "System variables", click "Path" → "Edit"
   - Click "New" and add your directory
   - Click "OK" and restart terminal

### Option 3: Portable

1. Download `commit-ai-windows-amd64.exe`
2. Rename to `commit-ai.exe`
3. Place in your project directory
4. Run `.\commit-ai.exe` from PowerShell

---

## macOS

### Option 1: Homebrew (Coming Soon)

```bash
brew install commit-ai
```

### Option 2: Direct Download

**For Intel Macs:**
```bash
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-darwin-amd64 -o /usr/local/bin/commit-ai
chmod +x /usr/local/bin/commit-ai
```

**For Apple Silicon (M1/M2/M3):**
```bash
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-darwin-arm64 -o /usr/local/bin/commit-ai
chmod +x /usr/local/bin/commit-ai
```

### Option 3: Using MacPorts

```bash
sudo port install commit-ai
```

---

## Linux

### Option 1: Direct Download

**AMD64:**
```bash
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-linux-amd64 -o ~/.local/bin/commit-ai
chmod +x ~/.local/bin/commit-ai
```

**ARM64:**
```bash
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-linux-arm64 -o ~/.local/bin/commit-ai
chmod +x ~/.local/bin/commit-ai
```

### Option 2: Package Managers

**Debian/Ubuntu:**
```bash
# Coming soon
sudo apt install commit-ai
```

**Fedora/RHEL:**
```bash
# Coming soon
sudo dnf install commit-ai
```

**Arch Linux:**
```bash
# Coming soon
yay -S commit-ai
```

### Option 3: Snap

```bash
snap install commit-ai
```

---

## Go Install

If you have Go 1.21+ installed:

```bash
go install github.com/NeelFrostrain/Commit-Ai@latest
```

This installs to `$GOPATH/bin/commit-ai` (usually `~/go/bin/commit-ai`).

Make sure `$GOPATH/bin` is in your PATH:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

---

## From Source

### Prerequisites
- Go 1.21 or higher
- Git
- Make (optional, but recommended)

### Build Steps

```bash
# Clone the repository
git clone https://github.com/NeelFrostrain/Commit-Ai.git
cd Commit-Ai

# Download dependencies
go mod download

# Build for your platform
make build

# Or build directly
go build -o commit-ai

# Install to PATH
sudo mv commit-ai /usr/local/bin/  # macOS/Linux
# or
move commit-ai.exe C:\Program Files\Commit-AI\  # Windows
```

### Build for Specific Platform

```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o commit-ai-windows-amd64.exe

# macOS Intel
GOOS=darwin GOARCH=amd64 go build -o commit-ai-darwin-amd64

# macOS Apple Silicon
GOOS=darwin GOARCH=arm64 go build -o commit-ai-darwin-arm64

# Linux AMD64
GOOS=linux GOARCH=amd64 go build -o commit-ai-linux-amd64

# Linux ARM64
GOOS=linux GOARCH=arm64 go build -o commit-ai-linux-arm64
```

---

## Verification

Verify your installation:

```bash
# Check version
commit-ai version

# Should output:
# Commit-AI
#   Version:    1.2.0
#   Build Date: 2024-01-15
#   Git Commit: abc1234def5678
#   Go Version: go1.21.0
#   OS/Arch:    linux/amd64
```

---

## Troubleshooting

### "commit-ai: command not found"

**Solution 1: Add to PATH**
```bash
# macOS/Linux
export PATH=$PATH:/path/to/commit-ai/directory
echo 'export PATH=$PATH:/path/to/commit-ai/directory' >> ~/.bashrc

# Windows PowerShell
$env:Path += ";C:\Program Files\Commit-AI"
```

**Solution 2: Use full path**
```bash
/path/to/commit-ai version
```

### "Permission denied"

```bash
# Make executable
chmod +x /path/to/commit-ai
```

### "Cannot find module"

```bash
# Reinstall from source
go install github.com/NeelFrostrain/Commit-Ai@latest
```

### "Wrong architecture"

Download the correct binary for your system:
```bash
# Check your architecture
uname -m  # macOS/Linux
# or
[System.Environment]::Is64BitOperatingSystem  # Windows PowerShell
```

### Installation on Corporate Network

If behind a proxy:
```bash
go env -w GOPROXY=https://proxy.example.com
go install github.com/NeelFrostrain/Commit-Ai@latest
```

---

## Next Steps

1. [Get API Key](CONFIGURATION.md#getting-api-key)
2. [Quick Start](QUICK_START.md)
3. [Usage Guide](USAGE.md)

---

## Support

- **Issues**: https://github.com/NeelFrostrain/Commit-Ai/issues
- **Discussions**: https://github.com/NeelFrostrain/Commit-Ai/discussions
- **Email**: support@commit-ai.dev (coming soon)

# Update Guide

Complete guide for updating Commit-AI to the latest version.

---

## 🚀 Quick Update

The easiest way to update Commit-AI is using the built-in update command:

```bash
commit-ai update
```

This will:
1. Check GitHub for the latest release
2. Show you what's new
3. Download the appropriate binary for your platform
4. Install it automatically with backup

---

## 📋 Update Commands

### Check for Updates

Check if a new version is available without installing:

```bash
commit-ai update --check
```

**Output:**
```
[Update] Checking for updates...

═══════════════════════════════════════════
🎉 Update Available!
═══════════════════════════════════════════
  Current Version: v1.2.0
  Latest Version:  v1.3.0
  Release Name:    Enhanced Features

📝 What's New:
  - Added auto-update functionality
  - Improved error handling
  - Better performance
  ...

[Info] Full release notes: https://github.com/NeelFrostrain/Commit-Ai/releases/tag/v1.3.0
═══════════════════════════════════════════

[Hint] Run 'commit-ai update' to install the latest version
```

### Install Update

Install the latest version:

```bash
commit-ai update
```

**Interactive Process:**
1. Checks for updates
2. Shows release information
3. Asks for confirmation
4. Downloads new version
5. Installs with automatic backup
6. Verifies installation

### Force Update

Force update even if you're on the latest version (useful for reinstalling):

```bash
commit-ai update --force
```

---

## 🔍 How It Works

### 1. Version Check

The update command:
- Connects to GitHub API
- Fetches latest release information
- Compares with your current version
- Skips draft and pre-release versions

### 2. Platform Detection

Automatically detects your platform and downloads the correct binary:

- **Windows**: `commit-ai-windows-amd64.exe`
- **macOS Intel**: `commit-ai-darwin-amd64`
- **macOS Apple Silicon**: `commit-ai-darwin-arm64`
- **Linux AMD64**: `commit-ai-linux-amd64`
- **Linux ARM64**: `commit-ai-linux-arm64`

### 3. Safe Installation

The update process is safe:

1. Downloads to temporary file (`.tmp`)
2. Backs up current binary (`.backup`)
3. Replaces with new version
4. Removes backup on success
5. Restores backup on failure

### 4. Verification

After update, verify the new version:

```bash
commit-ai version
```

---

## 🛠️ Manual Update

If automatic update fails, you can update manually:

### Windows

```powershell
# Download latest release
$url = "https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-windows-amd64.exe"
Invoke-WebRequest -Uri $url -OutFile commit-ai.exe

# Or use installer
$url = "https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/install-commit-ai.exe"
Invoke-WebRequest -Uri $url -OutFile install-commit-ai.exe
.\install-commit-ai.exe
```

### macOS

```bash
# Intel
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-darwin-amd64 -o /usr/local/bin/commit-ai
chmod +x /usr/local/bin/commit-ai

# Apple Silicon (M1/M2/M3)
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-darwin-arm64 -o /usr/local/bin/commit-ai
chmod +x /usr/local/bin/commit-ai
```

### Linux

```bash
# AMD64
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-linux-amd64 -o ~/.local/bin/commit-ai
chmod +x ~/.local/bin/commit-ai

# ARM64
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-linux-arm64 -o ~/.local/bin/commit-ai
chmod +x ~/.local/bin/commit-ai
```

### Go Install

```bash
go install github.com/NeelFrostrain/Commit-Ai@latest
```

---

## 🐛 Troubleshooting

### "Failed to check for updates"

**Cause**: Network connectivity issue or GitHub API rate limit

**Solution**:
```bash
# Check internet connection
ping github.com

# Try again later
commit-ai update --check

# Or update manually (see above)
```

### "No binary found for your platform"

**Cause**: Unsupported platform or architecture

**Solution**:
- Check supported platforms: Windows, macOS, Linux (amd64, arm64)
- Build from source: See [Build Guide](BUILD_GUIDE.md)

### "Failed to install update: permission denied"

**Cause**: Insufficient permissions to replace binary

**Solution**:

**Windows:**
```powershell
# Run as Administrator
Start-Process commit-ai.exe -ArgumentList "update" -Verb RunAs
```

**macOS/Linux:**
```bash
# Use sudo if installed in system directory
sudo commit-ai update

# Or install to user directory
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-[platform]-[arch] -o ~/.local/bin/commit-ai
chmod +x ~/.local/bin/commit-ai
```

### "Download incomplete"

**Cause**: Network interruption during download

**Solution**:
```bash
# Try again
commit-ai update

# Or download manually
```

### Update Stuck or Frozen

**Cause**: Network timeout or slow connection

**Solution**:
- Wait for timeout (30 seconds for check, 5 minutes for download)
- Cancel with Ctrl+C
- Try again with better connection
- Update manually

---

## 🔐 Security

### Verification

The update process:
- Downloads only from official GitHub releases
- Verifies file size matches expected size
- Uses HTTPS for all connections
- Creates backup before replacing binary

### Best Practices

1. **Always verify after update:**
   ```bash
   commit-ai version
   ```

2. **Check release notes:**
   ```bash
   commit-ai update --check
   # Read the release notes before updating
   ```

3. **Keep backups:**
   - Automatic backup created during update
   - Manual backup: `cp commit-ai commit-ai.backup`

4. **Use official sources:**
   - Only update from GitHub releases
   - Verify repository: `github.com/NeelFrostrain/Commit-Ai`

---

## 📊 Update Frequency

### Recommended Schedule

- **Check for updates**: Weekly
- **Install updates**: When new features are needed
- **Security updates**: Immediately

### Automatic Checks

Currently, Commit-AI does not check for updates automatically. You must run:

```bash
commit-ai update --check
```

**Future Enhancement**: Automatic update notifications (see [IMPROVEMENTS.md](IMPROVEMENTS.md))

---

## 🎯 Version Information

### Check Current Version

```bash
commit-ai version
```

**Output:**
```
Commit-AI
  Version:    v1.2.0
  Build Date: 2026-03-04 12:06:06
  Git Commit: 7a2d729
  Go Version: go1.25.6
  OS/Arch:    windows/amd64
```

### Version Numbering

Commit-AI follows [Semantic Versioning](https://semver.org/):

- **Major** (v2.0.0): Breaking changes
- **Minor** (v1.3.0): New features, backward compatible
- **Patch** (v1.2.1): Bug fixes, backward compatible

---

## 🔗 Related Documentation

- [Quick Start Guide](QUICK_START.md) - Installation and setup
- [Build Guide](BUILD_GUIDE.md) - Building from source
- [Changelog](CHANGELOG.md) - Version history
- [Improvements](IMPROVEMENTS.md) - Future enhancements

---

## 💡 Tips

1. **Check before important work:**
   ```bash
   commit-ai update --check
   ```

2. **Update during downtime:**
   - Updates are quick (usually < 1 minute)
   - No configuration changes needed

3. **Read release notes:**
   - Understand new features
   - Check for breaking changes
   - Learn about improvements

4. **Verify after update:**
   ```bash
   commit-ai version
   commit-ai --help
   ```

---

## 🆘 Need Help?

- **Issues**: https://github.com/NeelFrostrain/Commit-Ai/issues
- **Releases**: https://github.com/NeelFrostrain/Commit-Ai/releases
- **Documentation**: https://github.com/NeelFrostrain/Commit-Ai/tree/main/docs

---

**Last Updated**: 2026-03-04  
**Version**: 1.2.0

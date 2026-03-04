# Auto-Update Feature

Complete documentation for the new auto-update functionality in Commit-AI v1.2.0.

---

## 🎯 Overview

Commit-AI now includes a built-in auto-update feature that allows users to easily check for and install the latest version directly from the command line.

---

## ✨ Features

### 1. Check for Updates
```bash
commit-ai update --check
```

- Connects to GitHub API
- Fetches latest release information
- Compares with current version
- Shows release notes preview
- No installation, just information

### 2. Install Updates
```bash
commit-ai update
```

- Downloads appropriate binary for your platform
- Shows progress during download
- Creates automatic backup
- Installs safely with rollback on failure
- Verifies installation

### 3. Force Update
```bash
commit-ai update --force
```

- Reinstalls even if on latest version
- Useful for corrupted installations
- Same safe installation process

---

## 🏗️ Architecture

### Components

1. **internal/updater/updater.go**
   - Core update logic
   - GitHub API integration
   - Download management
   - Safe installation with backup

2. **cmd/root.go**
   - Update command implementation
   - User interaction
   - Error handling

### Key Functions

```go
// Check for updates
CheckForUpdate(currentVersion string) (*Release, bool, error)

// Get platform-specific binary
GetAssetForPlatform(release *Release) (*Asset, error)

// Download new version
DownloadUpdate(asset *Asset, destPath string) error

// Install with backup
InstallUpdate(currentPath string) error

// Get current executable path
GetExecutablePath() (string, error)
```

---

## 🔄 Update Flow

```
1. User runs: commit-ai update
   ↓
2. Check GitHub API for latest release
   ↓
3. Compare versions
   ↓
4. Display update information
   ↓
5. Ask for confirmation
   ↓
6. Detect platform (OS/Arch)
   ↓
7. Download appropriate binary
   ↓
8. Create backup of current binary
   ↓
9. Replace with new version
   ↓
10. Remove backup on success
    ↓
11. Show success message
```

---

## 🛡️ Safety Features

### 1. Automatic Backup
- Current binary backed up before replacement
- Restored automatically on failure
- Removed only after successful installation

### 2. Download Verification
- Checks file size matches expected
- Validates HTTP status codes
- Handles network interruptions

### 3. Platform Detection
- Automatically detects OS and architecture
- Downloads correct binary
- Prevents incompatible installations

### 4. Error Handling
- Clear error messages
- Helpful hints for resolution
- Graceful fallback to manual update

---

## 📊 Supported Platforms

| Platform | Architecture | Binary Name |
|----------|-------------|-------------|
| Windows | amd64 | commit-ai-windows-amd64.exe |
| macOS | amd64 (Intel) | commit-ai-darwin-amd64 |
| macOS | arm64 (Apple Silicon) | commit-ai-darwin-arm64 |
| Linux | amd64 | commit-ai-linux-amd64 |
| Linux | arm64 | commit-ai-linux-arm64 |

---

## 🔧 Technical Details

### GitHub API Integration

**Endpoint**: `https://api.github.com/repos/NeelFrostrain/Commit-Ai/releases/latest`

**Response Structure**:
```json
{
  "tag_name": "v1.3.0",
  "name": "Release v1.3.0",
  "body": "Release notes...",
  "draft": false,
  "prerelease": false,
  "assets": [
    {
      "name": "commit-ai-windows-amd64.exe",
      "browser_download_url": "https://...",
      "size": 8388608
    }
  ]
}
```

### Version Comparison

Simple string comparison (assumes semantic versioning):
```go
currentVersion = strings.TrimPrefix(currentVersion, "v")
latestVersion := strings.TrimPrefix(release.TagName, "v")
hasUpdate := latestVersion > currentVersion && currentVersion != "dev"
```

### Download Process

1. HTTP GET request to asset URL
2. Stream to temporary file (`.tmp`)
3. Verify size matches expected
4. Make executable (Unix-like systems)

### Installation Process

1. Rename current binary to `.backup`
2. Rename `.tmp` to current binary name
3. Remove `.backup` on success
4. Restore `.backup` on failure

---

## 🎨 User Experience

### Progress Indicators

```
[Update] Checking for updates...
[Update] Downloading commit-ai-windows-amd64.exe (8.0 MB)...
[✓] Download complete!
[Update] Installing update...
🎉 Successfully updated to v1.3.0!
```

### Release Information Display

```
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

[Info] Full release notes: https://github.com/...
═══════════════════════════════════════════
```

---

## 🧪 Testing

### Test Scenarios

1. **Check for updates (up-to-date)**
   ```bash
   commit-ai update --check
   # Expected: "You are already on the latest version"
   ```

2. **Check for updates (update available)**
   ```bash
   # Build with old version
   commit-ai update --check
   # Expected: Shows update information
   ```

3. **Install update**
   ```bash
   commit-ai update
   # Expected: Downloads and installs
   ```

4. **Force update**
   ```bash
   commit-ai update --force
   # Expected: Reinstalls current version
   ```

5. **Network failure**
   ```bash
   # Disconnect network
   commit-ai update --check
   # Expected: Clear error message with hint
   ```

6. **Permission denied**
   ```bash
   # Install in protected directory
   commit-ai update
   # Expected: Error with sudo hint
   ```

---

## 📝 Code Examples

### Check for Updates

```go
release, hasUpdate, err := updater.CheckForUpdate(version)
if err != nil {
    fmt.Printf("Failed to check: %v\n", err)
    return
}

if !hasUpdate {
    fmt.Println("Already on latest version")
    return
}

updater.PrintUpdateInfo(release, version)
```

### Download and Install

```go
// Get executable path
exePath, err := updater.GetExecutablePath()

// Get platform-specific asset
asset, err := updater.GetAssetForPlatform(release)

// Download
err = updater.DownloadUpdate(asset, exePath)

// Install
err = updater.InstallUpdate(exePath)
```

---

## 🚀 Future Enhancements

### Planned Features

1. **Automatic Update Checks**
   - Check on startup (configurable)
   - Notify user of available updates
   - Option to disable

2. **Update Channels**
   - Stable (default)
   - Beta
   - Nightly

3. **Rollback Command**
   ```bash
   commit-ai rollback
   ```

4. **Update History**
   ```bash
   commit-ai update --history
   ```

5. **Proxy Support**
   - HTTP/HTTPS proxy configuration
   - Corporate network support

6. **Checksum Verification**
   - SHA256 checksums
   - GPG signatures

See [IMPROVEMENTS.md](IMPROVEMENTS.md) for full roadmap.

---

## 🐛 Known Limitations

1. **No automatic checks**: User must manually run update command
2. **No rollback**: Cannot revert to previous version (yet)
3. **No proxy support**: Direct internet connection required
4. **No checksum verification**: Relies on HTTPS and file size
5. **Simple version comparison**: May not handle all edge cases

---

## 📚 Related Documentation

- [Update Guide](UPDATE_GUIDE.md) - User guide for updating
- [Build Guide](BUILD_GUIDE.md) - Building from source
- [Release Checklist](RELEASE_CHECKLIST.md) - Creating releases
- [Changelog](CHANGELOG.md) - Version history

---

## 🤝 Contributing

Want to improve the update feature? See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### Areas for Contribution

- Automatic update checks
- Rollback functionality
- Checksum verification
- Proxy support
- Better error handling
- Progress bars
- Update notifications

---

## 📊 Statistics

- **Code Added**: ~400 lines
- **New Files**: 2 (updater.go, UPDATE_GUIDE.md)
- **Commands Added**: 1 (update)
- **Flags Added**: 2 (--check, --force)
- **Platforms Supported**: 5 (Windows, macOS x2, Linux x2)

---

**Feature Added**: 2026-03-04  
**Version**: 1.2.0  
**Status**: Production Ready

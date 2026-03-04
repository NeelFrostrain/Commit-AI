package updater

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

const (
	githubAPIURL = "https://api.github.com/repos/NeelFrostrain/Commit-Ai/releases/latest"
	githubRepo   = "https://github.com/NeelFrostrain/Commit-Ai"
	timeout      = 30 * time.Second
)

var (
	green  = color.New(color.FgGreen, color.Bold).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed, color.Bold).SprintFunc()
	cyan   = color.New(color.FgCyan, color.Bold).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
)

// Release represents a GitHub release
type Release struct {
	TagName    string  `json:"tag_name"`
	Name       string  `json:"name"`
	Body       string  `json:"body"`
	Draft      bool    `json:"draft"`
	Prerelease bool    `json:"prerelease"`
	Assets     []Asset `json:"assets"`
}

// Asset represents a release asset
type Asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
	Size               int64  `json:"size"`
}

// CheckForUpdate checks if a new version is available
func CheckForUpdate(currentVersion string) (*Release, bool, error) {
	client := &http.Client{Timeout: timeout}

	req, err := http.NewRequest("GET", githubAPIURL, nil)
	if err != nil {
		return nil, false, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, false, fmt.Errorf("failed to fetch release info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, false, fmt.Errorf("failed to parse release info: %w", err)
	}

	// Skip drafts and prereleases
	if release.Draft || release.Prerelease {
		return nil, false, nil
	}

	// Compare versions (simple string comparison, assumes semantic versioning)
	currentVersion = strings.TrimPrefix(currentVersion, "v")
	latestVersion := strings.TrimPrefix(release.TagName, "v")

	hasUpdate := latestVersion > currentVersion && currentVersion != "dev"

	return &release, hasUpdate, nil
}

// GetAssetForPlatform returns the appropriate asset for the current platform
func GetAssetForPlatform(release *Release) (*Asset, error) {
	platform := runtime.GOOS
	arch := runtime.GOARCH

	// Determine the expected filename
	var expectedName string
	switch platform {
	case "windows":
		expectedName = fmt.Sprintf("commit-ai-%s-%s.exe", platform, arch)
	case "darwin", "linux":
		expectedName = fmt.Sprintf("commit-ai-%s-%s", platform, arch)
	default:
		return nil, fmt.Errorf("unsupported platform: %s/%s", platform, arch)
	}

	// Find matching asset
	for _, asset := range release.Assets {
		if asset.Name == expectedName {
			return &asset, nil
		}
	}

	return nil, fmt.Errorf("no binary found for %s/%s", platform, arch)
}

// DownloadUpdate downloads the new version
func DownloadUpdate(asset *Asset, destPath string) error {
	fmt.Printf("%s Downloading %s (%s)...\n", blue("[Update]"), asset.Name, formatSize(asset.Size))

	client := &http.Client{Timeout: 5 * time.Minute}

	resp, err := client.Get(asset.BrowserDownloadURL)
	if err != nil {
		return fmt.Errorf("failed to download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status %d", resp.StatusCode)
	}

	// Create temporary file
	tmpFile := destPath + ".tmp"
	out, err := os.Create(tmpFile)
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer out.Close()

	// Download with progress
	written, err := io.Copy(out, resp.Body)
	if err != nil {
		os.Remove(tmpFile)
		return fmt.Errorf("failed to write file: %w", err)
	}

	if written != asset.Size {
		os.Remove(tmpFile)
		return fmt.Errorf("incomplete download: got %d bytes, expected %d", written, asset.Size)
	}

	return nil
}

// InstallUpdate replaces the current binary with the new one
func InstallUpdate(currentPath string) error {
	tmpFile := currentPath + ".tmp"
	backupFile := currentPath + ".backup"

	// Make new file executable (Unix-like systems)
	if runtime.GOOS != "windows" {
		if err := os.Chmod(tmpFile, 0755); err != nil {
			return fmt.Errorf("failed to make executable: %w", err)
		}
	}

	// Backup current binary
	if err := os.Rename(currentPath, backupFile); err != nil {
		return fmt.Errorf("failed to backup current binary: %w", err)
	}

	// Move new binary to current location
	if err := os.Rename(tmpFile, currentPath); err != nil {
		// Restore backup on failure
		os.Rename(backupFile, currentPath)
		return fmt.Errorf("failed to install update: %w", err)
	}

	// Remove backup
	os.Remove(backupFile)

	return nil
}

// GetExecutablePath returns the path to the current executable
func GetExecutablePath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}

	// Resolve symlinks
	exePath, err = filepath.EvalSymlinks(exePath)
	if err != nil {
		return "", fmt.Errorf("failed to resolve symlinks: %w", err)
	}

	return exePath, nil
}

// formatSize formats bytes to human-readable size
func formatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// PrintUpdateInfo displays information about available update
func PrintUpdateInfo(release *Release, currentVersion string) {
	fmt.Printf("\n%s\n", cyan("═══════════════════════════════════════════"))
	fmt.Printf("%s Update Available!\n", green("🎉"))
	fmt.Printf("%s\n", cyan("═══════════════════════════════════════════"))
	fmt.Printf("  Current Version: %s\n", yellow(currentVersion))
	fmt.Printf("  Latest Version:  %s\n", green(release.TagName))
	fmt.Printf("  Release Name:    %s\n", release.Name)
	fmt.Printf("\n%s What's New:\n", blue("📝"))

	// Print first few lines of release notes
	lines := strings.Split(release.Body, "\n")
	maxLines := 10
	for i, line := range lines {
		if i >= maxLines {
			fmt.Printf("  ...\n")
			break
		}
		if strings.TrimSpace(line) != "" {
			fmt.Printf("  %s\n", line)
		}
	}

	fmt.Printf("\n%s Full release notes: %s/releases/tag/%s\n", blue("[Info]"), githubRepo, release.TagName)
	fmt.Printf("%s\n", cyan("═══════════════════════════════════════════"))
}

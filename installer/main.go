package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	repoOwner = "NeelFrostrain"
	repoName  = "Commit-AI"
	appName   = "commit-ai"
	version   = "1.2.0"
)

type GitHubRelease struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func main() {
	printBanner()

	if err := install(); err != nil {
		fmt.Printf("\n❌ Installation failed: %v\n", err)
		fmt.Println("\n💡 Troubleshooting:")
		fmt.Println("  1. Check your internet connection")
		fmt.Println("  2. Make sure you have write permissions")
		fmt.Println("  3. Close any running commit-ai instances")
		fmt.Println("  4. Try running as administrator")
		fmt.Print("\nPress Enter to exit...")
		fmt.Scanln()
		os.Exit(1)
	}

	printSuccess()
	fmt.Print("\nPress Enter to exit...")
	fmt.Scanln()
}

func printBanner() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║     Commit-AI Installer v" + version + "        ║")
	fmt.Println("║  AI-Powered Commit Message Generator   ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()
}

func printSuccess() {
	fmt.Println("\n╔════════════════════════════════════════╗")
	fmt.Println("║   ✓ Installation Successful!          ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println("\n📝 Next Steps:")
	fmt.Println("  1. RESTART your terminal")
	fmt.Println("  2. Run: commit-ai")
	fmt.Println("  3. Enter your GROQ API key when prompted")
	fmt.Println("\n🔑 Get API Key: https://console.groq.com/keys")
	fmt.Println("\n📚 Documentation: https://github.com/" + repoOwner + "/" + repoName)
}

func install() error {
	// Determine installation directory
	installDir, err := getInstallDir()
	if err != nil {
		return fmt.Errorf("failed to determine install directory: %w", err)
	}

	exePath := filepath.Join(installDir, appName+getExeExtension())

	fmt.Println("📦 Installation Details:")
	fmt.Printf("  → Install Directory: %s\n", installDir)
	fmt.Printf("  → Binary Path: %s\n", exePath)
	fmt.Println()

	// Create directory
	fmt.Println("📁 Creating installation directory...")
	if err := os.MkdirAll(installDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	fmt.Println("  ✓ Directory ready")

	// Get latest release info
	fmt.Println("\n🔍 Fetching latest release...")
	release, err := getLatestRelease()
	if err != nil {
		return fmt.Errorf("failed to fetch release info: %w", err)
	}
	fmt.Printf("  ✓ Found version: %s\n", release.TagName)

	// Find appropriate asset
	downloadURL, err := findAsset(release)
	if err != nil {
		return fmt.Errorf("failed to find download: %w", err)
	}

	// Download binary
	fmt.Println("\n⬇️  Downloading binary...")
	if err := downloadFile(downloadURL, exePath); err != nil {
		return fmt.Errorf("failed to download: %w", err)
	}
	fmt.Println("  ✓ Download complete")

	// Make executable (Unix-like systems)
	if runtime.GOOS != "windows" {
		if err := os.Chmod(exePath, 0755); err != nil {
			return fmt.Errorf("failed to make executable: %w", err)
		}
	}

	// Update PATH
	fmt.Println("\n🔧 Updating PATH...")
	if err := updatePATH(installDir); err != nil {
		fmt.Printf("  ⚠️  Warning: Could not update PATH automatically: %v\n", err)
		fmt.Printf("  💡 Please add manually: %s\n", installDir)
	} else {
		fmt.Println("  ✓ PATH updated")
	}

	return nil
}

func getInstallDir() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", "CommitAI"), nil
	case "darwin", "linux":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, ".local", "bin"), nil
	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

func getExeExtension() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}
	return ""
}

func getLatestRelease() (*GitHubRelease, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", repoOwner, repoName)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %s", resp.Status)
	}

	var release GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}

	return &release, nil
}

func findAsset(release *GitHubRelease) (string, error) {
	// Determine platform-specific binary name
	var assetName string
	switch runtime.GOOS {
	case "windows":
		assetName = appName + "-windows-amd64.exe"
	case "darwin":
		if runtime.GOARCH == "arm64" {
			assetName = appName + "-darwin-arm64"
		} else {
			assetName = appName + "-darwin-amd64"
		}
	case "linux":
		if runtime.GOARCH == "arm64" {
			assetName = appName + "-linux-arm64"
		} else {
			assetName = appName + "-linux-amd64"
		}
	default:
		return "", fmt.Errorf("unsupported platform: %s/%s", runtime.GOOS, runtime.GOARCH)
	}

	// Find matching asset
	for _, asset := range release.Assets {
		if asset.Name == assetName {
			return asset.BrowserDownloadURL, nil
		}
	}

	// Fallback to generic name
	fallbackName := appName + getExeExtension()
	for _, asset := range release.Assets {
		if asset.Name == fallbackName {
			return asset.BrowserDownloadURL, nil
		}
	}

	return "", fmt.Errorf("no compatible binary found for %s/%s", runtime.GOOS, runtime.GOARCH)
}

func downloadFile(url, filepath string) error {
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status: %s", resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Show progress
	fmt.Print("  Progress: ")
	written, err := io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%.2f MB\n", float64(written)/(1024*1024))

	return nil
}

func updatePATH(installDir string) error {
	switch runtime.GOOS {
	case "windows":
		return updateWindowsPATH(installDir)
	case "darwin", "linux":
		return updateUnixPATH(installDir)
	default:
		return fmt.Errorf("unsupported OS")
	}
}

func updateWindowsPATH(installDir string) error {
	pathEnv := os.Getenv("PATH")
	if strings.Contains(strings.ToLower(pathEnv), strings.ToLower(installDir)) {
		return nil // Already in PATH
	}

	cmd := exec.Command("setx", "PATH", pathEnv+";"+installDir)
	return cmd.Run()
}

func updateUnixPATH(installDir string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Try to update shell rc files
	shells := []string{
		filepath.Join(home, ".bashrc"),
		filepath.Join(home, ".zshrc"),
		filepath.Join(home, ".profile"),
	}

	exportLine := fmt.Sprintf("\n# Added by Commit-AI installer\nexport PATH=\"$PATH:%s\"\n", installDir)

	updated := false
	for _, shell := range shells {
		if _, err := os.Stat(shell); err == nil {
			// Check if already added
			content, err := os.ReadFile(shell)
			if err != nil {
				continue
			}

			if strings.Contains(string(content), installDir) {
				updated = true
				continue
			}

			// Append to file
			f, err := os.OpenFile(shell, os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				continue
			}
			defer f.Close()

			if _, err := f.WriteString(exportLine); err != nil {
				continue
			}
			updated = true
		}
	}

	if !updated {
		return fmt.Errorf("could not find shell rc file")
	}

	return nil
}

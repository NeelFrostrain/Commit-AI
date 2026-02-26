package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type GitHubRelease struct {
	TagName string `json:"tag_name"`
}

func getLatestTag(repoPath string) (string, error) {
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repoPath)
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(apiUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "unknown", fmt.Errorf("API returned status: %s", resp.Status)
	}
	var release GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}
	return release.TagName, nil
}

func main() {
	repoOwnerRepo := "NeelFrostrain/Commit-AI"
	repoURL := fmt.Sprintf("https://github.com/%s/releases/latest/download/commit-ai.exe", repoOwnerRepo)
	installDir := filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", "CommitAI")
	exePath := filepath.Join(installDir, "commit-ai.exe")

	fmt.Println("===========================================")
	fmt.Println("       COMMIT-AI INSTALLER / UPDATER       ")
	fmt.Println("===========================================")

	// 1. Version Check
	fmt.Print("[1/4] Checking for latest version... ")
	tag, err := getLatestTag(repoOwnerRepo)
	if err != nil {
		fmt.Printf("\n[!] Warning: %v\n", err)
	} else {
		fmt.Printf("Done (%s)\n", tag)
	}

	// 2. Preparation (Kill running process)
	fmt.Print("[2/4] Preparing environment... ")
	_ = exec.Command("taskkill", "/F", "/IM", "commit-ai.exe").Run() // Kill if running
	if err := os.MkdirAll(installDir, 0755); err != nil {
		fmt.Printf("\n[Error] Permission denied: %v\n", err)
		return
	}
	fmt.Println("Done")

	// 3. Download with simple progress
	fmt.Println("[3/4] Downloading binary...")
	if err := downloadFile(exePath, repoURL); err != nil {
		fmt.Printf("[Error] Download failed: %v\n", err)
		return
	}
	fmt.Println("      ✓ Download complete")

	// 4. Smart PATH Update
	fmt.Print("[4/4] Configuring System PATH... ")
	updatePath(installDir)

	fmt.Println("\n===========================================")
	fmt.Println(" SUCCESS: Commit-AI is now installed!")
	fmt.Println("===========================================")
	fmt.Println("\nIMPORTANT:")
	fmt.Println("1. Restart your Terminal/VS Code.")
	fmt.Println("2. Type 'commit-ai' to start.")
	fmt.Println("\nPress Enter to close...")
	fmt.Scanln()
}

// downloadFile saves the EXE and provides a basic spinner
func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// updatePath checks length limits before using setx
func updatePath(targetDir string) {
	currentPath := os.Getenv("PATH")
	
	// Check if already in PATH
	if strings.Contains(strings.ToLower(currentPath), strings.ToLower(targetDir)) {
		fmt.Println("Already configured.")
		return
	}

	// setx limit is 1024 characters. If PATH is longer, setx will corrupt it.
	if len(currentPath)+len(targetDir) > 1020 {
		fmt.Println("\n[!] Warning: PATH is too long for automatic update.")
		fmt.Printf("    Please manually add this to your PATH: %s\n", targetDir)
		return
	}

	newPath := currentPath + ";" + targetDir
	cmd := exec.Command("setx", "PATH", newPath)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed: %v\n", err)
	} else {
		fmt.Println("Done.")
	}
}
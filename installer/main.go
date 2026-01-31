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

// GitHubRelease struct to parse the API response
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
	// 1. Configuration
	repoOwnerRepo := "NeelFrostrain/Commit-AI" // Used for API
	repoURL := fmt.Sprintf("https://github.com/%s/releases/latest/download/commit-ai.exe", repoOwnerRepo)

	installDir := filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", "CommitAI")
	exePath := filepath.Join(installDir, "commit-ai.exe")

	fmt.Println("-------------------------------------------")
	fmt.Println("[Commit-AI] Fetching version info...")

	// --- NEW: Get and show the tag first ---
	tag, err := getLatestTag(repoOwnerRepo)
	if err == nil {
		fmt.Printf("[Info] Found Latest Version: %s\n", tag)
	} else {
		fmt.Printf("[Warning] Could not fetch tag: %v\n", err)
	}

	fmt.Println("[Commit-AI] Starting Installation...")
	fmt.Println("-------------------------------------------")

	// 2. Create Directory
	if _, err := os.Stat(installDir); os.IsNotExist(err) {
		err := os.MkdirAll(installDir, 0755)
		if err != nil {
			fmt.Printf("[Error] Failed to create folder: %v\n", err)
			return
		}
		fmt.Println("[Info] Created directory:", installDir)
	}

	// 3. Download the EXE from GitHub
	fmt.Println("[Info] Downloading binary...")
	resp, err := http.Get(repoURL)
	if err != nil {
		fmt.Printf("[Error] Failed to download: %v\n", err)
		return
	}
	defer resp.Body.Close()

	out, err := os.Create(exePath)
	if err != nil {
		fmt.Printf("[Error] Failed to create file (is the app running?): %v\n", err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("[Error] Failed to save file: %v\n", err)
		return
	}
	fmt.Println("[Success] Downloaded commit-ai.exe")

	// 4. Update PATH using setx
	pathEnv := os.Getenv("PATH")
	if !strings.Contains(strings.ToLower(pathEnv), strings.ToLower(installDir)) {
		cmd := exec.Command("setx", "PATH", pathEnv+";"+installDir)
		if err := cmd.Run(); err != nil {
			fmt.Printf("[Error] Failed to update PATH: %v\n", err)
		} else {
			fmt.Println("[Success] Added to PATH.")
		}
	} else {
		fmt.Println("[Info] Already in PATH.")
	}

	fmt.Println("\n-------------------------------------------")
	fmt.Println("Done! Please RESTART your terminal to use 'commit-ai'.")
	fmt.Print("Press Enter to exit...")
	fmt.Scanln()
}

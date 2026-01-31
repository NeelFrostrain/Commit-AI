package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/NeelFrostrain/Commit-Ai-Go/internal/ai"
	"github.com/NeelFrostrain/Commit-Ai-Go/internal/git"

	"github.com/algolyzer/groq-go"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	commitFlag bool
	yesFlag    bool
	magenta    = color.New(color.FgMagenta, color.Bold).SprintFunc()
	blue       = color.New(color.FgBlue).SprintFunc()
	red        = color.New(color.FgRed, color.Bold).SprintFunc()
	cyan       = color.New(color.FgCyan).SprintFunc()
)

var rootCmd = &cobra.Command{
	Use:     "commit-ai",
	Short:   "AI-powered git analysis and auto-committer",
	Version: "1.1.0",
	Run:     runCommitAI,
}

func init() {
	rootCmd.Flags().BoolVarP(&commitFlag, "commit", "c", false, "enable commit mode")
	rootCmd.Flags().BoolVarP(&yesFlag, "yes", "y", false, "skip confirmation prompt")
}

func runCommitAI(cmd *cobra.Command, args []string) {
	_ = godotenv.Load()

	apiKey := getAPIKey()
	if apiKey == "" {
		fmt.Printf("%s %s: GROQ_API_KEY is missing. Please set it and restart your terminal.\n", magenta("[Commit-AI]"), red("[Error]"))
		os.Exit(1)
	}

	fmt.Printf("%s %s: Analyzing modified files...\n", magenta("[Commit-AI]"), blue("[Info]"))
	exec.Command("git", "add", ".").Run()

	excludePatterns := git.GetIgnorePatterns()
	diffArgs := append([]string{"diff", "--cached", "--", "."}, excludePatterns...)
	out, _ := exec.Command("git", diffArgs...).Output()
	diff := string(out)

	if strings.TrimSpace(diff) == "" {
		fmt.Printf("%s %s: No changes detected.\n", magenta("[Commit-AI]"), color.GreenString("[Success]"))
		return
	}

	client := groq.NewClient(apiKey)
	temp := 0.7
	prompt := fmt.Sprintf(`Analyze this Git diff and provide a professional report.
    1. Provide a bulleted "REPORT" of technical changes.
    2. Provide a single-line "COMMIT_MESSAGE" (type: description).
    STRICT FORMAT:
    <report>* bullet points</report>
    <message>type: description</message>
    Diff: %s`, diff)

	resp, err := client.CreateChatCompletion(context.Background(), groq.ChatCompletionRequest{
		Model: "llama-3.1-8b-instant",
		Messages: []groq.ChatMessage{
			{Role: groq.RoleSystem, Content: "You are commit-ai. Follow tags strictly."},
			{Role: groq.RoleUser, Content: prompt},
		},
		Temperature: &temp,
	})

	if err != nil {
		fmt.Printf("AI failure: %v\n", err)
		return
	}

	title, report := ai.ParseResponse(resp.Choices[0].Message.Content)

	fmt.Printf("\n%s\nREPORT:\n%s\n\nCOMMIT_MESSAGE: %s\n%s\n\n", red("─── AI SUGGESTION ───"), report, title, red("─────────────────────"))

	if commitFlag {
		if yesFlag || confirm() {
			exec.Command("git", "commit", "-m", title, "-m", report).Run()
			fmt.Println("Committed!")
		}
	}
}

func getAPIKey() string {
	// 1. Check current terminal session environment
	key := os.Getenv("GROQ_API_KEY")
	if key != "" {
		return key
	}

	// 2. Check the hidden file in Home directory (~/.commit-ai-key)
	home, _ := os.UserHomeDir()
	keyFile := filepath.Join(home, ".commit-ai-key") // Use filepath for Windows compatibility
	data, err := os.ReadFile(keyFile)
	if err == nil && len(strings.TrimSpace(string(data))) > 0 {
		return strings.TrimSpace(string(data))
	}

	// 3. First-time setup: Ask user and save it
	fmt.Println(color.YellowString("\n[Setup] No GROQ_API_KEY found."))
	fmt.Println("You can get one at: https://console.groq.com/keys")
	fmt.Print("Paste your API Key here: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input != "" {
		// Save to file as backup
		_ = os.WriteFile(keyFile, []byte(input), 0600)

		// Try to set it permanently in Windows Registry so other terminals see it
		// This is the Go version of 'setx'
		_ = exec.Command("setx", "GROQ_API_KEY", input).Run()

		fmt.Println(color.GreenString("✔ Key saved!"))
		fmt.Println(color.CyanString("NOTE: Please close this terminal and open a NEW one for changes to take effect.\n"))
		return input
	}

	return ""
}

func confirm() bool {
	fmt.Printf("%s %s: Use this commit message? (y/n): ", magenta("[Commit-AI]"), color.YellowString("[Prompt]"))
	r := bufio.NewReader(os.Stdin)
	i, _ := r.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(i)) == "y"
}

func Execute() {
	rootCmd.Execute()
}

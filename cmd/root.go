package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/NeelFrostrain/Commit-Ai-Go/internal/ai"
	"github.com/NeelFrostrain/Commit-Ai-Go/internal/git"

	"github.com/AlecAivazis/survey/v2"
	"github.com/algolyzer/groq-go"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	magenta = color.New(color.FgMagenta, color.Bold).SprintFunc()
	blue    = color.New(color.FgBlue).SprintFunc()
	red     = color.New(color.FgRed, color.Bold).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
)

var commitFlag bool
var yesFlag bool

var rootCmd = &cobra.Command{
	Use:   "commit-ai",
	Short: "AI-powered semantic commit generator",
	Run:   runCommitAI,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolVarP(&commitFlag, "commit", "c", false, "Commit changes after selection")
	rootCmd.Flags().BoolVarP(&yesFlag, "yes", "y", false, "Skip confirmation prompts")
}

func runCommitAI(cmd *cobra.Command, args []string) {
	_ = godotenv.Load()
	apiKey := getAPIKey()

	// 1. Smart Staging Check
	staged, _ := exec.Command("git", "diff", "--cached", "--name-only").Output()
	if len(strings.TrimSpace(string(staged))) == 0 {
		var confirm bool
		fmt.Printf("%s No staged changes detected.\n", yellow("[!]"))
		survey.AskOne(&survey.Confirm{Message: "Stage all files and continue?"}, &confirm)
		if !confirm { return }
		exec.Command("git", "add", ".").Run()
	}

	// 2. Diff Optimization
	// We only send the summary + the diff to give the AI context of what was ADDED vs DELETED
	exclude := git.GetIgnorePatterns()
	diffArgs := append([]string{"diff", "--cached", "--stat", "--"}, exclude...)
	stats, _ := exec.Command("git", diffArgs...).Output()
	
	// Get actual code changes
	diffArgs = append([]string{"diff", "--cached", "-U2", "--"}, exclude...) // Lower context lines to save tokens
	out, _ := exec.Command("git", diffArgs...).Output()
	
	fullContext := fmt.Sprintf("Summary:\n%s\n\nDiff:\n%s", stats, out)
	if len(fullContext) > 12000 {
		fullContext = fullContext[:12000] + "\n...[diff truncated]"
	}

	// 3. AI Request with "Thinking" indicator
	fmt.Printf("%s %s: Brainstorming commit options...", magenta("[Commit-AI]"), blue("[Info]"))
	
	client := groq.NewClient(apiKey)
	temp := 0.5 // Slightly higher for better creativity in options
	resp, err := client.CreateChatCompletion(context.Background(), groq.ChatCompletionRequest{
		Model: "llama-3.1-8b-instant",
		Messages: []groq.ChatMessage{
			{
				Role: groq.RoleSystem,
				Content: `You are a Principal Engineer. Analyze the diff and provide 3 distinct semantic titles in <options> and a concise technical report in <report>.
				If files were deleted, use 'refactor' or 'chore'. If new files, use 'feat'.`,
			},
			{Role: groq.RoleUser, Content: ai.BuildPrompt(fullContext)},
		},
		Temperature: &temp,
	})
	fmt.Print("\r") // Clear the "Analyzing" line

	if err != nil {
		fmt.Printf("%s AI Error: %v\n", red("[Error]"), err)
		return
	}

	// 4. Interactive Selection (Title)
	options, aiReport := ai.ParseMultiResponse(resp.Choices[0].Message.Content)
	var selectedTitle string
	survey.AskOne(&survey.Select{
		Message: "Pick a Title:",
		Options: append(options, "Edit manually...", "Cancel"),
	}, &selectedTitle)

	if selectedTitle == "Cancel" { return }
	if selectedTitle == "Edit manually..." {
		survey.AskOne(&survey.Input{Message: "Custom Title:", Default: options[0]}, &selectedTitle)
	}

	// 5. Interactive Selection (Report)
	var selectedReport string
	reportChoices := []string{"Keep AI Report", "Edit Report", "No Report"}
	var reportAction string
	survey.AskOne(&survey.Select{Message: "Commit Details:", Options: reportChoices}, &reportAction)

	switch reportAction {
	case "Keep AI Report":
		selectedReport = aiReport
	case "Edit Report":
		survey.AskOne(&survey.Multiline{Message: "Edit Body:", Default: aiReport}, &selectedReport)
	case "No Report":
		selectedReport = ""
	}

	// 6. Execution
	if commitFlag {
		fmt.Printf("\n%s\n%s %s\n%s %s\n", magenta("─── FINAL ───"), blue("TITLE:"), selectedTitle, blue("BODY:"), selectedReport)
		if yesFlag || askConfirm("Execute git commit?") {
			args := []string{"commit", "-m", selectedTitle}
			if selectedReport != "" {
				args = append(args, "-m", selectedReport)
			}
			cmd := exec.Command("git", args...)
			cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
			if err := cmd.Run(); err == nil {
				fmt.Println(color.GreenString("✔ Commit created!"))
			}
		}
	}
}

func getAPIKey() string {
	if key := os.Getenv("GROQ_API_KEY"); key != "" { return key }
	fmt.Println(red("GROQ_API_KEY not found in environment.")); os.Exit(1)
	return ""
}

func askConfirm(msg string) bool {
	res := false
	survey.AskOne(&survey.Confirm{Message: msg}, &res)
	return res
}
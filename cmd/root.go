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

	// 1. Git Staging Check
	staged, _ := exec.Command("git", "diff", "--cached", "--name-only").Output()
	if len(strings.TrimSpace(string(staged))) == 0 {
		fmt.Printf("%s No staged changes. Stage all? (y/n): ", yellow("[?]"))
		var confirm bool
		survey.AskOne(&survey.Confirm{Message: "Stage all files?"}, &confirm)
		if confirm {
			exec.Command("git", "add", ".").Run()
		} else {
			return
		}
	}

	fmt.Printf("%s %s: Analyzing changes...\n", magenta("[Commit-AI]"), blue("[Info]"))

	// 2. Diff Extraction (Limit to 10k chars for context)
	exclude := git.GetIgnorePatterns()
	diffArgs := append([]string{"diff", "--cached", "--"}, exclude...)
	out, _ := exec.Command("git", diffArgs...).Output()
	diffText := string(out)
	if len(diffText) > 10000 {
		diffText = diffText[:10000] + "\n...[truncated]"
	}

	// 3. AI Request
	client := groq.NewClient(apiKey)
	temp := 0.3
	resp, err := client.CreateChatCompletion(context.Background(), groq.ChatCompletionRequest{
		Model: "llama-3.1-8b-instant",
		Messages: []groq.ChatMessage{
			{
				Role: groq.RoleSystem,
				Content: `You are a senior engineer. Provide 3 semantic commit options in <options> tags and a CONCISE summary in <report> tags.
				Use types: feat, fix, refactor, chore, docs, style, perf.`,
			},
			{Role: groq.RoleUser, Content: ai.BuildPrompt(diffText)},
		},
		Temperature: &temp,
	})

	if err != nil {
		fmt.Printf("%s AI Error: %v\n", red("[Error]"), err)
		return
	}

	// 4. Interaction Logic
	options, report := ai.ParseMultiResponse(resp.Choices[0].Message.Content)
	
	var selected string
	prompt := &survey.Select{
		Message: "Select a commit message:",
		Options: append(options, "Enter manually...", "Cancel"),
	}
	survey.AskOne(prompt, &selected)

	if selected == "Cancel" || selected == "" {
		return
	}

	if selected == "Enter manually..." {
		survey.AskOne(&survey.Input{Message: "Custom message:"}, &selected)
	}

	// 5. Finalize
	fmt.Printf("\n%s\nREPORT:\n%s\n%s\n", magenta("─── DETAILS ───"), report, magenta("───────────────"))

	if commitFlag {
		if yesFlag || askConfirm("Commit with this message?") {
			commitCmd := exec.Command("git", "commit", "-m", selected, "-m", report)
			commitCmd.Stdout, commitCmd.Stderr = os.Stdout, os.Stderr
			commitCmd.Run()
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
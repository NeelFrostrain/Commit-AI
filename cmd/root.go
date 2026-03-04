package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/NeelFrostrain/Commit-Ai/internal/ai"
	"github.com/NeelFrostrain/Commit-Ai/internal/config"
	"github.com/NeelFrostrain/Commit-Ai/internal/git"

	"github.com/AlecAivazis/survey/v2"
	"github.com/algolyzer/groq-go"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	magenta = color.New(color.FgMagenta, color.Bold).SprintFunc()
	blue    = color.New(color.FgBlue).SprintFunc()
	red     = color.New(color.FgRed, color.Bold).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
	green   = color.New(color.FgGreen, color.Bold).SprintFunc()
	cyan    = color.New(color.FgCyan, color.Bold).SprintFunc()
)

var (
	commitFlag  bool
	yesFlag     bool
	verboseFlag bool
	modelFlag   string
)

var (
	version   = "dev"
	buildDate = "unknown"
	gitCommit = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "commit-ai",
	Short: "AI-powered semantic commit generator",
	Long: `Commit-AI analyzes your git changes and generates professional,
conventional commit messages using AI. It helps maintain clean
git history with minimal effort.`,
	Run: runCommitAI,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  "Display detailed version information including build date and git commit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", cyan("Commit-AI"))
		fmt.Printf("  Version:    %s\n", version)
		fmt.Printf("  Build Date: %s\n", buildDate)
		fmt.Printf("  Git Commit: %s\n", gitCommit)
		fmt.Printf("  Go Version: %s\n", runtime.Version())
		fmt.Printf("  OS/Arch:    %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}

// SetVersion sets version information from main package
func SetVersion(v, date, commit string) {
	version = v
	buildDate = date
	gitCommit = commit
	rootCmd.Version = v
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().BoolVarP(&commitFlag, "commit", "c", false, "Commit changes after selection")
	rootCmd.Flags().BoolVarP(&yesFlag, "yes", "y", false, "Skip confirmation prompts")
	rootCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "Show detailed information")
	rootCmd.Flags().StringVarP(&modelFlag, "model", "m", "", "Override AI model (default: llama-3.1-8b-instant)")
}

func runCommitAI(cmd *cobra.Command, args []string) {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("%s %v\n", red("[Error]"), err)
		fmt.Printf("%s Visit https://console.groq.com/keys to get your API key\n", yellow("[Info]"))

		// Offer to save API key
		var apiKey string
		prompt := &survey.Input{Message: "Enter your GROQ API key (or press Enter to exit):"}
		if err := survey.AskOne(prompt, &apiKey); err != nil || apiKey == "" {
			return
		}

		if err := config.SaveAPIKey(apiKey); err != nil {
			fmt.Printf("%s Failed to save API key: %v\n", red("[Error]"), err)
			return
		}

		fmt.Printf("%s API key saved to ~/.commit-ai.env\n", green("[Success]"))
		cfg = &config.Config{APIKey: apiKey, Model: "llama-3.1-8b-instant", Temperature: 0.7, MaxTokens: 15000}
	}

	// Override model if flag is set
	if modelFlag != "" {
		cfg.Model = modelFlag
	}

	// Check for ANY changes (staged or unstaged)
	hasStaged, err := git.HasStagedChanges()
	if err != nil {
		fmt.Printf("%s Failed to check staged changes: %v\n", red("[Error]"), err)
		return
	}

	hasUnstaged, err := git.HasUnstagedChanges()
	if err != nil {
		fmt.Printf("%s Failed to check unstaged changes: %v\n", red("[Error]"), err)
		return
	}

	if !hasStaged && !hasUnstaged {
		fmt.Printf("%s No changes detected in repository.\n", yellow("[!]"))
		return
	}

	// If there are unstaged changes, offer to stage them
	if hasUnstaged {
		if !yesFlag {
			fmt.Printf("%s Found unstaged changes.\n", yellow("[!]"))
			var confirm bool
			survey.AskOne(&survey.Confirm{Message: "Stage all changes (git add .)?"}, &confirm)
			if !confirm {
				if !hasStaged {
					fmt.Printf("%s No staged changes to commit.\n", yellow("[!]"))
					return
				}
				fmt.Printf("%s Proceeding with only staged changes...\n", blue("[Info]"))
			} else {
				if err := git.StageAllFiles(); err != nil {
					fmt.Printf("%s Failed to stage files: %v\n", red("[Error]"), err)
					return
				}
				fmt.Printf("%s All changes staged\n", green("[✓]"))
			}
		} else {
			// Auto-stage in yes mode
			if err := git.StageAllFiles(); err != nil {
				fmt.Printf("%s Failed to stage files: %v\n", red("[Error]"), err)
				return
			}
			fmt.Printf("%s All changes staged\n", green("[✓]"))
		}
	}

	// Get staged files for scope suggestion
	stagedFiles, _ := git.GetStagedFiles()
	suggestedScope := ai.SuggestScope(stagedFiles)

	if verboseFlag {
		if suggestedScope != "" {
			fmt.Printf("%s Detected scope: %s\n", blue("[Info]"), suggestedScope)
		}
		fmt.Printf("%s Analyzing %d files\n", blue("[Info]"), len(stagedFiles))
	}

	// Get optimized diff
	exclude := git.GetIgnorePatterns()
	fullContext, err := git.GetStagedDiff(exclude, 8000) // Reduced for API limits
	if err != nil {
		fmt.Printf("%s Failed to get diff: %v\n", red("[Error]"), err)
		return
	}

	if verboseFlag {
		fmt.Printf("%s Diff size: %d characters\n", blue("[Info]"), len(fullContext))
	}

	// AI Request with progress indicator
	fmt.Printf("%s Analyzing changes with %s...\n", magenta("[Commit-AI]"), cfg.Model)

	client := groq.NewClient(cfg.APIKey)
	temp := cfg.Temperature
	resp, err := client.CreateChatCompletion(context.Background(), groq.ChatCompletionRequest{
		Model: cfg.Model,
		Messages: []groq.ChatMessage{
			{
				Role: groq.RoleSystem,
				Content: `You are a Principal Engineer and Git expert specializing in writing exceptional commit messages.

Your expertise includes:
- Deep understanding of Conventional Commits specification
- Ability to analyze code changes and understand their impact
- Writing clear, comprehensive technical documentation
- Organizing information into logical, scannable categories
- Providing actionable insights and metrics

When analyzing changes:
1. Identify the primary change type and scope accurately
2. Group related changes into logical categories (ARCHITECTURE, FEATURES, BUG FIXES, etc.)
3. Provide specific technical details with context
4. Include metrics, file counts, and measurements
5. Explain both WHAT changed and WHY it matters
6. Structure reports for easy scanning with clear headers
7. End with IMPACT section showing tangible benefits

Your commit messages should be:
- Professional and technically accurate
- Well-organized with clear categories
- Detailed enough to understand without reading code
- Focused on impact and reasoning, not just changes
- Following Conventional Commits format strictly`,
			},
			{Role: groq.RoleUser, Content: ai.BuildPrompt(fullContext)},
		},
		Temperature: &temp,
	})

	if err != nil {
		fmt.Printf("%s AI request failed: %v\n", red("[Error]"), err)
		fmt.Printf("%s Check your API key and network connection\n", yellow("[Hint]"))
		return
	}

	// Parse AI response
	rawResponse := resp.Choices[0].Message.Content
	options, aiReport := ai.ParseMultiResponse(rawResponse)

	if verboseFlag {
		fmt.Printf("%s Generated %d options\n", blue("[Info]"), len(options))
		// Debug: Show if AI didn't follow format
		if len(options) == 1 && options[0] == "chore: update project" {
			fmt.Printf("%s AI didn't follow format. First 300 chars of response:\n", yellow("[Debug]"))
			maxLen := 300
			if len(rawResponse) < maxLen {
				maxLen = len(rawResponse)
			}
			fmt.Printf("%s\n", rawResponse[:maxLen])
		}
	}

	// Validate options
	validOptions := make([]string, 0, len(options))
	for _, opt := range options {
		if ai.ValidateCommitMessage(opt) {
			validOptions = append(validOptions, opt)
		} else if verboseFlag {
			fmt.Printf("%s Skipping invalid option: %s\n", yellow("[Warning]"), opt)
		}
	}

	if len(validOptions) == 0 {
		validOptions = options // Fallback to all options if validation is too strict
	}

	// Interactive title selection
	var selectedTitle string
	titlePrompt := &survey.Select{
		Message: "Select commit title:",
		Options: append(validOptions, "✏️  Edit manually...", "❌ Cancel"),
	}

	if err := survey.AskOne(titlePrompt, &selectedTitle); err != nil {
		return
	}

	if selectedTitle == "❌ Cancel" {
		fmt.Println("Cancelled.")
		return
	}

	if selectedTitle == "✏️  Edit manually..." {
		inputPrompt := &survey.Input{
			Message: "Enter custom commit title:",
			Default: validOptions[0],
		}
		if err := survey.AskOne(inputPrompt, &selectedTitle); err != nil || selectedTitle == "" {
			return
		}
	}

	// Interactive report selection
	var selectedReport string
	reportChoices := []string{"📝 Keep AI report", "✏️  Edit report", "⊘  No report"}
	var reportAction string

	reportPrompt := &survey.Select{
		Message: "Commit body:",
		Options: reportChoices,
	}

	if err := survey.AskOne(reportPrompt, &reportAction); err != nil {
		return
	}

	switch reportAction {
	case "📝 Keep AI report":
		selectedReport = aiReport
	case "✏️  Edit report":
		editPrompt := &survey.Multiline{
			Message: "Edit commit body:",
			Default: aiReport,
		}
		survey.AskOne(editPrompt, &selectedReport)
	case "⊘  No report":
		selectedReport = ""
	}

	// Display final commit message
	fmt.Printf("\n%s\n", magenta("─────────────────────────────────"))
	fmt.Printf("%s %s\n", blue("Title:"), selectedTitle)
	if selectedReport != "" {
		fmt.Printf("%s\n%s\n", blue("Body:"), selectedReport)
	}
	fmt.Printf("%s\n", magenta("─────────────────────────────────"))

	// Execute commit
	if commitFlag {
		shouldCommit := yesFlag
		if !yesFlag {
			survey.AskOne(&survey.Confirm{Message: "Execute git commit?"}, &shouldCommit)
		}

		if shouldCommit {
			args := []string{"commit", "-m", selectedTitle}
			if selectedReport != "" {
				args = append(args, "-m", selectedReport)
			}

			cmd := exec.Command("git", args...)
			cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

			if err := cmd.Run(); err != nil {
				fmt.Printf("%s Commit failed: %v\n", red("[Error]"), err)
				return
			}

			fmt.Printf("\n%s Commit created successfully!\n", green("✔"))
		}
	} else {
		fmt.Printf("\n%s Use -c flag to commit automatically\n", yellow("[Hint]"))
	}
}

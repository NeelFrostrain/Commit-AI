package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
<<<<<<< Updated upstream
	"path/filepath"
	"strings"
=======
>>>>>>> Stashed changes

	"github.com/NeelFrostrain/Commit-Ai-Go/internal/ai"
	"github.com/NeelFrostrain/Commit-Ai-Go/internal/config"
	"github.com/NeelFrostrain/Commit-Ai-Go/internal/git"

	"github.com/algolyzer/groq-go"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
<<<<<<< Updated upstream
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
=======
	magenta = color.New(color.FgMagenta, color.Bold).SprintFunc()
	blue    = color.New(color.FgBlue).SprintFunc()
	red     = color.New(color.FgRed, color.Bold).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
	green   = color.New(color.FgGreen, color.Bold).SprintFunc()
)

var (
	commitFlag  bool
	yesFlag     bool
	verboseFlag bool
	modelFlag   string
)

var rootCmd = &cobra.Command{
	Use:   "commit-ai",
	Short: "AI-powered semantic commit generator",
	Long: `Commit-AI analyzes your git changes and generates professional,
conventional commit messages using AI. It helps maintain clean
git history with minimal effort.`,
	Run: runCommitAI,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
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
		cfg = &config.Config{APIKey: apiKey, Model: "llama-3.1-8b-instant", Temperature: 0.6, MaxTokens: 12000}
	}

	// Override model if flag is set
	if modelFlag != "" {
		cfg.Model = modelFlag
	}

	// Check for staged changes
	hasStaged, err := git.HasStagedChanges()
	if err != nil {
		fmt.Printf("%s Failed to check staged changes: %v\n", red("[Error]"), err)
		return
	}

	if !hasStaged {
		fmt.Printf("%s No staged changes detected.\n", yellow("[!]"))
		if !yesFlag {
			var confirm bool
			survey.AskOne(&survey.Confirm{Message: "Stage all files and continue?"}, &confirm)
			if !confirm {
				return
			}
		}

		if err := git.StageAllFiles(); err != nil {
			fmt.Printf("%s Failed to stage files: %v\n", red("[Error]"), err)
			return
		}
		fmt.Printf("%s All files staged\n", green("[✓]"))
	}

	// Get staged files for scope suggestion
	stagedFiles, _ := git.GetStagedFiles()
	suggestedScope := ai.SuggestScope(stagedFiles)

	if verboseFlag && suggestedScope != "" {
		fmt.Printf("%s Detected scope: %s\n", blue("[Info]"), suggestedScope)
	}

	// Get optimized diff
	exclude := git.GetIgnorePatterns()
	fullContext, err := git.GetStagedDiff(exclude, 15000) // Increased for comprehensive analysis
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
	options, aiReport := ai.ParseMultiResponse(resp.Choices[0].Message.Content)

	if verboseFlag {
		fmt.Printf("%s Generated %d options\n", blue("[Info]"), len(options))
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
>>>>>>> Stashed changes
		}
	} else {
		fmt.Printf("\n%s Use -c flag to commit automatically\n", yellow("[Hint]"))
	}
}
<<<<<<< Updated upstream

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
=======
>>>>>>> Stashed changes

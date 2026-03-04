package ai

import (
	"fmt"
	"strings"
)

// BuildPrompt creates an enhanced prompt with better context and instructions
func BuildPrompt(diff string, useEmojis bool) string {
	emojiInstructions := ""

	if useEmojis {
		emojiInstructions = `
EMOJI USAGE:
Add relevant emojis to make commits more visual and scannable:
- ✨ feat: New features
- 🐛 fix: Bug fixes
- 📝 docs: Documentation
- ♻️ refactor: Code refactoring
- ⚡ perf: Performance improvements
- 💄 style: UI/styling changes
- ✅ test: Tests
- 🔧 chore: Maintenance
- 🏗️ build: Build system
- 👷 ci: CI/CD changes
- 🔒 security: Security fixes
- 🌐 i18n: Internationalization
- ♿ a11y: Accessibility
- 🎨 design: Design improvements

Use emojis in:
1. Commit title: "✨ feat(api): add user authentication"
2. Category headers: "✨ FEATURES:", "🐛 BUG FIXES:", "📚 DOCUMENTATION:"
3. Bullet points for visual hierarchy`
	}

	return fmt.Sprintf(`You are a Principal Engineer analyzing git changes. Generate 3 professional commit message titles and a comprehensive technical report.

CRITICAL INSTRUCTIONS:
1. Analyze ACTUAL code changes in the diff - be SPECIFIC
2. Use exact XML-style tags: <options> and <report>
3. DO NOT invent features not in the diff
4. Include specific file names, function names, metrics
5. Provide detailed explanations with technical context

CHANGE TYPES:
- feat: New functionality (new functions, endpoints)
- fix: Bug fixes (fixing logic, errors)
- refactor: Code restructuring (renaming, reorganizing)
- perf: Performance improvements (optimization)
- style: Formatting/whitespace
- docs: Documentation only
- test: Tests
- chore: Build, dependencies, tooling
- build: Build system changes
- ci: CI/CD changes

OUTPUT FORMAT (STRICT):
<options>
1. type(scope): description of ACTUAL change
2. type(scope): alternative description
3. type(scope): different perspective
</options>
<report>
[CATEGORY]:
- Specific change with technical details FROM DIFF
- Implementation specifics with file/function names
- Impact or reasoning based on code

TECHNICAL DETAILS:
- Files changed: X files, Y insertions(+), Z deletions(-)
- Specific files: [list from diff]
- Key changes: [specific functions/variables from diff]
- Metrics: [measurements from diff]

IMPACT:
- Performance improvements (with metrics if evident)
- User/developer experience enhancements
- Code quality improvements
- Security/scalability improvements
</report>

IMPORTANT:
- Only include "BREAKING CHANGES:" if actual breaking changes exist
- Use categories: FEATURES, BUG FIXES, IMPROVEMENTS, ARCHITECTURE, etc.
- BE SPECIFIC about names and metrics from the diff
- EXPLAIN WHY changes matter, not just WHAT changed
- Include file counts and line statistics

EXAMPLE:
<options>
1. fix(parser): improve AI prompt accuracy and diff handling
2. fix(git): resolve diff retrieval and binary file filtering
3. refactor(ai): enhance prompt engineering for better analysis
</options>
<report>
IMPROVEMENTS:
- Enhanced AI prompt with critical instructions for accurate analysis
- Added warnings against generic/invented features
- Improved prompt clarity with 'FROM THE DIFF' annotations

BUG FIXES:
- Fixed GetStagedDiff to properly handle exclude patterns
- Corrected git diff command arguments to avoid filtering text files
- Fixed BuildPrompt test to pass emoji flag parameter

TECHNICAL DETAILS:
- 3 files changed: 122 insertions(+), 42 deletions(-)
- Files: internal/ai/parser.go, internal/ai/parser_test.go, internal/git/diff.go
- Improved diff retrieval from 17 to 6596+ characters
- Enhanced validation for emoji-prefixed commits

IMPACT:
- AI generates accurate commit messages from actual code changes
- Diff analysis no longer loses content
- Better accuracy and relevance
- Improved developer experience
</report>
%s
NOW ANALYZE THIS DIFF AND DESCRIBE ONLY WHAT YOU SEE:
%s`, emojiInstructions, diff)
}

// ParseResponse extracts title and report from AI response (legacy single-option format)
func ParseResponse(response string) (string, string) {
	title := "chore: update"
	report := ""

	// Extract message title
	msgStart := strings.Index(response, "<message>")
	msgEnd := strings.Index(response, "</message>")
	if msgStart != -1 && msgEnd != -1 {
		title = strings.TrimSpace(response[msgStart+9 : msgEnd])
	}

	// Extract report
	reportStart := strings.Index(response, "<report>")
	reportEnd := strings.Index(response, "</report>")
	if reportStart != -1 && reportEnd != -1 {
		report = strings.TrimSpace(response[reportStart+8 : reportEnd])
	}

	// Fallback: if no tags found, use entire response as report
	if report == "" && len(strings.TrimSpace(response)) > 0 {
		report = strings.TrimSpace(response)
		// If response has both message and report but no tags, try to split them
		if strings.Contains(response, "\n") {
			parts := strings.SplitN(response, "\n", 2)
			if len(parts) == 2 {
				title = strings.TrimSpace(parts[0])
				report = strings.TrimSpace(parts[1])
			}
		}
	}

	if report == "" {
		report = "Changes analyzed and staged for commit."
	}

	return title, report
}

// ParseMultiResponse extracts multiple commit options and report from AI response
func ParseMultiResponse(input string) ([]string, string) {
	report := "Changes analyzed."
	if start, end := strings.Index(input, "<report>"), strings.Index(input, "</report>"); start != -1 && end > start {
		report = strings.TrimSpace(input[start+8 : end])
	}

	var options []string
	if start, end := strings.Index(input, "<options>"), strings.Index(input, "</options>"); start != -1 && end > start {
		raw := strings.TrimSpace(input[start+9 : end])
		for _, line := range strings.Split(raw, "\n") {
			clean := strings.TrimSpace(line)
			if clean == "" {
				continue
			}

			// Smart Cleaning: Remove list numbers, quotes, and asterisks
			clean = strings.Map(func(r rune) rune {
				if r == '*' || r == '"' || r == '`' {
					return -1
				}
				return r
			}, clean)

			// Remove leading "1. " or "2. "
			if idx := strings.Index(clean, ". "); idx != -1 && idx < 4 {
				clean = clean[idx+2:]
			}

			if clean != "" {
				options = append(options, strings.TrimSpace(clean))
			}
		}
	}

	if len(options) == 0 {
		options = []string{"chore: update project"}
	}
	return options, report
}

// ValidateCommitMessage checks if a commit message follows conventional commit format
func ValidateCommitMessage(msg string) bool {
	validTypes := []string{"feat", "fix", "docs", "style", "refactor", "perf", "test", "chore", "build", "ci", "revert"}

	// Remove leading emojis and whitespace for validation
	cleaned := strings.TrimSpace(msg)
	// Skip emoji characters (they're typically 1-4 bytes in UTF-8)
	for len(cleaned) > 0 {
		r := []rune(cleaned)[0]
		// Check if first character is an emoji (Unicode range)
		if r > 127 {
			cleaned = string([]rune(cleaned)[1:])
			cleaned = strings.TrimSpace(cleaned)
		} else {
			break
		}
	}

	// Check for type at the beginning
	for _, t := range validTypes {
		if strings.HasPrefix(cleaned, t+":") || strings.HasPrefix(cleaned, t+"(") {
			return true
		}
	}
	return false
}

// SuggestScope analyzes file paths to suggest a scope
func SuggestScope(files []string) string {
	if len(files) == 0 {
		return ""
	}

	// Count directory occurrences
	dirCount := make(map[string]int)
	for _, file := range files {
		parts := strings.Split(file, "/")
		if len(parts) > 1 {
			dirCount[parts[0]]++
		}
	}

	// Find most common directory
	maxCount := 0
	scope := ""
	for dir, count := range dirCount {
		if count > maxCount {
			maxCount = count
			scope = dir
		}
	}

	return scope
}

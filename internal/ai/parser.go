package ai

import (
	"fmt"
	"strings"
)

// BuildPrompt creates an enhanced prompt with better context and instructions
func BuildPrompt(diff string, useEmojis bool) string {
	emojiInstructions := ""
	emojiExamples := ""

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

		emojiExamples = `
EXAMPLE WITH EMOJIS:
<options>
1. ✨ feat(api): add user authentication endpoints
2. 🔒 feat(auth): implement JWT-based login system
3. 🔧 chore(security): add authentication middleware
</options>
<report>
✨ FEATURES:
- 🔐 Added JWT authentication with RS256 signing
- 🚦 Implemented login endpoint with rate limiting (5 attempts/min)
- 🛡️ Created middleware for protected routes with role-based access

🔧 TECHNICAL DETAILS:
- 📊 5 files changed: 250 insertions, 10 deletions
- ✅ Test coverage: 90%%
- ⚡ Token validation: <1ms average

💡 IMPACT:
- 🔒 Improved security with industry-standard JWT
- 🚀 Better user experience with automatic token refresh
- 📉 Reduced server load with stateless authentication
</report>`
	}

	return fmt.Sprintf(`You are a Principal Engineer analyzing git changes. Generate 3 professional commit message titles and a comprehensive technical report.

CRITICAL INSTRUCTIONS:
1. You MUST analyze the ACTUAL code changes in the diff below
2. You MUST use the exact XML-style tags shown below
3. DO NOT make up features or changes that are not in the diff
4. DO NOT use generic examples - analyze the REAL changes
5. Be SPECIFIC about what files changed and what code was added/modified/deleted

CRITICAL: You MUST use the exact XML-style tags shown below. Do not use markdown, do not use any other format.
%s
ANALYSIS REQUIREMENTS:
1. READ THE DIFF CAREFULLY - identify what actually changed
2. Identify PRIMARY change type and scope from file paths and code changes
3. Analyze WHAT changed (specific functions, variables, logic)
4. Explain WHY it matters based on the code context
5. Use imperative mood ("add" not "added")
6. Keep titles under 72 characters
7. Provide detailed, structured report based on ACTUAL changes

CHANGE TYPES:
- feat: New functionality or capability (new functions, endpoints, features)
- fix: Bug fixes or error corrections (fixing logic, handling errors)
- refactor: Code restructuring without behavior change (renaming, reorganizing)
- perf: Performance improvements (optimization, caching)
- style: Formatting, whitespace (code style only)
- docs: Documentation only (comments, README, docs)
- test: Adding/updating tests (test files, test cases)
- chore: Build, dependencies, tooling (package.json, go.mod, config)
- build: Build system changes (Makefile, build scripts)
- ci: CI/CD changes (GitHub Actions, CI config)

OUTPUT FORMAT (STRICT - USE THESE EXACT TAGS):
<options>
1. type(scope): concise description of ACTUAL change
2. type(scope): alternative description of ACTUAL change
3. type(scope): different perspective on ACTUAL change
</options>
<report>
[CATEGORY NAME]:
- Specific change with technical details FROM THE DIFF
- Another change with context FROM THE DIFF
- Impact or reasoning based on ACTUAL code

[ANOTHER CATEGORY]:
- Detailed technical change FROM THE DIFF
- Implementation specifics FROM THE DIFF
- Benefits or improvements based on ACTUAL changes

TECHNICAL DETAILS:
- Files changed statistics (count from diff)
- Key metrics or measurements FROM THE DIFF
- Test results if applicable FROM THE DIFF

IMPACT:
- Performance improvements (if evident from code)
- User experience enhancements (if evident from code)
- Developer experience improvements (if evident from code)
</report>

IMPORTANT NOTES:
- Only include "BREAKING CHANGES:" section if there are actual breaking changes in the diff
- If no breaking changes, do NOT include that section at all
- Use clear category names like FEATURES, BUG FIXES, IMPROVEMENTS, etc.
- Keep categories relevant to the ACTUAL changes in the diff
- DO NOT invent features or changes that are not in the diff
- BE SPECIFIC about function names, variable names, file names from the diff

EXAMPLE OUTPUT (NO BREAKING CHANGES):
<options>
1. fix(cache): resolve memory leak in cleanup method
2. fix(memory): prevent memory retention in cache
3. chore(cache): improve cleanup implementation
</options>
<report>
BUG FIXES:
- Fixed memory leak in cache cleanup method
- Changed cleanup to delete items individually
- Added proper resource cleanup in defer statements

TECHNICAL DETAILS:
- 1 file changed: 15 insertions, 8 deletions
- Memory usage reduced by 40%% in tests

IMPACT:
- Prevents memory leak in production
- Improves application stability
</report>
%s
NOW ANALYZE THIS DIFF CAREFULLY AND DESCRIBE ONLY WHAT YOU SEE:
%s`, emojiInstructions, emojiExamples, diff)
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

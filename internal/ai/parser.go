package ai

import (
	"fmt"
	"strings"
)

// BuildPrompt creates an enhanced prompt with better context and instructions
func BuildPrompt(diff string) string {
	return fmt.Sprintf(`You are a Principal Engineer analyzing git changes. Generate 3 professional commit message titles and a comprehensive technical report.

CRITICAL: You MUST use the exact XML-style tags shown below. Do not use markdown, do not use any other format.

ANALYSIS REQUIREMENTS:
1. Identify PRIMARY change type and scope from file paths
2. Analyze WHAT changed, WHY it matters, and the IMPACT
3. Use imperative mood ("add" not "added")
4. Keep titles under 72 characters
5. Provide detailed, structured report

CHANGE TYPES:
- feat: New functionality or capability
- fix: Bug fixes or error corrections  
- refactor: Code restructuring without behavior change
- perf: Performance improvements
- style: Formatting, whitespace
- docs: Documentation only
- test: Adding/updating tests
- chore: Build, dependencies, tooling
- build: Build system changes
- ci: CI/CD changes

OUTPUT FORMAT (STRICT - USE THESE EXACT TAGS):
<options>
1. type(scope): concise description
2. type(scope): alternative description  
3. type(scope): different perspective
</options>
<report>
[CATEGORY NAME]:
- Specific change with technical details
- Another change with context
- Impact or reasoning

[ANOTHER CATEGORY]:
- Detailed technical change
- Implementation specifics
- Benefits or improvements

TECHNICAL DETAILS:
- Files changed statistics
- Key metrics or measurements
- Test results if applicable

IMPACT:
- Performance improvements
- User experience enhancements
- Developer experience improvements
</report>

IMPORTANT NOTES:
- Only include "BREAKING CHANGES:" section if there are actual breaking changes
- If no breaking changes, do NOT include that section at all
- Use clear category names like FEATURES, BUG FIXES, IMPROVEMENTS, etc.
- Keep categories relevant to the actual changes

EXAMPLE OUTPUT (WITH BREAKING CHANGES):
<options>
1. feat(api): add user authentication endpoints
2. feat(auth): implement JWT-based login system
3. chore(security): add authentication middleware
</options>
<report>
BREAKING CHANGES:
- API endpoint /login now requires POST instead of GET
- Authentication token format changed from Bearer to JWT

FEATURES:
- Added JWT authentication with RS256 signing
- Implemented login endpoint with rate limiting
- Created middleware for protected routes

TECHNICAL DETAILS:
- 5 files changed: 250 insertions, 10 deletions
- Test coverage: 90%%

IMPACT:
- Improved security with industry-standard JWT
- Better user experience with automatic token refresh
</report>

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

NOW ANALYZE THIS DIFF:
%s`, diff)
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

	// Check for type at the beginning
	for _, t := range validTypes {
		if strings.HasPrefix(msg, t+":") || strings.HasPrefix(msg, t+"(") {
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

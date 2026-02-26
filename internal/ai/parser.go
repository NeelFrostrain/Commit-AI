package ai

import (
	"fmt"
	"strings"
)

func BuildPrompt(diff string) string {
    return fmt.Sprintf(`Analyze this git diff and provide exactly 3 distinct semantic commit message options.
Each option MUST follow the Conventional Commits format: <type>(optional scope): <description>

Change types to consider:
- feat: A new feature
- fix: A bug fix
- refactor: A code change that neither fixes a bug nor adds a feature
- perf: A code change that improves performance
- style: Changes that do not affect the meaning of the code (white-space, formatting, etc)
- docs: Documentation only changes
- chore: Updating build tasks, package manager configs, etc.
- test: Adding missing tests or correcting existing tests

STRICT FORMAT:
<options>
1. type: description
2. type: description
3. type: description
</options>
<report>
* technical change bullet points
</report>

Diff:
%s`, diff)
}

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
			if clean == "" { continue }

			// Smart Cleaning: Remove list numbers, quotes, and asterisks
			clean = strings.Map(func(r rune) rune {
				if r == '*' || r == '"' || r == '`' { return -1 }
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

	if len(options) == 0 { options = []string{"chore: update project"} }
	return options, report
}
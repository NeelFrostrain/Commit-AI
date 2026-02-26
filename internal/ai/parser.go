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
	// 1. Extract the Report
	report := ""
	reportStart := strings.Index(input, "<report>")
	reportEnd := strings.Index(input, "</report>")
	if reportStart != -1 && reportEnd != -1 {
		report = strings.TrimSpace(input[reportStart+8 : reportEnd])
	}

	// 2. Extract Options
	options := []string{}
	optsStart := strings.Index(input, "<options>")
	optsEnd := strings.Index(input, "</options>")
	
	if optsStart != -1 && optsEnd != -1 {
		rawOptions := strings.TrimSpace(input[optsStart+9 : optsEnd])
		lines := strings.Split(rawOptions, "\n")
		for _, line := range lines {
			// Remove numbering like "1. ", "2. ", etc.
			cleanLine := line
			if len(line) > 3 && (line[1] == '.' || line[2] == '.') {
				parts := strings.SplitN(line, " ", 2)
				if len(parts) > 1 {
					cleanLine = parts[1]
				}
			}
			
			// Clean up "type: " prefix if AI adds it literally
			cleanLine = strings.TrimPrefix(cleanLine, "type: ")
			cleanLine = strings.TrimSpace(cleanLine)
			
			if cleanLine != "" {
				options = append(options, cleanLine)
			}
		}
	}

	// Ensure we have 3 options; if AI fails, provide fallbacks
	for len(options) < 3 {
		options = append(options, fmt.Sprintf("chore: update files (fallback %d)", len(options)+1))
	}

	return options, report
}
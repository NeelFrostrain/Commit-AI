package ai

import (
	"fmt"
	"strings"
)

func ParseResponse(input string) (string, string) {
	extract := func(tag string) string {
		startTag := "<" + tag + ">"
		endTag := "</" + tag + ">"
		start := strings.Index(input, startTag)
		end := strings.Index(input, endTag)
		if start == -1 || end == -1 {
			return ""
		}
		return strings.TrimSpace(input[start+len(startTag) : end])
	}

	report := extract("report")
	title := extract("message")

	if title == "" {
		lines := strings.Split(input, "\n")
		for _, line := range lines {
			upperLine := strings.ToUpper(line)
			if strings.Contains(upperLine, "COMMIT_MESSAGE:") {
				parts := strings.SplitN(line, ":", 2)
				if len(parts) > 1 {
					title = strings.TrimSpace(parts[1])
				}
				break
			}
		}
	}

	title = strings.ReplaceAll(title, "**", "")
	title = strings.TrimPrefix(title, "type: ")
	title = strings.TrimPrefix(title, "summary: ")
	title = strings.TrimSuffix(title, ".")
	title = strings.TrimSpace(title)

	if !strings.Contains(title, ":") {
		title = "feat: " + title
	}

	parts := strings.SplitN(title, ":", 2)
	if len(parts) > 1 {
		typePart := strings.TrimSpace(parts[0])
		descPart := strings.TrimSpace(parts[1])
		if len(descPart) > 0 {
			descPart = strings.ToLower(string(descPart[0])) + descPart[1:]
		}
		title = fmt.Sprintf("%s: %s", typePart, descPart)
	}

	if report == "" {
		report = "Modified project files."
	}

	return title, report
}

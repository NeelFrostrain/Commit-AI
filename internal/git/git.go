package git

import (
	"fmt"
	"os"
	"strings"
)

func GetIgnorePatterns() []string {
	// Base patterns for safety
	patterns := []string{
		"package-lock.json", "bun.lockb", "yarn.lock", "pnpm-lock.yaml",
		"node_modules", "dist", "vendor", ".git", "commit-ai.exe",
		"*.log", "*.exe", "*.bin", "*.pyc", ".DS_Store",
	}

	data, err := os.ReadFile(".gitignore")
	if err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)

			// 1. Skip empty lines and comments
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}

			// 2. Safety: Remove leading slashes
			// Git pathspec :(exclude)/Library/ often fails; :(exclude)Library/ works better.
			line = strings.TrimPrefix(line, "/")

			// 3. Safety: Remove trailing slashes
			line = strings.TrimSuffix(line, "/")

			patterns = append(patterns, line)
		}
	}

	// Remove duplicates (optional but recommended)
	patterns = unique(patterns)

	var excludeArgs []string
	for _, p := range patterns {
		// Use the "magic" pathspec syntax
		// We wrap in quotes in the actual EXEC call later,
		// but here we just format the string.
		excludeArgs = append(excludeArgs, fmt.Sprintf(":(exclude)%s", p))
	}
	return excludeArgs
}

// Helper to prevent redundant arguments
func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

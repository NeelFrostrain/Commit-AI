package git

import (
	"os"
	"path/filepath"
	"strings"
)

func GetIgnorePatterns() []string {
	// Defaults to skip heavy/auto-generated files
	patterns := []string{
		"*.lock", "*-lock.json", "*-lock.yaml", "package-lock.json", "yarn.lock", "pnpm-lock.yaml",
		"vendor/*", "dist/*", ".git/*", "node_modules/*", "target/*", "build/*", "out/*",
		"*.exe", "*.dll", "*.so", "*.a", "*.o", "*.out",
	}

	if data, err := os.ReadFile(".gitignore"); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" && !strings.HasPrefix(line, "#") {
				// Convert Windows backslashes to Git-friendly forward slashes
				patterns = append(patterns, filepath.ToSlash(line))
			}
		}
	}
	
	var excludeArgs []string
	for _, p := range patterns {
		excludeArgs = append(excludeArgs, ":(exclude)"+p)
	}
	return excludeArgs
}
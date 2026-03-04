package git

import (
	"fmt"
	"os/exec"
	"strings"
)

// DiffStats represents git diff statistics
type DiffStats struct {
	FilesChanged int
	Insertions   int
	Deletions    int
	Files        []string
}

// GetStagedDiff retrieves the staged diff with optimizations
func GetStagedDiff(excludePatterns []string, maxSize int) (string, error) {
	// Get summary stats
	diffArgs := append([]string{"diff", "--cached", "--stat", "--"}, excludePatterns...)
	stats, err := exec.Command("git", diffArgs...).Output()
	if err != nil {
		return "", fmt.Errorf("failed to get diff stats: %w", err)
	}

	// Get actual code changes with reduced context
	diffArgs = append([]string{"diff", "--cached", "-U2", "--"}, excludePatterns...)
	diff, err := exec.Command("git", diffArgs...).Output()
	if err != nil {
		return "", fmt.Errorf("failed to get diff: %w", err)
	}

	fullContext := fmt.Sprintf("Summary:\n%s\n\nDiff:\n%s", stats, diff)

	// Truncate if too large
	if len(fullContext) > maxSize {
		fullContext = fullContext[:maxSize] + "\n...[diff truncated for token limit]"
	}

	return fullContext, nil
}

// GetStagedFiles returns list of staged file names
func GetStagedFiles() ([]string, error) {
	out, err := exec.Command("git", "diff", "--cached", "--name-only").Output()
	if err != nil {
		return nil, err
	}

	files := strings.Split(strings.TrimSpace(string(out)), "\n")
	var result []string
	for _, f := range files {
		if f != "" {
			result = append(result, f)
		}
	}
	return result, nil
}

// HasStagedChanges checks if there are any staged changes
func HasStagedChanges() (bool, error) {
	out, err := exec.Command("git", "diff", "--cached", "--name-only").Output()
	if err != nil {
		return false, err
	}
	return len(strings.TrimSpace(string(out))) > 0, nil
}

// HasUnstagedChanges checks if there are any unstaged changes
func HasUnstagedChanges() (bool, error) {
	out, err := exec.Command("git", "diff", "--name-only").Output()
	if err != nil {
		return false, err
	}

	// Also check for untracked files
	untracked, err := exec.Command("git", "ls-files", "--others", "--exclude-standard").Output()
	if err != nil {
		return false, err
	}

	return len(strings.TrimSpace(string(out))) > 0 || len(strings.TrimSpace(string(untracked))) > 0, nil
}

// StageAllFiles stages all changes
func StageAllFiles() error {
	return exec.Command("git", "add", ".").Run()
}

// ParseDiffStats extracts statistics from git diff output
func ParseDiffStats(diffOutput string) DiffStats {
	stats := DiffStats{}
	lines := strings.Split(diffOutput, "\n")

	for _, line := range lines {
		// Parse summary line: "3 files changed, 45 insertions(+), 12 deletions(-)"
		if strings.Contains(line, "file") && strings.Contains(line, "changed") {
			fmt.Sscanf(line, "%d file", &stats.FilesChanged)
			if strings.Contains(line, "insertion") {
				fmt.Sscanf(line, "%d file%*s %d insertion", &stats.FilesChanged, &stats.Insertions)
			}
			if strings.Contains(line, "deletion") {
				fmt.Sscanf(line, "%*s %d deletion", &stats.Deletions)
			}
		}
	}

	return stats
}

package ai

import (
	"testing"
)

func TestValidateCommitMessage(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		expected bool
	}{
		{"valid feat", "feat: add new feature", true},
		{"valid fix with scope", "fix(api): resolve bug", true},
		{"valid refactor", "refactor(auth): improve logic", true},
		{"invalid no type", "add new feature", false},
		{"invalid wrong type", "added: new feature", false},
		{"valid chore", "chore: update dependencies", true},
		{"valid docs", "docs(readme): update installation", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateCommitMessage(tt.message)
			if result != tt.expected {
				t.Errorf("ValidateCommitMessage(%q) = %v, want %v", tt.message, result, tt.expected)
			}
		})
	}
}

func TestSuggestScope(t *testing.T) {
	tests := []struct {
		name     string
		files    []string
		expected string
	}{
		{
			name:     "api files",
			files:    []string{"api/handler.go", "api/routes.go", "api/middleware.go"},
			expected: "api",
		},
		{
			name:     "mixed files",
			files:    []string{"ui/component.tsx", "api/handler.go", "ui/styles.css"},
			expected: "ui",
		},
		{
			name:     "root files",
			files:    []string{"main.go", "README.md"},
			expected: "",
		},
		{
			name:     "empty",
			files:    []string{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SuggestScope(tt.files)
			if result != tt.expected {
				t.Errorf("SuggestScope(%v) = %q, want %q", tt.files, result, tt.expected)
			}
		})
	}
}

func TestParseMultiResponse(t *testing.T) {
	input := `<options>
1. feat(api): add user authentication
2. feat(auth): implement login system
3. chore(security): add auth middleware
</options>
<report>
- Added JWT authentication
- Implemented login endpoint
- Added middleware for protected routes
</report>`

	options, report := ParseMultiResponse(input)

	if len(options) != 3 {
		t.Errorf("Expected 3 options, got %d", len(options))
	}

	expectedOptions := []string{
		"feat(api): add user authentication",
		"feat(auth): implement login system",
		"chore(security): add auth middleware",
	}

	for i, expected := range expectedOptions {
		if i < len(options) && options[i] != expected {
			t.Errorf("Option %d: got %q, want %q", i, options[i], expected)
		}
	}

	if report == "" {
		t.Error("Expected non-empty report")
	}
}

func TestParseMultiResponseWithInvalidFormat(t *testing.T) {
	input := "Some random text without proper tags"

	options, report := ParseMultiResponse(input)

	if len(options) == 0 {
		t.Error("Expected fallback option")
	}

	if report == "" {
		t.Error("Expected fallback report")
	}
}

func TestBuildPrompt(t *testing.T) {
	diff := "diff --git a/file.go b/file.go\n+func newFunction() {}"

	prompt := BuildPrompt(diff, false)

	// Check that prompt contains key instructions
	requiredPhrases := []string{
		"commit",
		"feat",
		"fix",
		"<options>",
		"<report>",
	}

	for _, phrase := range requiredPhrases {
		if !contains(prompt, phrase) {
			t.Errorf("Prompt missing required phrase: %q", phrase)
		}
	}

	// Check that diff is included
	if !contains(prompt, diff) {
		t.Error("Prompt does not contain the diff")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

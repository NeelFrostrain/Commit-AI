package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config holds application configuration
type Config struct {
	APIKey      string
	Model       string
	Temperature float64
	MaxTokens   int
}

// Load loads configuration from environment and files
func Load() (*Config, error) {
	// Try loading from .env in current directory
	_ = godotenv.Load()

	// Try loading from user home directory
	home, err := os.UserHomeDir()
	if err == nil {
		_ = godotenv.Load(filepath.Join(home, ".commit-ai.env"))
	}

	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GROQ_API_KEY not found in environment")
	}

	model := os.Getenv("COMMIT_AI_MODEL")
	if model == "" {
		model = "llama-3.1-8b-instant" // Default model
	}

	return &Config{
		APIKey:      apiKey,
		Model:       model,
		Temperature: 0.7,   // Higher for more detailed, creative responses
		MaxTokens:   15000, // Increased for comprehensive reports
	}, nil
}

// SaveAPIKey saves the API key to user's home directory
func SaveAPIKey(apiKey string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	envPath := filepath.Join(home, ".commit-ai.env")
	content := fmt.Sprintf("GROQ_API_KEY=%s\n", apiKey)

	return os.WriteFile(envPath, []byte(content), 0600)
}

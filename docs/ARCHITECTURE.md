# Architecture Guide

Understanding Commit-AI's design and structure.

## Project Structure

```
Commit-AI/
├── main.go                 # Entry point
├── cmd/
│   └── root.go            # CLI command handler
├── internal/
│   ├── ai/
│   │   ├── parser.go      # AI response parsing
│   │   └── parser_test.go # AI tests
│   ├── config/
│   │   └── config.go      # Configuration management
│   ├── git/
│   │   ├── git.go         # Git utilities
│   │   └── diff.go        # Diff operations
│   └── updater/
│       └── updater.go     # Auto-update functionality
├── installer/
│   └── main.go            # Installer
├── docs/                  # Documentation
├── go.mod                 # Dependencies
├── Makefile               # Build automation
├── build.ps1              # Windows build script
└── .gitattributes         # Binary file marking
```

## Core Modules

### 1. Main Entry Point (`main.go`)

```go
func main() {
    // Set version info from build flags
    cmd.SetVersion(version, buildDate, gitCommit)
    
    // Execute CLI command
    if err := cmd.Execute(); err != nil {
        os.Exit(1)
    }
}
```

**Responsibilities:**
- Application entry point
- Version information setup
- Error handling

---

### 2. CLI Handler (`cmd/root.go`)

**Main Function:** `runCommitAI()`

**Workflow:**
1. Load configuration
2. Check for staged/unstaged changes
3. Offer to stage all changes if needed
4. Get staged files for scope detection
5. Retrieve staged diff with binary filtering
6. Send to Groq AI with enhanced prompt
7. Parse AI response (3 options + report)
8. Interactive selection
9. Execute git commit or display

**Key Features:**
- Interactive prompts using `survey` library
- Color-coded output using `fatih/color`
- Verbose mode for debugging
- Emoji support
- Model override capability

---

### 3. AI Module (`internal/ai/parser.go`)

**Key Functions:**

#### `BuildPrompt(diff string, useEmojis bool) string`
- Constructs enhanced AI prompt
- Includes critical instructions
- Adds emoji guidance if enabled
- Returns formatted prompt for Groq API

#### `ParseMultiResponse(input string) ([]string, string)`
- Extracts 3 commit options from AI response
- Extracts detailed report
- Handles invalid formats with fallback
- Returns options and report

#### `ValidateCommitMessage(msg string) bool`
- Validates Conventional Commits format
- Strips leading emojis for validation
- Checks for valid commit types
- Returns validation result

#### `SuggestScope(files []string) string`
- Analyzes file paths
- Detects most common directory
- Returns suggested scope
- Used for auto-detection

---

### 4. Configuration Module (`internal/config/config.go`)

**Key Functions:**

#### `Load() (*Config, error)`
- Loads from `.env` or `~/.commit-ai.env`
- Returns Config struct with:
  - APIKey
  - Model (default: `llama-3.1-8b-instant`)
  - Temperature (default: 0.7)
  - MaxTokens (default: 15000)

#### `SaveAPIKey(key string) error`
- Saves API key to `~/.commit-ai.env`
- Creates file if not exists
- Sets secure permissions (600)

---

### 5. Git Module (`internal/git/`)

#### `git.go` - Git Utilities

**Key Functions:**
- `GetIgnorePatterns()` - Parses `.gitignore`
- `GetIgnorePatternsFromFile()` - Reads ignore file

#### `diff.go` - Diff Operations

**Key Functions:**

#### `GetStagedDiff(excludePatterns []string, maxSize int) (string, error)`
- Retrieves staged changes
- Filters binary files (50+ types)
- Applies exclude patterns
- Truncates to maxSize
- Returns formatted diff

#### `GetStagedFiles() ([]string, error)`
- Lists all staged files
- Used for scope detection

#### `HasStagedChanges() (bool, error)`
- Checks for staged changes

#### `HasUnstagedChanges() (bool, error)`
- Checks for unstaged changes
- Includes untracked files

#### `StageAllFiles() error`
- Runs `git add .`
- Stages all changes

#### `isBinaryFile(filename string, patterns []string) bool`
- Checks if file matches binary patterns
- Uses filepath.Match for pattern matching

---

### 6. Updater Module (`internal/updater/updater.go`)

**Key Functions:**

#### `CheckForUpdate(currentVersion string) (*Release, bool, error)`
- Queries GitHub API
- Compares versions
- Returns latest release info

#### `GetAssetForPlatform(release *Release) (*Asset, error)`
- Selects correct binary for platform
- Supports Windows, macOS (Intel/ARM), Linux (amd64/arm64)

#### `DownloadUpdate(asset *Asset, exePath string) error`
- Downloads binary from GitHub
- Saves to temporary location
- Verifies download

#### `InstallUpdate(exePath string) error`
- Backs up current binary
- Replaces with new version
- Handles Windows/Unix differences

---

### 7. Installer (`installer/main.go`)

**Functionality:**
- Downloads latest binary from GitHub
- Creates installation directory
- Updates PATH environment variable
- Platform-specific installation logic

---

## Data Flow

```
User Input
    ↓
CLI Handler (cmd/root.go)
    ↓
Git Operations (internal/git/)
    ├─ Check for changes
    ├─ Stage changes if needed
    └─ Get staged diff
    ↓
AI Module (internal/ai/)
    ├─ Build prompt
    ├─ Send to Groq API
    └─ Parse response
    ↓
Interactive Selection
    ├─ Choose commit message
    ├─ Choose report format
    └─ Review final message
    ↓
Git Commit
    └─ Execute git commit
```

---

## Key Design Decisions

### 1. Modular Architecture
- Separate concerns into dedicated packages
- Clean interfaces between modules
- Easy to test and extend

### 2. Binary File Filtering
- Prevents AI from analyzing compiled code
- Reduces API token usage
- Improves commit message accuracy

### 3. Interactive Prompts
- User-friendly CLI experience
- Multiple options for flexibility
- Edit capability for customization

### 4. Emoji Support
- Optional visual enhancement
- Improves commit history readability
- Configurable via flag

### 5. Auto-Update
- Built-in update mechanism
- Safe update with backup
- Cross-platform support

---

## Dependencies

### External Libraries

| Library | Purpose | Version |
|---------|---------|---------|
| `github.com/spf13/cobra` | CLI framework | Latest |
| `github.com/AlecAivazis/survey/v2` | Interactive prompts | v2.x |
| `github.com/algolyzer/groq-go` | Groq API client | Latest |
| `github.com/fatih/color` | Colored output | Latest |
| `github.com/joho/godotenv` | .env parsing | Latest |

### Standard Library

- `context` - Context management
- `fmt` - Formatting
- `os` - OS operations
- `os/exec` - Command execution
- `strings` - String operations
- `net/http` - HTTP requests
- `encoding/json` - JSON parsing

---

## Configuration Flow

```
Environment Variables
    ↓
.env (project)
    ↓
~/.commit-ai.env (global)
    ↓
Default Values
    ↓
Config Struct
```

---

## Error Handling

### Error Types

1. **Configuration Errors**
   - Missing API key
   - Invalid configuration

2. **Git Errors**
   - Failed to get diff
   - Failed to stage changes
   - Failed to commit

3. **AI Errors**
   - API request failed
   - Invalid response format
   - Rate limit exceeded

4. **System Errors**
   - File I/O errors
   - Permission errors
   - Network errors

### Error Recovery

- Helpful error messages
- Suggestions for resolution
- Graceful degradation
- Fallback mechanisms

---

## Performance Considerations

### Optimization Strategies

1. **Diff Size Limiting**
   - Truncate to 8,000 characters
   - Reduces API token usage
   - Maintains accuracy

2. **Binary File Filtering**
   - Prevents unnecessary analysis
   - Reduces API calls
   - Improves performance

3. **Caching**
   - Configuration caching
   - Scope detection caching
   - Pattern matching caching

4. **Concurrency**
   - Parallel file analysis (future)
   - Async API calls (future)

---

## Security Considerations

### API Key Management

- Stored in `~/.commit-ai.env`
- File permissions: 600 (read/write owner only)
- Never logged or displayed
- Validated before use

### Input Validation

- Commit message validation
- File path validation
- API response validation

### Binary File Handling

- Prevents accidental analysis of sensitive data
- Filters 50+ binary file types
- Extensible pattern matching

---

## Testing Strategy

### Test Coverage

- **Unit Tests**: 85%+ coverage
- **Integration Tests**: Git operations
- **End-to-End Tests**: Full workflow

### Test Files

- `internal/ai/parser_test.go` - AI module tests
- Test cases for validation, parsing, scope detection

---

## Future Architecture Improvements

1. **Plugin System**
   - Custom AI providers
   - Custom commit templates
   - Custom validation rules

2. **Caching Layer**
   - Cache AI responses
   - Cache scope detection
   - Reduce API calls

3. **Async Processing**
   - Parallel file analysis
   - Background updates
   - Non-blocking UI

4. **Multi-Model Support**
   - OpenAI integration
   - Anthropic integration
   - Local model support

5. **IDE Integration**
   - VSCode extension
   - JetBrains plugin
   - Vim/Neovim plugin

---

## Next Steps

- [Build Guide](BUILD.md)
- [Contributing Guide](CONTRIBUTING.md)
- [Testing Guide](TESTING.md)

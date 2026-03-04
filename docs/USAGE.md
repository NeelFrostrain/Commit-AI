# Usage Guide

Complete guide to using Commit-AI.

## Basic Workflow

### 1. Make Changes
```bash
# Edit your files
vim src/api.go
vim src/config.go
```

### 2. Stage Changes
```bash
# Stage specific files
git add src/api.go src/config.go

# Or stage all changes
git add .
```

### 3. Generate Commit Message
```bash
# Generate and review
commit-ai

# Or generate and commit
commit-ai -c
```

### 4. Select and Commit
```bash
# Choose from 3 AI-generated options
# Select report format (keep, edit, skip)
# Review final message
# Confirm commit
```

---

## Command Reference

### Main Command

```bash
commit-ai [flags]
```

### Flags

| Flag | Short | Description | Example |
|------|-------|-------------|---------|
| `--commit` | `-c` | Commit after selection | `commit-ai -c` |
| `--yes` | `-y` | Skip confirmations | `commit-ai -y` |
| `--verbose` | `-v` | Show analysis details | `commit-ai -v` |
| `--emoji` | `-e` | Add emojis | `commit-ai -e` |
| `--model` | `-m` | Override AI model | `commit-ai -m llama-3.1-70b-versatile` |
| `--help` | `-h` | Show help | `commit-ai -h` |

### Subcommands

```bash
# Show version information
commit-ai version

# Check for updates
commit-ai update

# Check for updates without installing
commit-ai update --check

# Force update
commit-ai update --force
```

---

## Common Workflows

### Quick Commit
```bash
# Generate and commit in one step
commit-ai -c
```

### Detailed Analysis
```bash
# See what the AI is analyzing
commit-ai -v

# Output:
# [Info] Detected scope: api
# [Info] Analyzing 3 files
# [Info] Diff size: 2,456 characters
# [Info] Generated 3 options
```

### Visual Commits
```bash
# Add emojis for better readability
commit-ai -e

# Output includes emojis like:
# ✨ feat(api): add user authentication
# 🐛 BUG FIXES:
# 📊 TECHNICAL DETAILS:
```

### Automated Workflow
```bash
# Skip all confirmations
commit-ai -c -y

# Useful for scripts and CI/CD
```

### Combined Flags
```bash
# Emoji + Verbose + Auto-commit
commit-ai -cev

# Emoji + Auto-commit + Skip confirmations
commit-ai -ecy
```

---

## Interactive Prompts

### Stage Changes Prompt
```
? Stage all changes (git add .)?
> Yes
  No
```

Choose "Yes" to automatically stage all changes, or "No" to proceed with only staged changes.

### Commit Title Selection
```
? Select commit title:
> feat(api): add user authentication
  feat(auth): implement JWT-based login
  chore(security): add authentication middleware
  ✏️  Edit manually...
  ❌ Cancel
```

Choose from AI suggestions or edit manually.

### Report Format Selection
```
? Commit body:
> 📝 Keep AI report
  ✏️  Edit report
  ⊘  No report
```

- **Keep AI report**: Use AI-generated report as-is
- **Edit report**: Modify the report before committing
- **No report**: Commit with title only

---

## Output Examples

### Verbose Mode Output
```
[Info] Detected scope: api
[Info] Analyzing 3 files
[Info] Diff size: 2,456 characters
[Commit-AI] Analyzing changes with llama-3.1-8b-instant...
[Info] Generated 3 options
```

### Final Commit Display
```
─────────────────────────────────
Title: feat(api): add user authentication
Body:
FEATURES:
- Implemented JWT-based authentication
- Added login endpoint with rate limiting
- Created middleware for protected routes

TECHNICAL DETAILS:
- 3 files changed: 150 insertions(+), 20 deletions(-)
- Test coverage: 95%

IMPACT:
- Improved security with industry-standard JWT
- Better user experience with automatic token refresh
─────────────────────────────────
```

---

## Tips and Tricks

### 1. Stage Related Changes Only
```bash
# Better commit messages when changes are grouped
git add src/auth.go src/middleware.go
commit-ai
```

### 2. Use Verbose Mode During Development
```bash
# See what the AI is analyzing
commit-ai -v

# Helps understand scope detection and diff analysis
```

### 3. Edit Reports for Clarity
```bash
# Generate commit message
commit-ai

# Select "✏️  Edit report"
# Customize the report for your team
```

### 4. Combine with Git Aliases
```bash
# Add to ~/.gitconfig
[alias]
    ai = "!commit-ai -c"
    ai-verbose = "!commit-ai -cv"
    ai-emoji = "!commit-ai -ce"

# Usage
git ai
git ai-verbose
git ai-emoji
```

### 5. Use in Scripts
```bash
#!/bin/bash
# Auto-commit script
git add .
commit-ai -c -y
```

### 6. Pre-commit Hook
```bash
#!/bin/bash
# .git/hooks/pre-commit
# Validate commit message format
commit-ai -v
```

---

## Troubleshooting

### "No changes detected"
```bash
# Check git status
git status

# Stage changes
git add .

# Try again
commit-ai
```

### "AI request failed"
```bash
# Check internet connection
ping google.com

# Verify API key
echo $GROQ_API_KEY

# Try again
commit-ai
```

### "Rate limit exceeded"
```bash
# Wait a moment
sleep 10

# Try again
commit-ai

# Or stage fewer files
git reset
git add src/api.go
commit-ai
```

### "Invalid commit message"
```bash
# Use edit option to fix
commit-ai

# Select "✏️  Edit manually..."
# Fix the message format
```

---

## Advanced Usage

### Custom AI Models
```bash
# Use more powerful model
commit-ai -m llama-3.1-70b-versatile

# Use faster model
commit-ai -m llama-3.1-8b-instant
```

### Environment Variables
```bash
# Override model
export COMMIT_AI_MODEL=llama-3.1-70b-versatile
commit-ai

# Override temperature
export COMMIT_AI_TEMPERATURE=0.5
commit-ai
```

### Batch Processing
```bash
# Process multiple commits
for branch in feature-1 feature-2 feature-3; do
  git checkout $branch
  git add .
  commit-ai -c -y
done
```

---

## Best Practices

1. **Stage related changes together**
   - Better commit messages
   - Cleaner git history
   - Easier to review

2. **Use verbose mode during development**
   - Understand scope detection
   - See diff analysis
   - Debug issues

3. **Review AI suggestions**
   - Choose the best option
   - Edit if needed
   - Ensure accuracy

4. **Use emojis for visual commits**
   - Better readability
   - Easier to scan history
   - More engaging

5. **Commit frequently**
   - Smaller, focused commits
   - Better git history
   - Easier to revert if needed

---

## Next Steps

- [Examples](EXAMPLES.md)
- [Configuration](CONFIGURATION.md)
- [Troubleshooting](TROUBLESHOOTING.md)
- [FAQ](FAQ.md)

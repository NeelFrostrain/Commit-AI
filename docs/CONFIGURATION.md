# Configuration Guide

Configure Commit-AI for your workflow.

## Environment Variables

### Required

**GROQ_API_KEY**
- Your Groq API key from https://console.groq.com/keys
- Required for AI analysis
- Stored in `~/.commit-ai.env` or `.env`

```bash
export GROQ_API_KEY=gsk_your_key_here
```

### Optional

**COMMIT_AI_MODEL**
- AI model to use
- Default: `llama-3.1-8b-instant`
- Options: `llama-3.1-70b-versatile`, `mixtral-8x7b-32768`, etc.

```bash
export COMMIT_AI_MODEL=llama-3.1-70b-versatile
```

## Configuration Files

### Global Configuration

**Location**: `~/.commit-ai.env`

```bash
GROQ_API_KEY=gsk_your_key_here
COMMIT_AI_MODEL=llama-3.1-8b-instant
```

### Project Configuration

**Location**: `.env` (in project root)

```bash
GROQ_API_KEY=gsk_your_key_here
COMMIT_AI_MODEL=llama-3.1-8b-instant
```

Project configuration overrides global configuration.

### YAML Configuration (Future)

**Location**: `.commit-ai.yaml`

```yaml
# AI Settings
model: "llama-3.1-8b-instant"
temperature: 0.7
max_tokens: 15000

# Commit Rules
rules:
  max_title_length: 72
  require_scope: false
  allowed_types:
    - feat
    - fix
    - docs
    - refactor
    - test
    - chore
    - build
    - ci
    - perf
    - style

# Scope Settings
scopes:
  auto_detect: true
  aliases:
    frontend: ui
    backend: api
    database: db

# Ignore Patterns
ignore:
  - "*.log"
  - "node_modules/"
  - ".git/"
```

## Getting API Key

### Step 1: Visit Groq Console
Go to https://console.groq.com/keys

### Step 2: Sign Up or Log In
- Create a free account or log in
- No credit card required for free tier

### Step 3: Create API Key
- Click "Create New API Key"
- Copy the key
- Save it securely

### Step 4: Configure Commit-AI
```bash
# Option 1: Interactive setup
commit-ai
# Enter key when prompted

# Option 2: Manual setup
echo "GROQ_API_KEY=gsk_your_key_here" > ~/.commit-ai.env
```

## Command-Line Flags

### Model Override
```bash
# Use different model
commit-ai -m llama-3.1-70b-versatile

# Use local model (future)
commit-ai -m local:llama2
```

### Verbose Mode
```bash
# See detailed analysis
commit-ai -v

# Output includes:
# - Detected scope
# - Number of files analyzed
# - Diff size
# - Number of options generated
```

### Emoji Mode
```bash
# Add emojis to commit messages
commit-ai -e

# Emojis added to:
# - Commit title
# - Category headers
# - Bullet points
```

### Auto-Commit
```bash
# Commit automatically after selection
commit-ai -c

# Skip confirmation prompts
commit-ai -y

# Combine flags
commit-ai -cev
```

## Advanced Configuration

### Custom Ignore Patterns

Add to `.gitignore`:
```
# Binary files
*.exe
*.dll
*.so

# Build artifacts
dist/
build/

# Dependencies
node_modules/
vendor/
```

Commit-AI automatically respects `.gitignore` patterns.

### Temperature Settings

Control AI creativity (0.0 = deterministic, 1.0 = creative):

```bash
# In .env
COMMIT_AI_TEMPERATURE=0.7  # Balanced (default)
COMMIT_AI_TEMPERATURE=0.3  # More consistent
COMMIT_AI_TEMPERATURE=0.9  # More creative
```

### Token Limits

Control response length:

```bash
# In .env
COMMIT_AI_MAX_TOKENS=8000   # Shorter responses
COMMIT_AI_MAX_TOKENS=15000  # Longer responses (default)
```

## Troubleshooting

### "API key not found"
```bash
# Check if key is set
echo $GROQ_API_KEY

# Set key
export GROQ_API_KEY=gsk_your_key_here

# Or create .env file
echo "GROQ_API_KEY=gsk_your_key_here" > ~/.commit-ai.env
```

### "Invalid API key"
- Verify key from https://console.groq.com/keys
- Check for typos
- Regenerate key if needed

### "Rate limit exceeded"
- Upgrade Groq plan
- Wait before retrying
- Stage fewer files at once

### "Model not found"
- Check available models at https://console.groq.com/docs/models
- Use default model: `llama-3.1-8b-instant`

## Best Practices

1. **Store API key securely**
   - Use `~/.commit-ai.env` for global config
   - Add `.env` to `.gitignore` for project config
   - Never commit API keys

2. **Use project-specific config**
   - Create `.env` in project root
   - Override global settings as needed
   - Share `.commit-ai.yaml` with team

3. **Optimize for your workflow**
   - Use `-e` flag for visual commits
   - Use `-v` flag during development
   - Use `-c` flag for automation

4. **Monitor API usage**
   - Check usage at https://console.groq.com/usage
   - Adjust token limits if needed
   - Consider upgrading plan if needed

## Environment Variables Reference

| Variable | Default | Description |
|----------|---------|-------------|
| `GROQ_API_KEY` | - | Groq API key (required) |
| `COMMIT_AI_MODEL` | `llama-3.1-8b-instant` | AI model to use |
| `COMMIT_AI_TEMPERATURE` | `0.7` | AI creativity (0.0-1.0) |
| `COMMIT_AI_MAX_TOKENS` | `15000` | Max response length |

## Next Steps

- [Usage Guide](USAGE.md)
- [Examples](EXAMPLES.md)
- [Troubleshooting](TROUBLESHOOTING.md)

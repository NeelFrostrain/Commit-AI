# Quick Start Guide

Get started with Commit-AI in 5 minutes.

## Prerequisites

- Git installed and configured
- Groq API key (free at https://console.groq.com/keys)

## Step 1: Install Commit-AI

### Windows
```powershell
# Download and run installer
# Or use Go
go install github.com/NeelFrostrain/Commit-Ai@latest
```

### macOS
```bash
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-darwin-arm64 -o /usr/local/bin/commit-ai
chmod +x /usr/local/bin/commit-ai
```

### Linux
```bash
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-linux-amd64 -o ~/.local/bin/commit-ai
chmod +x ~/.local/bin/commit-ai
```

## Step 2: Get API Key

1. Visit https://console.groq.com/keys
2. Sign up or log in
3. Create a new API key
4. Copy the key

## Step 3: Configure

Run commit-ai for the first time:

```bash
commit-ai
```

When prompted, enter your API key. It will be saved to `~/.commit-ai.env`.

## Step 4: Make Changes

```bash
# Make your code changes
echo "console.log('hello')" > app.js

# Stage changes
git add .
```

## Step 5: Generate Commit Message

```bash
# Generate commit message
commit-ai

# Select from 3 AI-generated options
# Choose report format (keep, edit, or skip)
# Review and confirm
```

## Step 6: Commit

```bash
# Option 1: Use commit-ai to commit
commit-ai -c

# Option 2: Manual commit
git commit -m "feat(app): add hello world"
```

---

## Common Commands

```bash
# Generate and commit in one step
commit-ai -c

# Verbose mode (see analysis)
commit-ai -v

# Add emojis
commit-ai -e

# Skip confirmations
commit-ai -y

# Combine flags
commit-ai -cev
```

---

## Next Steps

- [Full Usage Guide](USAGE.md)
- [Configuration](CONFIGURATION.md)
- [Examples](EXAMPLES.md)
- [Troubleshooting](TROUBLESHOOTING.md)

---

## Tips

1. **Stage only related changes** - Better commit messages when changes are grouped logically
2. **Use verbose mode** - See what the AI is analyzing with `-v` flag
3. **Try emoji mode** - Use `-e` flag for more visual commits
4. **Check examples** - See [Examples](EXAMPLES.md) for inspiration

---

## Troubleshooting

### "API key not found"
```bash
# Set API key
echo "GROQ_API_KEY=your_key_here" > ~/.commit-ai.env
```

### "No changes detected"
```bash
# Make sure you have staged changes
git status
git add .
```

### "Rate limit exceeded"
- Wait a moment and try again
- Upgrade your Groq plan
- Stage fewer files at once

For more help, see [Troubleshooting Guide](TROUBLESHOOTING.md).

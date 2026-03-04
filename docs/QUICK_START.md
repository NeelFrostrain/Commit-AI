# Commit-AI Quick Start Guide

## 🎯 5-Minute Setup

### Step 1: Install
```bash
# If you have Go installed:
go install github.com/NeelFrostrain/Commit-Ai@latest

# Or download the binary from releases
```

### Step 2: Get API Key
1. Visit https://console.groq.com/keys
2. Create a free account
3. Generate an API key

### Step 3: First Run
```bash
commit-ai
# Paste your API key when prompted
```

Done! You're ready to generate AI-powered commits.

---

## 📖 Common Usage Patterns

### Basic Workflow
```bash
# Make your changes
git add .

# Generate commit message
commit-ai

# Select from AI suggestions
# Choose to keep, edit, or skip the report
```

### Auto-Commit Mode
```bash
# Generate and commit in one step
commit-ai -c

# Skip all confirmations (careful!)
commit-ai -cy
```

### Verbose Mode
```bash
# See detailed information
commit-ai -v

# Useful for debugging or learning
```

### Custom Model
```bash
# Use a different AI model
commit-ai -m llama-3.1-70b-versatile

# More powerful but slower
```

---

## 🎨 Example Scenarios

### Scenario 1: New Feature
```bash
# You added a login page
git add src/pages/login.tsx src/api/auth.ts

commit-ai -c
# AI suggests: "feat(auth): add user login page"
```

### Scenario 2: Bug Fix
```bash
# You fixed a memory leak
git add src/cache/manager.go

commit-ai -c
# AI suggests: "fix(cache): resolve memory leak in cleanup"
```

### Scenario 3: Refactoring
```bash
# You reorganized code
git add src/utils/*.ts

commit-ai -c
# AI suggests: "refactor(utils): reorganize helper functions"
```

---

## ⚙️ Configuration

### Environment Variables
```bash
# Set in .env or ~/.commit-ai.env
GROQ_API_KEY=your_key_here
COMMIT_AI_MODEL=llama-3.1-8b-instant
```

### Project Configuration
Create `.commit-ai.yaml` in your project:
```yaml
model: "llama-3.1-8b-instant"
temperature: 0.6
rules:
  max_title_length: 72
  require_scope: false
```

---

## 🔧 Troubleshooting

### "No staged changes detected"
```bash
# Stage your changes first
git add .

# Or let commit-ai do it
commit-ai  # Answer "yes" when prompted
```

### "GROQ_API_KEY not found"
```bash
# Set it in your environment
echo "GROQ_API_KEY=your_key" > ~/.commit-ai.env

# Or export it
export GROQ_API_KEY=your_key
```

### "AI request failed"
- Check your internet connection
- Verify your API key is valid
- Check Groq service status

### Generated messages are too generic
```bash
# Use verbose mode for better context
commit-ai -v

# Try a more powerful model
commit-ai -m llama-3.1-70b-versatile

# Make sure you have meaningful changes staged
```

---

## 💡 Pro Tips

1. **Stage selectively**: Only stage related changes for better AI suggestions
   ```bash
   git add src/auth/*  # Only auth-related files
   commit-ai
   ```

2. **Use with git hooks**: Automate commit message generation
   ```bash
   # Add to .git/hooks/prepare-commit-msg
   #!/bin/bash
   commit-ai --auto > $1
   ```

3. **Combine with conventional commits**: AI follows the standard automatically
   - `feat:` for new features
   - `fix:` for bug fixes
   - `refactor:` for code improvements

4. **Review before committing**: Always review AI suggestions
   ```bash
   commit-ai  # Review first
   # Then commit manually if needed
   ```

5. **Use verbose mode to learn**: See what the AI is analyzing
   ```bash
   commit-ai -v
   ```

---

## 📚 Next Steps

- Read [IMPROVEMENTS.md](IMPROVEMENTS.md) for advanced features
- Check [CONTRIBUTING.md](CONTRIBUTING.md) to contribute
- Star the repo if you find it useful!
- Share with your team

---

## 🆘 Need Help?

- 📖 Read the full [README.md](README.md)
- 🐛 Report issues on GitHub
- 💬 Start a discussion for questions
- ⭐ Star the repo to show support

Happy committing! 🚀

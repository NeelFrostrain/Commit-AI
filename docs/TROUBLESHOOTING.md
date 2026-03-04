# Troubleshooting Guide

Solutions to common issues.

## Installation Issues

### "commit-ai: command not found"

**Problem**: Command not found after installation

**Solutions**:

1. **Check PATH**
```bash
# macOS/Linux
echo $PATH

# Windows PowerShell
$env:Path
```

2. **Add to PATH**
```bash
# macOS/Linux
export PATH=$PATH:/path/to/commit-ai
echo 'export PATH=$PATH:/path/to/commit-ai' >> ~/.bashrc

# Windows PowerShell
$env:Path += ";C:\Program Files\Commit-AI"
```

3. **Use full path**
```bash
/path/to/commit-ai version
```

### "Permission denied"

**Problem**: Cannot execute binary

**Solution**:
```bash
# Make executable
chmod +x /path/to/commit-ai
```

### "Wrong architecture"

**Problem**: Binary doesn't work on your system

**Solution**:
```bash
# Check your architecture
uname -m  # macOS/Linux
# or
[System.Environment]::Is64BitOperatingSystem  # Windows

# Download correct binary
# amd64 for 64-bit Intel
# arm64 for Apple Silicon or ARM64
```

---

## Configuration Issues

### "API key not found"

**Problem**: Groq API key not configured

**Solutions**:

1. **Check if key is set**
```bash
echo $GROQ_API_KEY
```

2. **Set key globally**
```bash
# macOS/Linux
export GROQ_API_KEY=gsk_your_key_here
echo 'export GROQ_API_KEY=gsk_your_key_here' >> ~/.bashrc

# Windows PowerShell
$env:GROQ_API_KEY = "gsk_your_key_here"
```

3. **Create .env file**
```bash
# Global
echo "GROQ_API_KEY=gsk_your_key_here" > ~/.commit-ai.env

# Project
echo "GROQ_API_KEY=gsk_your_key_here" > .env
```

4. **Get API key**
   - Visit https://console.groq.com/keys
   - Create new API key
   - Copy and save

### "Invalid API key"

**Problem**: API key is invalid or expired

**Solutions**:

1. **Verify key format**
   - Should start with `gsk_`
   - Should be long string

2. **Check key validity**
   - Visit https://console.groq.com/keys
   - Verify key is active
   - Regenerate if needed

3. **Update key**
```bash
# Update .env file
echo "GROQ_API_KEY=gsk_new_key_here" > ~/.commit-ai.env
```

---

## Git Issues

### "No changes detected"

**Problem**: Commit-AI says no changes found

**Solutions**:

1. **Check git status**
```bash
git status
```

2. **Stage changes**
```bash
# Stage specific files
git add src/file.go

# Or stage all
git add .
```

3. **Try again**
```bash
commit-ai
```

### "Failed to get diff"

**Problem**: Cannot retrieve git diff

**Solutions**:

1. **Check git installation**
```bash
git --version
```

2. **Check repository**
```bash
git status
```

3. **Check permissions**
```bash
# Ensure you have read permissions
ls -la .git
```

### "Failed to stage files"

**Problem**: Cannot stage changes

**Solutions**:

1. **Check git status**
```bash
git status
```

2. **Try manual staging**
```bash
git add .
```

3. **Check permissions**
```bash
# Ensure write permissions
ls -la .git
```

---

## AI Issues

### "AI request failed"

**Problem**: API request to Groq failed

**Solutions**:

1. **Check internet connection**
```bash
ping google.com
```

2. **Check API key**
```bash
echo $GROQ_API_KEY
```

3. **Check Groq status**
   - Visit https://status.groq.com
   - Check for service outages

4. **Try again**
```bash
commit-ai
```

### "Rate limit exceeded"

**Problem**: Too many API requests

**Solutions**:

1. **Wait and retry**
```bash
sleep 10
commit-ai
```

2. **Stage fewer files**
```bash
git reset
git add src/api.go
commit-ai
```

3. **Upgrade Groq plan**
   - Visit https://console.groq.com/settings/billing
   - Upgrade to higher tier

### "Invalid response format"

**Problem**: AI response doesn't match expected format

**Solutions**:

1. **Try again**
```bash
commit-ai
```

2. **Use verbose mode**
```bash
commit-ai -v
```

3. **Check diff size**
```bash
git diff --cached | wc -c
```

4. **Reduce diff size**
```bash
git reset
git add src/file.go
commit-ai
```

---

## Output Issues

### "Garbled characters in output"

**Problem**: Special characters not displaying correctly

**Solutions**:

1. **Check terminal encoding**
```bash
# macOS/Linux
echo $LANG

# Windows PowerShell
[System.Console]::OutputEncoding
```

2. **Set UTF-8 encoding**
```bash
# macOS/Linux
export LANG=en_US.UTF-8

# Windows PowerShell
[System.Console]::OutputEncoding = [System.Text.Encoding]::UTF8
```

### "Colors not showing"

**Problem**: Output not colored

**Solutions**:

1. **Check terminal support**
   - Use modern terminal (iTerm2, Windows Terminal, etc.)
   - Ensure 256-color support

2. **Force colors**
```bash
# Set environment variable
export FORCE_COLOR=1
commit-ai
```

---

## Performance Issues

### "Slow response"

**Problem**: Commit-AI takes too long

**Solutions**:

1. **Check internet speed**
```bash
ping -c 5 api.groq.com
```

2. **Reduce diff size**
```bash
# Stage fewer files
git reset
git add src/file.go
commit-ai
```

3. **Use faster model**
```bash
commit-ai -m llama-3.1-8b-instant
```

### "High memory usage"

**Problem**: Commit-AI uses too much memory

**Solutions**:

1. **Reduce diff size**
```bash
git reset
git add src/file.go
commit-ai
```

2. **Close other applications**
   - Free up system memory

3. **Restart terminal**
```bash
exit
# Reopen terminal
commit-ai
```

---

## Update Issues

### "Update failed"

**Problem**: Cannot update to latest version

**Solutions**:

1. **Check internet connection**
```bash
ping github.com
```

2. **Try manual update**
   - Download latest binary from GitHub
   - Replace current binary

3. **Check permissions**
```bash
# Ensure write permissions
ls -la /usr/local/bin/commit-ai
```

### "Update stuck"

**Problem**: Update process hangs

**Solutions**:

1. **Cancel update**
```bash
Ctrl+C
```

2. **Try again**
```bash
commit-ai update
```

3. **Manual update**
   - Download from GitHub
   - Replace binary manually

---

## Commit Issues

### "Commit failed"

**Problem**: Git commit fails

**Solutions**:

1. **Check git configuration**
```bash
git config user.name
git config user.email
```

2. **Configure git**
```bash
git config --global user.name "Your Name"
git config --global user.email "your@email.com"
```

3. **Check permissions**
```bash
# Ensure write permissions
ls -la .git
```

### "Invalid commit message"

**Problem**: Commit message format is invalid

**Solutions**:

1. **Use edit option**
```bash
commit-ai
# Select "✏️  Edit manually..."
# Fix the message
```

2. **Check format**
   - Should be: `type(scope): description`
   - Example: `feat(api): add authentication`

3. **Validate manually**
```bash
git commit -m "feat(api): add authentication"
```

---

## Debugging

### Enable Verbose Mode

```bash
# See detailed analysis
commit-ai -v

# Output includes:
# - Detected scope
# - Number of files
# - Diff size
# - Number of options
```

### Check Configuration

```bash
# View current configuration
echo $GROQ_API_KEY
echo $COMMIT_AI_MODEL

# View .env file
cat ~/.commit-ai.env
```

### Check Git Status

```bash
# View staged changes
git diff --cached

# View all changes
git status

# View file list
git diff --cached --name-only
```

---

## Getting Help

### Resources

- **GitHub Issues**: https://github.com/NeelFrostrain/Commit-Ai/issues
- **GitHub Discussions**: https://github.com/NeelFrostrain/Commit-Ai/discussions
- **Documentation**: https://github.com/NeelFrostrain/Commit-Ai/tree/main/docs

### Reporting Issues

When reporting issues, include:

1. **Version**
```bash
commit-ai version
```

2. **System info**
```bash
uname -a  # macOS/Linux
# or
[System.Environment]::OSVersion  # Windows
```

3. **Error message**
```bash
commit-ai -v 2>&1
```

4. **Steps to reproduce**
   - Exact commands run
   - Expected vs actual behavior

---

## FAQ

**Q: Is my API key secure?**
A: Yes, API key is stored locally in `~/.commit-ai.env` with secure permissions (600).

**Q: Can I use Commit-AI offline?**
A: No, Commit-AI requires internet connection for AI analysis.

**Q: Can I use different AI models?**
A: Yes, use `-m` flag: `commit-ai -m llama-3.1-70b-versatile`

**Q: How much does it cost?**
A: Groq offers free tier with rate limits. Upgrade for higher limits.

**Q: Can I use Commit-AI in CI/CD?**
A: Yes, use `-c -y` flags for automation: `commit-ai -c -y`

---

## Next Steps

- [Usage Guide](USAGE.md)
- [Configuration](CONFIGURATION.md)
- [FAQ](FAQ.md)

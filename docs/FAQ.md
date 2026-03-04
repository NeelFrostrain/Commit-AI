# Frequently Asked Questions

Common questions about Commit-AI.

## General Questions

### What is Commit-AI?

Commit-AI is an intelligent Git commit message generator that uses AI to analyze your code changes and create professional, detailed commit messages following the Conventional Commits specification.

### How does it work?

1. You make code changes and stage them with `git add`
2. Run `commit-ai`
3. Commit-AI retrieves your staged changes
4. Sends the diff to Groq's AI for analysis
5. AI generates 3 commit message options
6. You select one and optionally edit it
7. Commit-AI creates the git commit

### Is it free?

Yes! Commit-AI is free and open-source (MIT license). Groq also offers a free tier for API usage.

### What platforms are supported?

- Windows (amd64)
- macOS (Intel and Apple Silicon)
- Linux (amd64, arm64)

### Can I use it offline?

No, Commit-AI requires an internet connection to communicate with Groq's AI API.

---

## Installation & Setup

### How do I install Commit-AI?

See [Installation Guide](INSTALLATION.md) for detailed instructions for your platform.

Quick install:
```bash
# macOS/Linux
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-darwin-arm64 -o /usr/local/bin/commit-ai
chmod +x /usr/local/bin/commit-ai

# Windows
# Download installer from GitHub releases
```

### How do I get an API key?

1. Visit https://console.groq.com/keys
2. Sign up or log in
3. Create a new API key
4. Copy and save the key

### How do I configure Commit-AI?

See [Configuration Guide](CONFIGURATION.md).

Quick setup:
```bash
commit-ai
# Enter API key when prompted
```

---

## Usage Questions

### How do I use Commit-AI?

Basic workflow:
```bash
git add .
commit-ai
# Select from 3 options
# Review and confirm
```

Or commit automatically:
```bash
commit-ai -c
```

### What flags are available?

| Flag | Description |
|------|-------------|
| `-c` | Commit automatically |
| `-y` | Skip confirmations |
| `-v` | Verbose mode |
| `-e` | Add emojis |
| `-m` | Override AI model |

### Can I edit the commit message?

Yes! When selecting the commit title, choose "✏️  Edit manually..." to customize it.

### Can I edit the commit body?

Yes! When selecting the report format, choose "✏️  Edit report" to customize the body.

### Can I use Commit-AI in CI/CD?

Yes! Use the `-c -y` flags for automation:
```bash
commit-ai -c -y
```

### Can I use Commit-AI with pre-commit hooks?

Yes! Add to `.git/hooks/pre-commit`:
```bash
#!/bin/bash
commit-ai -v
```

---

## AI & Model Questions

### What AI model does Commit-AI use?

Default: `llama-3.1-8b-instant` (fast and accurate)

Other options:
- `llama-3.1-70b-versatile` (more powerful)
- `mixtral-8x7b-32768` (alternative)

### Can I use a different AI model?

Yes! Use the `-m` flag:
```bash
commit-ai -m llama-3.1-70b-versatile
```

### How accurate is the AI?

Very accurate! The AI understands code logic and generates meaningful commit messages. Test it with your own code to see.

### Can I use my own AI model?

Not yet, but it's planned for v2.0.0. Currently, only Groq models are supported.

### How much does the AI cost?

Groq offers a free tier with rate limits. Upgrade your plan for higher limits.

---

## Security & Privacy

### Is my API key secure?

Yes! Your API key is stored locally in `~/.commit-ai.env` with secure permissions (600). It's never logged or displayed.

### Does Commit-AI send my code to external servers?

Only the staged diff is sent to Groq's API for analysis. No other data is transmitted.

### Can I use Commit-AI in a corporate environment?

Yes! You can:
- Use your own Groq account
- Configure proxy settings
- Store API key securely
- Use in CI/CD pipelines

### What data does Commit-AI collect?

None! Commit-AI doesn't collect any usage data or analytics.

---

## Performance & Limits

### Why is Commit-AI slow?

Possible reasons:
- Slow internet connection
- Large diff size
- Groq API rate limit
- System resource constraints

Solutions:
- Check internet speed
- Stage fewer files
- Wait and retry
- Close other applications

### What's the maximum diff size?

Default: 8,000 characters (configurable)

Larger diffs are truncated to stay within API limits.

### What's the rate limit?

Groq's free tier has rate limits. Check your usage at https://console.groq.com/usage.

Upgrade your plan for higher limits.

### Can I increase the diff size?

Yes! Set environment variable:
```bash
export COMMIT_AI_MAX_TOKENS=20000
```

---

## Troubleshooting

### "API key not found"

Set your API key:
```bash
echo "GROQ_API_KEY=gsk_your_key_here" > ~/.commit-ai.env
```

### "No changes detected"

Stage your changes:
```bash
git add .
commit-ai
```

### "Rate limit exceeded"

Wait a moment and try again, or upgrade your Groq plan.

### "AI request failed"

Check your internet connection and API key validity.

For more troubleshooting, see [Troubleshooting Guide](TROUBLESHOOTING.md).

---

## Features & Roadmap

### What commit types are supported?

- feat (features)
- fix (bug fixes)
- docs (documentation)
- style (formatting)
- refactor (code restructuring)
- perf (performance)
- test (tests)
- chore (maintenance)
- build (build system)
- ci (CI/CD)

### Can I customize commit types?

Not yet, but it's planned for future versions.

### What's the roadmap?

See [Roadmap](ROADMAP.md) for planned features.

### Can I contribute?

Yes! See [Contributing Guide](CONTRIBUTING.md).

---

## Comparison

### How is Commit-AI different from other tools?

| Feature | Commit-AI | Others |
|---------|-----------|--------|
| AI-powered | ✅ | ✅ |
| Free | ✅ | ❌ |
| Open-source | ✅ | ❌ |
| Multiple options | ✅ | ❌ |
| Emoji support | ✅ | ❌ |
| Detailed reports | ✅ | ❌ |
| Cross-platform | ✅ | ✅ |

### Why should I use Commit-AI?

- Free and open-source
- AI-powered analysis
- Multiple commit options
- Detailed reports
- Emoji support
- Easy to use
- Cross-platform

---

## Getting Help

### Where can I get help?

- [Troubleshooting Guide](TROUBLESHOOTING.md)
- [GitHub Issues](https://github.com/NeelFrostrain/Commit-Ai/issues)
- [GitHub Discussions](https://github.com/NeelFrostrain/Commit-Ai/discussions)
- [Documentation](https://github.com/NeelFrostrain/Commit-Ai/tree/main/docs)

### How do I report a bug?

1. Check [Troubleshooting Guide](TROUBLESHOOTING.md)
2. Search [GitHub Issues](https://github.com/NeelFrostrain/Commit-Ai/issues)
3. Create new issue with:
   - Version: `commit-ai version`
   - System info: `uname -a`
   - Error message
   - Steps to reproduce

### How do I request a feature?

1. Check [Roadmap](ROADMAP.md)
2. Search [GitHub Issues](https://github.com/NeelFrostrain/Commit-Ai/issues)
3. Create new issue with feature description

---

## More Questions?

- **GitHub Issues**: https://github.com/NeelFrostrain/Commit-Ai/issues
- **GitHub Discussions**: https://github.com/NeelFrostrain/Commit-Ai/discussions
- **Documentation**: https://github.com/NeelFrostrain/Commit-Ai/tree/main/docs

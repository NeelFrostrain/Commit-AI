# Commit-AI 🤖

<div align="center">

**Transform your Git commits with AI-powered, professional commit messages**

[![Version](https://img.shields.io/badge/version-1.2.0-blue?style=flat-square)](https://github.com/NeelFrostrain/Commit-Ai/releases)
[![Go](https://img.shields.io/badge/go-%3E%3D1.21-00ADD8?style=flat-square&logo=go&logoColor=white)](https://golang.org)
[![Groq](https://img.shields.io/badge/Groq-AI-cyan?style=flat-square)](https://groq.com)
[![License](https://img.shields.io/badge/license-MIT-green?style=flat-square)](LICENSE)

[Features](#-features) • [Installation](#-installation) • [Usage](#-usage) • [Documentation](#-documentation) • [Contributing](#-contributing)

</div>

---

## 🎯 What is Commit-AI?

Commit-AI is an intelligent Git commit message generator that analyzes your code changes and creates professional, detailed commit messages following the [Conventional Commits](https://www.conventionalcommits.org/) specification.

**Stop writing vague commits like:**
```
git commit -m "fixed stuff"
git commit -m "updates"
git commit -m "wip"
```

**Start writing professional commits like:**
```
feat(api): add JWT authentication system

FEATURES:
- Implemented JWT-based authentication with RS256 signing
- Added login endpoint with rate limiting
- Created middleware for protected routes

TECHNICAL DETAILS:
- 8 files changed: 450 insertions, 20 deletions
- Test coverage: 95% on auth module

IMPACT:
- Improved security with industry-standard JWT
- Better user experience with automatic token refresh
```

---

## ✨ Features

### 🧠 Intelligent Analysis
- **Deep Diff Analysis**: Understands code logic, not just file metadata
- **Scope Detection**: Automatically detects scope from file paths (api, ui, auth, etc.)
- **Change Classification**: Identifies feat, fix, refactor, perf, docs, and more

### 📝 Professional Output
- **Conventional Commits**: Strictly follows the `type(scope): description` standard
- **Structured Reports**: Generates detailed, categorized commit bodies
- **Multiple Options**: Provides 3 distinct commit message suggestions

### 🚀 Developer Experience
- **Auto-Stage**: Detects and offers to stage all changes automatically
- **Interactive**: Choose from AI suggestions or edit manually
- **Verbose Mode**: See detailed analysis and debugging information
- **Fast**: Powered by Groq's Llama 3.1 for near-instant results

### 🔧 Customizable
- **Model Selection**: Override AI model with `-m` flag
- **Configuration**: Support for project-specific settings
- **Templates**: Customizable commit message templates

---

## � Installation

### Windows

#### Option 1: Installer (Recommended)
1. Download `install-commit-ai.exe` from [latest release](https://github.com/NeelFrostrain/Commit-Ai/releases/latest)
2. Run the installer
3. Restart your terminal
4. Run `commit-ai`

#### Option 2: Manual
1. Download `commit-ai-windows-amd64.exe` from [releases](https://github.com/NeelFrostrain/Commit-Ai/releases/latest)
2. Rename to `commit-ai.exe`
3. Add to your PATH

### macOS

```bash
# Download and install
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-darwin-amd64 -o /usr/local/bin/commit-ai
chmod +x /usr/local/bin/commit-ai

# For Apple Silicon (M1/M2)
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-darwin-arm64 -o /usr/local/bin/commit-ai
chmod +x /usr/local/bin/commit-ai
```

### Linux

```bash
# Download and install
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-linux-amd64 -o ~/.local/bin/commit-ai
chmod +x ~/.local/bin/commit-ai

# For ARM64
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-linux-arm64 -o ~/.local/bin/commit-ai
chmod +x ~/.local/bin/commit-ai
```

### Go Install

```bash
go install github.com/NeelFrostrain/Commit-Ai@latest
```

---

## 🚀 Quick Start

### 1. Get API Key
Visit [Groq Console](https://console.groq.com/keys) and create a free API key.

### 2. First Run
```bash
commit-ai
# Enter your API key when prompted
# It will be saved to ~/.commit-ai.env
```

### 3. Make Changes and Commit
```bash
# Make your changes
git add .

# Generate commit message
commit-ai

# Or generate and commit in one step
commit-ai -c
```

---

## 💡 Usage

### Basic Commands

```bash
# Generate commit message
commit-ai

# Generate and commit
commit-ai -c

# Verbose mode (see details)
commit-ai -v

# Skip confirmations
commit-ai -y

# Use different AI model
commit-ai -m llama-3.1-70b-versatile

# Check version
commit-ai version

# Check for updates
commit-ai update --check

# Update to latest version
commit-ai update

# Combine flags
commit-ai -cv  # Verbose + auto-commit
```

### Command Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--commit` | `-c` | Commit changes after selection |
| `--yes` | `-y` | Skip confirmation prompts |
| `--verbose` | `-v` | Show detailed information |
| `--model` | `-m` | Override AI model |
| `--help` | `-h` | Show help message |

### Commands

| Command | Description |
|---------|-------------|
| `commit-ai` | Generate commit message (default) |
| `commit-ai version` | Show version information |
| `commit-ai update` | Update to latest version |
| `commit-ai update --check` | Check for updates without installing |
| `commit-ai update --force` | Force update even if on latest version |

### Workflow Example

```bash
# 1. Make your changes
vim src/api/auth.go src/middleware/jwt.go

# 2. Run commit-ai with verbose mode
commit-ai -v

# Output:
# [!] Found unstaged changes.
# ? Stage all changes (git add .)? Yes
# [✓] All changes staged
# [Info] Detected scope: api
# [Info] Analyzing 2 files
# [Info] Diff size: 1,234 characters
# [Commit-AI] Analyzing changes with llama-3.1-8b-instant...
# [Info] Generated 3 options

# 3. Select from AI suggestions
# ? Select commit title:
# > feat(api): add JWT authentication system
#   feat(auth): implement token-based authentication
#   chore(security): add authentication middleware

# 4. Choose report format
# ? Commit body:
# > 📝 Keep AI report
#   ✏️  Edit report
#   ⊘  No report

# 5. Review and commit
# ─────────────────────────────────
# Title: feat(api): add JWT authentication system
# Body:
# FEATURES:
# - Implemented JWT-based authentication
# - Added login endpoint with validation
# - Created middleware for protected routes
# ─────────────────────────────────
```

---

## 📚 Documentation

### User Guides
- [Quick Start Guide](docs/QUICK_START.md) - Get started in 5 minutes
- [Demo & Examples](docs/DEMO.md) - Real-world usage examples
- [Enhanced Prompts](docs/ENHANCED_PROMPTS.md) - How AI generates messages

### Developer Guides
- [Contributing Guide](docs/CONTRIBUTING.md) - How to contribute
- [Improvements Roadmap](docs/IMPROVEMENTS.md) - Future enhancements
- [Upgrade Summary](docs/UPGRADE_SUMMARY.md) - Latest changes

### Reference
- [Changelog](docs/CHANGELOG.md) - Version history
- [Before & After](docs/BEFORE_AFTER.md) - Visual comparisons
- [Test Report](docs/TEST_REPORT.md) - Test results

---

## 🔧 Configuration

### Environment Variables

Create `.env` in your project or `~/.commit-ai.env` globally:

```bash
GROQ_API_KEY=your_key_here
COMMIT_AI_MODEL=llama-3.1-8b-instant
```

### Project Configuration

Create `.commit-ai.yaml` in your project:

```yaml
model: "llama-3.1-8b-instant"
temperature: 0.7
max_tokens: 8000

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

scopes:
  auto_detect: true
  aliases:
    frontend: ui
    backend: api
```

---

## 🎨 Examples

### Feature Addition
```
feat(api): add user authentication endpoints

FEATURES:
- Implemented JWT-based authentication with RS256 signing
- Added login endpoint with rate limiting (5 attempts/minute)
- Created middleware for protected routes with role-based access

SECURITY:
- Added secure password hashing with bcrypt
- Implemented token refresh mechanism
- Configured secure headers (HSTS, CSP)

TECHNICAL DETAILS:
- 8 files changed: 450 insertions, 20 deletions
- Test coverage: 95% on auth module
- Performance: Token validation <1ms average

IMPACT:
- Improved security with industry-standard JWT
- Better user experience with automatic token refresh
- Reduced server load with stateless authentication
```

### Bug Fix
```
fix(cache): resolve memory leak in cleanup method

BUG FIXES:
- Changed cleanup to delete items individually
- Fixed map reallocation causing memory retention
- Added proper resource cleanup in defer statements

TECHNICAL DETAILS:
- 1 file changed: 15 insertions, 8 deletions
- Memory usage reduced by 40% in long-running tests

IMPACT:
- Prevents memory leak in production
- Improves application stability
- Reduces memory footprint over time
```

### Refactoring
```
refactor(internal): restructure codebase into modular packages

ARCHITECTURE:
- Separated concerns into dedicated modules (config, git, ai)
- Created clean interfaces between components
- Improved code organization and maintainability

IMPROVEMENTS:
- Enhanced testability with dependency injection
- Reduced coupling between modules
- Improved code reusability

TECHNICAL DETAILS:
- 15 files changed: 800 insertions, 300 deletions
- Test coverage increased from 60% to 85%

IMPACT:
- Easier to maintain and extend
- Better code quality and organization
- Improved developer experience
```

### With Breaking Changes
```
feat(api): redesign authentication system

BREAKING CHANGES:
- API endpoint /login now requires POST instead of GET
- Authentication token format changed from Bearer to JWT
- Old tokens will be invalidated after migration

FEATURES:
- Implemented new JWT-based authentication
- Added refresh token mechanism
- Enhanced security with RS256 signing

MIGRATION GUIDE:
- Update API calls from GET to POST
- Replace Bearer tokens with JWT format
- Re-authenticate users after deployment

TECHNICAL DETAILS:
- 12 files changed: 600 insertions, 200 deletions
- Migration script provided in /scripts

IMPACT:
- Significantly improved security
- Better scalability with stateless auth
- Modern authentication standard
```

---

## 🛠️ Development

### Prerequisites
- Go 1.21 or higher
- Git
- Groq API key

### Setup
```bash
# Clone repository
git clone https://github.com/NeelFrostrain/Commit-Ai.git
cd Commit-Ai

# Install dependencies
go mod download

# Build
make build

# Run tests
make test

# Run with verbose output
make run-verbose
```

### Build Commands
```bash
make build          # Build for development
make build-prod     # Build production binary to bin/
make installer      # Build installer
make build-all      # Build for all platforms
make release        # Full release build
```

### Testing
```bash
make test           # Run tests
make test-coverage  # Generate coverage report
make lint           # Run linter
make fmt            # Format code
make check          # Run all checks
```

---

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](docs/CONTRIBUTING.md) for details.

### Quick Contribution Steps
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`commit-ai -c` 😉)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 🐛 Troubleshooting

### "No changes detected"
```bash
# Make sure you have changes
git status

# Stage your changes
git add .

# Or let commit-ai do it
commit-ai  # Answer "yes" when prompted
```

### "API key not found"
```bash
# Set it globally
echo "GROQ_API_KEY=your_key" > ~/.commit-ai.env

# Or in your project
echo "GROQ_API_KEY=your_key" > .env
```

### "AI request failed"
- Check your internet connection
- Verify your API key is valid
- Check Groq service status
- Try reducing diff size (stage fewer files)

### "Rate limit exceeded"
- Wait a moment and try again
- Upgrade your Groq plan
- Stage fewer files at once

---

## 📊 Project Stats

- **Language**: Go
- **AI Model**: Groq Llama 3.1
- **Test Coverage**: 85%+
- **Platforms**: Windows, macOS, Linux
- **License**: MIT

---

## 🙏 Acknowledgments

- [Groq](https://groq.com) for lightning-fast AI inference
- [Cobra](https://github.com/spf13/cobra) for CLI framework
- [Survey](https://github.com/AlecAivazis/survey) for interactive prompts
- All contributors and users

---

## 📄 License

MIT © [Neel Frostrain](https://github.com/NeelFrostrain)

See [LICENSE](LICENSE) file for details.

---

## 🔗 Links

- **GitHub**: https://github.com/NeelFrostrain/Commit-Ai
- **Issues**: https://github.com/NeelFrostrain/Commit-Ai/issues
- **Releases**: https://github.com/NeelFrostrain/Commit-Ai/releases
- **Groq Console**: https://console.groq.com/keys

---

<div align="center">

**Made with ❤️ by developers, for developers**

⭐ Star us on GitHub if you find this useful!

</div>

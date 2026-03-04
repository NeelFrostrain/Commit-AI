# Commit-AI 🤖

<div align="center">

**Transform your Git commits with AI-powered, professional commit messages**

[![Version](https://img.shields.io/badge/version-1.2.0-blue?style=flat-square)](https://github.com/NeelFrostrain/Commit-Ai/releases)
[![Go](https://img.shields.io/badge/go-%3E%3D1.21-00ADD8?style=flat-square&logo=go&logoColor=white)](https://golang.org)
[![Groq](https://img.shields.io/badge/Groq-AI-cyan?style=flat-square)](https://groq.com)
[![License](https://img.shields.io/badge/license-MIT-green?style=flat-square)](LICENSE)
[![Test Coverage](https://img.shields.io/badge/coverage-85%25-brightgreen?style=flat-square)](docs/TESTING.md)

[Features](#-features) • [Installation](#-installation) • [Quick Start](#-quick-start) • [Documentation](#-documentation) • [Contributing](#-contributing)

</div>

---

## 🎯 What is Commit-AI?

Commit-AI is an intelligent Git commit message generator that analyzes your code changes and creates professional, detailed commit messages following the [Conventional Commits](https://www.conventionalcommits.org/) specification.

**Stop writing vague commits:**
```
git commit -m "fixed stuff"
git commit -m "updates"
git commit -m "wip"
```

**Start writing professional commits:**
```
✨ feat(api): add JWT authentication system

FEATURES:
- Implemented JWT-based authentication with RS256 signing
- Added login endpoint with rate limiting (5 attempts/minute)
- Created middleware for protected routes with role-based access

TECHNICAL DETAILS:
- 8 files changed: 450 insertions(+), 20 deletions(-)
- Test coverage: 95% on auth module
- Performance: Token validation <1ms average

IMPACT:
- Improved security with industry-standard JWT
- Better user experience with automatic token refresh
- Reduced server load with stateless authentication
```

---

## ✨ Features

### 🧠 Intelligent Analysis
- **Deep Diff Analysis**: Understands code logic, not just file metadata
- **Scope Detection**: Automatically detects scope from file paths (api, ui, auth, etc.)
- **Change Classification**: Identifies feat, fix, refactor, perf, docs, and more
- **Binary Filtering**: Excludes 50+ binary file types (executables, images, audio, video, 3D models, etc.)

### 📝 Professional Output
- **Conventional Commits**: Strictly follows the `type(scope): description` standard
- **Structured Reports**: Generates detailed, categorized commit bodies
- **Multiple Options**: Provides 3 distinct commit message suggestions
- **Emoji Support**: Optional emoji prefixes for visual enhancement

### 🚀 Developer Experience
- **Auto-Stage**: Detects and offers to stage all changes automatically
- **Interactive**: Choose from AI suggestions or edit manually
- **Verbose Mode**: See detailed analysis and debugging information
- **Fast**: Powered by Groq's Llama 3.1 for near-instant results

### 🔧 Customizable
- **Model Selection**: Override AI model with `-m` flag
- **Configuration**: Support for project-specific settings
- **Auto-Update**: Built-in update checker and installer
- **Cross-Platform**: Windows, macOS, Linux support

---

## 📦 Installation

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
# Intel Macs
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-darwin-amd64 -o /usr/local/bin/commit-ai
chmod +x /usr/local/bin/commit-ai

# Apple Silicon (M1/M2/M3)
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-darwin-arm64 -o /usr/local/bin/commit-ai
chmod +x /usr/local/bin/commit-ai
```

### Linux

```bash
# AMD64
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-linux-amd64 -o ~/.local/bin/commit-ai
chmod +x ~/.local/bin/commit-ai

# ARM64
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

# Generate and commit automatically
commit-ai -c

# Verbose mode (see analysis details)
commit-ai -v

# Add emojis to commit messages
commit-ai -e

# Skip confirmation prompts
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
commit-ai -cev  # Emoji + Verbose + Auto-commit
```

### Command Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--commit` | `-c` | Commit changes after selection |
| `--yes` | `-y` | Skip confirmation prompts |
| `--verbose` | `-v` | Show detailed analysis information |
| `--emoji` | `-e` | Add emojis to commit messages |
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

---

## 📚 Documentation

### Getting Started
- [Installation Guide](docs/INSTALLATION.md) - Detailed setup instructions
- [Quick Start](docs/QUICK_START.md) - 5-minute setup guide
- [Configuration](docs/CONFIGURATION.md) - Environment variables and settings

### Usage Guides
- [Usage Guide](docs/USAGE.md) - Complete usage documentation
- [Examples](docs/EXAMPLES.md) - Real-world usage examples
- [Emoji Reference](docs/EMOJI_REFERENCE.md) - Emoji types and usage

### Developer Guides
- [Architecture](docs/ARCHITECTURE.md) - Project structure and design
- [Building from Source](docs/BUILD.md) - Build instructions
- [Contributing](docs/CONTRIBUTING.md) - Contribution guidelines
- [Testing](docs/TESTING.md) - Test suite and coverage

### Reference
- [Changelog](docs/CHANGELOG.md) - Version history
- [Supported Files](docs/SUPPORTED_FILES.md) - File type filtering
- [Troubleshooting](docs/TROUBLESHOOTING.md) - Common issues and solutions
- [FAQ](docs/FAQ.md) - Frequently asked questions
- [Roadmap](docs/ROADMAP.md) - Future enhancements

---

## 🎨 Examples

### Feature Addition with Emojis
```
✨ feat(api): add user authentication endpoints

✨ FEATURES:
- 🔐 Implemented JWT-based authentication with RS256 signing
- 🚦 Added login endpoint with rate limiting (5 attempts/minute)
- 🛡️ Created middleware for protected routes with role-based access

🔒 SECURITY:
- 🔑 Added secure password hashing with bcrypt
- 🔄 Implemented token refresh mechanism
- 🛡️ Configured secure headers (HSTS, CSP)

🔧 TECHNICAL DETAILS:
- 📊 8 files changed: 450 insertions(+), 20 deletions(-)
- ✅ Test coverage: 95% on auth module
- ⚡ Performance: Token validation <1ms average

💡 IMPACT:
- 🔒 Improved security with industry-standard JWT
- 🚀 Better user experience with automatic token refresh
- 📉 Reduced server load with stateless authentication
```

### Bug Fix
```
🐛 fix(cache): resolve memory leak in cleanup method

🐛 BUG FIXES:
- Fixed memory leak in cache cleanup method
- Changed cleanup to delete items individually
- Added proper resource cleanup in defer statements

🔧 TECHNICAL DETAILS:
- 1 file changed: 15 insertions(+), 8 deletions(-)
- Memory usage reduced by 40% in long-running tests

💡 IMPACT:
- Prevents memory leak in production
- Improves application stability
- Reduces memory footprint over time
```

### Refactoring
```
♻️ refactor(internal): restructure codebase into modular packages

♻️ ARCHITECTURE:
- Separated concerns into dedicated modules (config, git, ai)
- Created clean interfaces between components
- Improved code organization and maintainability

📈 IMPROVEMENTS:
- Enhanced testability with dependency injection
- Reduced coupling between modules
- Improved code reusability

🔧 TECHNICAL DETAILS:
- 15 files changed: 800 insertions(+), 300 deletions(-)
- Test coverage increased from 60% to 85%

💡 IMPACT:
- Easier to maintain and extend
- Better code quality and organization
- Improved developer experience
```

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
max_tokens: 15000

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
./commit-ai -v
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

For more troubleshooting, see [Troubleshooting Guide](docs/TROUBLESHOOTING.md).

---

## 📊 Project Stats

- **Language**: Go 1.21+
- **AI Model**: Groq Llama 3.1
- **Test Coverage**: 85%+
- **Platforms**: Windows, macOS, Linux
- **Binary Size**: 8-10 MB (optimized)
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

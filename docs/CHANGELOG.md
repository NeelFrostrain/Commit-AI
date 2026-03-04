# Changelog

All notable changes to Commit-AI are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [1.2.0] - 2024-01-15

### Added

- **Auto-Update Command**: New `commit-ai update` command
  - `--check` flag to check for updates without installing
  - `--force` flag to force update even if on latest version
  - Automatic platform detection (Windows, macOS, Linux)
  - Progress display during download
  - Safe update with automatic backup

- **Version Command**: New `commit-ai version` command
  - Shows detailed version information
  - Displays build date and git commit hash
  - Shows Go version and OS/Architecture

- **Emoji Support**: New `-e/--emoji` flag
  - Adds contextual emojis to commit messages
  - 14+ emoji types for different categories
  - Visual hierarchy in commit bodies
  - Improves commit history readability

- **Enhanced AI Prompt**: Improved prompt engineering
  - Critical instructions for accurate analysis
  - Warnings against generic/invented features
  - Detailed guidance for comprehensive reports
  - Better token efficiency

- **Binary File Filtering**: Comprehensive binary file support
  - `.gitattributes` for marking binary files
  - 50+ binary file patterns (executables, images, audio, video, 3D models, etc.)
  - Prevents AI from analyzing compiled code
  - Reduces API token usage

- **Comprehensive Documentation**: 10+ documentation files
  - Installation guide for all platforms
  - Quick start guide (5 minutes)
  - Complete usage guide
  - Configuration guide
  - Real-world examples
  - Architecture documentation
  - Troubleshooting guide
  - FAQ

### Changed

- **Project Rename**: Changed from `Commit-Ai-Go` to `Commit-Ai`
- **Module Path**: Updated to `github.com/NeelFrostrain/Commit-Ai`
- **Documentation Structure**: Reorganized and expanded documentation
- **Build System**: Enhanced with better version tracking
- **AI Prompts**: Improved to generate more accurate commit messages
- **Diff Handling**: Fixed to properly retrieve full diff content

### Fixed

- **Diff Retrieval**: Fixed issue where diff was truncated to 17 characters
- **Binary File Analysis**: AI no longer analyzes binary files
- **Emoji Validation**: Commit messages with leading emojis now validate correctly
- **BuildPrompt Test**: Fixed missing emoji flag parameter
- **Git Diff Arguments**: Corrected command arguments to avoid filtering text files

### Removed

- Temporary documentation files
- Redundant upgrade and test result files
- Unused code and variables

---

## [1.1.0] - 2024-01-10

### Added

- **Auto-Stage Changes**: Automatically detect and offer to stage all changes
- **Verbose Mode**: `-v` flag for detailed analysis information
- **Model Override**: `-m` flag to use different AI models
- **Scope Detection**: Automatic scope detection from file paths
- **Enhanced AI Prompts**: Improved prompt engineering for better commit messages
- **Comprehensive Testing**: 85%+ test coverage on core modules
- **CI/CD Pipeline**: GitHub Actions for automated testing
- **Multi-Platform Builds**: Support for Windows, Linux, macOS (amd64, arm64)

### Changed

- **Modular Architecture**: Restructured into `config`, `git`, and `ai` packages
- **Improved Error Handling**: Better error messages with helpful hints
- **Optimized Diff Size**: Reduced to 8,000 chars to avoid API limits
- **Better User Experience**: Interactive prompts with emojis and colors

### Fixed

- Memory leak in cleanup method
- Edge cases in AI response parsing
- Nil pointer issues in git operations
- Invalid commit message format validation

---

## [1.0.0] - 2024-01-05

### Added

- Initial release of Commit-AI
- AI-powered commit message generation using Groq
- Conventional Commits format support
- Interactive commit message selection
- Basic configuration support
- Environment variable configuration
- Git integration for staged changes
- Cross-platform support (Windows, Linux, macOS)

### Features

- Generate 3 commit message options
- Structured commit bodies with categories
- Support for `.env` configuration
- Cross-platform support

---

## Version History

| Version | Date | Status |
|---------|------|--------|
| 1.2.0 | 2024-01-15 | Latest |
| 1.1.0 | 2024-01-10 | Stable |
| 1.0.0 | 2024-01-05 | Initial |

---

## Upgrade Guide

### From v1.1.0 to v1.2.0

No breaking changes. Simply download the new version:

```bash
# Windows
# Download new install-commit-ai.exe and run

# macOS/Linux
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/latest/download/commit-ai-[platform]-[arch] -o commit-ai
chmod +x commit-ai
```

### From v1.0.0 to v1.1.0

No breaking changes. Configuration files remain compatible.

---

## Future Releases

### Planned for v1.3.0

- `--dry-run` flag for preview without committing
- Commit templates for team standards
- Better error messages with suggestions
- `--undo` functionality to revert commits
- Breaking change detection

### Planned for v2.0.0

- Multi-model support (OpenAI, Anthropic, local models)
- Context-aware learning from commit history
- Pre-commit hooks integration
- Branch analysis and suggestions
- IDE plugins (VSCode, JetBrains)
- Commit analytics and metrics
- Team collaboration features
- Multi-language support

---

## Links

- [GitHub Repository](https://github.com/NeelFrostrain/Commit-Ai)
- [GitHub Releases](https://github.com/NeelFrostrain/Commit-Ai/releases)
- [Issue Tracker](https://github.com/NeelFrostrain/Commit-Ai/issues)
- [Contributing Guide](CONTRIBUTING.md)

# Changelog

All notable changes to Commit-AI will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [1.2.0] - 2024-12-XX

### Added
- **Version Command**: New `commit-ai version` command showing detailed build info
- **Enhanced Versioning**: Build date and git commit embedded in binaries
- **Better Build System**: Improved build scripts with version tracking
- **Comprehensive Documentation**: Organized docs with build and release guides
- **Production Binaries**: Pre-built binaries in `bin/` folder for releases

### Changed
- **Project Rename**: Changed from `Commit-Ai-Go` to `Commit-Ai`
- **Module Path**: Updated to `github.com/NeelFrostrain/Commit-Ai`
- **Documentation Structure**: Cleaned up and reorganized documentation
- **Build Scripts**: Enhanced with better version information display

### Fixed
- Character encoding issues in PowerShell build script
- Module import paths throughout the codebase

### Removed
- Temporary documentation files (summaries, reports)
- Redundant upgrade and test result files

---

## [1.1.0] - 2024-11-XX

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

## [1.0.0] - 2024-10-XX

### Added
- Initial release of Commit-AI
- AI-powered commit message generation using Groq
- Conventional Commits format support
- Interactive commit message selection
- Basic configuration support
- Environment variable configuration
- Git integration for staged changes

### Features
- Generate 3 commit message options
- Structured commit bodies with categories
- Support for `.env` configuration
- Cross-platform support (Windows, Linux, macOS)

---

## Version History

- **v1.2.0** - Enhanced versioning and documentation
- **v1.1.0** - Major architecture overhaul with testing
- **v1.0.0** - Initial release

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

See [IMPROVEMENTS.md](IMPROVEMENTS.md) for planned features and enhancements.

---

## Links

- [GitHub Releases](https://github.com/NeelFrostrain/Commit-Ai/releases)
- [Contributing Guide](CONTRIBUTING.md)
- [Build Guide](BUILD_GUIDE.md)

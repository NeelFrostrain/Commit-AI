# Changelog

All notable changes to Commit-AI will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- **Smart Scope Detection**: Automatically detects scope from file paths (e.g., `api`, `ui`, `auth`)
- **Commit Message Validation**: Validates generated messages against Conventional Commits format
- **Verbose Mode**: `-v` flag for detailed operation information
- **Model Selection**: `-m` flag to override AI model
- **Configuration Module**: Centralized configuration management with support for custom settings
- **Git Operations Module**: Dedicated module for git operations with better error handling
- **Enhanced AI Prompts**: Improved prompt engineering for better, more accurate commit messages
- **Interactive API Key Setup**: Guided setup flow when API key is missing
- **Comprehensive Tests**: Unit tests for AI parsing, validation, and scope detection
- **Development Tools**: Makefile for common development tasks
- **CI/CD Pipeline**: GitHub Actions workflow for automated testing and building
- **Documentation**: 
  - IMPROVEMENTS.md with detailed enhancement suggestions
  - CONTRIBUTING.md for contributors
  - QUICK_START.md for new users
  - Example configuration file (.commit-ai.example.yaml)

### Changed
- **Better Error Messages**: More specific error handling with helpful hints
- **Improved User Experience**: Enhanced visual feedback with emojis and better formatting
- **Optimized Diff Handling**: Better token management and diff truncation
- **Enhanced Parsing**: More robust AI response parsing with fallbacks

### Fixed
- Better handling of edge cases in commit message parsing
- Improved validation logic for conventional commit format

## [1.1.0] - Previous Release

### Features
- AI-powered commit message generation
- Support for Conventional Commits format
- Interactive selection of commit messages
- Technical report generation
- Global API key configuration
- Git diff analysis with ignore patterns

---

## Future Releases

See [IMPROVEMENTS.md](IMPROVEMENTS.md) for planned features and enhancements.

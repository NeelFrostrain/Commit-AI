# TODO - Commit-AI

Current action items and future enhancements for Commit-AI.

---

## 🔥 High Priority

### Version 1.2.0 Release
- [x] Add version command
- [x] Enhance build system with versioning
- [x] Rename project from Commit-Ai-Go to Commit-Ai
- [x] Clean up documentation
- [ ] Update all GitHub references
- [ ] Create v1.2.0 release
- [ ] Test installer on clean systems

### Documentation
- [x] Consolidate redundant docs
- [x] Create comprehensive README
- [x] Add build and release guides
- [ ] Add video tutorial/demo
- [ ] Create troubleshooting guide

---

## 📋 Medium Priority

### Features
- [ ] Add `--dry-run` flag to preview without committing
- [ ] Support for commit message templates
- [ ] Add `--amend` flag to amend last commit
- [ ] Configuration wizard for first-time setup
- [ ] Support for custom AI prompts
- [ ] Add `--stats` flag to show usage statistics

### User Experience
- [ ] Add progress bar for AI requests
- [ ] Colorized diff preview
- [ ] Commit message history/favorites
- [ ] Interactive tutorial mode
- [ ] Shell completion (bash, zsh, fish)

### Testing
- [ ] Add integration tests
- [ ] Test on more platforms
- [ ] Performance benchmarks
- [ ] Load testing with large diffs

---

## 🎯 Low Priority

### Advanced Features
- [ ] Support for multiple AI providers (OpenAI, Anthropic)
- [ ] Local AI model support (Ollama)
- [ ] Commit message linting
- [ ] Git hooks integration
- [ ] Team conventions enforcement
- [ ] Commit message analytics

### Developer Experience
- [ ] GoReleaser integration
- [ ] Automated changelog generation
- [ ] Docker image
- [ ] Homebrew formula
- [ ] Chocolatey package
- [ ] Snap package

### Documentation
- [ ] API documentation
- [ ] Architecture diagrams
- [ ] Performance optimization guide
- [ ] Security best practices
- [ ] Internationalization (i18n)

---

## 🐛 Known Issues

### Minor
- [ ] Long diffs may hit API rate limits
- [ ] Some special characters in commit messages need escaping
- [ ] Verbose mode output could be more structured

### Enhancement Requests
- [ ] Support for conventional commit scopes from config
- [ ] Better handling of merge commits
- [ ] Support for co-authored commits
- [ ] Integration with issue trackers (Jira, GitHub Issues)

---

## 🚀 Future Versions

### v1.3.0 - Enhanced Configuration
- Custom AI prompt templates
- Team-wide configuration sharing
- Commit message validation rules
- Scope and type customization

### v1.4.0 - Advanced Features
- Multiple AI provider support
- Commit message history
- Analytics and insights
- Git hooks integration

### v2.0.0 - Major Overhaul
- Plugin system
- Web UI for configuration
- Team collaboration features
- Enterprise features

---

## 📝 Notes

### Development Workflow
1. Create feature branch
2. Implement changes with tests
3. Update documentation
4. Run `make check` to verify
5. Create pull request
6. Use commit-ai for commit messages!

### Release Workflow
1. Update version in build scripts
2. Update CHANGELOG.md
3. Run `make release` or `.\build.ps1 release`
4. Create git tag
5. Push to GitHub
6. Create GitHub release with binaries
7. Update documentation

---

## 🤝 Contributing

Want to help? Pick an item from this list and:

1. Comment on the related issue (or create one)
2. Fork the repository
3. Create a feature branch
4. Implement the feature
5. Add tests
6. Update documentation
7. Submit a pull request

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

---

## 📊 Progress Tracking

- **v1.2.0**: 90% complete
- **Documentation**: 95% complete
- **Test Coverage**: 85%
- **Platform Support**: 100% (Windows, macOS, Linux)

---

Last Updated: 2024-12-XX

# Release Checklist

Complete checklist for creating a new Commit-AI release.

## Pre-Release

### Code Quality

- [ ] All tests passing (`make test` or `.\build.ps1 test`)
- [ ] Code formatted (`make fmt`)
- [ ] No linter warnings (`make lint`)
- [ ] Test coverage > 85%
- [ ] All diagnostics resolved

### Documentation

- [ ] README.md updated with new features
- [ ] CHANGELOG.md updated with version changes
- [ ] Version numbers updated in:
  - [ ] `build.ps1` ($Version variable)
  - [ ] Git tag
  - [ ] Documentation references
- [ ] All documentation links working
- [ ] Examples tested and verified

### Testing

- [ ] Manual testing on Windows
- [ ] Manual testing on Linux (if available)
- [ ] Manual testing on macOS (if available)
- [ ] Test with real repositories
- [ ] Test installer functionality
- [ ] Test all CLI flags:
  - [ ] `-c` (commit)
  - [ ] `-v` (verbose)
  - [ ] `-y` (yes)
  - [ ] `-m` (model override)
  - [ ] `-h` (help)

## Build Process

### Version Update

```bash
# 1. Update version in build.ps1
# Edit: $Version = "v1.3.0"

# 2. Create git tag
git tag -a v1.3.0 -m "Release v1.3.0: [Brief description]"

# 3. Verify tag
git describe --tags
```

### Build Binaries

```bash
# Windows
.\build.ps1 release

# Unix/Linux/macOS
make release
```

### Verify Builds

- [ ] Production binary works: `.\bin\commit-ai.exe --help`
- [ ] Installer works: `.\bin\install-commit-ai.exe --help`
- [ ] All platform binaries created in `dist/`:
  - [ ] commit-ai-windows-amd64.exe
  - [ ] commit-ai-linux-amd64
  - [ ] commit-ai-linux-arm64
  - [ ] commit-ai-darwin-amd64
  - [ ] commit-ai-darwin-arm64

### Generate Checksums

```bash
# Windows
cd dist
Get-FileHash * -Algorithm SHA256 | Format-List > checksums.txt

# Unix/Linux/macOS
cd dist
sha256sum * > checksums.txt
```

## GitHub Release

### Create Release

1. [ ] Go to https://github.com/NeelFrostrain/Commit-Ai/releases
2. [ ] Click "Draft a new release"
3. [ ] Choose tag: v1.3.0
4. [ ] Release title: "v1.3.0 - [Brief Title]"

### Upload Binaries

- [ ] `bin/commit-ai.exe` (Windows production binary)
- [ ] `bin/install-commit-ai.exe` (Windows installer)
- [ ] `dist/commit-ai-windows-amd64.exe`
- [ ] `dist/commit-ai-linux-amd64`
- [ ] `dist/commit-ai-linux-arm64`
- [ ] `dist/commit-ai-darwin-amd64`
- [ ] `dist/commit-ai-darwin-arm64`
- [ ] `dist/checksums.txt`

### Release Notes Template

```markdown
## 🎉 What's New in v1.3.0

[Brief overview of major changes]

### ✨ New Features
- Feature 1 description
- Feature 2 description

### 🐛 Bug Fixes
- Fix 1 description
- Fix 2 description

### 🔧 Improvements
- Improvement 1 description
- Improvement 2 description

### 📚 Documentation
- Documentation updates

### 🔨 Technical Changes
- Technical change 1
- Technical change 2

## 📦 Installation

### Windows
Download `install-commit-ai.exe` and run it.

### macOS
```bash
# Intel
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/download/v1.3.0/commit-ai-darwin-amd64 -o /usr/local/bin/commit-ai
chmod +x /usr/local/bin/commit-ai

# Apple Silicon
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/download/v1.3.0/commit-ai-darwin-arm64 -o /usr/local/bin/commit-ai
chmod +x /usr/local/bin/commit-ai
```

### Linux
```bash
# AMD64
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/download/v1.3.0/commit-ai-linux-amd64 -o ~/.local/bin/commit-ai
chmod +x ~/.local/bin/commit-ai

# ARM64
curl -L https://github.com/NeelFrostrain/Commit-Ai/releases/download/v1.3.0/commit-ai-linux-arm64 -o ~/.local/bin/commit-ai
chmod +x ~/.local/bin/commit-ai
```

## 📊 Stats
- **Files Changed**: X
- **Insertions**: +X
- **Deletions**: -X
- **Test Coverage**: X%
- **Binary Size**: ~X MB

## 🙏 Contributors
Thanks to all contributors who made this release possible!

## 🔗 Links
- [Full Changelog](https://github.com/NeelFrostrain/Commit-Ai/compare/v1.2.0...v1.3.0)
- [Documentation](https://github.com/NeelFrostrain/Commit-Ai/tree/main/docs)
- [Issues](https://github.com/NeelFrostrain/Commit-Ai/issues)
```

### Publish Release

- [ ] Review all information
- [ ] Check "Set as latest release"
- [ ] Click "Publish release"

## Post-Release

### Verification

- [ ] Release appears on GitHub
- [ ] All binaries downloadable
- [ ] Checksums match
- [ ] Installation instructions work
- [ ] README badges updated (if needed)

### Communication

- [ ] Announce on project README
- [ ] Update project description (if needed)
- [ ] Share on social media (optional)
- [ ] Notify users (if applicable)

### Git Cleanup

```bash
# Push tags
git push origin v1.3.0

# Push any remaining changes
git push origin main

# Clean local build artifacts (optional)
make clean
```

### Update Development

- [ ] Bump version for next development cycle
- [ ] Create new TODO items for next release
- [ ] Update project roadmap

## Rollback Plan

If issues are discovered after release:

### Minor Issues

1. [ ] Document workaround in release notes
2. [ ] Plan fix for next patch release

### Critical Issues

1. [ ] Delete release from GitHub
2. [ ] Delete git tag: `git tag -d v1.3.0 && git push origin :refs/tags/v1.3.0`
3. [ ] Fix issue
4. [ ] Restart release process

## Version Numbering

Follow Semantic Versioning (SemVer):

- **Major** (v2.0.0): Breaking changes
- **Minor** (v1.3.0): New features, backward compatible
- **Patch** (v1.2.1): Bug fixes, backward compatible

### Examples

- `v1.2.0` → `v1.3.0`: New features added
- `v1.2.0` → `v1.2.1`: Bug fixes only
- `v1.2.0` → `v2.0.0`: Breaking API changes

## Release Frequency

Recommended schedule:

- **Major**: Every 6-12 months
- **Minor**: Every 1-2 months
- **Patch**: As needed for critical bugs

## Automation (Future)

Consider automating with:

- **GoReleaser**: Automated release builds
- **GitHub Actions**: CI/CD pipeline
- **Semantic Release**: Automated versioning

Example GoReleaser config:

```yaml
# .goreleaser.yaml
builds:
  - env:
      - CGO_ENABLED=0
    goos:
      - linux
      - windows
      - darwin
    goarch:
      - amd64
      - arm64
    ldflags:
      - -s -w
      - -X main.Version={{.Version}}

archives:
  - format: tar.gz
    format_overrides:
      - goos: windows
        format: zip

checksum:
  name_template: 'checksums.txt'

release:
  github:
    owner: NeelFrostrain
    name: Commit-Ai
```

## Emergency Contacts

- **Maintainer**: [Your contact info]
- **Backup**: [Backup contact]
- **GitHub**: https://github.com/NeelFrostrain/Commit-Ai/issues

## Notes

- Always test on a clean system before release
- Keep release notes clear and user-focused
- Include migration guides for breaking changes
- Maintain backward compatibility when possible
- Document all known issues in release notes

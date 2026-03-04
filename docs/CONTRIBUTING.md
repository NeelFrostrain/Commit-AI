# Contributing Guide

Thank you for your interest in contributing to Commit-AI!

## Code of Conduct

Be respectful, inclusive, and professional. We're all here to make Commit-AI better.

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git
- Make (optional)

### Setup Development Environment

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
```

## Development Workflow

### 1. Create Feature Branch

```bash
git checkout -b feature/your-feature-name
```

### 2. Make Changes

- Write clean, idiomatic Go code
- Follow existing code style
- Add tests for new functionality
- Update documentation

### 3. Test Your Changes

```bash
# Run tests
make test

# Run with coverage
make test-coverage

# Run linter
make lint

# Format code
make fmt
```

### 4. Commit Changes

```bash
# Use Commit-AI to generate commit message
commit-ai -c

# Or manually commit
git commit -m "feat(module): description"
```

### 5. Push and Create Pull Request

```bash
git push origin feature/your-feature-name
```

Then create a Pull Request on GitHub.

## Code Style

### Go Code Style

- Follow [Effective Go](https://golang.org/doc/effective_go)
- Use `gofmt` for formatting
- Use `golint` for linting
- Write clear, descriptive variable names
- Add comments for exported functions

### Example

```go
// ParseResponse extracts title and report from AI response
func ParseResponse(response string) (string, string) {
    // Implementation
}
```

## Testing

### Write Tests

```go
func TestMyFunction(t *testing.T) {
    // Arrange
    input := "test"
    
    // Act
    result := MyFunction(input)
    
    // Assert
    if result != "expected" {
        t.Errorf("Expected 'expected', got '%s'", result)
    }
}
```

### Run Tests

```bash
# Run all tests
make test

# Run specific test
go test -run TestMyFunction ./internal/ai

# Run with coverage
make test-coverage
```

### Coverage Requirements

- Aim for 85%+ coverage
- Test edge cases
- Test error conditions

## Documentation

### Update Documentation

- Update relevant `.md` files in `docs/`
- Update README.md if needed
- Add examples if applicable
- Keep documentation up-to-date

### Documentation Style

- Use clear, concise language
- Include code examples
- Add links to related docs
- Use proper markdown formatting

## Commit Message Format

Follow Conventional Commits:

```
type(scope): description

CATEGORY:
- Detailed change 1
- Detailed change 2

TECHNICAL DETAILS:
- Files changed: X files
- Key changes: description

IMPACT:
- Impact 1
- Impact 2
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation
- `style`: Code style
- `refactor`: Code refactoring
- `perf`: Performance improvement
- `test`: Tests
- `chore`: Maintenance

### Example

```
feat(ai): improve prompt engineering for better analysis

IMPROVEMENTS:
- Enhanced AI prompt with critical instructions
- Added warnings against generic features
- Improved prompt clarity

TECHNICAL DETAILS:
- 1 file changed: 50 insertions(+), 20 deletions(-)

IMPACT:
- Better commit message accuracy
- Improved developer experience
```

## Pull Request Process

### Before Submitting

1. ✅ Tests pass: `make test`
2. ✅ Code formatted: `make fmt`
3. ✅ Linter passes: `make lint`
4. ✅ Documentation updated
5. ✅ Commit messages follow format

### PR Description

Include:

- **What**: What does this PR do?
- **Why**: Why is this change needed?
- **How**: How does it work?
- **Testing**: How was it tested?
- **Related**: Related issues or PRs

### Example PR Description

```markdown
## What
Adds emoji support to commit messages

## Why
Improves commit history readability and visual appeal

## How
- Added `-e/--emoji` flag
- Enhanced AI prompt with emoji instructions
- Updated validation to accept emoji-prefixed commits

## Testing
- Tested with various commit types
- Verified emoji display in different terminals
- All tests passing

## Related
Closes #123
```

## Review Process

### Code Review

- Maintainers will review your PR
- Provide feedback and suggestions
- Request changes if needed
- Approve when ready

### Addressing Feedback

1. Make requested changes
2. Commit with descriptive message
3. Push changes
4. Request re-review

## Reporting Issues

### Bug Reports

Include:

- **Version**: `commit-ai version`
- **System**: `uname -a`
- **Error**: Full error message
- **Steps**: How to reproduce
- **Expected**: What should happen
- **Actual**: What actually happened

### Feature Requests

Include:

- **Description**: What feature?
- **Use Case**: Why is it needed?
- **Example**: How would it work?
- **Alternative**: Any alternatives?

## Areas for Contribution

### Code

- Bug fixes
- Performance improvements
- New features
- Code refactoring
- Test coverage

### Documentation

- Improve existing docs
- Add new guides
- Fix typos
- Add examples
- Translate docs

### Community

- Answer questions
- Help with issues
- Share feedback
- Promote project

## Development Tips

### Useful Commands

```bash
# Build
make build

# Build production
make build-prod

# Run tests
make test

# Generate coverage
make test-coverage

# Format code
make fmt

# Run linter
make lint

# Run all checks
make check

# Clean build artifacts
make clean
```

### Debugging

```bash
# Run with verbose output
./commit-ai -v

# Run with specific model
./commit-ai -m llama-3.1-70b-versatile

# Check configuration
echo $GROQ_API_KEY
```

### Testing Locally

```bash
# Build locally
make build

# Test with real git repo
cd /tmp
mkdir test-repo
cd test-repo
git init
echo "test" > file.txt
git add .
/path/to/commit-ai -v
```

## Release Process

### Version Numbering

Follow [Semantic Versioning](https://semver.org/):
- MAJOR: Breaking changes
- MINOR: New features
- PATCH: Bug fixes

### Release Steps

1. Update version in code
2. Update CHANGELOG.md
3. Create git tag
4. Build release binaries
5. Create GitHub release
6. Upload binaries

## Questions?

- **GitHub Issues**: https://github.com/NeelFrostrain/Commit-Ai/issues
- **GitHub Discussions**: https://github.com/NeelFrostrain/Commit-Ai/discussions
- **Documentation**: https://github.com/NeelFrostrain/Commit-Ai/tree/main/docs

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to Commit-AI! 🎉

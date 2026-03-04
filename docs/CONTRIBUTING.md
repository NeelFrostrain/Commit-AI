# Contributing to Commit-AI

Thank you for your interest in contributing to Commit-AI! This document provides guidelines and instructions for contributing.

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or higher
- Git
- A Groq API key (for testing)

### Setup Development Environment

1. Fork and clone the repository:
```bash
git clone https://github.com/YOUR_USERNAME/Commit-Ai.git
cd Commit-Ai
```

2. Install dependencies:
```bash
go mod download
```

3. Set up your API key:
```bash
echo "GROQ_API_KEY=your_key_here" > .env
```

4. Build and test:
```bash
make build
make test
```

## 📝 Development Workflow

### Making Changes

1. Create a new branch:
```bash
git checkout -b feature/your-feature-name
```

2. Make your changes following our coding standards

3. Run tests and linting:
```bash
make check
```

4. Commit your changes (use commit-ai if you have it installed!):
```bash
git add .
commit-ai -c
```

5. Push and create a pull request

### Code Style

- Follow standard Go conventions
- Run `go fmt` before committing
- Use meaningful variable and function names
- Add comments for complex logic
- Keep functions small and focused

### Testing

- Write tests for new features
- Maintain or improve code coverage
- Run `make test-coverage` to check coverage
- Test on multiple platforms if possible

## 🏗️ Project Structure

```
Commit-Ai/
├── cmd/              # Command-line interface
│   └── root.go       # Main command logic
├── internal/         # Internal packages
│   ├── ai/          # AI integration and parsing
│   ├── config/      # Configuration management
│   └── git/         # Git operations
├── scripts/         # Build and installation scripts
└── main.go          # Entry point
```

## 🐛 Reporting Bugs

When reporting bugs, please include:

1. Your operating system and Go version
2. Steps to reproduce the issue
3. Expected vs actual behavior
4. Relevant error messages or logs
5. Sample git diff (if applicable)

## 💡 Suggesting Features

We welcome feature suggestions! Please:

1. Check existing issues first
2. Describe the use case
3. Explain why it would be useful
4. Provide examples if possible

## 🔍 Code Review Process

1. All submissions require review
2. Maintainers will review within 1-2 weeks
3. Address feedback promptly
4. Once approved, maintainers will merge

## 📋 Commit Message Guidelines

We use Conventional Commits (of course!):

- `feat:` New features
- `fix:` Bug fixes
- `docs:` Documentation changes
- `refactor:` Code refactoring
- `test:` Test additions or changes
- `chore:` Maintenance tasks

Examples:
```
feat(ai): add support for custom models
fix(git): handle empty repositories correctly
docs(readme): update installation instructions
```

## 🎯 Areas for Contribution

Looking for where to start? Check out:

- Issues labeled `good first issue`
- Issues labeled `help wanted`
- Documentation improvements
- Test coverage improvements
- Performance optimizations

## 📜 License

By contributing, you agree that your contributions will be licensed under the MIT License.

## 🤝 Code of Conduct

- Be respectful and inclusive
- Welcome newcomers
- Focus on constructive feedback
- Help others learn and grow

## 💬 Getting Help

- Open an issue for bugs or features
- Start a discussion for questions
- Check existing issues and discussions first

## 🙏 Thank You

Your contributions make Commit-AI better for everyone. We appreciate your time and effort!

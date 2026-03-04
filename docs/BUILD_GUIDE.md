# Build Guide

Complete guide for building Commit-AI from source and creating releases.

## Prerequisites

- Go 1.21 or higher
- Git
- Make (for Unix-like systems)
- PowerShell (for Windows)

## Quick Build

### Development Build

```bash
# Unix/Linux/macOS
make build

# Windows
.\build.ps1 build
```

This creates `commit-ai.exe` in the project root for quick testing.

### Production Build

```bash
# Unix/Linux/macOS
make build-prod

# Windows
.\build.ps1 build-prod
```

This creates an optimized binary in `bin/commit-ai.exe` with:
- Stripped debug symbols (`-s -w`)
- Trimmed paths for reproducible builds
- Version information embedded

## Build Targets

### Windows (PowerShell)

```powershell
# Show all available targets
.\build.ps1 help

# Development build
.\build.ps1 build

# Production build to bin/
.\build.ps1 build-prod

# Build installer
.\build.ps1 installer

# Build for all platforms
.\build.ps1 build-all

# Run tests
.\build.ps1 test

# Clean artifacts
.\build.ps1 clean

# Full release build
.\build.ps1 release
```

### Unix/Linux/macOS (Make)

```bash
# Show all available targets
make help

# Development build
make build

# Production build to bin/
make build-prod

# Build installer
make installer

# Build for all platforms
make build-all

# Run tests
make test

# Test with coverage
make test-coverage

# Run linter
make lint

# Format code
make fmt

# Clean artifacts
make clean

# Full release build
make release
```

## Platform Builds

### Supported Platforms

The project supports building for:

- Windows (amd64)
- Linux (amd64, arm64)
- macOS (amd64, arm64/Apple Silicon)

### Cross-Platform Build

```bash
# Unix/Linux/macOS
make build-all

# Windows
.\build.ps1 build-all
```

This creates binaries in `dist/` for all platforms:
- `commit-ai-windows-amd64.exe`
- `commit-ai-linux-amd64`
- `commit-ai-linux-arm64`
- `commit-ai-darwin-amd64`
- `commit-ai-darwin-arm64`

## Release Process

### 1. Prepare Release

```bash
# Update version in files
# - build.ps1: $Version = "v1.3.0"
# - Makefile: VERSION variable (auto from git tags)

# Create git tag
git tag -a v1.3.0 -m "Release v1.3.0"
git push origin v1.3.0
```

### 2. Build Release

```bash
# Unix/Linux/macOS
make release

# Windows
.\build.ps1 release
```

This will:
1. Clean previous builds
2. Run all tests
3. Build production binary (`bin/commit-ai.exe`)
4. Build installer (`bin/install-commit-ai.exe`)
5. Build all platform binaries (`dist/`)

### 3. Verify Builds

```bash
# Test production binary
./bin/commit-ai.exe --help

# Test installer
./bin/install-commit-ai.exe --help

# Check all platform builds exist
ls -la dist/
```

### 4. Create GitHub Release

1. Go to GitHub Releases page
2. Click "Draft a new release"
3. Select the tag (e.g., v1.3.0)
4. Upload binaries:
   - `bin/commit-ai.exe` (Windows production)
   - `bin/install-commit-ai.exe` (Windows installer)
   - `dist/commit-ai-windows-amd64.exe`
   - `dist/commit-ai-linux-amd64`
   - `dist/commit-ai-linux-arm64`
   - `dist/commit-ai-darwin-amd64`
   - `dist/commit-ai-darwin-arm64`
5. Write release notes
6. Publish release

## Build Flags

### LDFLAGS

The build uses these linker flags:

- `-X main.Version=$(VERSION)` - Embeds version string
- `-s` - Strips symbol table
- `-w` - Strips DWARF debug info

Result: Smaller binary size (typically 30-40% reduction)

### GOFLAGS

- `-trimpath` - Removes absolute paths for reproducible builds

## Testing

### Run Tests

```bash
# Unix/Linux/macOS
make test

# Windows
.\build.ps1 test
```

### Coverage Report

```bash
# Unix/Linux/macOS only
make test-coverage

# Opens coverage.html in browser
```

### Manual Testing

```bash
# Build and test
make build
./commit-ai.exe -v

# Test with real changes
git add .
./commit-ai.exe -v
```

## Troubleshooting

### "go: command not found"

Install Go from https://golang.org/dl/

### "make: command not found" (Windows)

Use PowerShell build script instead:
```powershell
.\build.ps1 build
```

### Build fails with "package not found"

```bash
go mod download
go mod tidy
```

### Cross-compilation fails

Ensure you have the target platform toolchain:
```bash
# For CGO-enabled builds
CGO_ENABLED=0 make build-all
```

### Version not embedded

Check that git tags exist:
```bash
git describe --tags --always
```

## Development Workflow

### 1. Make Changes

```bash
# Edit code
vim cmd/root.go

# Format code
make fmt
```

### 2. Test Changes

```bash
# Run tests
make test

# Build and test manually
make build
./commit-ai.exe -v
```

### 3. Commit Changes

```bash
# Use commit-ai itself!
./commit-ai.exe -c
```

### 4. Create Pull Request

```bash
git push origin feature/my-feature
# Open PR on GitHub
```

## CI/CD

The project uses GitHub Actions for automated builds and tests.

### Workflow Triggers

- Push to main branch
- Pull requests
- Tag creation (releases)

### What Gets Built

- All platform binaries
- Test coverage reports
- Release artifacts (on tags)

## Binary Sizes

Typical binary sizes (with `-s -w` flags):

- Windows: ~8-10 MB
- Linux: ~8-10 MB
- macOS: ~8-10 MB

Without optimization flags: ~12-15 MB

## Security

### Code Signing (Future)

For production releases, consider:

- Windows: Authenticode signing
- macOS: Apple Developer ID signing
- Linux: GPG signatures

### Checksums

Generate checksums for releases:

```bash
# Unix/Linux/macOS
cd dist
sha256sum * > checksums.txt

# Windows
cd dist
Get-FileHash * -Algorithm SHA256 > checksums.txt
```

## Performance

### Build Time

Typical build times on modern hardware:

- Single platform: 5-10 seconds
- All platforms: 30-60 seconds
- Full release (with tests): 1-2 minutes

### Optimization

For faster builds during development:

```bash
# Skip tests
make build

# Use cached builds
go build -i
```

## Additional Resources

- [Go Build Documentation](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies)
- [Cross Compilation](https://golang.org/doc/install/source#environment)
- [Release Best Practices](https://goreleaser.com/quick-start/)

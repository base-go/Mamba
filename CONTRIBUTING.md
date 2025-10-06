# Contributing to Mamba

Thank you for your interest in contributing to Mamba.

## Code of Conduct

By participating in this project, you agree to abide by our Code of Conduct (see CODE_OF_CONDUCT.md).

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the existing issues to avoid duplicates. When creating a bug report, include:

- **Clear title and description**
- **Steps to reproduce** the behavior
- **Expected behavior** and what actually happened
- **Code samples** or test cases if applicable
- **Go version** and operating system
- **Mamba version**

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, include:

- **Clear title and description**
- **Use case** - why is this enhancement useful?
- **Possible implementation** - if you have ideas
- **Examples** - code samples showing how it would work

### Pull Requests

1. **Fork the repo** and create your branch from `main`
2. **Make your changes** following our coding standards
3. **Add tests** if you've added code that should be tested
4. **Update documentation** if you've changed APIs
5. **Run tests** to ensure they all pass
6. **Format your code** with `gofmt`
7. **Submit a pull request**

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Git

### Getting Started

1. Fork and clone the repository:
   ```bash
   git clone https://github.com/YOUR_USERNAME/mamba.git
   cd mamba
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run tests:
   ```bash
   go test ./...
   ```

4. Build the project:
   ```bash
   go build ./...
   ```

5. Try the demo:
   ```bash
   cd examples/basic
   go run main.go
   ```

## Coding Standards

### Go Style

- Follow the [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use `gofmt` to format your code
- Run `go vet` to check for common mistakes
- Keep functions small and focused
- Write descriptive variable and function names

### Documentation

- Add godoc comments to all exported types, functions, and methods
- Include examples in godoc comments when helpful
- Keep README.md up to date
- Update CHANGELOG.md following [Keep a Changelog](https://keepachangelog.com/)

### Testing

- Write table-driven tests when applicable
- Aim for high test coverage
- Test edge cases and error conditions
- Use descriptive test names: `TestFunctionName_Scenario_ExpectedBehavior`

### Commits

- Write clear, descriptive commit messages
- Use present tense ("Add feature" not "Added feature")
- Reference issues and PRs when relevant
- Keep commits focused on a single change

Example:
```
Add spinner support for long-running operations

- Implement WithSpinner helper function
- Add spinner model with bubbletea
- Update documentation with spinner examples

Fixes #123
```

## Project Structure

```
mamba/
├── command.go           # Core Command struct and methods
├── help.go             # Modern help generation
├── doc.go              # Package documentation
├── pkg/
│   ├── style/          # Terminal styling utilities
│   ├── interactive/    # Interactive prompt helpers
│   └── spinner/        # Spinner and progress bar components
├── examples/
│   └── basic/          # Demo application
├── .github/
│   └── workflows/      # CI/CD workflows
└── README.md
```

## Adding New Features

When adding new features:

1. **Maintain Cobra compatibility** - Ensure existing Cobra code still works
2. **Add to the demo** - Include examples in `examples/basic/main.go`
3. **Write tests** - Add comprehensive test coverage
4. **Document it** - Update README.md and add godoc comments
5. **Update CHANGELOG** - Add entry under [Unreleased]

## Maintaining Cobra Compatibility

Mamba is a **100% drop-in replacement** for Cobra. When making changes:

- All Cobra APIs must continue to work
- Behavior should match Cobra unless enhancing UX
- New features should be additive, not breaking
- Test with real Cobra code to verify compatibility

## Release Process

Releases are managed by project maintainers:

1. Update CHANGELOG.md with release notes
2. Tag the release: `git tag -a v1.x.x -m "Release v1.x.x"`
3. Push the tag: `git push origin v1.x.x`
4. GitHub Actions will build and publish

## Questions?

Feel free to open an issue for questions or join discussions in GitHub Discussions.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

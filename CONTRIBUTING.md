
# Contributing to Filer

Thank you for your interest in contributing to Filer! This document provides guidelines and instructions for contributing.

## Code of Conduct

By participating in this project, you agree to maintain a respectful and inclusive environment for everyone.

## How to Contribute

### Reporting Bugs

Before creating bug reports, please check existing issues to avoid duplicates. When creating a bug report, include:

- **Clear title and description**
- **Steps to reproduce** the issue
- **Expected behavior** vs actual behavior
- **Environment details** (OS, Go version, etc.)
- **Code samples** or error messages if applicable

### Suggesting Enhancements

Enhancement suggestions are welcome! Please provide:

- **Clear use case** for the enhancement
- **Detailed description** of the proposed functionality
- **Examples** of how it would work
- **Potential implementation approach** (optional)

### Pull Requests

1. **Fork the repository** and create your branch from `filer-cli-push`
2. **Make your changes** following the coding standards below
3. **Add tests** if applicable
4. **Update documentation** as needed
5. **Ensure tests pass** with `go test ./...`
6. **Run linters** with `go vet ./...`
7. **Commit with clear messages** following conventional commits
8. **Submit the pull request**

## Development Setup

### Prerequisites

- Go 1.19 or higher
- Git
- Make (optional, for using Makefile commands)

### Setting Up Your Development Environment

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/Go-command-Line.git
cd Go-command-Line

# Add upstream remote
git remote add upstream https://github.com/PranshuS007/Go-command-Line.git

# Create a feature branch
git checkout -b feature/your-feature-name

# Install dependencies
go mod download

# Build the project
go build -o filer
```

## Coding Standards

### Go Style Guide

- Follow the [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use `gofmt` to format your code
- Run `go vet` to catch common mistakes
- Use meaningful variable and function names
- Add comments for exported functions and types

### Code Organization

```go
// Package documentation
package packagename

import (
    // Standard library imports
    "fmt"
    "os"
    
    // Third-party imports
    "github.com/spf13/cobra"
    
    // Local imports
    "github.com/user/filer/internal/models"
)

// Exported function with documentation
// FunctionName does something useful.
// It takes parameter x and returns y.
func FunctionName(x int) int {
    // Implementation
    return x * 2
}
```

### Error Handling

Always handle errors appropriately:

```go
// Good
result, err := someOperation()
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}

// Avoid
result, _ := someOperation() // Don't ignore errors
```

### Testing

- Write unit tests for new functionality
- Aim for good test coverage
- Use table-driven tests where appropriate
- Test edge cases and error conditions

Example test structure:

```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name    string
        input   int
        want    int
        wantErr bool
    }{
        {"positive number", 5, 10, false},
        {"zero", 0, 0, false},
        {"negative number", -5, -10, false},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := FunctionName(tt.input)
            if got != tt.want {
                t.Errorf("FunctionName() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Commit Messages

Follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>(<scope>): <subject>

<body>

<footer>
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

Examples:
```
feat(search): add support for regex patterns

Add regex pattern matching to the search command, allowing users
to use more complex search patterns beyond simple glob matching.

Closes #123
```

```
fix(list): handle permission errors gracefully

Previously, the list command would crash when encountering files
without read permissions. Now it logs a warning and continues.

Fixes #456
```

## Documentation

- Update README.md for user-facing changes
- Add inline comments for complex logic
- Update command help text if adding/modifying commands
- Include examples in documentation

## Review Process

1. All submissions require review before merging
2. Reviewers will check:
   - Code quality and style
   - Test coverage
   - Documentation completeness
   - Functionality and correctness
3. Address review feedback promptly
4. Once approved, maintainers will merge your PR

## Questions?

If you have questions about contributing:

- Open an issue with the `question` label
- Check existing documentation
- Review closed issues for similar questions

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

Thank you for contributing to Filer! 🎉

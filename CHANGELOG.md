
# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Comprehensive README.md with detailed documentation
- LICENSE file (MIT License)
- CONTRIBUTING.md with contribution guidelines
- GitHub Actions CI/CD workflow
- Makefile for common development tasks
- Unit tests for models package
- CHANGELOG.md for tracking changes
- Better error handling throughout the codebase
- Improved code documentation and comments

### Changed
- Enhanced .gitignore with more comprehensive exclusions
- Improved organize command with better user confirmation flow
- Better safety checks in file operations

### Removed
- Removed unrelated addingOps.go file that didn't belong in the project

### Fixed
- Fixed potential panic in hidden file detection
- Completed organize command implementation
- Improved error messages throughout the application

## [1.0.0] - 2025-09-24

### Added
- Initial release of Filer CLI tool
- List command for displaying files and directories
- Search command for finding files with various criteria
- Stats command for directory analysis
- Organize command for automatic file organization
- Support for multiple output formats (table, JSON, CSV)
- Recursive file operations
- File filtering by size, extension, and date
- Comprehensive sorting options

[Unreleased]: https://github.com/PranshuS007/Go-command-Line/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/PranshuS007/Go-command-Line/releases/tag/v1.0.0

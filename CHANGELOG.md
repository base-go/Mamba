# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2025-01-04

### Added
- Initial release of Mamba
- 100% drop-in replacement for Cobra with the same API
- Modern styled output with lipgloss integration
  - `PrintSuccess()`, `PrintError()`, `PrintWarning()`, `PrintInfo()`
  - `PrintHeader()`, `PrintSubHeader()`, `PrintBullet()`
  - `PrintCode()`, `PrintBox()` for formatted output
- Interactive prompts with charmbracelet/huh
  - `AskString()` - text input
  - `AskConfirm()` - yes/no confirmation
  - `AskSelect()` - single selection
  - `AskMultiSelect()` - multiple selection
- Loading spinners for long-running operations
  - `WithSpinner()` helper function
  - Automatic success/error indication
- Progress bars for batch operations
  - `WithProgress()` helper function
  - Visual progress tracking
- Modern styled help messages
  - Colored command names, flags, and descriptions
  - Enhanced formatting for better readability
  - Automatic `-h` and `--help` flag handling
- Style package for custom formatting
  - Pre-defined color schemes
  - Text styling utilities (bold, italic, underline, dim)
  - Box and frame rendering
- Full Cobra API compatibility
  - Command structure (Use, Short, Long, Example)
  - Subcommands (AddCommand, RemoveCommand)
  - Flags (local, persistent, shorthand via pflag)
  - Lifecycle hooks (PreRun, PostRun, PersistentPreRun, etc.)
  - Argument validators (NoArgs, ExactArgs, MinimumNArgs, etc.)
  - Error handling (RunE, PreRunE, PostRunE)
  - Context support
- Comprehensive documentation
  - Package-level godoc
  - Example applications
  - Migration guide from Cobra
- Demo application showcasing all features

### Fixed
- Help flag (`-h`, `--help`) now works correctly for subcommands
- Argument validation properly bypassed when help is requested
- Flag parsing integrated with pflag error handling

[Unreleased]: https://github.com/base-go/mamba/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/base-go/mamba/releases/tag/v1.0.0

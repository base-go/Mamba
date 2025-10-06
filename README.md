# Mamba

[![Go Reference](https://pkg.go.dev/badge/github.com/base-go/mamba.svg)](https://pkg.go.dev/github.com/base-go/mamba)
[![Go Report Card](https://goreportcard.com/badge/github.com/base-go/mamba)](https://goreportcard.com/report/github.com/base-go/mamba)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A modern, **drop-in replacement** for [Cobra](https://github.com/spf13/cobra) with enhanced terminal features including beautiful colored output, interactive prompts, loading spinners, and progress bars.

> **Fast, elegant, and powerful** - just like the snake.

## Features

**Core Cobra Compatibility:**
- **Subcommands & Nesting** - Full support for complex CLI hierarchies
- **POSIX Flags** - Short & long flags via [pflag](https://github.com/spf13/pflag)
- **Command Aliases** - Multiple names for the same command
- **Lifecycle Hooks** - PreRun, PostRun, and persistent variants
- **Argument Validation** - Built-in validators (ExactArgs, MinimumNArgs, etc.)
- **Help Generation** - Automatic -h, --help with enhanced styling

**Modern Terminal Features:**
- **Styled Output** - Beautiful colors and formatting using [lipgloss](https://github.com/charmbracelet/lipgloss)
- **Interactive Prompts** - User-friendly input with [huh](https://github.com/charmbracelet/huh)
- **Loading Spinners** - Visual feedback for long-running operations
- **Progress Bars** - Track progress for batch operations
- **Enhanced Help** - Colored, structured help messages

## Demo

See Mamba in action! Run the included demo to experience all features:

```bash
go get github.com/base-go/mamba
cd examples/basic
go run main.go
```

The demo showcases:
- Styled output (success, error, warning, info)
- Loading spinners with animations
- Progress bars for batch operations
- Interactive prompts
- Modern help messages

## Installation

```bash
go get github.com/base-go/mamba
```

## Quick Start

```go
package main

import (
    "github.com/base-go/mamba"
)

func main() {
    rootCmd := &mamba.Command{
        Use:   "myapp",
        Short: "My awesome CLI application",
        Run: func(cmd *mamba.Command, args []string) {
            cmd.PrintSuccess("Welcome to Mamba!")
        },
    }

    rootCmd.Execute()
}
```

## Migration from Cobra

Migrating from Cobra to Mamba is incredibly simple - just change your imports!

### Before (Cobra)

```go
import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "myapp",
    Short: "My application",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello from Cobra")
    },
}
```

### After (Mamba)

```go
import (
    "github.com/base-go/mamba"
)

var rootCmd = &mamba.Command{
    Use:   "myapp",
    Short: "My application",
    Run: func(cmd *mamba.Command, args []string) {
        cmd.PrintSuccess("Hello from Mamba!")
    },
}
```

That's it! Your existing Cobra code will work seamlessly with Mamba, but you now have access to modern terminal features.

## Cobra Compatibility

Mamba implements all core Cobra features:

**Fully Supported:**
- Subcommand-based CLIs (app server, app fetch, etc.)
- POSIX-compliant flags (short & long versions via pflag)
- Nested subcommands
- Global, local, and cascading flags
- Automatic help generation with -h, --help
- Command aliases
- Lifecycle hooks (PreRun, PostRun, etc.)
- Argument validators (ExactArgs, MinimumNArgs, etc.)
- Custom help and usage functions
- Context support

**Not Yet Implemented:**
- Intelligent suggestions ("did you mean...?")
- Shell autocomplete generation (bash, zsh, fish, powershell)
- Man page generation
- Command grouping in help
- Viper integration for config files

## Mamba vs Cobra

| Feature | Cobra | Mamba |
|---------|-------|-------|
| Command structure | Yes | Yes |
| Subcommands & nesting | Yes | Yes |
| POSIX flags (pflag) | Yes | Yes |
| Lifecycle hooks | Yes | Yes |
| Argument validation | Yes | Yes |
| Command aliases | Yes | Yes |
| Help generation | Yes | Yes (Enhanced & Styled) |
| Shell autocomplete | Yes | No (planned) |
| Intelligent suggestions | Yes | No (planned) |
| Man page generation | Yes | No (planned) |
| Viper integration | Optional | No (planned) |
| Colored output | No | Yes |
| Loading spinners | No | Yes |
| Progress bars | No | Yes |
| Interactive prompts | No | Yes |
| Modern terminal UX | No | Yes |

## Enhanced Features

### Styled Output

Mamba provides beautiful, pre-styled output methods:

```go
cmd := &mamba.Command{
    Use: "demo",
    Run: func(cmd *mamba.Command, args []string) {
        cmd.PrintSuccess("Operation completed successfully")
        cmd.PrintError("Something went wrong")
        cmd.PrintWarning("This is a warning")
        cmd.PrintInfo("Here's some information")
        cmd.PrintHeader("Section Header")
        cmd.PrintBullet("Bullet point item")
        cmd.PrintCode("go run main.go")
        cmd.PrintBox("Title", "Important message in a box")
    },
}
```

### Interactive Prompts

Create beautiful interactive prompts with ease:

```go
import "github.com/base-go/mamba/pkg/interactive"

// Simple text input
name, err := interactive.AskString("What's your name?", "Enter your name")

// Yes/No confirmation
confirmed, err := interactive.AskConfirm("Do you want to continue?", true)

// Selection from list
options := []interactive.SelectOption{
    {Key: "red", Value: "Red"},
    {Key: "blue", Value: "Blue"},
    {Key: "green", Value: "Green"},
}
color, err := interactive.AskSelect("Choose a color:", options)

// Multi-selection
selected, err := interactive.AskMultiSelect("Choose features:", options, 0)
```

### Loading Spinners

Show progress for long-running operations:

```go
import "github.com/base-go/mamba/pkg/spinner"

err := spinner.WithSpinner("Processing...", func() error {
    // Your long-running operation here
    time.Sleep(2 * time.Second)
    return nil
})
```

### Progress Bars

Track progress for batch operations:

```go
import "github.com/base-go/mamba/pkg/spinner"

spinner.WithProgress("Processing items...", total, func(update func()) {
    for i := 0; i < total; i++ {
        // Process item
        time.Sleep(100 * time.Millisecond)
        update() // Update progress bar
    }
})
```

### Custom Styling

Use the style package directly for custom formatting:

```go
import "github.com/base-go/mamba/pkg/style"

fmt.Println(style.Success("Success!"))
fmt.Println(style.Error("Error!"))
fmt.Println(style.Bold("Bold text"))
fmt.Println(style.Code("code snippet"))
fmt.Println(style.Header("Header"))
fmt.Println(style.Box("Title", "Content"))
```

## Full Example

See the [examples/basic](examples/basic/main.go) directory for a complete demonstration of all features:

```bash
# Run the example
cd examples/basic
go run main.go

# Or use the pre-built binary
./mamba styled          # View styled output demo
./mamba greet --name Alice  # Try command with flags
./mamba process --count 5   # See loading spinner
./mamba progress --count 10 # See progress bar
./mamba interactive         # Try interactive prompts
./mamba list                # View feature list
```

## API Compatibility

Mamba maintains 100% API compatibility with Cobra. All these features work identically:

- Command structure (`Use`, `Short`, `Long`, `Example`)
- Subcommands (`AddCommand`, `RemoveCommand`)
- Flags (local, persistent, shorthand)
- Lifecycle hooks (`PreRun`, `PostRun`, `PersistentPreRun`, etc.)
- Argument validators (`NoArgs`, `ExactArgs`, `MinimumNArgs`, etc.)
- Error handling (`RunE`, `PreRunE`, `PostRunE`)
- Context support
- Custom help and usage templates

## Advanced Usage

### Disabling Modern Features

If you want to disable modern styling:

```go
disableColors := false
cmd := &mamba.Command{
    Use:          "myapp",
    EnableColors: &disableColors,
}
```

### Custom IO Writers

Like Cobra, Mamba supports custom IO writers:

```go
cmd.SetOutput(customWriter)
cmd.SetErr(customErrWriter)
cmd.SetIn(customReader)
```

### Working with Existing Cobra Projects

For large projects with many files:

1. Update imports globally:
   ```bash
   find . -name "*.go" -type f -exec sed -i '' 's|github.com/spf13/cobra|github.com/base-go/mamba|g' {} +
   ```

2. Run `go mod tidy`:
   ```bash
   go mod tidy
   ```

3. Rebuild your project:
   ```bash
   go build
   ```

That's it! Your application now has modern terminal features.

## Architecture

Mamba is built on top of excellent libraries:

- [spf13/pflag](https://github.com/spf13/pflag) - POSIX-style flag parsing
- [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling
- [charmbracelet/huh](https://github.com/charmbracelet/huh) - Interactive forms
- [charmbracelet/bubbles](https://github.com/charmbracelet/bubbles) - TUI components

## Roadmap

Planned features to achieve full Cobra parity:

- **v1.1.0**
  - Intelligent command suggestions ("did you mean...?")
  - Custom help templates and functions
  - Command grouping in help output

- **v1.2.0**
  - Shell completion generation (bash, zsh, fish, powershell)
  - Man page generation
  - Enhanced error messages with suggestions

- **v1.3.0**
  - Viper integration for configuration management
  - Advanced validation and hooks
  - Performance optimizations

Want to contribute? Check out our [issues](https://github.com/base-go/mamba/issues) or submit a PR!

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

MIT License - see LICENSE file for details

## Why Mamba?

Cobra is excellent, but modern CLI applications deserve modern UX. Mamba brings:

- **Better UX** - Colors, spinners, and interactive prompts make CLIs more user-friendly
- **Zero Migration Cost** - Drop-in replacement means you can adopt it incrementally
- **Progressive Enhancement** - Use as much or as little of the modern features as you want
- **Battle-tested** - Built on proven libraries from the Charm ecosystem

Named after the black mamba snake - fast, elegant, and powerful.

## Credits

Mamba is inspired by [Cobra](https://github.com/spf13/cobra) and uses the fantastic [Charm](https://charm.sh/) libraries for terminal UI.

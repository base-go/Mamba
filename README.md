# 🐍 Mamba

A modern, drop-in replacement for [Cobra](https://github.com/spf13/cobra) with enhanced terminal features including beautiful colored output, interactive prompts, loading spinners, and progress bars.

## Features

- ✅ **100% Cobra Compatible** - Drop-in replacement with the same API
- 🎨 **Modern Styled Output** - Beautiful colors and formatting using [lipgloss](https://github.com/charmbracelet/lipgloss)
- 💬 **Interactive Prompts** - User-friendly input prompts with [huh](https://github.com/charmbracelet/huh)
- ⏳ **Loading Spinners** - Visual feedback for long-running operations
- 📊 **Progress Bars** - Track progress for batch operations
- 🎯 **Styled Help Messages** - Enhanced help output with colors and structure
- 🔧 **Full Flag Support** - Compatible with [pflag](https://github.com/spf13/pflag)

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

- ✅ Command structure (`Use`, `Short`, `Long`, `Example`)
- ✅ Subcommands (`AddCommand`, `RemoveCommand`)
- ✅ Flags (local, persistent, shorthand)
- ✅ Lifecycle hooks (`PreRun`, `PostRun`, `PersistentPreRun`, etc.)
- ✅ Argument validators (`NoArgs`, `ExactArgs`, `MinimumNArgs`, etc.)
- ✅ Error handling (`RunE`, `PreRunE`, `PostRunE`)
- ✅ Context support
- ✅ Custom help and usage templates

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

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details

## Why Mamba?

Cobra is excellent, but modern CLI applications deserve modern UX. Mamba brings:

- **Better UX** - Colors, spinners, and interactive prompts make CLIs more user-friendly
- **Zero Migration Cost** - Drop-in replacement means you can adopt it incrementally
- **Progressive Enhancement** - Use as much or as little of the modern features as you want
- **Battle-tested** - Built on proven libraries from the Charm ecosystem

Named after the black mamba snake - fast, elegant, and powerful! 🐍

## Credits

Mamba is inspired by [Cobra](https://github.com/spf13/cobra) and uses the fantastic [Charm](https://charm.sh/) libraries for terminal UI.

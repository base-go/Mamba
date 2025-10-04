package mamba

import (
	"fmt"
	"strings"

	"github.com/base-go/mamba/pkg/style"
	"github.com/spf13/pflag"
)

// ModernHelp generates a modern styled help message
func (c *Command) ModernHelp() string {
	var sb strings.Builder

	// Header
	if c.Long != "" {
		sb.WriteString(style.Header(c.Name()))
		sb.WriteString("\n\n")
		sb.WriteString(style.Muted(c.Long))
		sb.WriteString("\n\n")
	} else if c.Short != "" {
		sb.WriteString(style.Header(c.Name()))
		sb.WriteString("\n\n")
		sb.WriteString(style.Muted(c.Short))
		sb.WriteString("\n\n")
	}

	// Usage
	sb.WriteString(style.SubHeader("Usage"))
	sb.WriteString("\n  ")
	sb.WriteString(style.Command(c.UseLine()))
	sb.WriteString("\n\n")

	// Examples
	if c.Example != "" {
		sb.WriteString(style.SubHeader("Examples"))
		sb.WriteString("\n")
		examples := strings.Split(c.Example, "\n")
		for _, example := range examples {
			if strings.TrimSpace(example) != "" {
				sb.WriteString("  ")
				sb.WriteString(style.Dim(example))
				sb.WriteString("\n")
			}
		}
		sb.WriteString("\n")
	}

	// Available Commands
	if len(c.commands) > 0 {
		sb.WriteString(style.SubHeader("Available Commands"))
		sb.WriteString("\n")

		maxLen := 0
		visibleCmds := []*Command{}
		for _, cmd := range c.commands {
			if !cmd.Hidden {
				visibleCmds = append(visibleCmds, cmd)
				if len(cmd.Name()) > maxLen {
					maxLen = len(cmd.Name())
				}
			}
		}

		for _, cmd := range visibleCmds {
			sb.WriteString("  ")
			sb.WriteString(style.Command(fmt.Sprintf("%-*s", maxLen, cmd.Name())))
			sb.WriteString("  ")
			sb.WriteString(style.Muted(cmd.Short))
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}

	// Flags
	if c.Flags().HasFlags() {
		sb.WriteString(style.SubHeader("Flags"))
		sb.WriteString("\n")
		sb.WriteString(c.modernFlagUsages())
		sb.WriteString("\n")
	}

	// Global/Persistent Flags (if not root command)
	if c.HasParent() && c.parent.PersistentFlags().HasFlags() {
		sb.WriteString(style.SubHeader("Global Flags"))
		sb.WriteString("\n")
		sb.WriteString(c.modernPersistentFlagUsages())
		sb.WriteString("\n")
	}

	// Additional help
	if c.HasSubCommands() {
		sb.WriteString(style.Dim(fmt.Sprintf("Use \"%s [command] --help\" for more information about a command.", c.Root().Name())))
		sb.WriteString("\n")
	}

	return sb.String()
}

// modernFlagUsages returns modern styled flag usages
func (c *Command) modernFlagUsages() string {
	var sb strings.Builder

	maxLen := 0
	c.Flags().VisitAll(func(f *pflag.Flag) {
		flagLen := len(f.Name) + 6 // "--" + name + "  "
		if f.Shorthand != "" {
			flagLen += 4 // "-X, "
		}
		if flagLen > maxLen {
			maxLen = flagLen
		}
	})

	c.Flags().VisitAll(func(f *pflag.Flag) {
		sb.WriteString("  ")

		flagStr := ""
		if f.Shorthand != "" {
			flagStr = style.Flag(fmt.Sprintf("-%s, --%s", f.Shorthand, f.Name))
		} else {
			flagStr = style.Flag(fmt.Sprintf("    --%s", f.Name))
		}

		// Pad to align descriptions
		padding := maxLen - len(f.Name) - 6
		if f.Shorthand != "" {
			padding -= 4
		}

		sb.WriteString(flagStr)
		sb.WriteString(strings.Repeat(" ", padding))
		sb.WriteString("  ")

		// Add type hint for non-boolean flags
		if f.Value.Type() != "bool" {
			sb.WriteString(style.Argument(fmt.Sprintf("<%s>", f.Value.Type())))
			sb.WriteString("  ")
		}

		sb.WriteString(style.Muted(f.Usage))

		// Show default value if it's not empty and not "false" for bools
		if f.DefValue != "" && !(f.Value.Type() == "bool" && f.DefValue == "false") {
			sb.WriteString(style.Dim(fmt.Sprintf(" (default: %s)", f.DefValue)))
		}

		sb.WriteString("\n")
	})

	return sb.String()
}

// modernPersistentFlagUsages returns modern styled persistent flag usages
func (c *Command) modernPersistentFlagUsages() string {
	var sb strings.Builder

	maxLen := 0
	c.parent.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		flagLen := len(f.Name) + 6
		if f.Shorthand != "" {
			flagLen += 4
		}
		if flagLen > maxLen {
			maxLen = flagLen
		}
	})

	c.parent.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		// Skip if already shown in local flags
		if c.Flags().Lookup(f.Name) != nil {
			return
		}

		sb.WriteString("  ")

		flagStr := ""
		if f.Shorthand != "" {
			flagStr = style.Flag(fmt.Sprintf("-%s, --%s", f.Shorthand, f.Name))
		} else {
			flagStr = style.Flag(fmt.Sprintf("    --%s", f.Name))
		}

		padding := maxLen - len(f.Name) - 6
		if f.Shorthand != "" {
			padding -= 4
		}

		sb.WriteString(flagStr)
		sb.WriteString(strings.Repeat(" ", padding))
		sb.WriteString("  ")

		if f.Value.Type() != "bool" {
			sb.WriteString(style.Argument(fmt.Sprintf("<%s>", f.Value.Type())))
			sb.WriteString("  ")
		}

		sb.WriteString(style.Muted(f.Usage))

		if f.DefValue != "" && !(f.Value.Type() == "bool" && f.DefValue == "false") {
			sb.WriteString(style.Dim(fmt.Sprintf(" (default: %s)", f.DefValue)))
		}

		sb.WriteString("\n")
	})

	return sb.String()
}

// PrintSuccess prints a success message
func (c *Command) PrintSuccess(msg string) {
	fmt.Fprintln(c.OutOrStdout(), style.Success(msg))
}

// PrintError prints an error message
func (c *Command) PrintError(msg string) {
	fmt.Fprintln(c.ErrOrStderr(), style.Error(msg))
}

// PrintWarning prints a warning message
func (c *Command) PrintWarning(msg string) {
	fmt.Fprintln(c.OutOrStdout(), style.Warning(msg))
}

// PrintInfo prints an info message
func (c *Command) PrintInfo(msg string) {
	fmt.Fprintln(c.OutOrStdout(), style.Info(msg))
}

// PrintHeader prints a header
func (c *Command) PrintHeader(msg string) {
	fmt.Fprintln(c.OutOrStdout(), style.Header(msg))
}

// PrintSubHeader prints a sub-header
func (c *Command) PrintSubHeader(msg string) {
	fmt.Fprintln(c.OutOrStdout(), style.SubHeader(msg))
}

// PrintBullet prints a bullet point
func (c *Command) PrintBullet(msg string) {
	fmt.Fprintln(c.OutOrStdout(), style.Bullet(msg))
}

// PrintBox prints text in a box
func (c *Command) PrintBox(title, content string) {
	fmt.Fprintln(c.OutOrStdout(), style.Box(title, content))
}

// PrintCode prints code or technical text
func (c *Command) PrintCode(code string) {
	fmt.Fprintln(c.OutOrStdout(), style.Code(code))
}

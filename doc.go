/*
Package mamba provides a modern, drop-in replacement for Cobra with enhanced terminal features.

Mamba maintains 100% API compatibility with Cobra while adding beautiful colored output,
interactive prompts, loading spinners, progress bars, and modern styled help messages.

# Quick Start

Create a simple command:

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

# Modern Features

Mamba adds several modern terminal features on top of Cobra's functionality:

Styled Output - Beautiful, pre-styled output methods:

	cmd.PrintSuccess("Operation completed!")
	cmd.PrintError("Something went wrong")
	cmd.PrintWarning("Be careful")
	cmd.PrintInfo("Here's some info")
	cmd.PrintHeader("Section Title")

Interactive Prompts - User-friendly input with the interactive package:

	import "github.com/base-go/mamba/pkg/interactive"

	name, err := interactive.AskString("What's your name?", "Enter name")
	confirmed, err := interactive.AskConfirm("Continue?", true)

Loading Spinners - Visual feedback for operations:

	import "github.com/base-go/mamba/pkg/spinner"

	spinner.WithSpinner("Processing...", func() error {
		// Your operation here
		return nil
	})

Progress Bars - Track batch operations:

	spinner.WithProgress("Loading...", total, func(update func()) {
		for i := 0; i < total; i++ {
			// Process item
			update()
		}
	})

# Migration from Cobra

Migrating from Cobra is as simple as changing imports:

Before:

	import "github.com/spf13/cobra"

After:

	import "github.com/base-go/mamba"

All Cobra features continue to work:
  - Command structure (Use, Short, Long, Example)
  - Subcommands (AddCommand, RemoveCommand)
  - Flags (local, persistent, shorthand)
  - Lifecycle hooks (PreRun, PostRun, PersistentPreRun, etc.)
  - Argument validators (NoArgs, ExactArgs, MinimumNArgs, etc.)
  - Error handling (RunE, PreRunE, PostRunE)
  - Context support
  - Custom help and usage templates

# Styling

Use the style package for custom formatting:

	import "github.com/base-go/mamba/pkg/style"

	fmt.Println(style.Success("Success!"))
	fmt.Println(style.Error("Error!"))
	fmt.Println(style.Bold("Bold text"))
	fmt.Println(style.Code("code snippet"))
	fmt.Println(style.Box("Title", "Content"))

# Disabling Modern Features

If needed, modern styling can be disabled:

	disableColors := false
	cmd := &mamba.Command{
		Use:          "myapp",
		EnableColors: &disableColors,
	}

# Architecture

Mamba is built on proven libraries:
  - spf13/pflag - POSIX-style flag parsing
  - charmbracelet/lipgloss - Terminal styling
  - charmbracelet/huh - Interactive forms
  - charmbracelet/bubbles - TUI components
*/
package mamba

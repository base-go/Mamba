package main

import (
	"fmt"
	"os"
	"time"

	"github.com/base-go/mamba"
	"github.com/base-go/mamba/pkg/interactive"
	"github.com/base-go/mamba/pkg/spinner"
	"github.com/base-go/mamba/pkg/style"
)

var (
	verbose bool
)

func main() {
	// Root command
	rootCmd := &mamba.Command{
		Use:   "demo",
		Short: "A demo CLI application showcasing Mamba's modern terminal features",
		Long: `Mamba Demo is a CLI application that demonstrates the modern terminal features
available in Mamba, including colored output, interactive prompts, spinners, and more.

Mamba is a drop-in replacement for Cobra with enhanced UX features.`,
		Example: `  demo greet --name John
  demo process --count 10
  demo interactive
  demo styled`,
		Run: func(cmd *mamba.Command, args []string) {
			// Welcome header
			cmd.PrintHeader("🐍 Welcome to Mamba!")
			fmt.Println()
			cmd.PrintInfo("A modern CLI framework with beautiful terminal features")
			fmt.Println()
			time.Sleep(500 * time.Millisecond)

			// Show styled messages
			cmd.PrintSubHeader("Status Messages:")
			cmd.PrintSuccess("Success messages look like this")
			cmd.PrintWarning("Warning messages look like this")
			cmd.PrintError("Error messages look like this")
			cmd.PrintInfo("Info messages look like this")
			fmt.Println()
			time.Sleep(800 * time.Millisecond)

			// Show spinner demo
			cmd.PrintSubHeader("Loading Spinner:")
			spinner.WithSpinner("Initializing demo environment...", func() error {
				time.Sleep(2 * time.Second)
				return nil
			})
			fmt.Println()
			time.Sleep(300 * time.Millisecond)

			// Show progress bar demo
			cmd.PrintSubHeader("Progress Bar:")
			spinner.WithProgress("Loading components...", 8, func(update func()) {
				components := []string{"Colors", "Styles", "Prompts", "Spinners", "Progress", "Tables", "Boxes", "Icons"}
				for range components {
					time.Sleep(300 * time.Millisecond)
					update()
				}
			})
			fmt.Println()
			time.Sleep(300 * time.Millisecond)

			// Interactive prompt
			cmd.PrintSubHeader("Interactive Prompt:")
			name, err := interactive.AskString("What's your name?", "Enter your name")
			if err != nil {
				name = "Friend"
			}
			fmt.Println()

			// Personalized greeting
			cmd.PrintSuccess(fmt.Sprintf("Hello, %s! 👋", name))
			fmt.Println()
			time.Sleep(500 * time.Millisecond)

			// Show feature list
			cmd.PrintSubHeader("Available Features:")
			cmd.PrintBullet("Beautiful colored output")
			cmd.PrintBullet("Interactive prompts")
			cmd.PrintBullet("Loading spinners")
			cmd.PrintBullet("Progress bars")
			cmd.PrintBullet("Styled help messages")
			cmd.PrintBullet("100% Cobra compatible")
			fmt.Println()
			time.Sleep(500 * time.Millisecond)

			// Show code example
			cmd.PrintSubHeader("Quick Start:")
			cmd.PrintCode("go get github.com/base-go/mamba")
			fmt.Println()
			time.Sleep(300 * time.Millisecond)

			// Show box with next steps
			cmd.PrintBox("Next Steps",
				"Try these commands:\n"+
					"  • demo styled       - See all styling options\n"+
					"  • demo interactive  - Try interactive prompts\n"+
					"  • demo process      - Watch a spinner in action\n"+
					"  • demo --help       - See all available commands")
			fmt.Println()

			// Final message
			cmd.PrintInfo("Run 'demo --help' to see all available commands")
		},
	}

	// Global flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	// Greet command - basic command with flags
	greetCmd := &mamba.Command{
		Use:   "greet",
		Short: "Greet someone with style",
		Long:  "The greet command demonstrates styled output and flag handling.",
		Example: `  demo greet --name John
  demo greet -n Alice --enthusiastic`,
		RunE: func(cmd *mamba.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			enthusiastic, _ := cmd.Flags().GetBool("enthusiastic")

			if name == "" {
				cmd.PrintWarning("No name provided, using default")
				name = "World"
			}

			greeting := fmt.Sprintf("Hello, %s!", name)
			if enthusiastic {
				greeting += " It's great to see you!"
			}

			cmd.PrintSuccess(greeting)

			if verbose {
				cmd.PrintInfo("Verbose mode is enabled")
			}

			return nil
		},
	}
	greetCmd.Flags().StringP("name", "n", "", "Name to greet")
	greetCmd.Flags().BoolP("enthusiastic", "e", false, "Be more enthusiastic")

	// Process command - demonstrates spinner
	processCmd := &mamba.Command{
		Use:   "process",
		Short: "Process items with a loading spinner",
		Long:  "The process command demonstrates loading spinners for long-running operations.",
		RunE: func(cmd *mamba.Command, args []string) error {
			count, _ := cmd.Flags().GetInt("count")

			cmd.PrintHeader("Processing Items")

			err := spinner.WithSpinner("Processing items...", func() error {
				time.Sleep(time.Duration(count*100) * time.Millisecond)
				return nil
			})

			if err != nil {
				return err
			}

			cmd.PrintSuccess(fmt.Sprintf("Successfully processed %d items", count))
			return nil
		},
	}
	processCmd.Flags().IntP("count", "c", 5, "Number of items to process")

	// Progress command - demonstrates progress bar
	progressCmd := &mamba.Command{
		Use:   "progress",
		Short: "Show a progress bar",
		Long:  "The progress command demonstrates progress bars for batch operations.",
		RunE: func(cmd *mamba.Command, args []string) error {
			count, _ := cmd.Flags().GetInt("count")

			cmd.PrintHeader("Batch Processing")

			spinner.WithProgress("Processing batch...", count, func(update func()) {
				for i := 0; i < count; i++ {
					time.Sleep(200 * time.Millisecond)
					update()
				}
			})

			cmd.PrintSuccess("Batch processing completed")
			return nil
		},
	}
	progressCmd.Flags().IntP("count", "c", 10, "Number of items in batch")

	// Interactive command - demonstrates prompts
	interactiveCmd := &mamba.Command{
		Use:   "interactive",
		Short: "Interactive prompt demo",
		Long:  "The interactive command demonstrates various types of interactive prompts.",
		RunE: func(cmd *mamba.Command, args []string) error {
			cmd.PrintHeader("Interactive Demo")

			// Text input
			name, err := interactive.AskString("What's your name?", "Enter your name")
			if err != nil {
				return err
			}

			// Confirmation
			likes, err := interactive.AskConfirm(fmt.Sprintf("Nice to meet you, %s! Do you like Mamba?", name), true)
			if err != nil {
				return err
			}

			if likes {
				cmd.PrintSuccess("That's great! We're glad you like it!")
			} else {
				cmd.PrintInfo("We'd love to hear your feedback on how to improve!")
			}

			// Selection
			options := []interactive.SelectOption{
				{Key: "cobra", Value: "Cobra"},
				{Key: "click", Value: "Click (Python)"},
				{Key: "commander", Value: "Commander.js"},
				{Key: "other", Value: "Other"},
			}

			framework, err := interactive.AskSelect("What CLI framework do you currently use?", options)
			if err != nil {
				return err
			}

			cmd.PrintInfo(fmt.Sprintf("You selected: %s", framework))

			return nil
		},
	}

	// Styled command - demonstrates various styling options
	styledCmd := &mamba.Command{
		Use:   "styled",
		Short: "Showcase various styling options",
		Long:  "The styled command demonstrates different text styling and formatting options.",
		Run: func(cmd *mamba.Command, args []string) {
			cmd.PrintHeader("Styling Showcase")
			fmt.Println()

			// Status messages
			cmd.PrintSuccess("This is a success message")
			cmd.PrintError("This is an error message")
			cmd.PrintWarning("This is a warning message")
			cmd.PrintInfo("This is an info message")
			fmt.Println()

			// Text formatting
			fmt.Println(style.Bold("Bold text"))
			fmt.Println(style.Italic("Italic text"))
			fmt.Println(style.Underline("Underlined text"))
			fmt.Println(style.Dim("Dimmed text"))
			fmt.Println(style.Muted("Muted text"))
			fmt.Println()

			// Bullet points
			cmd.PrintSubHeader("Features:")
			cmd.PrintBullet("Modern colored output")
			cmd.PrintBullet("Interactive prompts")
			cmd.PrintBullet("Loading spinners")
			cmd.PrintBullet("Progress bars")
			fmt.Println()

			// Code
			cmd.PrintCode("go get github.com/base-go/mamba")
			fmt.Println()

			// Box
			cmd.PrintBox("Important", "Mamba is a drop-in replacement for Cobra!\nJust change your imports and enjoy the modern UX.")
		},
	}

	// List command - demonstrates complex styling
	listCmd := &mamba.Command{
		Use:   "list",
		Short: "List items with styling",
		Long:  "The list command demonstrates styled list output.",
		Run: func(cmd *mamba.Command, args []string) {
			cmd.PrintHeader("Available Features")
			fmt.Println()

			features := []struct {
				name        string
				description string
				status      string
			}{
				{"Colored Output", "Beautiful terminal colors using lipgloss", "ready"},
				{"Interactive Prompts", "User-friendly prompts with huh", "ready"},
				{"Spinners", "Loading indicators for operations", "ready"},
				{"Progress Bars", "Visual progress tracking", "ready"},
				{"Styled Help", "Modern help messages", "ready"},
			}

			for _, f := range features {
				fmt.Print(style.Command(f.name))
				fmt.Print(" - ")
				fmt.Print(style.Muted(f.description))
				fmt.Print(" ")
				if f.status == "ready" {
					fmt.Print(style.Success("✓"))
				}
				fmt.Println()
			}
		},
	}

	// Error command - demonstrates error handling
	errorCmd := &mamba.Command{
		Use:   "error",
		Short: "Demonstrate error handling",
		Long:  "The error command shows how errors are displayed with modern styling.",
		RunE: func(cmd *mamba.Command, args []string) error {
			cmd.PrintWarning("About to trigger an error...")
			time.Sleep(500 * time.Millisecond)
			return fmt.Errorf("this is a demonstration error")
		},
	}

	// Add all subcommands
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(processCmd)
	rootCmd.AddCommand(progressCmd)
	rootCmd.AddCommand(interactiveCmd)
	rootCmd.AddCommand(styledCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(errorCmd)

	// Execute
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

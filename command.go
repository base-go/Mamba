package mamba

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

// Command represents a CLI command with modern terminal features
type Command struct {
	// Use is the one-line usage message
	Use string

	// Aliases is an array of aliases that can be used instead of the first word in Use
	Aliases []string

	// Short is the short description shown in the 'help' output
	Short string

	// Long is the long message shown in the 'help <this-command>' output
	Long string

	// Example is examples of how to use the command
	Example string

	// Run is the function to call when this command is executed
	// If both Run and RunE are defined, RunE takes precedence
	Run func(cmd *Command, args []string)

	// RunE is the function to call when this command is executed, with error handling
	RunE func(cmd *Command, args []string) error

	// PreRun is called before Run
	PreRun func(cmd *Command, args []string)

	// PreRunE is called before RunE, with error handling
	PreRunE func(cmd *Command, args []string) error

	// PostRun is called after Run
	PostRun func(cmd *Command, args []string)

	// PostRunE is called after RunE, with error handling
	PostRunE func(cmd *Command, args []string) error

	// PersistentPreRun is called before PreRun and inherited by children
	PersistentPreRun func(cmd *Command, args []string)

	// PersistentPreRunE is called before PreRunE and inherited by children
	PersistentPreRunE func(cmd *Command, args []string) error

	// PersistentPostRun is called after PostRun and inherited by children
	PersistentPostRun func(cmd *Command, args []string)

	// PersistentPostRunE is called after PostRunE and inherited by children
	PersistentPostRunE func(cmd *Command, args []string) error

	// SilenceErrors prevents error messages from being displayed
	SilenceErrors bool

	// SilenceUsage prevents usage from being displayed on errors
	SilenceUsage bool

	// DisableFlagParsing disables flag parsing
	DisableFlagParsing bool

	// DisableAutoGenTag prevents auto-generation tag in help
	DisableAutoGenTag bool

	// Hidden hides this command from help output
	Hidden bool

	// Args defines expected arguments
	Args PositionalArgs

	// ValidArgs is list of all valid non-flag arguments
	ValidArgs []string

	// ValidArgsFunction is an optional function for custom argument completion
	ValidArgsFunction func(cmd *Command, args []string, toComplete string) ([]string, error)

	// Version is the version for this command
	Version string

	// commands is the list of subcommands
	commands []*Command

	// parent is a parent command for this command
	parent *Command

	// flags is the full flagset for this command
	flags *pflag.FlagSet

	// pflags is the persistent flagset for this command
	pflags *pflag.FlagSet

	// lflags is the local flagset for this command
	lflags *pflag.FlagSet

	// output is the writer to write output to
	output io.Writer

	// input is the reader to read input from
	input io.Reader

	// errOutput is the writer to write errors to
	errOutput io.Writer

	// ctx holds context for the command execution
	ctx interface{}

	// Modern terminal features
	// EnableColors enables colored output (default: auto-detect)
	EnableColors *bool

	// EnableInteractive enables interactive prompts
	EnableInteractive bool

	// ShowSpinner enables loading spinners
	ShowSpinner bool
}

// PositionalArgs defines a validation function for positional arguments
type PositionalArgs func(cmd *Command, args []string) error

// Standard validators for positional arguments
var (
	NoArgs = func(cmd *Command, args []string) error {
		if len(args) > 0 {
			return fmt.Errorf("unknown command %q for %q", args[0], cmd.Use)
		}
		return nil
	}

	ArbitraryArgs = func(cmd *Command, args []string) error {
		return nil
	}

	MinimumNArgs = func(n int) PositionalArgs {
		return func(cmd *Command, args []string) error {
			if len(args) < n {
				return fmt.Errorf("requires at least %d arg(s), only received %d", n, len(args))
			}
			return nil
		}
	}

	MaximumNArgs = func(n int) PositionalArgs {
		return func(cmd *Command, args []string) error {
			if len(args) > n {
				return fmt.Errorf("accepts at most %d arg(s), received %d", n, len(args))
			}
			return nil
		}
	}

	ExactArgs = func(n int) PositionalArgs {
		return func(cmd *Command, args []string) error {
			if len(args) != n {
				return fmt.Errorf("accepts %d arg(s), received %d", n, len(args))
			}
			return nil
		}
	}

	RangeArgs = func(min, max int) PositionalArgs {
		return func(cmd *Command, args []string) error {
			if len(args) < min || len(args) > max {
				return fmt.Errorf("accepts between %d and %d arg(s), received %d", min, max, len(args))
			}
			return nil
		}
	}
)

// Execute runs the command
func (c *Command) Execute() error {
	return c.ExecuteContext(nil)
}

// ExecuteContext runs the command with context
func (c *Command) ExecuteContext(ctx interface{}) error {
	c.ctx = ctx

	args := os.Args[1:]
	return c.execute(args)
}

func (c *Command) execute(args []string) error {
	// Parse flags if enabled
	if !c.DisableFlagParsing {
		if err := c.ParseFlags(args); err != nil {
			return err
		}
		args = c.Flags().Args()
	}

	// Find the command to execute
	cmd, flags, err := c.Find(args)
	if err != nil {
		return err
	}

	// Validate arguments
	if cmd.Args != nil {
		if err := cmd.Args(cmd, flags); err != nil {
			return err
		}
	}

	// Execute persistent pre-run
	if err := cmd.executePersistentPreRun(flags); err != nil {
		return err
	}

	// Execute pre-run
	if err := cmd.executePreRun(flags); err != nil {
		return err
	}

	// Execute main run
	if err := cmd.executeRun(flags); err != nil {
		if !c.SilenceErrors {
			fmt.Fprintln(c.ErrOrStderr(), err)
		}
		if !c.SilenceUsage {
			c.Usage()
		}
		return err
	}

	// Execute post-run
	if err := cmd.executePostRun(flags); err != nil {
		return err
	}

	// Execute persistent post-run
	if err := cmd.executePersistentPostRun(flags); err != nil {
		return err
	}

	return nil
}

func (c *Command) executePersistentPreRun(args []string) error {
	if c.PersistentPreRunE != nil {
		return c.PersistentPreRunE(c, args)
	}
	if c.PersistentPreRun != nil {
		c.PersistentPreRun(c, args)
	}
	return nil
}

func (c *Command) executePreRun(args []string) error {
	if c.PreRunE != nil {
		return c.PreRunE(c, args)
	}
	if c.PreRun != nil {
		c.PreRun(c, args)
	}
	return nil
}

func (c *Command) executeRun(args []string) error {
	if c.RunE != nil {
		return c.RunE(c, args)
	}
	if c.Run != nil {
		c.Run(c, args)
	}
	return nil
}

func (c *Command) executePostRun(args []string) error {
	if c.PostRunE != nil {
		return c.PostRunE(c, args)
	}
	if c.PostRun != nil {
		c.PostRun(c, args)
	}
	return nil
}

func (c *Command) executePersistentPostRun(args []string) error {
	if c.PersistentPostRunE != nil {
		return c.PersistentPostRunE(c, args)
	}
	if c.PersistentPostRun != nil {
		c.PersistentPostRun(c, args)
	}
	return nil
}

// AddCommand adds one or more subcommands
func (c *Command) AddCommand(cmds ...*Command) {
	for _, cmd := range cmds {
		if cmd == c {
			panic("command can't be a child of itself")
		}
		cmd.parent = c
		c.commands = append(c.commands, cmd)
	}
}

// RemoveCommand removes one or more subcommands
func (c *Command) RemoveCommand(cmds ...*Command) {
	commands := []*Command{}
main:
	for _, command := range c.commands {
		for _, cmd := range cmds {
			if command == cmd {
				command.parent = nil
				continue main
			}
		}
		commands = append(commands, command)
	}
	c.commands = commands
}

// Find finds the command to execute
func (c *Command) Find(args []string) (*Command, []string, error) {
	if len(args) == 0 {
		return c, args, nil
	}

	// Check for subcommand
	for _, cmd := range c.commands {
		if cmd.Name() == args[0] || cmd.HasAlias(args[0]) {
			return cmd.Find(args[1:])
		}
	}

	return c, args, nil
}

// Name returns the command's name
func (c *Command) Name() string {
	name := c.Use
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

// HasAlias checks if a string is an alias
func (c *Command) HasAlias(s string) bool {
	for _, a := range c.Aliases {
		if a == s {
			return true
		}
	}
	return false
}

// Commands returns all subcommands
func (c *Command) Commands() []*Command {
	return c.commands
}

// Parent returns the parent command
func (c *Command) Parent() *Command {
	return c.parent
}

// Flags returns the complete FlagSet
func (c *Command) Flags() *pflag.FlagSet {
	if c.flags == nil {
		c.flags = pflag.NewFlagSet(c.Name(), pflag.ContinueOnError)
		c.flags.SetOutput(c.ErrOrStderr())
	}
	return c.flags
}

// LocalFlags returns the local FlagSet
func (c *Command) LocalFlags() *pflag.FlagSet {
	if c.lflags == nil {
		c.lflags = pflag.NewFlagSet(c.Name(), pflag.ContinueOnError)
		c.lflags.SetOutput(c.ErrOrStderr())
	}
	return c.lflags
}

// PersistentFlags returns the persistent FlagSet
func (c *Command) PersistentFlags() *pflag.FlagSet {
	if c.pflags == nil {
		c.pflags = pflag.NewFlagSet(c.Name(), pflag.ContinueOnError)
		c.pflags.SetOutput(c.ErrOrStderr())
	}
	return c.pflags
}

// ParseFlags parses the flags
func (c *Command) ParseFlags(args []string) error {
	c.mergePersistentFlags()

	return c.Flags().Parse(args)
}

func (c *Command) mergePersistentFlags() {
	// Merge parent persistent flags
	if c.parent != nil {
		c.parent.mergePersistentFlags()
		c.parent.PersistentFlags().VisitAll(func(f *pflag.Flag) {
			if c.Flags().Lookup(f.Name) == nil {
				c.Flags().AddFlag(f)
			}
		})
	}

	// Merge local persistent flags
	c.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		if c.Flags().Lookup(f.Name) == nil {
			c.Flags().AddFlag(f)
		}
	})

	// Merge local flags
	c.LocalFlags().VisitAll(func(f *pflag.Flag) {
		if c.Flags().Lookup(f.Name) == nil {
			c.Flags().AddFlag(f)
		}
	})
}

// SetOutput sets the output writer
func (c *Command) SetOutput(output io.Writer) {
	c.output = output
}

// SetErr sets the error output writer
func (c *Command) SetErr(err io.Writer) {
	c.errOutput = err
}

// SetIn sets the input reader
func (c *Command) SetIn(in io.Reader) {
	c.input = in
}

// OutOrStdout returns the output writer or stdout
func (c *Command) OutOrStdout() io.Writer {
	if c.output != nil {
		return c.output
	}
	if c.parent != nil {
		return c.parent.OutOrStdout()
	}
	return os.Stdout
}

// ErrOrStderr returns the error output writer or stderr
func (c *Command) ErrOrStderr() io.Writer {
	if c.errOutput != nil {
		return c.errOutput
	}
	if c.parent != nil {
		return c.parent.ErrOrStderr()
	}
	return os.Stderr
}

// InOrStdin returns the input reader or stdin
func (c *Command) InOrStdin() io.Reader {
	if c.input != nil {
		return c.input
	}
	if c.parent != nil {
		return c.parent.InOrStdin()
	}
	return os.Stdin
}

// Usage prints the usage message
func (c *Command) Usage() error {
	if c.shouldUseModernHelp() {
		fmt.Fprintln(c.OutOrStdout(), c.ModernHelp())
	} else {
		fmt.Fprintln(c.OutOrStdout(), c.UsageString())
	}
	return nil
}

// UsageString returns the usage string (plain version)
func (c *Command) UsageString() string {
	var sb strings.Builder

	sb.WriteString("Usage:\n")
	sb.WriteString("  ")
	sb.WriteString(c.UseLine())
	sb.WriteString("\n\n")

	if c.Long != "" {
		sb.WriteString(c.Long)
		sb.WriteString("\n\n")
	} else if c.Short != "" {
		sb.WriteString(c.Short)
		sb.WriteString("\n\n")
	}

	if len(c.commands) > 0 {
		sb.WriteString("Available Commands:\n")
		for _, cmd := range c.commands {
			if !cmd.Hidden {
				sb.WriteString(fmt.Sprintf("  %-12s %s\n", cmd.Name(), cmd.Short))
			}
		}
		sb.WriteString("\n")
	}

	if c.Flags().HasFlags() {
		sb.WriteString("Flags:\n")
		sb.WriteString(c.Flags().FlagUsages())
	}

	return sb.String()
}

// UseLine returns the usage line
func (c *Command) UseLine() string {
	var useline string
	if c.parent != nil {
		useline = c.parent.UseLine() + " " + c.Use
	} else {
		useline = c.Use
	}
	return useline
}

// Help prints the help message
func (c *Command) Help() error {
	if c.shouldUseModernHelp() {
		fmt.Fprintln(c.OutOrStdout(), c.ModernHelp())
	} else {
		fmt.Fprintln(c.OutOrStdout(), c.UsageString())
	}
	return nil
}

// shouldUseModernHelp determines if modern help should be used
func (c *Command) shouldUseModernHelp() bool {
	// Use modern help by default unless explicitly disabled
	if c.EnableColors != nil {
		return *c.EnableColors
	}
	// Auto-detect: use modern help if output is a terminal
	// For now, default to true
	return true
}

// Context returns the command context
func (c *Command) Context() interface{} {
	return c.ctx
}

// SetContext sets the command context
func (c *Command) SetContext(ctx interface{}) {
	c.ctx = ctx
}

// HasSubCommands returns true if the command has subcommands
func (c *Command) HasSubCommands() bool {
	return len(c.commands) > 0
}

// HasParent returns true if the command has a parent
func (c *Command) HasParent() bool {
	return c.parent != nil
}

// Root returns the root command
func (c *Command) Root() *Command {
	if c.parent != nil {
		return c.parent.Root()
	}
	return c
}

// SetVersionTemplate sets the version template
func (c *Command) SetVersionTemplate(s string) {
	// TODO: implement version templating
}

// SetHelpCommand sets the help command
func (c *Command) SetHelpCommand(cmd *Command) {
	// TODO: implement custom help command
}

// SetHelpFunc sets the help function
func (c *Command) SetHelpFunc(f func(*Command, []string)) {
	// TODO: implement custom help function
}

// SetUsageFunc sets the usage function
func (c *Command) SetUsageFunc(f func(*Command) error) {
	// TODO: implement custom usage function
}

// SetUsageTemplate sets the usage template
func (c *Command) SetUsageTemplate(s string) {
	// TODO: implement usage templating
}

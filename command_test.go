package mamba

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestCommand_Execute(t *testing.T) {
	var output string
	cmd := &Command{
		Use: "test",
		Run: func(cmd *Command, args []string) {
			output = "executed"
		},
	}

	err := cmd.Execute()
	if err != nil {
		t.Errorf("Execute() error = %v", err)
	}
	if output != "executed" {
		t.Errorf("Expected output 'executed', got '%s'", output)
	}
}

func TestCommand_ExecuteWithError(t *testing.T) {
	cmd := &Command{
		Use: "test",
		RunE: func(cmd *Command, args []string) error {
			return nil
		},
	}

	err := cmd.Execute()
	if err != nil {
		t.Errorf("Execute() error = %v", err)
	}
}

func TestCommand_Subcommands(t *testing.T) {
	rootCmd := &Command{
		Use: "root",
	}

	subCmd := &Command{
		Use: "sub",
		Run: func(cmd *Command, args []string) {},
	}

	rootCmd.AddCommand(subCmd)

	if len(rootCmd.Commands()) != 1 {
		t.Errorf("Expected 1 subcommand, got %d", len(rootCmd.Commands()))
	}

	if subCmd.Parent() != rootCmd {
		t.Error("Expected subcommand parent to be root command")
	}
}

func TestCommand_Flags(t *testing.T) {
	cmd := &Command{
		Use: "test",
	}

	cmd.Flags().String("name", "", "Name flag")
	cmd.Flags().Bool("verbose", false, "Verbose flag")

	if !cmd.Flags().HasFlags() {
		t.Error("Expected command to have flags")
	}

	nameFlag := cmd.Flags().Lookup("name")
	if nameFlag == nil {
		t.Error("Expected 'name' flag to exist")
	}
}

func TestCommand_PersistentFlags(t *testing.T) {
	rootCmd := &Command{
		Use: "root",
	}

	rootCmd.PersistentFlags().Bool("verbose", false, "Verbose flag")

	subCmd := &Command{
		Use: "sub",
	}

	rootCmd.AddCommand(subCmd)

	// Persistent flags should be available on subcommand after merging
	if rootCmd.PersistentFlags().Lookup("verbose") == nil {
		t.Error("Expected persistent flag on root command")
	}
}

func TestCommand_Help(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := &Command{
		Use:   "test",
		Short: "Test command",
		Long:  "This is a test command for unit testing",
	}
	cmd.SetOutput(buf)

	err := cmd.Help()
	if err != nil {
		t.Errorf("Help() error = %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "test") {
		t.Errorf("Expected help to contain command name 'test', got: %s", output)
	}
}

func TestCommand_UsageString(t *testing.T) {
	cmd := &Command{
		Use:   "test [flags]",
		Short: "Test command",
	}

	usage := cmd.UsageString()
	if !strings.Contains(usage, "test") {
		t.Errorf("Expected usage to contain 'test', got: %s", usage)
	}
	if !strings.Contains(usage, "Usage:") {
		t.Errorf("Expected usage to contain 'Usage:', got: %s", usage)
	}
}

func TestCommand_LifecycleHooks(t *testing.T) {
	var executed []string

	cmd := &Command{
		Use: "test",
		PreRun: func(cmd *Command, args []string) {
			executed = append(executed, "pre")
		},
		Run: func(cmd *Command, args []string) {
			executed = append(executed, "run")
		},
		PostRun: func(cmd *Command, args []string) {
			executed = append(executed, "post")
		},
	}

	cmd.Execute()

	expected := []string{"pre", "run", "post"}
	if len(executed) != len(expected) {
		t.Errorf("Expected %d hooks executed, got %d", len(expected), len(executed))
	}
	for i, e := range expected {
		if i >= len(executed) || executed[i] != e {
			t.Errorf("Expected hook %d to be '%s', got '%s'", i, e, executed[i])
		}
	}
}

func TestCommand_Args_ExactArgs(t *testing.T) {
	cmd := &Command{
		Use:  "test <arg>",
		Args: ExactArgs(1),
		Run:  func(cmd *Command, args []string) {},
	}

	tests := []struct {
		args    []string
		wantErr bool
	}{
		{[]string{"arg1"}, false},
		{[]string{}, true},
		{[]string{"arg1", "arg2"}, true},
	}

	for _, tt := range tests {
		err := cmd.Args(cmd, tt.args)
		if (err != nil) != tt.wantErr {
			t.Errorf("ExactArgs() with args %v, wantErr %v, got err %v", tt.args, tt.wantErr, err)
		}
	}
}

func TestCommand_Args_MinimumNArgs(t *testing.T) {
	cmd := &Command{
		Use:  "test <args...>",
		Args: MinimumNArgs(2),
		Run:  func(cmd *Command, args []string) {},
	}

	tests := []struct {
		args    []string
		wantErr bool
	}{
		{[]string{"arg1", "arg2"}, false},
		{[]string{"arg1", "arg2", "arg3"}, false},
		{[]string{"arg1"}, true},
		{[]string{}, true},
	}

	for _, tt := range tests {
		err := cmd.Args(cmd, tt.args)
		if (err != nil) != tt.wantErr {
			t.Errorf("MinimumNArgs() with args %v, wantErr %v, got err %v", tt.args, tt.wantErr, err)
		}
	}
}

func TestCommand_Args_NoArgs(t *testing.T) {
	cmd := &Command{
		Use:  "test",
		Args: NoArgs,
		Run:  func(cmd *Command, args []string) {},
	}

	if err := cmd.Args(cmd, []string{}); err != nil {
		t.Errorf("NoArgs() with no args should not error, got: %v", err)
	}

	if err := cmd.Args(cmd, []string{"arg"}); err == nil {
		t.Error("NoArgs() with args should error")
	}
}

func TestCommand_Find(t *testing.T) {
	rootCmd := &Command{Use: "root"}
	subCmd := &Command{Use: "sub"}
	rootCmd.AddCommand(subCmd)

	tests := []struct {
		args     []string
		wantName string
	}{
		{[]string{}, "root"},
		{[]string{"sub"}, "sub"},
		{[]string{"sub", "arg"}, "sub"},
		{[]string{"unknown"}, "root"},
	}

	for _, tt := range tests {
		cmd, _, _ := rootCmd.Find(tt.args)
		if cmd.Name() != tt.wantName {
			t.Errorf("Find(%v) = %s, want %s", tt.args, cmd.Name(), tt.wantName)
		}
	}
}

func TestCommand_Aliases(t *testing.T) {
	cmd := &Command{
		Use:     "list",
		Aliases: []string{"ls", "l"},
	}

	if !cmd.HasAlias("ls") {
		t.Error("Expected command to have alias 'ls'")
	}
	if !cmd.HasAlias("l") {
		t.Error("Expected command to have alias 'l'")
	}
	if cmd.HasAlias("unknown") {
		t.Error("Expected command not to have alias 'unknown'")
	}
}

func TestCommand_HelpFlag(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd := &Command{
		Use:   "test",
		Short: "Test command",
	}
	rootCmd.SetOutput(buf)

	// Simulate -h flag
	rootCmd.Flags().BoolP("help", "h", false, "help")
	rootCmd.Flags().Set("help", "true")

	if !rootCmd.helpFlagSet() {
		t.Error("Expected help flag to be set")
	}
}

func TestCommand_IO(t *testing.T) {
	inBuf := bytes.NewBufferString("input")
	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)

	cmd := &Command{Use: "test"}
	cmd.SetIn(inBuf)
	cmd.SetOutput(outBuf)
	cmd.SetErr(errBuf)

	if cmd.InOrStdin() != inBuf {
		t.Error("Expected InOrStdin to return set input")
	}
	if cmd.OutOrStdout() != outBuf {
		t.Error("Expected OutOrStdout to return set output")
	}
	if cmd.ErrOrStderr() != errBuf {
		t.Error("Expected ErrOrStderr to return set error output")
	}
}

func TestCommand_Root(t *testing.T) {
	rootCmd := &Command{Use: "root"}
	subCmd := &Command{Use: "sub"}
	subSubCmd := &Command{Use: "subsub"}

	rootCmd.AddCommand(subCmd)
	subCmd.AddCommand(subSubCmd)

	if subSubCmd.Root() != rootCmd {
		t.Error("Expected Root() to return root command")
	}
	if subCmd.Root() != rootCmd {
		t.Error("Expected Root() to return root command")
	}
	if rootCmd.Root() != rootCmd {
		t.Error("Expected Root() to return itself for root command")
	}
}

func TestCommand_HasParent(t *testing.T) {
	rootCmd := &Command{Use: "root"}
	subCmd := &Command{Use: "sub"}
	rootCmd.AddCommand(subCmd)

	if rootCmd.HasParent() {
		t.Error("Expected root command to have no parent")
	}
	if !subCmd.HasParent() {
		t.Error("Expected subcommand to have parent")
	}
}

func TestCommand_HasSubCommands(t *testing.T) {
	rootCmd := &Command{Use: "root"}
	subCmd := &Command{Use: "sub"}

	if rootCmd.HasSubCommands() {
		t.Error("Expected command to have no subcommands initially")
	}

	rootCmd.AddCommand(subCmd)

	if !rootCmd.HasSubCommands() {
		t.Error("Expected command to have subcommands after adding")
	}
}

func TestCommand_RemoveCommand(t *testing.T) {
	rootCmd := &Command{Use: "root"}
	subCmd1 := &Command{Use: "sub1"}
	subCmd2 := &Command{Use: "sub2"}

	rootCmd.AddCommand(subCmd1, subCmd2)

	if len(rootCmd.Commands()) != 2 {
		t.Errorf("Expected 2 subcommands, got %d", len(rootCmd.Commands()))
	}

	rootCmd.RemoveCommand(subCmd1)

	if len(rootCmd.Commands()) != 1 {
		t.Errorf("Expected 1 subcommand after removal, got %d", len(rootCmd.Commands()))
	}
	if rootCmd.Commands()[0] != subCmd2 {
		t.Error("Expected remaining command to be subCmd2")
	}
	if subCmd1.HasParent() {
		t.Error("Expected removed command to have no parent")
	}
}

func TestCommand_ExecuteWithArgs(t *testing.T) {
	var receivedArgs []string
	cmd := &Command{
		Use: "test",
		Run: func(cmd *Command, args []string) {
			receivedArgs = args
		},
	}

	// Can't easily test with os.Args, but we can test the execute path
	err := cmd.execute([]string{"arg1", "arg2"})
	if err != nil {
		t.Errorf("execute() error = %v", err)
	}
	if len(receivedArgs) != 2 || receivedArgs[0] != "arg1" || receivedArgs[1] != "arg2" {
		t.Errorf("Expected args [arg1, arg2], got %v", receivedArgs)
	}
}

func TestCommand_ErrorHandling(t *testing.T) {
	errBuf := new(bytes.Buffer)
	cmd := &Command{
		Use: "test",
		RunE: func(cmd *Command, args []string) error {
			return fmt.Errorf("test error")
		},
	}
	cmd.SetErr(errBuf)

	err := cmd.Execute()
	if err == nil {
		t.Error("Expected error from RunE")
	}
}

func TestCommand_SilenceErrors(t *testing.T) {
	errBuf := new(bytes.Buffer)
	cmd := &Command{
		Use:           "test",
		SilenceErrors: true,
		RunE: func(cmd *Command, args []string) error {
			return fmt.Errorf("test error")
		},
	}
	cmd.SetErr(errBuf)

	err := cmd.Execute()
	if err == nil {
		t.Error("Expected error to be returned even when silenced")
	}
	if errBuf.String() != "" {
		t.Error("Expected no error output when SilenceErrors is true")
	}
}

func TestCommand_PersistentPreRunE(t *testing.T) {
	var executed []string
	rootCmd := &Command{
		Use: "root",
		PersistentPreRunE: func(cmd *Command, args []string) error {
			executed = append(executed, "persistent-pre")
			return nil
		},
		Run: func(cmd *Command, args []string) {
			executed = append(executed, "run")
		},
	}

	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Execute() error = %v", err)
	}

	if len(executed) != 2 {
		t.Errorf("Expected 2 hooks, got %d: %v", len(executed), executed)
	}
	if executed[0] != "persistent-pre" {
		t.Errorf("Expected first hook to be persistent-pre, got %s", executed[0])
	}
}

func TestCommand_PersistentPostRunE(t *testing.T) {
	var executed []string
	rootCmd := &Command{
		Use: "root",
		Run: func(cmd *Command, args []string) {
			executed = append(executed, "run")
		},
		PersistentPostRunE: func(cmd *Command, args []string) error {
			executed = append(executed, "persistent-post")
			return nil
		},
	}

	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Execute() error = %v", err)
	}

	if len(executed) != 2 {
		t.Errorf("Expected 2 hooks, got %d: %v", len(executed), executed)
	}
	if executed[1] != "persistent-post" {
		t.Errorf("Expected last hook to be persistent-post, got %s", executed[1])
	}
}

func TestCommand_IOFallbackToParent(t *testing.T) {
	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	inBuf := bytes.NewBufferString("input")

	rootCmd := &Command{Use: "root"}
	rootCmd.SetOutput(outBuf)
	rootCmd.SetErr(errBuf)
	rootCmd.SetIn(inBuf)

	subCmd := &Command{Use: "sub"}
	rootCmd.AddCommand(subCmd)

	if subCmd.OutOrStdout() != outBuf {
		t.Error("Expected subcommand to inherit parent output")
	}
	if subCmd.ErrOrStderr() != errBuf {
		t.Error("Expected subcommand to inherit parent error output")
	}
	if subCmd.InOrStdin() != inBuf {
		t.Error("Expected subcommand to inherit parent input")
	}
}

func TestCommand_Context(t *testing.T) {
	ctx := map[string]string{"key": "value"}
	cmd := &Command{Use: "test"}

	cmd.SetContext(ctx)

	result := cmd.Context()
	if result == nil {
		t.Error("Expected context to be set")
	}
	if resultMap, ok := result.(map[string]string); !ok || resultMap["key"] != "value" {
		t.Error("Expected context to contain correct data")
	}
}

func TestCommand_DisableFlagParsing(t *testing.T) {
	var receivedArgs []string
	cmd := &Command{
		Use:                "test",
		DisableFlagParsing: true,
		Run: func(cmd *Command, args []string) {
			receivedArgs = args
		},
	}

	err := cmd.execute([]string{"--flag", "value"})
	if err != nil {
		t.Errorf("execute() error = %v", err)
	}
	// With flag parsing disabled, flags are treated as args
	if len(receivedArgs) != 2 {
		t.Errorf("Expected 2 args with disabled flag parsing, got %d", len(receivedArgs))
	}
}

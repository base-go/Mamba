package mamba

import (
	"bytes"
	"strings"
	"testing"
)

func TestCommand_ModernHelp(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := &Command{
		Use:     "test",
		Short:   "Test command",
		Long:    "This is a long description of the test command",
		Example: "test --flag value\ntest arg1 arg2",
	}
	cmd.Flags().String("name", "", "Name flag")
	cmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	cmd.SetOutput(buf)

	result := cmd.ModernHelp()

	// Check for required sections
	if !strings.Contains(result, "test") {
		t.Error("ModernHelp should contain command name")
	}
	if !strings.Contains(result, "Usage") {
		t.Error("ModernHelp should contain Usage section")
	}
	if !strings.Contains(result, "Examples") {
		t.Error("ModernHelp should contain Examples section")
	}
	if !strings.Contains(result, "Flags") {
		t.Error("ModernHelp should contain Flags section")
	}
	if !strings.Contains(result, "--name") {
		t.Error("ModernHelp should contain flag names")
	}
}

func TestCommand_ModernHelpWithSubcommands(t *testing.T) {
	rootCmd := &Command{
		Use:   "root",
		Short: "Root command",
	}

	subCmd1 := &Command{
		Use:   "sub1",
		Short: "First subcommand",
	}

	subCmd2 := &Command{
		Use:    "sub2",
		Short:  "Second subcommand",
		Hidden: true,
	}

	rootCmd.AddCommand(subCmd1, subCmd2)

	result := rootCmd.ModernHelp()

	if !strings.Contains(result, "Available Commands") {
		t.Error("ModernHelp should contain Available Commands section")
	}
	if !strings.Contains(result, "sub1") {
		t.Error("ModernHelp should list visible subcommands")
	}
	if strings.Contains(result, "sub2") {
		t.Error("ModernHelp should not list hidden subcommands")
	}
}

func TestCommand_ModernHelpWithPersistentFlags(t *testing.T) {
	rootCmd := &Command{
		Use: "root",
	}
	rootCmd.PersistentFlags().Bool("debug", false, "Debug mode")

	subCmd := &Command{
		Use: "sub",
	}
	subCmd.Flags().String("local", "", "Local flag")

	rootCmd.AddCommand(subCmd)

	result := subCmd.ModernHelp()

	if !strings.Contains(result, "Flags") {
		t.Error("ModernHelp should show local flags")
	}
}

func TestCommand_PrintSuccess(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := &Command{Use: "test"}
	cmd.SetOutput(buf)

	cmd.PrintSuccess("Success message")

	output := buf.String()
	if !strings.Contains(output, "Success message") {
		t.Errorf("PrintSuccess should output message, got: %s", output)
	}
}

func TestCommand_PrintError(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := &Command{Use: "test"}
	cmd.SetErr(buf)

	cmd.PrintError("Error message")

	output := buf.String()
	if !strings.Contains(output, "Error message") {
		t.Errorf("PrintError should output message, got: %s", output)
	}
}

func TestCommand_PrintWarning(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := &Command{Use: "test"}
	cmd.SetOutput(buf)

	cmd.PrintWarning("Warning message")

	output := buf.String()
	if !strings.Contains(output, "Warning message") {
		t.Errorf("PrintWarning should output message, got: %s", output)
	}
}

func TestCommand_PrintInfo(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := &Command{Use: "test"}
	cmd.SetOutput(buf)

	cmd.PrintInfo("Info message")

	output := buf.String()
	if !strings.Contains(output, "Info message") {
		t.Errorf("PrintInfo should output message, got: %s", output)
	}
}

func TestCommand_PrintHeader(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := &Command{Use: "test"}
	cmd.SetOutput(buf)

	cmd.PrintHeader("Header Text")

	output := buf.String()
	if !strings.Contains(output, "Header Text") {
		t.Errorf("PrintHeader should output message, got: %s", output)
	}
}

func TestCommand_PrintSubHeader(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := &Command{Use: "test"}
	cmd.SetOutput(buf)

	cmd.PrintSubHeader("SubHeader Text")

	output := buf.String()
	if !strings.Contains(output, "SubHeader Text") {
		t.Errorf("PrintSubHeader should output message, got: %s", output)
	}
}

func TestCommand_PrintBullet(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := &Command{Use: "test"}
	cmd.SetOutput(buf)

	cmd.PrintBullet("Bullet item")

	output := buf.String()
	if !strings.Contains(output, "Bullet item") {
		t.Errorf("PrintBullet should output message, got: %s", output)
	}
}

func TestCommand_PrintBox(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := &Command{Use: "test"}
	cmd.SetOutput(buf)

	cmd.PrintBox("Title", "Content")

	output := buf.String()
	if !strings.Contains(output, "Title") {
		t.Errorf("PrintBox should output title, got: %s", output)
	}
	if !strings.Contains(output, "Content") {
		t.Errorf("PrintBox should output content, got: %s", output)
	}
}

func TestCommand_PrintCode(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := &Command{Use: "test"}
	cmd.SetOutput(buf)

	cmd.PrintCode("go run main.go")

	output := buf.String()
	if !strings.Contains(output, "go run main.go") {
		t.Errorf("PrintCode should output code, got: %s", output)
	}
}

func TestCommand_Usage(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := &Command{
		Use:   "test",
		Short: "Test command",
	}
	cmd.SetOutput(buf)

	err := cmd.Usage()
	if err != nil {
		t.Errorf("Usage() error = %v", err)
	}

	output := buf.String()
	if output == "" {
		t.Error("Usage() should produce output")
	}
}

func TestCommand_UsageWithModernHelp(t *testing.T) {
	buf := new(bytes.Buffer)
	enabled := true
	cmd := &Command{
		Use:          "test",
		Short:        "Test command",
		EnableColors: &enabled,
	}
	cmd.SetOutput(buf)

	err := cmd.Usage()
	if err != nil {
		t.Errorf("Usage() error = %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "test") {
		t.Errorf("Usage should contain command name, got: %s", output)
	}
}

func TestCommand_UsageWithPlainHelp(t *testing.T) {
	buf := new(bytes.Buffer)
	disabled := false
	cmd := &Command{
		Use:          "test",
		Short:        "Test command",
		EnableColors: &disabled,
	}
	cmd.SetOutput(buf)

	err := cmd.Usage()
	if err != nil {
		t.Errorf("Usage() error = %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "Usage:") {
		t.Errorf("Plain usage should contain 'Usage:', got: %s", output)
	}
}

package style

import (
	"github.com/charmbracelet/lipgloss"
)

// Theme colors
var (
	// Primary colors
	PrimaryColor   = lipgloss.Color("#7C3AED") // Purple
	SecondaryColor = lipgloss.Color("#06B6D4") // Cyan
	AccentColor    = lipgloss.Color("#F59E0B") // Amber

	// Status colors
	SuccessColor = lipgloss.Color("#10B981") // Green
	ErrorColor   = lipgloss.Color("#EF4444") // Red
	WarningColor = lipgloss.Color("#F59E0B") // Amber
	InfoColor    = lipgloss.Color("#3B82F6") // Blue

	// Text colors
	TextColor       = lipgloss.Color("#F3F4F6") // Light gray
	MutedColor      = lipgloss.Color("#9CA3AF") // Gray
	HighlightColor  = lipgloss.Color("#FBBF24") // Yellow
	DimColor        = lipgloss.Color("#6B7280") // Dark gray
	SubtleColor     = lipgloss.Color("#4B5563") // Darker gray
	BrightTextColor = lipgloss.Color("#FFFFFF") // White
)

// Base styles
var (
	// Text styles
	BoldStyle      = lipgloss.NewStyle().Bold(true)
	ItalicStyle    = lipgloss.NewStyle().Italic(true)
	UnderlineStyle = lipgloss.NewStyle().Underline(true)
	DimStyle       = lipgloss.NewStyle().Foreground(DimColor)
	MutedStyle     = lipgloss.NewStyle().Foreground(MutedColor)

	// Header styles
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(PrimaryColor).
			MarginBottom(1)

	SubHeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(SecondaryColor)

	// Status styles
	SuccessStyle = lipgloss.NewStyle().
			Foreground(SuccessColor).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(ErrorColor).
			Bold(true)

	WarningStyle = lipgloss.NewStyle().
			Foreground(WarningColor).
			Bold(true)

	InfoStyle = lipgloss.NewStyle().
			Foreground(InfoColor).
			Bold(true)

	// Box styles
	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(PrimaryColor).
			Padding(1, 2)

	HighlightBoxStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(AccentColor).
				Padding(1, 2).
				Background(lipgloss.Color("#1F2937"))

	// Command styles
	CommandStyle = lipgloss.NewStyle().
			Foreground(AccentColor).
			Bold(true)

	FlagStyle = lipgloss.NewStyle().
			Foreground(InfoColor)

	ArgumentStyle = lipgloss.NewStyle().
			Foreground(HighlightColor).
			Italic(true)

	// List styles
	BulletStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true)

	ListItemStyle = lipgloss.NewStyle().
			Foreground(TextColor).
			PaddingLeft(2)

	// Code/technical styles
	CodeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A78BFA")).
			Background(lipgloss.Color("#1F2937")).
			Padding(0, 1)

	// Prompt styles
	PromptStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true)

	InputStyle = lipgloss.NewStyle().
			Foreground(AccentColor)
)

// Status icons
const (
	SuccessIcon  = "✓"
	ErrorIcon    = "✗"
	WarningIcon  = "⚠"
	InfoIcon     = "ℹ"
	QuestionIcon = "?"
	ArrowIcon    = "→"
	BulletIcon   = "•"
	CheckIcon    = "✔"
	CrossIcon    = "✖"
)

// Render functions

// Success renders a success message
func Success(msg string) string {
	return SuccessStyle.Render(SuccessIcon+" ") + SuccessStyle.Render(msg)
}

// Error renders an error message
func Error(msg string) string {
	return ErrorStyle.Render(ErrorIcon+" ") + ErrorStyle.Render(msg)
}

// Warning renders a warning message
func Warning(msg string) string {
	return WarningStyle.Render(WarningIcon+" ") + WarningStyle.Render(msg)
}

// Info renders an info message
func Info(msg string) string {
	return InfoStyle.Render(InfoIcon+" ") + InfoStyle.Render(msg)
}

// Header renders a header
func Header(msg string) string {
	return HeaderStyle.Render(msg)
}

// SubHeader renders a sub-header
func SubHeader(msg string) string {
	return SubHeaderStyle.Render(msg)
}

// Command renders a command name
func Command(cmd string) string {
	return CommandStyle.Render(cmd)
}

// Flag renders a flag
func Flag(flag string) string {
	return FlagStyle.Render(flag)
}

// Argument renders an argument
func Argument(arg string) string {
	return ArgumentStyle.Render(arg)
}

// Code renders code or technical text
func Code(code string) string {
	return CodeStyle.Render(code)
}

// Bullet renders a bullet point
func Bullet(msg string) string {
	return BulletStyle.Render(BulletIcon+" ") + ListItemStyle.Render(msg)
}

// Box renders text in a box
func Box(title, content string) string {
	if title != "" {
		title = HeaderStyle.Render(title) + "\n\n"
	}
	return BoxStyle.Render(title + content)
}

// HighlightBox renders text in a highlighted box
func HighlightBox(title, content string) string {
	if title != "" {
		title = HeaderStyle.Render(title) + "\n\n"
	}
	return HighlightBoxStyle.Render(title + content)
}

// Bold renders bold text
func Bold(msg string) string {
	return BoldStyle.Render(msg)
}

// Italic renders italic text
func Italic(msg string) string {
	return ItalicStyle.Render(msg)
}

// Underline renders underlined text
func Underline(msg string) string {
	return UnderlineStyle.Render(msg)
}

// Dim renders dimmed text
func Dim(msg string) string {
	return DimStyle.Render(msg)
}

// Muted renders muted text
func Muted(msg string) string {
	return MutedStyle.Render(msg)
}

// Prompt renders a prompt
func Prompt(msg string) string {
	return PromptStyle.Render(msg + " " + ArrowIcon + " ")
}

// Input renders user input
func Input(msg string) string {
	return InputStyle.Render(msg)
}

// Colorize applies a color to text
func Colorize(msg string, color lipgloss.Color) string {
	return lipgloss.NewStyle().Foreground(color).Render(msg)
}

// WithBackground applies a background color to text
func WithBackground(msg string, fg, bg lipgloss.Color) string {
	return lipgloss.NewStyle().
		Foreground(fg).
		Background(bg).
		Render(msg)
}

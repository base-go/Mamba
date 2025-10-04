package style

import (
	"strings"
	"testing"
)

func TestSuccess(t *testing.T) {
	result := Success("test message")
	if !strings.Contains(result, "test message") {
		t.Errorf("Success() should contain message, got: %s", result)
	}
	if !strings.Contains(result, SuccessIcon) {
		t.Errorf("Success() should contain success icon, got: %s", result)
	}
}

func TestError(t *testing.T) {
	result := Error("test error")
	if !strings.Contains(result, "test error") {
		t.Errorf("Error() should contain message, got: %s", result)
	}
	if !strings.Contains(result, ErrorIcon) {
		t.Errorf("Error() should contain error icon, got: %s", result)
	}
}

func TestWarning(t *testing.T) {
	result := Warning("test warning")
	if !strings.Contains(result, "test warning") {
		t.Errorf("Warning() should contain message, got: %s", result)
	}
	if !strings.Contains(result, WarningIcon) {
		t.Errorf("Warning() should contain warning icon, got: %s", result)
	}
}

func TestInfo(t *testing.T) {
	result := Info("test info")
	if !strings.Contains(result, "test info") {
		t.Errorf("Info() should contain message, got: %s", result)
	}
	if !strings.Contains(result, InfoIcon) {
		t.Errorf("Info() should contain info icon, got: %s", result)
	}
}

func TestHeader(t *testing.T) {
	result := Header("Test Header")
	if !strings.Contains(result, "Test Header") {
		t.Errorf("Header() should contain text, got: %s", result)
	}
}

func TestSubHeader(t *testing.T) {
	result := SubHeader("Test SubHeader")
	if !strings.Contains(result, "Test SubHeader") {
		t.Errorf("SubHeader() should contain text, got: %s", result)
	}
}

func TestCommand(t *testing.T) {
	result := Command("mycommand")
	if !strings.Contains(result, "mycommand") {
		t.Errorf("Command() should contain text, got: %s", result)
	}
}

func TestFlag(t *testing.T) {
	result := Flag("--flag")
	if !strings.Contains(result, "--flag") {
		t.Errorf("Flag() should contain text, got: %s", result)
	}
}

func TestArgument(t *testing.T) {
	result := Argument("<arg>")
	if !strings.Contains(result, "<arg>") {
		t.Errorf("Argument() should contain text, got: %s", result)
	}
}

func TestCode(t *testing.T) {
	result := Code("go run main.go")
	if !strings.Contains(result, "go run main.go") {
		t.Errorf("Code() should contain text, got: %s", result)
	}
}

func TestBullet(t *testing.T) {
	result := Bullet("bullet item")
	if !strings.Contains(result, "bullet item") {
		t.Errorf("Bullet() should contain message, got: %s", result)
	}
	if !strings.Contains(result, BulletIcon) {
		t.Errorf("Bullet() should contain bullet icon, got: %s", result)
	}
}

func TestBox(t *testing.T) {
	result := Box("Title", "Content goes here")
	if !strings.Contains(result, "Title") {
		t.Errorf("Box() should contain title, got: %s", result)
	}
	if !strings.Contains(result, "Content goes here") {
		t.Errorf("Box() should contain content, got: %s", result)
	}
}

func TestBoxWithoutTitle(t *testing.T) {
	result := Box("", "Just content")
	if !strings.Contains(result, "Just content") {
		t.Errorf("Box() should contain content, got: %s", result)
	}
}

func TestHighlightBox(t *testing.T) {
	result := HighlightBox("Important", "This is important")
	if !strings.Contains(result, "Important") {
		t.Errorf("HighlightBox() should contain title, got: %s", result)
	}
	if !strings.Contains(result, "This is important") {
		t.Errorf("HighlightBox() should contain content, got: %s", result)
	}
}

func TestBold(t *testing.T) {
	result := Bold("bold text")
	if !strings.Contains(result, "bold text") {
		t.Errorf("Bold() should contain text, got: %s", result)
	}
}

func TestItalic(t *testing.T) {
	result := Italic("italic text")
	if !strings.Contains(result, "italic text") {
		t.Errorf("Italic() should contain text, got: %s", result)
	}
}

func TestUnderline(t *testing.T) {
	result := Underline("underlined text")
	if !strings.Contains(result, "underlined text") {
		t.Errorf("Underline() should contain text, got: %s", result)
	}
}

func TestDim(t *testing.T) {
	result := Dim("dimmed text")
	if !strings.Contains(result, "dimmed text") {
		t.Errorf("Dim() should contain text, got: %s", result)
	}
}

func TestMuted(t *testing.T) {
	result := Muted("muted text")
	if !strings.Contains(result, "muted text") {
		t.Errorf("Muted() should contain text, got: %s", result)
	}
}

func TestPrompt(t *testing.T) {
	result := Prompt("Enter value")
	if !strings.Contains(result, "Enter value") {
		t.Errorf("Prompt() should contain message, got: %s", result)
	}
	if !strings.Contains(result, ArrowIcon) {
		t.Errorf("Prompt() should contain arrow icon, got: %s", result)
	}
}

func TestInput(t *testing.T) {
	result := Input("user input")
	if !strings.Contains(result, "user input") {
		t.Errorf("Input() should contain text, got: %s", result)
	}
}

func TestColorize(t *testing.T) {
	result := Colorize("colored text", PrimaryColor)
	if !strings.Contains(result, "colored text") {
		t.Errorf("Colorize() should contain text, got: %s", result)
	}
}

func TestWithBackground(t *testing.T) {
	result := WithBackground("text", TextColor, PrimaryColor)
	if !strings.Contains(result, "text") {
		t.Errorf("WithBackground() should contain text, got: %s", result)
	}
}

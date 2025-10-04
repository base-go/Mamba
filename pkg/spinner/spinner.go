package spinner

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Spinner represents a loading spinner
type Spinner struct {
	message string
	style   lipgloss.Style
	spinner spinner.Model
	done    bool
	err     error
	output  io.Writer
	program *tea.Program
}

type spinnerModel struct {
	spinner spinner.Model
	message string
	style   lipgloss.Style
	done    bool
	err     error
}

func (m spinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m spinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		return m, nil
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case doneMsg:
		m.done = true
		return m, tea.Quit
	case errMsg:
		m.err = msg.err
		m.done = true
		return m, tea.Quit
	default:
		return m, nil
	}
}

func (m spinnerModel) View() string {
	if m.done {
		if m.err != nil {
			return m.style.Foreground(lipgloss.Color("#EF4444")).Render("✗ " + m.message + ": " + m.err.Error())
		}
		return m.style.Foreground(lipgloss.Color("#10B981")).Render("✓ " + m.message)
	}
	return m.spinner.View() + " " + m.style.Render(m.message)
}

type doneMsg struct{}
type errMsg struct{ err error }

// New creates a new spinner
func New(message string) *Spinner {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#7C3AED"))

	return &Spinner{
		message: message,
		style:   lipgloss.NewStyle().Foreground(lipgloss.Color("#F3F4F6")),
		spinner: s,
		output:  os.Stdout,
	}
}

// SetMessage updates the spinner message
func (s *Spinner) SetMessage(message string) {
	s.message = message
}

// SetOutput sets the output writer
func (s *Spinner) SetOutput(w io.Writer) {
	s.output = w
}

// Start starts the spinner
func (s *Spinner) Start() *Spinner {
	model := spinnerModel{
		spinner: s.spinner,
		message: s.message,
		style:   s.style,
	}
	s.program = tea.NewProgram(model, tea.WithOutput(s.output))
	go s.program.Run()
	return s
}

// Stop stops the spinner
func (s *Spinner) Stop() {
	if s.program != nil {
		s.program.Send(doneMsg{})
		time.Sleep(50 * time.Millisecond) // Give it time to render
	}
}

// Fail stops the spinner with an error
func (s *Spinner) Fail(err error) {
	if s.program != nil {
		s.program.Send(errMsg{err: err})
		time.Sleep(50 * time.Millisecond) // Give it time to render
	}
}

// Wait waits for the spinner to finish
func (s *Spinner) Wait() {
	if s.program != nil {
		s.program.Wait()
	}
}

// WithSpinner runs a function with a spinner
func WithSpinner(message string, fn func() error) error {
	s := New(message)
	s.Start()

	err := fn()

	if err != nil {
		s.Fail(err)
	} else {
		s.Stop()
	}

	s.Wait()
	return err
}

// Progress represents a progress bar
type Progress struct {
	total   int
	current int
	message string
	prog    progress.Model
	output  io.Writer
	program *tea.Program
}

type progressModel struct {
	progress progress.Model
	current  float64
	total    float64
	message  string
	done     bool
}

func (m progressModel) Init() tea.Cmd {
	return nil
}

func (m progressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		return m, nil
	case progressMsg:
		m.current = msg.current
		if m.current >= m.total {
			m.done = true
			return m, tea.Quit
		}
		return m, nil
	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - 4
		if m.progress.Width > 80 {
			m.progress.Width = 80
		}
		return m, nil
	default:
		return m, nil
	}
}

func (m progressModel) View() string {
	if m.done {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color("#10B981")).
			Render("✓ " + m.message + " (100%)")
	}

	percent := m.current / m.total
	return fmt.Sprintf("%s\n%s %.0f%%",
		m.message,
		m.progress.ViewAs(percent),
		percent*100,
	)
}

type progressMsg struct {
	current float64
}

// NewProgress creates a new progress bar
func NewProgress(message string, total int) *Progress {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(80),
	)

	return &Progress{
		total:   total,
		current: 0,
		message: message,
		prog:    p,
		output:  os.Stdout,
	}
}

// SetOutput sets the output writer
func (p *Progress) SetOutput(w io.Writer) {
	p.output = w
}

// Start starts the progress bar
func (p *Progress) Start() *Progress {
	model := progressModel{
		progress: p.prog,
		current:  0,
		total:    float64(p.total),
		message:  p.message,
	}
	p.program = tea.NewProgram(model, tea.WithOutput(p.output))
	go p.program.Run()
	return p
}

// Increment increments the progress
func (p *Progress) Increment() {
	if p.program != nil {
		p.program.Send(progressMsg{current: float64(p.current + 1)})
		p.current++
	}
}

// Set sets the progress to a specific value
func (p *Progress) Set(current int) {
	if p.program != nil {
		p.program.Send(progressMsg{current: float64(current)})
		p.current = current
	}
}

// Wait waits for the progress bar to finish
func (p *Progress) Wait() {
	if p.program != nil {
		p.program.Wait()
	}
}

// WithProgress runs a function with a progress bar
func WithProgress(message string, total int, fn func(update func())) error {
	p := NewProgress(message, total)
	p.Start()

	current := 0
	update := func() {
		current++
		p.program.Send(progressMsg{current: float64(current)})
	}

	fn(update)

	// Ensure we reach 100%
	p.program.Send(progressMsg{current: float64(total)})
	p.Wait()

	return nil
}

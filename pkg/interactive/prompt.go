package interactive

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

// Prompt represents a simple text input prompt
type Prompt struct {
	Title       string
	Description string
	Placeholder string
	Value       *string
	Required    bool
}

// Run executes the prompt
func (p *Prompt) Run() error {
	input := huh.NewInput().
		Title(p.Title).
		Description(p.Description).
		Placeholder(p.Placeholder).
		Value(p.Value)

	if p.Required {
		input = input.Validate(func(s string) error {
			if s == "" {
				return fmt.Errorf("this field is required")
			}
			return nil
		})
	}

	return input.Run()
}

// Confirm represents a yes/no confirmation prompt
type Confirm struct {
	Title       string
	Description string
	Value       *bool
	Affirmative string
	Negative    string
}

// Run executes the confirmation prompt
func (c *Confirm) Run() error {
	confirm := huh.NewConfirm().
		Title(c.Title).
		Description(c.Description).
		Value(c.Value)

	if c.Affirmative != "" {
		confirm = confirm.Affirmative(c.Affirmative)
	}
	if c.Negative != "" {
		confirm = confirm.Negative(c.Negative)
	}

	return confirm.Run()
}

// Select represents a selection prompt
type Select struct {
	Title       string
	Description string
	Options     []SelectOption
	Value       *string
}

// SelectOption represents an option in a select prompt
type SelectOption struct {
	Key   string
	Value string
}

// Run executes the select prompt
func (s *Select) Run() error {
	options := make([]huh.Option[string], len(s.Options))
	for i, opt := range s.Options {
		options[i] = huh.NewOption(opt.Value, opt.Key)
	}

	return huh.NewSelect[string]().
		Title(s.Title).
		Description(s.Description).
		Options(options...).
		Value(s.Value).
		Run()
}

// MultiSelect represents a multi-selection prompt
type MultiSelect struct {
	Title       string
	Description string
	Options     []SelectOption
	Value       *[]string
	Limit       int
}

// Run executes the multi-select prompt
func (m *MultiSelect) Run() error {
	options := make([]huh.Option[string], len(m.Options))
	for i, opt := range m.Options {
		options[i] = huh.NewOption(opt.Value, opt.Key)
	}

	multiSelect := huh.NewMultiSelect[string]().
		Title(m.Title).
		Description(m.Description).
		Options(options...).
		Value(m.Value)

	if m.Limit > 0 {
		multiSelect = multiSelect.Limit(m.Limit)
	}

	return multiSelect.Run()
}

// Text represents a multi-line text input prompt
type Text struct {
	Title       string
	Description string
	Placeholder string
	Value       *string
	CharLimit   int
	Required    bool
}

// Run executes the text prompt
func (t *Text) Run() error {
	text := huh.NewText().
		Title(t.Title).
		Description(t.Description).
		Placeholder(t.Placeholder).
		Value(t.Value)

	if t.CharLimit > 0 {
		text = text.CharLimit(t.CharLimit)
	}

	if t.Required {
		text = text.Validate(func(s string) error {
			if s == "" {
				return fmt.Errorf("this field is required")
			}
			return nil
		})
	}

	return text.Run()
}

// Form represents a group of prompts
type Form struct {
	Title       string
	Description string
	Groups      []FormGroup
}

// FormGroup represents a group of form fields
type FormGroup struct {
	Title  string
	Fields []huh.Field
}

// Run executes the form
func (f *Form) Run() error {
	groups := make([]*huh.Group, len(f.Groups))
	for i, g := range f.Groups {
		group := huh.NewGroup(g.Fields...)
		if g.Title != "" {
			group = group.Title(g.Title)
		}
		groups[i] = group
	}

	return huh.NewForm(groups...).Run()
}

// Helper functions for common prompts

// AskString prompts for a string input
func AskString(title, placeholder string) (string, error) {
	var value string
	p := &Prompt{
		Title:       title,
		Placeholder: placeholder,
		Value:       &value,
		Required:    true,
	}
	err := p.Run()
	return value, err
}

// AskConfirm prompts for a yes/no confirmation
func AskConfirm(title string, defaultValue bool) (bool, error) {
	value := defaultValue
	c := &Confirm{
		Title: title,
		Value: &value,
	}
	err := c.Run()
	return value, err
}

// AskSelect prompts for a selection from a list
func AskSelect(title string, options []SelectOption) (string, error) {
	var value string
	s := &Select{
		Title:   title,
		Options: options,
		Value:   &value,
	}
	err := s.Run()
	return value, err
}

// AskMultiSelect prompts for multiple selections from a list
func AskMultiSelect(title string, options []SelectOption, limit int) ([]string, error) {
	var value []string
	m := &MultiSelect{
		Title:   title,
		Options: options,
		Value:   &value,
		Limit:   limit,
	}
	err := m.Run()
	return value, err
}

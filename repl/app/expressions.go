package app

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type Expressions struct {
	textArea textarea.Model
}

func NewExpressions() *Expressions {
	return &Expressions{
		textArea: textarea.New(),
	}
}

func (m *Expressions) Focus() {
	m.textArea.Focus()
}

func (m *Expressions) Blur() {
	m.textArea.Blur()
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *Expressions) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m *Expressions) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.textArea, cmd = m.textArea.Update(msg)

	return m, cmd
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m *Expressions) View() string {
	return m.textArea.View()
}

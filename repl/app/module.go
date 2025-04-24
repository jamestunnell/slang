package app

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type Module struct {
	textArea textarea.Model
}

func NewModule() *Module {
	return &Module{
		textArea: textarea.New(),
	}
}

func (m *Module) Focus() {
	m.textArea.Focus()
}

func (m *Module) Blur() {
	m.textArea.Blur()
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *Module) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m *Module) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.textArea, cmd = m.textArea.Update(msg)

	return m, cmd
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m *Module) View() string {
	return m.textArea.View()
}

package app

import (
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type Module struct {
	focused  bool
	textArea textarea.Model
}

func NewModule() *Module {
	return &Module{
		textArea: newTextArea(12),
		focused:  false,
	}
}

func (m *Module) GetName() string {
	return "Module"
}

func (m *Module) IsFocused() bool {
	return m.focused
}

func (m *Module) Focus() tea.Cmd {
	m.focused = true

	return m.textArea.Focus()
}

func (m *Module) Blur() {
	m.textArea.Blur()

	m.focused = false
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *Module) Init() tea.Cmd {
	return cursor.Blink
}

func (m *Module) Resize(width, height int) {
	m.textArea.SetWidth(width - 10)
	m.textArea.SetHeight(height - 2)
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m *Module) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.(type) {
	default:
		m.textArea, cmd = m.textArea.Update(msg)
	}

	return m, cmd
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m *Module) View() string {
	return boxStyle(true).Render(m.textArea.View())
}

package app

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/lipgloss"
)

var cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))

func newTextArea(height int) textarea.Model {
	ta := textarea.New()

	ta.ShowLineNumbers = true
	ta.Cursor.Style = cursorStyle

	ta.SetHeight(height)

	return ta
}

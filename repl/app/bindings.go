package app

import "github.com/charmbracelet/bubbles/key"

var (
	bindingQuit = key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	)
	bindingNextTab = key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next tab"),
	)
)

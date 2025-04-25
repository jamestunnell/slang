package app

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit     key.Binding
	NextTab  key.Binding
	Evaluate key.Binding
	NavUp    key.Binding
	NavDown  key.Binding
}

func NewKeyMap() KeyMap {
	return KeyMap{
		NextTab: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "next tab"),
		),
		Evaluate: key.NewBinding(
			key.WithKeys("ctrl+w"),
			key.WithHelp("ctrl+w", "evaluate"),
		),
		// Help: key.NewBinding(
		// 	key.WithKeys("?"),
		// 	key.WithHelp("?", "toggle help"),
		// ),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+x"),
			key.WithHelp("ctrl+x", "exit"),
		),
		NavUp: key.NewBinding(
			key.WithKeys("shift+up"),
			key.WithHelp("shift+↑", "nav up"),
		),
		NavDown: key.NewBinding(
			key.WithKeys("shift+down"),
			key.WithHelp("shift+↓", "nav down"),
		),
	}
}

// ShortHelp returns a slice of bindings to be displayed in the short
// version of the help. The help bubble will render help in the order in
// which the help items are returned here.
func (km KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{km.Quit, km.NextTab, km.NavUp, km.NavDown, km.Evaluate}
}

// FullHelp returns an extended group of help items, grouped by columns.
// The help bubble will render the help in the order in which the help
// items are returned here.
func (km KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{km.Quit},
		{km.NextTab, km.NavUp, km.NavDown},
		{km.Evaluate},
	}
}

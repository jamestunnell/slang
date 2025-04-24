package app

import "github.com/charmbracelet/lipgloss"

var (
	borderColorActive   = lipgloss.Color("36")
	borderColorInactive = lipgloss.Color("238")
)

func boxStyle(active bool) lipgloss.Style {
	borderColor := borderColorInactive
	if active {
		borderColor = borderColorActive
	}

	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(borderColor).
		Padding(1)
}

package app

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

const tickDur = 100 * time.Millisecond

func tick() tea.Cmd {
	return tea.Tick(tickDur, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

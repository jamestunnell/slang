package app

import (
	"fmt"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/virtualmachine"
	"github.com/mistakenelf/teacup/statusbar"
)

type App struct {
	width  int
	height int

	tabs      []Tab
	tabIdx    int
	help      help.Model
	replPath  string
	keyMap    KeyMap
	statusBar statusbar.Model
}

type Args struct {
	VMID    string
	RPCAddr string
}

type Tab interface {
	GetName() string
	IsFocused() bool

	Focus() tea.Cmd
	Blur()

	tea.Model
}

type EvaluateMsg struct{}
type NavDownMsg struct{}
type NavUpMsg struct{}

// var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

func New(
	c virtualmachine.Client,
	info slang.VMInfo,
) *App {
	app := &App{
		replPath:  fmt.Sprintf("VM: name=%s id=%s ", info.Name, info.ID),
		statusBar: newStatusBar(),
		tabs:      []Tab{NewExpressions(c), NewModule()},
		tabIdx:    0,
		keyMap:    NewKeyMap(),
		help:      help.New(),
	}

	app.currentTab().Focus()

	return app
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (app *App) Init() tea.Cmd {
	return cursor.Blink
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (app *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	switch mm := msg.(type) {
	case tea.WindowSizeMsg:
		cmds = append(cmds, app.handleSize(mm)...)
	case tea.KeyMsg:
		cmds = append(cmds, app.handleKey(mm)...)
	default:
		_, cmd := app.currentTab().Update(msg)
		if cmd != nil {
			cmds = []tea.Cmd{cmd}
		}
	}

	if len(cmds) == 0 {
		return app, nil
	}

	return app, tea.Batch(cmds...)
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (app *App) View() string {
	tabBoxes := make([]string, len(app.tabs))

	for i, tab := range app.tabs {
		tabBoxes[i] = boxStyle(tab.IsFocused()).Render(tab.GetName())
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(lipgloss.Left, tabBoxes...),
		app.currentTab().View(),
		app.help.View(app.keyMap),
		app.statusBar.View(),
	)
}

func (app *App) handleSize(msg tea.WindowSizeMsg) []tea.Cmd {
	if msg.Width == app.width && msg.Height == app.height {
		return []tea.Cmd{}
	}

	app.height = msg.Height
	app.width = msg.Width

	app.statusBar.SetSize(msg.Width)

	cmds := []tea.Cmd{}

	for _, tab := range app.tabs {
		if _, cmd := tab.Update(msg); cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	app.updateStatusBar()

	return cmds
}

func (app *App) currentTab() Tab {
	return app.tabs[app.tabIdx]
}

func evaluate() tea.Msg {
	return EvaluateMsg{}
}

func navUp() tea.Msg {
	return NavUpMsg{}
}

func navDown() tea.Msg {
	return NavDownMsg{}
}

func (app *App) handleKey(msg tea.KeyMsg) []tea.Cmd {
	switch {
	case key.Matches(msg, app.keyMap.Quit):
		return []tea.Cmd{tea.Quit}
	case key.Matches(msg, app.keyMap.Evaluate):
		return []tea.Cmd{evaluate}
	case key.Matches(msg, app.keyMap.NavUp):
		return []tea.Cmd{navUp}
	case key.Matches(msg, app.keyMap.NavDown):
		return []tea.Cmd{navDown}
	case key.Matches(msg, app.keyMap.NextTab):
		app.currentTab().Blur()

		app.tabIdx = (app.tabIdx + 1) % len(app.tabs)

		return []tea.Cmd{app.currentTab().Focus()}
	default:
		_, cmd := app.currentTab().Update(msg)
		if cmd != nil {
			return []tea.Cmd{cmd}
		}
	}

	return []tea.Cmd{}
}

func (app *App) updateStatusBar() {
	app.statusBar.SetContent(
		fmt.Sprintf("%d x %d", app.width, app.height),
		app.replPath,
		"",
		"",
	)
}

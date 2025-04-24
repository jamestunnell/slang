package app

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mistakenelf/teacup/statusbar"
)

type App struct {
	width  int
	height int

	tabs      []Tab
	tabIdx    int
	help      help.Model
	replPath  string
	statusBar statusbar.Model
}

type Args struct {
	VMID    string
	RPCAddr string
}

type Tab struct {
	Name  string
	Model TabModel
}

type TabModel interface {
	Focus()
	Blur()

	tea.Model
}

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

func New(args *Args) *App {
	fmt.Printf("running REPL with args %#v\n", args)

	app := &App{
		replPath:  "VM: " + args.VMID,
		statusBar: newStatusBar(),
		tabs: []Tab{
			{Name: "Expressions", Model: NewExpressions()},
			{Name: "Module", Model: NewModule()},
		},
		tabIdx: 0,
	}

	app.tabs[app.tabIdx].Model.Focus()

	return app
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (app *App) Init() tea.Cmd {
	return nil
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
		tabBoxes[i] = boxStyle(i == app.tabIdx).Render(tab.Name)
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(lipgloss.Top, tabBoxes...),
		app.tabs[app.tabIdx].Model.View(),
		helpStyle.Render(app.help.ShortHelpView([]key.Binding{bindingNextTab, bindingQuit})),
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

	app.updateStatusBar()

	return []tea.Cmd{}
}

func (app *App) handleKey(msg tea.KeyMsg) []tea.Cmd {

	switch {
	case key.Matches(msg, bindingQuit):
		return []tea.Cmd{tea.Quit}
	case key.Matches(msg, bindingNextTab):
		app.tabs[app.tabIdx].Model.Blur()

		app.tabIdx = (app.tabIdx + 1) % len(app.tabs)

		app.tabs[app.tabIdx].Model.Focus()

		return []tea.Cmd{}
	}

	_, cmd := app.tabs[app.tabIdx].Model.Update(msg)
	if cmd != nil {
		return []tea.Cmd{cmd}
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

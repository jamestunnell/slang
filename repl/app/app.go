package app

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/archives"
	"github.com/mistakenelf/teacup/statusbar"
	"github.com/psanford/memfs"
)

type App struct {
	width  int
	height int

	editor       textarea.Model
	help         help.Model
	keyMap       KeyMap
	inputDigest  string
	packageState slang.PackageState
	replPath     string
	statusBar    statusbar.Model
	version      *Version
	vm           slang.VirtualMachine
}

type Args struct {
	VMID    string
	RPCAddr string
}

type EvaluateMsg struct{}

// var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

func New(
	vm slang.VirtualMachine,
) *App {
	app := &App{
		editor:       newTextArea(),
		help:         help.New(),
		keyMap:       NewKeyMap(),
		inputDigest:  "",
		packageState: slang.PkgAdded,
		replPath:     fmt.Sprintf("VM: name=%s id=%s ", vm.GetName(), vm.GetID()),
		statusBar:    newStatusBar(),
		version:      &Version{Major: 0, Minor: 0, Patch: 0},
		vm:           vm,
	}

	app.evaluate()
	app.editor.Focus()

	return app
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (app *App) Init() tea.Cmd {
	return tea.Batch(cursor.Blink, tick())
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (app *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	switch mm := msg.(type) {
	case tea.WindowSizeMsg:
		app.handleSize(mm)
	case tea.KeyMsg:
		cmds = append(cmds, app.handleKey(mm)...)
	case tickMsg:
		state, found := app.vm.GetPackageState(app.makePkgAddr())
		if found {
			app.packageState = state

			app.updateStatusBar()
		}

		cmds = append(cmds, tick())
	default:
		var cmd tea.Cmd

		app.editor, cmd = app.editor.Update(msg)
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
	return lipgloss.JoinVertical(
		lipgloss.Top,
		app.editor.View(),
		app.help.View(app.keyMap),
		app.statusBar.View(),
	)
}

func (app *App) makePkgAddr() slang.PackageAddress {
	return slang.PackageAddress{
		Path:    app.vm.GetName(),
		Version: app.version.String(),
	}
}

func (app *App) handleSize(msg tea.WindowSizeMsg) {
	const (
		helpHeight      = 1
		statusBarHeight = 1
	)

	if msg.Width == app.width && msg.Height == app.height {
		return
	}

	log.Printf("REPL: resize to %dx%d\n", msg.Width, msg.Height)

	app.height = msg.Height
	app.width = msg.Width

	app.statusBar.SetSize(msg.Width)
	app.editor.SetWidth(msg.Width)
	app.editor.SetHeight(msg.Height - (helpHeight + statusBarHeight))

	app.updateStatusBar()
}

func (app *App) evaluate() {
	input := app.editor.Value()

	digest := archives.MakeSHA256Digest([]byte(input))
	if digest == app.inputDigest {
		return
	}

	_, err := parseInput(input)
	if err != nil {
		log.Printf("REPL: failed to parse statements: %v\n", err)

		return
	}

	app.inputDigest = digest

	app.version.RevMinor()

	meta := slang.PackageMeta{
		Address:      app.makePkgAddr(),
		Dependencies: []slang.PackageAddress{},
	}
	tgz := archives.NewTarGz(meta.Address.String())
	archiveFs := memfs.New()

	if err = archiveFs.WriteFile("module.sl", []byte(input), 0666); err != nil {
		log.Printf("REPL: failed to write archive file: %v\n", err)

		return
	}

	if err = tgz.Pack(archiveFs); err != nil {
		log.Printf("REPL: failed to pack archive: %v\n", err)

		return
	}

	app.vm.UpsertPackage(meta, tgz)
}

func (app *App) handleKey(msg tea.KeyMsg) []tea.Cmd {
	switch {
	case key.Matches(msg, app.keyMap.Quit):
		return []tea.Cmd{tea.Quit}
	case key.Matches(msg, app.keyMap.Evaluate):
		app.evaluate()
	default:
		var cmd tea.Cmd

		app.editor, cmd = app.editor.Update(msg)
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
		app.packageState.String(),
		"",
	)
}

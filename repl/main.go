package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mistakenelf/teacup/statusbar"
	"golang.org/x/exp/maps"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexing"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/parsing/parsers"
	"github.com/jamestunnell/slang/sliceutil"
)

type focusArea int

type Model struct {
	focus            focusArea
	width            int
	height           int
	help             help.Model
	addStmts         textarea.Model
	viewStmts        viewport.Model
	statementsByName map[string]*statements.Statement
	statusBar        statusbar.Model
}

const (
	focusAddStmtsInput focusArea = iota
	focusParseStmtsBtn
	focusViewStmts
)

var (
	bindingParseAdd = key.NewBinding(
		key.WithKeys("ctrl+a"),
		key.WithHelp("ctrl+a", "parse+add"),
	)
	bindingQuit = key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	)
	bindingSwitch = key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch"),
	)

	borderColorActive   = lipgloss.Color("36")
	borderColorInactive = lipgloss.Color("238")

	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))

	// cursorLineStyle = lipgloss.NewStyle().
	// 		Background(lipgloss.Color("57")).
	// 		Foreground(lipgloss.Color("230"))

	// placeholderStyle = lipgloss.NewStyle().
	// 			Foreground(lipgloss.Color("238"))

	// endOfBufferStyle = lipgloss.NewStyle().
	// 			Foreground(lipgloss.Color("235"))

	// focusedPlaceholderStyle = lipgloss.NewStyle().
	// 			Foreground(lipgloss.Color("99"))

	// focusedBorderStyle = lipgloss.NewStyle().
	// 			Border(lipgloss.RoundedBorder()).
	// 			BorderForeground(borderColorActive)

	// blurredBorderStyle = lipgloss.NewStyle().
	// 			Border(lipgloss.RoundedBorder()).
	// 			BorderForeground(borderColorInactive)

	activeBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderColorActive)

	inactiveBoxStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(borderColorActive)
)

func newTextarea(placeholder string) textarea.Model {
	t := textarea.New()
	t.Prompt = ""
	t.Placeholder = placeholder
	t.ShowLineNumbers = true
	t.Cursor.Style = cursorStyle
	// t.FocusedStyle.Placeholder = focusedPlaceholderStyle
	// t.BlurredStyle.Placeholder = placeholderStyle
	// t.FocusedStyle.CursorLine = cursorLineStyle
	// t.FocusedStyle.Base = focusedBorderStyle
	// t.BlurredStyle.Base = blurredBorderStyle
	// t.FocusedStyle.EndOfBuffer = endOfBufferStyle
	// t.BlurredStyle.EndOfBuffer = endOfBufferStyle
	t.KeyMap.DeleteWordBackward.SetEnabled(false)
	t.KeyMap.LineNext = key.NewBinding(key.WithKeys("down"))
	t.KeyMap.LinePrevious = key.NewBinding(key.WithKeys("up"))

	t.Blur()

	return t
}

func newModel() *Model {
	m := &Model{
		focus:    focusAddStmtsInput,
		addStmts: newTextarea("Type something"),
		help:     help.New(),
		// viewStmts: viewport.New(,),
		statementsByName: map[string]*statements.Statement{},
		statusBar: statusbar.New(
			statusbar.ColorConfig{
				Foreground: lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#ffffff"},
				Background: lipgloss.AdaptiveColor{Light: "#F25D94", Dark: "#F25D94"},
			},
			statusbar.ColorConfig{
				Foreground: lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
				Background: lipgloss.AdaptiveColor{Light: "#3c3836", Dark: "#3c3836"},
			},
			statusbar.ColorConfig{
				Foreground: lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
				Background: lipgloss.AdaptiveColor{Light: "#A550DF", Dark: "#A550DF"},
			},
			statusbar.ColorConfig{
				Foreground: lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
				Background: lipgloss.AdaptiveColor{Light: "#6124DF", Dark: "#6124DF"},
			},
		),
	}

	m.addStmts.Focus()

	return m
}

func (f focusArea) String() string {
	switch f {
	case focusAddStmtsInput:
		return "INPUTSTMTS"
	case focusViewStmts:
		return "VIEWSTMTS"
	case focusParseStmtsBtn:
		return "PARSESTMTS"
	}

	return ""
}

func (m *Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m *Model) switchFocus() {
	switch m.focus {
	case focusAddStmtsInput:
		m.addStmts.Blur()
		m.viewStmts.Style.BorderForeground(borderColorActive)

		m.focus = focusViewStmts
	case focusViewStmts:
		m.addStmts.Focus()
		m.viewStmts.Style.BorderForeground(borderColorInactive)

		m.focus = focusAddStmtsInput
	}
}

func (m *Model) parseAndAdd() {
	if m.focus != focusAddStmtsInput {
		return
	}

	newStmts, err := parseInput(m.addStmts.Value())
	if err != nil {
		log.Printf("failed to parse input: %v", err)

		return
	}

	log.Printf("parsed %d statements", len(newStmts))

	for _, stmt := range newStmts {
		name, _ := stmt.GetName()

		m.statementsByName[name] = stmt
	}

	allStmts := maps.Values(m.statementsByName)

	m.addStmts.SetValue("")
	m.viewStmts.SetContent(renderStatements(allStmts))
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	const (
		helpHeight   = 10
		statusHeight = 1
	)

	var cmds []tea.Cmd

	log.Printf("got msg: %#v", msg)

	switch mm := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(mm, bindingParseAdd):
			m.parseAndAdd()
		case key.Matches(mm, bindingSwitch):
			m.switchFocus()
		case key.Matches(mm, bindingQuit):
			m.addStmts.Blur()

			cmds = append(cmds, tea.Quit)
		}
	case tea.WindowSizeMsg:
		m.height = mm.Height
		m.width = mm.Width

		mainHeight := mm.Height - helpHeight

		m.addStmts.SetHeight(mainHeight)
		m.addStmts.SetWidth(mm.Width / 2)

		m.viewStmts = viewport.New(mm.Width/2-10, mainHeight)

		m.statusBar.SetSize(mm.Width)

		// cmds = append(cmds, m.addStmts.Focus())
	}

	var update func(msg tea.Msg) []tea.Cmd

	switch m.focus {
	case focusAddStmtsInput:
		update = m.updateAddStmts
	case focusViewStmts:
		update = m.updateViewStmts
	}

	cmds = append(cmds, update(msg)...)

	return m, tea.Batch(cmds...)
}

func (m *Model) updateAddStmts(msg tea.Msg) []tea.Cmd {
	var cmd tea.Cmd

	m.addStmts, cmd = m.addStmts.Update(msg)
	if cmd != nil {
		return []tea.Cmd{cmd}
	}

	lineInfo := m.addStmts.LineInfo()
	rowStatus := fmt.Sprintf("row: %d/%d", 1+m.addStmts.Line(), m.addStmts.LineCount())
	columnStatus := fmt.Sprintf("col: %d/%d", 1+lineInfo.ColumnOffset, lineInfo.Width)

	m.setStatusBar(rowStatus, columnStatus)

	return []tea.Cmd{cmd}
}

func (m *Model) updateViewStmts(msg tea.Msg) []tea.Cmd {
	var cmd tea.Cmd

	m.viewStmts, cmd = m.viewStmts.Update(msg)
	if cmd != nil {
		return []tea.Cmd{cmd}
	}

	m.setStatusBar("", "")

	return []tea.Cmd{cmd}
}

func (m *Model) setStatusBar(rowStatus, colStatus string) {
	wd, _ := os.Getwd()

	m.statusBar.SetContent(m.focus.String(), wd, rowStatus, colStatus)
}

func (m *Model) currentBindings() []key.Binding {
	var bindings []key.Binding

	switch m.focus {
	case focusAddStmtsInput:
		bindings = []key.Binding{bindingParseAdd, bindingSwitch, bindingQuit}
	case focusViewStmts:
		bindings = []key.Binding{bindingSwitch, bindingQuit}
	}

	return bindings
}

func (m *Model) boxStyle(tgtFocus focusArea) lipgloss.Style {
	if m.focus == tgtFocus {
		return lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(borderColorActive).
			Padding(1)
	}

	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(borderColorInactive).
		Padding(1)
}

func (m *Model) View() string {
	mainBody := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.boxStyle(focusAddStmtsInput).Render(m.addStmts.View()),
			m.boxStyle(focusViewStmts).Render(m.viewStmts.View()),
		),
		m.help.ShortHelpView(m.currentBindings()),
	)

	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.NewStyle().Height(m.height-statusbar.Height).Render(mainBody),
		m.statusBar.View(),
	)
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("Failed to set up debug file logging:", err)
		os.Exit(1)
	}
	defer f.Close()

	p := tea.NewProgram(newModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Failed to run:", err)
		os.Exit(1)
	}
}

func parseInput(input string) ([]*statements.Statement, error) {
	runes := lexing.NewRuneSource(strings.NewReader(input))
	l := lexing.NewLexer(runes)
	p := parsers.NewFileParser()

	if err := parsing.RunParser(l, p); err != nil {
		return []*statements.Statement{}, err
	}

	return p.GetStatements(), nil
}

var groupTypes = []slang.StatementType{
	slang.StatementUSE,
	slang.StatementINTERFACE,
	slang.StatementSTRUCT,
	slang.StatementCONST,
	slang.StatementVAR,
	slang.StatementFUNC,
}

func renderStatements(stmts []*statements.Statement) string {
	w := slang.NewCodeWriter()
	writePreNewline := false

	for _, stmtType := range groupTypes {
		group := sliceutil.Where(stmts, func(s *statements.Statement) bool {
			return s.GetType() == stmtType
		})
		if len(group) == 0 {
			continue
		}

		slices.SortFunc(group, func(a, b *statements.Statement) int {
			nameA, _ := a.GetName()
			nameB, _ := b.GetName()

			return cmp.Compare(nameA, nameB)
		})

		if writePreNewline {
			w.WriteNewline()
		}

		for _, stmt := range group {
			stmt.Render(0, w)
		}

		writePreNewline = true
	}

	return w.String()
}

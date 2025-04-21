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
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/mistakenelf/teacup/statusbar"
	"github.com/mrusme/neonmodem/ui/helpers"
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
	inputArea        textarea.Model
	viewerArea       textarea.Model
	statementsByName map[string]*statements.Statement
	statusBar        statusbar.Model
	confirmForm      *huh.Form
}

type addStmtsMsg struct {
	Statements []*statements.Statement
}

type focusOnInputMsg struct{}

const (
	confirmHeight = 10
	confirmWidth  = 30

	focusInput focusArea = iota
	focusViewer
	focusDialog
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
)

func newTextarea(placeholder string) textarea.Model {
	t := textarea.New()
	t.Prompt = ""
	t.Placeholder = placeholder
	t.ShowLineNumbers = true
	t.Cursor.Style = cursorStyle

	return t
}

func newModel() *Model {
	m := &Model{
		focus:            focusInput,
		inputArea:        newTextarea("Type something"),
		viewerArea:       newTextarea(""),
		help:             help.New(),
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
		confirmForm: newConfirmForm(),
	}

	m.inputArea.Focus()

	m.viewerArea.Blur()
	m.viewerArea.KeyMap.DeleteWordBackward.Unbind()
	m.viewerArea.KeyMap.DeleteWordForward.Unbind()
	m.viewerArea.KeyMap.DeleteAfterCursor.Unbind()
	m.viewerArea.KeyMap.DeleteBeforeCursor.Unbind()
	m.viewerArea.KeyMap.InsertNewline.Unbind()
	m.viewerArea.KeyMap.DeleteCharacterBackward.Unbind()
	m.viewerArea.KeyMap.DeleteCharacterForward.Unbind()
	m.viewerArea.KeyMap.Paste.Unbind()
	m.viewerArea.KeyMap.TransposeCharacterBackward.Unbind()

	return m
}

func newConfirmForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Key("confirm").
				Title("Override existing?").
				Affirmative("Yes").
				Negative("No"),
		),
	)
}

func (f focusArea) String() string {
	switch f {
	case focusInput:
		return "INPUT"
	case focusDialog:
		return "DIALOG"
	case focusViewer:
		return "VIEWER"
	}

	return ""
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, m.confirmForm.Init())
}

func (m *Model) focusOnInput() {
	m.inputArea.Focus()
	m.viewerArea.Blur()

	m.focus = focusInput
}

func (m *Model) focusOnViewer() {
	m.inputArea.Blur()
	m.viewerArea.Focus()

	m.focus = focusViewer
}

func (m *Model) cycleFocus() {
	switch m.focus {
	case focusInput:
		m.focusOnViewer()
	case focusViewer:
		m.focusOnInput()
	}
}

func (m *Model) parseStatements() {
	if m.focus != focusInput {
		return
	}

	newStmts, err := parseInput(m.inputArea.Value())
	if err != nil {
		log.Printf("failed to parse input: %v", err)

		return
	}

	log.Printf("parsed %d statements", len(newStmts))

	names := sliceutil.Map(newStmts, func(s *statements.Statement) string {
		name, _ := s.GetName()

		return name
	})
	conflicts := sliceutil.Where(names, func(name string) bool {
		_, found := m.statementsByName[name]

		return found
	})

	if len(conflicts) == 0 {
		m.addStatements(newStmts)

		return
	}

	m.focus = focusDialog

	m.confirmForm.WithWidth(confirmWidth)
	m.confirmForm.WithHeight(confirmHeight)
	m.confirmForm.SubmitCmd = tea.Batch(
		m.confirmForm.Init(),
		func() tea.Msg { return addStmtsMsg{Statements: newStmts} },
		func() tea.Msg { return focusOnInputMsg{} },
	)
	m.confirmForm.CancelCmd = func() tea.Msg { return focusOnInputMsg{} }
}

func (m *Model) addStatements(stmts []*statements.Statement) {
	for _, stmt := range stmts {
		name, _ := stmt.GetName()

		m.statementsByName[name] = stmt
	}

	m.inputArea.SetValue("")
	m.viewerArea.SetValue(renderStatements(maps.Values(m.statementsByName)))
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	const (
		helpHeight   = 10
		statusHeight = 1
	)

	var cmds []tea.Cmd

	switch mm := msg.(type) {
	case focusOnInputMsg:
		m.focusOnInput()
	case addStmtsMsg:
		if m.confirmForm.GetBool("confirm") {
			m.addStatements(mm.Statements)
		}

		m.confirmForm = newConfirmForm()

		cmds = append(cmds, m.confirmForm.Init())
	case tea.KeyMsg:
		switch {
		case key.Matches(mm, bindingParseAdd):
			m.parseStatements()
		case key.Matches(mm, bindingSwitch):
			m.cycleFocus()
		case key.Matches(mm, bindingQuit):
			m.inputArea.Blur()
			m.viewerArea.Blur()

			cmds = append(cmds, tea.Quit)
		}
	case tea.WindowSizeMsg:
		if mm.Width == m.width && mm.Height == m.height {
			break
		}

		m.height = mm.Height
		m.width = mm.Width

		areaHeight := mm.Height - helpHeight
		areaWidth := (mm.Width / 2) - 5

		m.inputArea.SetHeight(areaHeight)
		m.inputArea.SetWidth(areaWidth)

		m.viewerArea.SetHeight(areaHeight)
		m.viewerArea.SetWidth(areaWidth)

		m.statusBar.SetSize(mm.Width)
	}

	var update func(msg tea.Msg) []tea.Cmd

	switch m.focus {
	case focusInput:
		update = m.updateInputArea
	case focusViewer:
		update = m.updateViewerArea
	case focusDialog:
		update = m.updateDialog
	}

	cmds = append(cmds, update(msg)...)

	return m, tea.Batch(cmds...)
}

func (m *Model) updateTextAreaStatusBar(ta textarea.Model) {
	lineInfo := ta.LineInfo()
	rowStatus := fmt.Sprintf("row: %d/%d", 1+ta.Line(), ta.LineCount())
	columnStatus := fmt.Sprintf("col: %d/%d", 1+lineInfo.ColumnOffset, lineInfo.Width)

	m.updateStatusBar(rowStatus, columnStatus)
}

func (m *Model) updateInputArea(msg tea.Msg) []tea.Cmd {
	var cmd tea.Cmd

	m.inputArea, cmd = m.inputArea.Update(msg)

	m.updateTextAreaStatusBar(m.inputArea)

	if cmd != nil {
		return []tea.Cmd{cmd}
	}

	return []tea.Cmd{}
}

func (m *Model) updateViewerArea(msg tea.Msg) []tea.Cmd {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		if len(keyMsg.Runes) > 0 {
			return []tea.Cmd{}
		}
	}

	var cmd tea.Cmd

	m.viewerArea, cmd = m.viewerArea.Update(msg)

	m.updateTextAreaStatusBar(m.viewerArea)

	if cmd != nil {
		return []tea.Cmd{cmd}
	}
	return []tea.Cmd{}
}

func (m *Model) updateDialog(msg tea.Msg) []tea.Cmd {
	form, cmd := m.confirmForm.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.confirmForm = f
	}

	m.updateStatusBar("", "")

	if cmd != nil {
		return []tea.Cmd{cmd}
	}

	return []tea.Cmd{}
}

func (m *Model) updateStatusBar(rowStatus, colStatus string) {
	wd, _ := os.Getwd()

	m.statusBar.SetContent(m.focus.String(), wd, rowStatus, colStatus)
}

func (m *Model) currentBindings() []key.Binding {
	var bindings []key.Binding

	switch m.focus {
	case focusInput:
		bindings = []key.Binding{bindingParseAdd, bindingSwitch, bindingQuit}
	case focusViewer:
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
			m.boxStyle(focusInput).Render(m.inputArea.View()),
			m.boxStyle(focusViewer).Render(m.viewerArea.View()),
		),
		m.help.ShortHelpView(m.currentBindings()),
	)

	everything := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.NewStyle().Height(m.height-statusbar.Height).Render(mainBody),
		m.statusBar.View(),
	)

	if m.focus != focusDialog {
		return everything
	}

	return helpers.PlaceOverlay(
		m.width/2-confirmWidth/2,
		m.height/2-confirmHeight/2,
		m.boxStyle(focusDialog).Render(m.confirmForm.View()),
		everything, false)
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

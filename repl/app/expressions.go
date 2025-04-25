package app

import (
	"log"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/lexing"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/parsing/parsers"
	"github.com/jamestunnell/slang/sliceutil"
	"github.com/jamestunnell/slang/virtualmachine"
)

type Expressions struct {
	focused bool
	client  virtualmachine.Client
	history []exprHistoryItem

	exprEditor   textarea.Model
	historyTable table.Model
}

type Focusable interface {
	Focused() bool
	Focus() tea.Cmd
	Blur()
}

type GetFocusableFunc func() Focusable

type exprHistoryItem struct {
	Time          time.Time
	Input, Result string
}

func NewExpressions(client virtualmachine.Client) *Expressions {
	exprEditor := newTextArea()
	historyTable := table.New(
		table.WithColumns([]table.Column{
			{Title: "Time", Width: 25},
			{Title: "Expression", Width: 25},
			{Title: "Result", Width: 50},
		}),
		table.WithFocused(false),
		table.WithHeight(10),
	)

	return &Expressions{
		exprEditor:   exprEditor,
		historyTable: historyTable,
		focused:      false,
		history:      []exprHistoryItem{},
		client:       client,
	}
}

func (m *Expressions) GetName() string {
	return "Expressions"
}

func (m *Expressions) IsFocused() bool {
	return m.focused
}

func (m *Expressions) Focus() tea.Cmd {
	m.focused = true

	return m.exprEditor.Focus()
}

func (m *Expressions) Blur() {
	m.exprEditor.Blur()

	m.focused = false
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m *Expressions) Init() tea.Cmd {
	return cursor.Blink
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m *Expressions) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch mm := msg.(type) {
	case tea.WindowSizeMsg:
		m.exprEditor.SetWidth(mm.Width)
	case EvaluateMsg:
		t := time.Now()
		input := m.exprEditor.Value()

		expr, err := parseExpr(input)
		if err != nil {
			log.Printf("failed to parse expression '%s': %v", input, err)

			m.addHistoryRow(t, input, err.Error())

			break
		}

		obj, err := m.client.EvaluateExpr(expr)
		if err != nil {
			log.Printf("failed to eval expression '%s': %v", input, err)

			m.addHistoryRow(t, input, err.Error())

			break
		}

		m.addHistoryRow(t, input, obj.Inspect())
	case NavUpMsg, NavDownMsg:
		if m.exprEditor.Focused() {
			m.exprEditor.Blur()
			m.historyTable.Focus()
		} else {
			cmd = m.exprEditor.Focus()
			m.historyTable.Blur()
		}
	default:
		if m.exprEditor.Focused() {
			m.exprEditor, cmd = m.exprEditor.Update(msg)
		} else if len(m.history) > 0 {
			if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "enter" {
				log.Printf("selected row: %v", m.historyTable.SelectedRow())
			}

			m.historyTable, cmd = m.historyTable.Update(msg)
		}
	}

	return m, cmd
}

func (m *Expressions) addHistoryRow(t time.Time, input, result string) {
	item := exprHistoryItem{Time: t, Input: input, Result: result}
	m.history = append(m.history, item)

	rows := sliceutil.Map(m.history, func(item exprHistoryItem) table.Row {
		return []string{item.Time.Format(time.RFC3339), item.Input, item.Result}
	})

	m.historyTable.SetRows(rows)
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m *Expressions) View() string {
	vertElems := []string{
		m.exprEditor.View(),
	}

	if len(m.history) != 0 {
		vertElems = append(vertElems, m.historyTable.View())
	}

	return lipgloss.JoinVertical(lipgloss.Left, vertElems...)
}

func parseExpr(s string) (*expressions.Expression, error) {
	r := strings.NewReader(s)
	l := lexing.NewLexer(lexing.NewRuneSource(r))
	p := parsers.NewExprParser(parsing.PrecedenceLOWEST)

	if err := parsing.RunParser(l, p); err != nil {
		return nil, err
	}

	return p.Expr, nil
}

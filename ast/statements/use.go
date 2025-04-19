package statements

import (
	"path"

	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/tokens"
)

type Use struct {
	Rename    string   `json:"rename,omitempty"`
	PathParts []string `json:"pathParts"`
}

func NewUse(rename string, parts []string) *Statement {
	core := &Use{Rename: rename, PathParts: parts}

	return NewStatement(slang.StatementUSE, core)
}

func (u *Use) GetName() (string, bool) {
	if u.Rename != "" {
		return u.Rename, true
	}

	n := len(u.PathParts)
	if n == 0 {
		return "", false
	}

	return u.PathParts[n-1], true
}

func (u *Use) IsEqual(other Core) bool {
	u2, ok := other.(*Use)
	if !ok {
		return false
	}

	if u.Rename != u2.Rename {
		return false
	}

	return slices.Equal(u.PathParts, u2.PathParts)
}

func (u *Use) Render(level int, w slang.CodeWriter) {
	w.WriteString(tokens.StrUSE)
	w.WriteString(" ")

	if u.Rename != "" {
		w.WriteString(u.Rename)
		w.WriteString(" ")
	}

	w.WriteString(`"`)
	w.WriteString(path.Join(u.PathParts...))
	w.WriteString(`"`)
}

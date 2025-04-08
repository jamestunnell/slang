package statements

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
)

type Use struct {
	Rename    string   `json:"rename,omitempty"`
	PathParts []string `json:"pathParts"`
}

func NewUse(rename string, parts []string) *Statement {
	core := &Use{Rename: rename, PathParts: parts}

	return NewStatement(slang.StatementUSE, core)
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

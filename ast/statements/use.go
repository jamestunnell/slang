package statements

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
)

type Use struct {
	PathParts []string `json:"pathParts"`
}

func NewUse(parts ...string) *Statement {
	core := &Use{PathParts: parts}

	return NewStatement(slang.StatementUSE, core)
}

func (u *Use) IsEqual(other Core) bool {
	u2, ok := other.(*Use)
	if !ok {
		return false
	}

	return slices.Equal(u.PathParts, u2.PathParts)
}

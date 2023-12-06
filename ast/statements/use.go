package statements

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
)

type Use struct {
	*Base

	PathParts []string `json:"pathParts"`
}

func NewUse(parts ...string) *Use {
	return &Use{
		Base:      NewBase(slang.StatementUSE),
		PathParts: parts,
	}
}

func (u *Use) Equal(other slang.Statement) bool {
	u2, ok := other.(*Use)
	if !ok {
		return false
	}

	return slices.Equal(u.PathParts, u2.PathParts)
}

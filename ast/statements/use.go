package statements

import (
	"github.com/jamestunnell/slang"
)

type Use struct {
	Path string `json:"path"`
}

func NewUse(path string) *Use {
	return &Use{Path: path}
}

func (u *Use) Type() slang.StatementType {
	return slang.StatementUSE
}

func (u *Use) Equal(other slang.Statement) bool {
	u2, ok := other.(*Use)
	if !ok {
		return false
	}

	return u.Path == u2.Path
}

package statements

import (
	"github.com/jamestunnell/slang"
)

type Base struct {
	StmtType slang.StatementType `json:"type"`
}

func NewBase(typ slang.StatementType) *Base {
	return &Base{StmtType: typ}
}

func (b *Base) GetType() slang.StatementType {
	return b.StmtType
}

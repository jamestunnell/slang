package expressions

import (
	"github.com/jamestunnell/slang"
)

type Base struct {
	ExprType slang.ExprType `json:"expressionType"`
}

func NewBase(typ slang.ExprType) *Base {
	return &Base{ExprType: typ}
}

func (b *Base) Type() slang.ExprType {
	return b.ExprType
}

package expressions

import (
	"github.com/jamestunnell/slang"
)

func NewBool(val bool) *Expression {
	return NewConst(slang.ExprBOOL, val)
}

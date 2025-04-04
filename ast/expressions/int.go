package expressions

import (
	"github.com/jamestunnell/slang"
)

func NewInt(val int64) *Expression {
	return NewConst(slang.ExprINT, val)
}

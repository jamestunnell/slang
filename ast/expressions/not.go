package expressions

import (
	"github.com/jamestunnell/slang"
)

type Not struct {
	*UnaryOperation
}

func NewNot(val slang.Expression) slang.Expression {
	return NewUnaryOperation(slang.ExprNOT, slang.MethodNOT, val)
}

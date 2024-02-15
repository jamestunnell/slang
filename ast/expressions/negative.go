package expressions

import (
	"github.com/jamestunnell/slang"
)

type Negative struct {
	*UnaryOperation
}

func NewNegative(val slang.Expression) slang.Expression {
	return NewUnaryOperation(slang.ExprNEGATIVE, slang.MethodNEG, val)
}

package expressions

import "github.com/jamestunnell/slang"

func NewAnd(left, right slang.Expression) slang.Expression {
	return NewBinaryOperation(slang.ExprAND, left, right)
}

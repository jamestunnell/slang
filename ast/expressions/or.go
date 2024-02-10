package expressions

import "github.com/jamestunnell/slang"

func NewOr(left, right slang.Expression) slang.Expression {
	return NewBinaryOperation(slang.ExprOR, left, right)
}

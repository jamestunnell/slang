package expressions

import "github.com/jamestunnell/slang"

func NewMultiply(left, right slang.Expression) *Expression {
	return NewBinaryOperation(slang.ExprMULTIPLY, left, right)
}

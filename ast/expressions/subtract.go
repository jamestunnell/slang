package expressions

import "github.com/jamestunnell/slang"

func NewSubtract(left, right slang.Expression) *Expression {
	return NewBinaryOperation(slang.ExprSUBTRACT, left, right)
}

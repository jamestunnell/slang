package expressions

import "github.com/jamestunnell/slang"

func NewSubtract(left, right slang.Expression) slang.Expression {
	return NewBinaryOperation(slang.ExprSUBTRACT, slang.MethodSUB, left, right)
}

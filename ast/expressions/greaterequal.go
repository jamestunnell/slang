package expressions

import "github.com/jamestunnell/slang"

func NewGreaterEqual(left, right slang.Expression) slang.Expression {
	return NewBinaryOperation(slang.ExprGREATEREQUAL, left, right)
}

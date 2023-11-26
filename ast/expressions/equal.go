package expressions

import "github.com/jamestunnell/slang"

func NewEqual(left, right slang.Expression) slang.Expression {
	return NewBinaryOperation(slang.ExprEQUAL, left, right)
}

package expressions

import "github.com/jamestunnell/slang"

func NewLessEqual(left, right slang.Expression) *Expression {
	return NewBinaryOperation(slang.ExprLESSEQUAL, left, right)
}

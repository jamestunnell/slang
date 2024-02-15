package expressions

import "github.com/jamestunnell/slang"

func NewAdd(left, right slang.Expression) slang.Expression {
	return NewBinaryOperation(slang.ExprADD, slang.MethodADD, left, right)
}

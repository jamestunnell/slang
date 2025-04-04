package expressions

import "github.com/jamestunnell/slang"

func NewAdd(left, right slang.Expression) *Expression {
	return NewBinaryOperation(slang.ExprADD, left, right)
}

package expressions

import "github.com/jamestunnell/slang"

func NewGreater(left, right slang.Expression) *Expression {
	return NewBinaryOperation(slang.ExprGREATER, left, right)
}

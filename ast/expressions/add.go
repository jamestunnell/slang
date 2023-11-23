package expressions

import "github.com/jamestunnell/slang"

type Add struct {
	*BinaryOperation
}

func NewAdd(left, right slang.Expression) slang.Expression {
	return &Add{
		BinaryOperation: NewBinaryOperation(AddOperator, left, right),
	}
}

func (expr *Add) Type() slang.ExprType { return slang.ExprADD }

func (expr *Add) Equal(other slang.Expression) bool {
	a2, ok := other.(*Add)
	if !ok {
		return false
	}

	return expr.BinaryOperation.Equal(a2.BinaryOperation)
}

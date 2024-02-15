package expressions

import (
	"github.com/jamestunnell/slang"
)

type BinaryOperation struct {
	*Base

	Left  slang.Expression `json:"left"`
	Right slang.Expression `json:"right"`

	opMethod string
}

func NewBinaryOperation(typ slang.ExprType, opMethod string, Left, Right slang.Expression) *BinaryOperation {
	return &BinaryOperation{
		Base:     NewBase(typ),
		Left:     Left,
		Right:    Right,
		opMethod: opMethod,
	}
}

func (binop *BinaryOperation) Equal(other slang.Expression) bool {
	binop2, ok := other.(*BinaryOperation)
	if !ok {
		return false
	}

	if binop.ExprType != binop2.ExprType {
		return false
	}

	return binop.Left.Equal(binop2.Left) && binop.Right.Equal(binop2.Right)
}

func (bo *BinaryOperation) Eval(env slang.Environment) (slang.Object, error) {
	l, err := bo.Left.Eval(env)
	if err != nil {
		return nil, err
	}

	r, err := bo.Right.Eval(env)
	if err != nil {
		return nil, err
	}

	return l.Send(bo.opMethod, r)
}

// // func (bo *BinaryOperation) String() string {
// // 	return fmt.Sprintf("%s %s %s", bo.Left, bo.Operator, bo.Right)
// // }

// func (Operator BinaryOperator) MethodName() string {
// 	var str string

// 	switch Operator {
// 	case AddOperator:
// 		str = slang.MethodADD
// 	case SubtractOperator:
// 		str = slang.MethodSUB
// 	case MultiplyOperator:
// 		str = slang.MethodMUL
// 	case DivideOperator:
// 		str = slang.MethodDIV
// 	case EqualOperator:
// 		str = slang.MethodEQ
// 	case NotEqualOperator:
// 		str = slang.MethodNEQ
// 	case LessOperator:
// 		str = slang.MethodLT
// 	case LessEqualOperator:
// 		str = slang.MethodLEQ
// 	case GreaterOperator:
// 		str = slang.MethodGT
// 	case GreaterEqualOperator:
// 		str = slang.MethodGEQ
// 	default:
// 		log.Fatal().Msgf("unexpected Operator %d", Operator)
// 	}

// 	return str
// }

// func (Operator BinaryOperator) MakeExpression(l, r slang.Expression) slang.Expression {
// 	var expr slang.Expression
// 	switch Operator {
// 	case AddOperator:
// 		expr = NewAdd(l, r)
// 	case SubtractOperator:
// 		expr = NewSubtract(l, r)
// 	case MultiplyOperator:
// 		expr = NewMultiply(l, r)
// 	case DivideOperator:
// 		expr = NewDivide(l, r)
// 	case EqualOperator:
// 		expr = NewEqual(l, r)
// 	case NotEqualOperator:
// 		expr = NewNotEqual(l, r)
// 	case LessOperator:
// 		expr = NewLess(l, r)
// 	case LessEqualOperator:
// 		expr = NewLessEqual(l, r)
// 	case GreaterOperator:
// 		expr = NewGreater(l, r)
// 	case GreaterEqualOperator:
// 		expr = NewGreaterEqual(l, r)
// 	default:
// 		log.Fatal().Msgf("unexpected Operator %d", Operator)
// 	}

// 	return expr
// }

package expressions

import (
	"github.com/jamestunnell/slang"
)

type UnaryOperation struct {
	*Base

	Value slang.Expression `json:"value"`
}

func NewUnaryOperation(typ slang.ExprType, val slang.Expression) *UnaryOperation {
	return &UnaryOperation{
		Base:  NewBase(typ),
		Value: val,
	}
}

func (op *UnaryOperation) Equal(other slang.Expression) bool {
	op2, ok := other.(*UnaryOperation)
	if !ok {
		return false
	}

	if op.ExprType != op2.ExprType {
		return false
	}

	return op.Value.Equal(op2.Value)
}

// func (bo *UnaryOperation) Eval(env *slang.Environment) (slang.Object, error) {
// 	a, err := bo.Left.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	b, err := bo.Right.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	return a.Send(bo.Operator.MethodName(), b)
// }

// // func (bo *UnaryOperation) String() string {
// // 	return fmt.Sprintf("%s %s %s", bo.Left, bo.Operator, bo.Right)
// // }

// func (Operator UnaryOperator) MethodName() string {
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

// func (Operator UnaryOperator) MakeExpression(l, r slang.Expression) slang.Expression {
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

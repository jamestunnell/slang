package expressions

import (
	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/objects"
)

type BinaryOperator int

type BinaryOperation struct {
	Left, Right slang.Expression
	Operator    BinaryOperator
}

const (
	AddOperator BinaryOperator = iota
	SubtractOperator
	MultiplyOperator
	DivideOperator
	EqualOperator
	NotEqualOperator
	LessOperator
	LessEqualOperator
	GreaterOperator
	GreaterEqualOperator
)

func NewBinaryOperation(op BinaryOperator, Left, Right slang.Expression) *BinaryOperation {
	return &BinaryOperation{
		Operator: op,
		Left:     Left,
		Right:    Right,
	}
}

func (bo *BinaryOperation) Equal(other *BinaryOperation) bool {
	if other.Operator != bo.Operator {
		return false
	}

	return other.Left.Equal(bo.Left) && other.Right.Equal(bo.Right)
}

func (bo *BinaryOperation) Eval(env *slang.Environment) (slang.Object, error) {
	a, err := bo.Left.Eval(env)
	if err != nil {
		return objects.NULL(), err
	}

	b, err := bo.Right.Eval(env)
	if err != nil {
		return objects.NULL(), err
	}

	return a.Send(bo.Operator.MethodName(), b)
}

// func (bo *BinaryOperation) String() string {
// 	return fmt.Sprintf("%s %s %s", bo.Left, bo.Operator, bo.Right)
// }

func (Operator BinaryOperator) MethodName() string {
	var str string

	switch Operator {
	case AddOperator:
		str = slang.MethodADD
	case SubtractOperator:
		str = slang.MethodSUB
	case MultiplyOperator:
		str = slang.MethodMUL
	case DivideOperator:
		str = slang.MethodDIV
	case EqualOperator:
		str = slang.MethodEQ
	case NotEqualOperator:
		str = slang.MethodNEQ
	case LessOperator:
		str = slang.MethodLT
	case LessEqualOperator:
		str = slang.MethodLEQ
	case GreaterOperator:
		str = slang.MethodGT
	case GreaterEqualOperator:
		str = slang.MethodGEQ
	default:
		log.Fatal().Msgf("unexpected Operator %d", Operator)
	}

	return str
}

func (Operator BinaryOperator) MakeExpression(l, r slang.Expression) slang.Expression {
	var expr slang.Expression
	switch Operator {
	case AddOperator:
		expr = NewAdd(l, r)
	case SubtractOperator:
		expr = NewSubtract(l, r)
	case MultiplyOperator:
		expr = NewMultiply(l, r)
	case DivideOperator:
		expr = NewDivide(l, r)
	case EqualOperator:
		expr = NewEqual(l, r)
	case NotEqualOperator:
		expr = NewNotEqual(l, r)
	case LessOperator:
		expr = NewLess(l, r)
	case LessEqualOperator:
		expr = NewLessEqual(l, r)
	case GreaterOperator:
		expr = NewGreater(l, r)
	case GreaterEqualOperator:
		expr = NewGreaterEqual(l, r)
	default:
		log.Fatal().Msgf("unexpected Operator %d", Operator)
	}

	return expr
}

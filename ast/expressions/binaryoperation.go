package expressions

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/tokens"
)

type BinaryOperation struct {
	Left  *Expression `json:"left"`
	Right *Expression `json:"right"`

	typ slang.ExprType
}

func NewBinaryOperation(typ slang.ExprType, Left, Right *Expression) *Expression {
	return NewExpression(typ, &BinaryOperation{
		Left:  Left,
		Right: Right,
		typ:   typ,
	})
}

func (binop *BinaryOperation) IsEqual(other Core) bool {
	binop2, ok := other.(*BinaryOperation)
	if !ok {
		return false
	}

	return binop.Left.IsEqual(binop2.Left) && binop.Right.IsEqual(binop2.Right)
}

func (binop *BinaryOperation) Render(level int, w slang.CodeWriter) {
	binop.Left.Render(level, w)

	switch binop.typ {
	case slang.ExprADD:
		w.WriteString(tokens.StrPLUS)
	case slang.ExprAND:
		w.WriteString(tokens.StrAND)
	case slang.ExprDIVIDE:
		w.WriteString(tokens.StrSLASH)
	case slang.ExprEQUAL:
		w.WriteString(tokens.StrEQUALEQUAL)
	case slang.ExprGREATER:
		w.WriteString(tokens.StrGREATER)
	case slang.ExprGREATEREQUAL:
		w.WriteString(tokens.StrGREATEREQ)
	case slang.ExprLESS:
		w.WriteString(tokens.StrLESS)
	case slang.ExprLESSEQUAL:
		w.WriteString(tokens.StrLESSEQ)
	case slang.ExprMULTIPLY:
		w.WriteString(tokens.StrSTAR)
	case slang.ExprNOTEQUAL:
		w.WriteString(tokens.StrNOTEQ)
	case slang.ExprOR:
		w.WriteString(tokens.StrOR)
	case slang.ExprSUBTRACT:
		w.WriteString(tokens.StrMINUS)
	default:
		w.WriteString("???")
	}

	binop.Right.Render(level, w)
}

// func (bo *BinaryOperation) Eval(env *slang.Environment) (slang.Object, error) {
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

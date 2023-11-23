package expressions

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
)

type FuncCall struct {
	Function  slang.Expression // Identifier or FunctionLiteral
	Arguments []slang.Expression
}

func NewFuncCall(fn slang.Expression, args ...slang.Expression) slang.Expression {
	return &FuncCall{
		Function:  fn,
		Arguments: args,
	}
}

func (c *FuncCall) Type() slang.ExprType {
	return slang.ExprFUNCCALL
}

func (c *FuncCall) Equal(other slang.Expression) bool {
	c2, ok := other.(*FuncCall)
	if !ok {
		return false
	}

	if !c2.Function.Equal(c.Function) {
		return false
	}

	return slices.EqualFunc(c.Arguments, c2.Arguments, expressionsEqual)
}

// func (expr *FuncCall) Eval(env *slang.Environment) (slang.Object, error) {
// 	obj, err := expr.Function.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	vals := make([]slang.Object, len(expr.Arguments))
// 	for i := 0; i < len(expr.Arguments); i++ {
// 		val, err := expr.Arguments[i].Eval(env)
// 		if err != nil {
// 			return objects.NULL(), err
// 		}

// 		vals[i] = val
// 	}

// 	return obj.Send(slang.MethodCALL, vals...)
// }

func expressionsEqual(a, b slang.Expression) bool {
	return a.Equal(b)
}

package expressions

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/objects"
	"golang.org/x/exp/slices"
)

type FunctionCall struct {
	Function  slang.Expression // Identifier or FunctionLiteral
	Arguments []slang.Expression
}

func NewFunctionCall(fn slang.Expression, args ...slang.Expression) slang.Expression {
	return &FunctionCall{
		Function:  fn,
		Arguments: args,
	}
}

func (c *FunctionCall) Type() slang.ExprType {
	return slang.ExprFUNCTIONCALL
}

func (c *FunctionCall) Equal(other slang.Expression) bool {
	c2, ok := other.(*FunctionCall)
	if !ok {
		return false
	}

	if !c2.Function.Equal(c.Function) {
		return false
	}

	return slices.EqualFunc(c.Arguments, c2.Arguments, expressionsEqual)
}

func (expr *FunctionCall) Eval(env *slang.Environment) (slang.Object, error) {
	obj, err := expr.Function.Eval(env)
	if err != nil {
		return objects.NULL(), err
	}

	vals := make([]slang.Object, len(expr.Arguments))
	for i := 0; i < len(expr.Arguments); i++ {
		val, err := expr.Arguments[i].Eval(env)
		if err != nil {
			return objects.NULL(), err
		}

		vals[i] = val
	}

	return obj.Send(slang.MethodCALL, vals...)
}

func expressionsEqual(a, b slang.Expression) bool {
	return a.Equal(b)
}

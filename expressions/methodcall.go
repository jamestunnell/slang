package expressions

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/objects"
	"golang.org/x/exp/slices"
)

type MethodCall struct {
	Object     slang.Expression
	MethodName *Identifier
	Arguments  []slang.Expression
}

func NewMethodCall(
	obj slang.Expression,
	mname *Identifier,
	args ...slang.Expression) slang.Expression {
	return &MethodCall{
		Object:     obj,
		MethodName: mname,
		Arguments:  args,
	}
}

func (c *MethodCall) Type() slang.ExprType {
	return slang.ExprMETHODCALL
}

func (c *MethodCall) Equal(other slang.Expression) bool {
	c2, ok := other.(*MethodCall)
	if !ok {
		return false
	}

	return c2.Object.Equal(c.Object) &&
		c2.MethodName.Equal(c.MethodName) &&
		slices.EqualFunc(c.Arguments, c2.Arguments, expressionsEqual)
}

func (c *MethodCall) Eval(env *slang.Environment) (slang.Object, error) {
	obj, err := c.Object.Eval(env)
	if err != nil {
		return objects.NULL(), err
	}

	vals := make([]slang.Object, len(c.Arguments))
	for i := 0; i < len(c.Arguments); i++ {
		val, err := c.Arguments[i].Eval(env)
		if err != nil {
			return objects.NULL(), err
		}

		vals[i] = val
	}

	return obj.Send(c.MethodName.Name, vals...)
}

package expressions

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
)

type MethodCall struct {
	*Base

	Subject   slang.Expression   `json:"subject"`
	Method    string             `json:"method"`
	Arguments []slang.Expression `json:"args"`
}

func NewMethodCall(
	subject slang.Expression,
	method string,
	args ...slang.Expression) slang.Expression {
	return &MethodCall{
		Base:      NewBase(slang.ExprMETHODCALL),
		Subject:   subject,
		Method:    method,
		Arguments: args,
	}
}

func (c *MethodCall) Equal(other slang.Expression) bool {
	c2, ok := other.(*MethodCall)
	if !ok {
		return false
	}

	return c2.Subject.Equal(c.Subject) &&
		c2.Method == c.Method &&
		slices.EqualFunc(c.Arguments, c2.Arguments, expressionsEqual)
}

// func (c *MethodCall) Eval(env *slang.Environment) (slang.Object, error) {
// 	obj, err := c.Object.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	vals := make([]slang.Object, len(c.Arguments))
// 	for i := 0; i < len(c.Arguments); i++ {
// 		val, err := c.Arguments[i].Eval(env)
// 		if err != nil {
// 			return objects.NULL(), err
// 		}

// 		vals[i] = val
// 	}

// 	return obj.Send(c.MethodName.Name, vals...)
// }

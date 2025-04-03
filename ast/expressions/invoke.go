package expressions

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
)

type InvokePos struct {
	*Base

	Subject slang.Expression `json:"subject"`
	Args    []*InvokeArg     `json:"args"`
}

type InvokeArg struct {
	Name  string           `json:"name,omitempty"`
	Value slang.Expression `json:"value"`
}

func NewInvoke(
	subject slang.Expression,
	args ...*InvokeArg,
) slang.Expression {
	return &InvokePos{
		Base:    NewBase(slang.ExprINVOKE),
		Subject: subject,
		Args:    args,
	}
}

func NewInvokeArgKW(name string, val slang.Expression) *InvokeArg {
	return &InvokeArg{Name: name, Value: val}
}

func NewInvokeArgPos(val slang.Expression) *InvokeArg {
	return &InvokeArg{Name: "", Value: val}
}

func (c *InvokePos) Equal(other slang.Expression) bool {
	c2, ok := other.(*InvokePos)
	if !ok {
		return false
	}

	if !c2.Subject.Equal(c.Subject) {
		return false
	}

	if !slices.EqualFunc(c.Args, c2.Args, invokeArgsEqual) {
		return false
	}

	return true
}

// func (expr *Call) Eval(env *slang.Environment) (slang.Object, error) {
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

func invokeArgsEqual(a, b *InvokeArg) bool {
	return a.Name == b.Name && a.Value.Equal(b.Value)
}

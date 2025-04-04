package expressions

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
)

type Invoke struct {
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
) *Expression {
	return NewExpression(slang.ExprINVOKE, &Invoke{
		Subject: subject,
		Args:    args,
	})
}

func NewInvokeArgKW(name string, val slang.Expression) *InvokeArg {
	return &InvokeArg{Name: name, Value: val}
}

func NewInvokeArgPos(val slang.Expression) *InvokeArg {
	return &InvokeArg{Name: "", Value: val}
}

func (c *Invoke) IsEqual(other Core) bool {
	c2, ok := other.(*Invoke)
	if !ok {
		return false
	}

	if !c2.Subject.IsEqual(c.Subject) {
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
	return a.Name == b.Name && a.Value.IsEqual(b.Value)
}

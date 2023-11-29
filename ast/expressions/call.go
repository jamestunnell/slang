package expressions

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
)

type Call struct {
	*Base

	Function slang.Expression `json:"function"`
	Args     []*Arg           `json:"args"`
}

type Arg struct {
	Name  string           `json:"name,omitempty"`
	Value slang.Expression `json:"value"`
}

func NewPositionalArg(val slang.Expression) *Arg {
	return &Arg{Name: "", Value: val}
}

func NewKeywordArg(name string, val slang.Expression) *Arg {
	return &Arg{Name: name, Value: val}
}

func NewCall(
	fn slang.Expression,
	args ...*Arg,
) slang.Expression {
	return &Call{
		Base:     NewBase(slang.ExprCALL),
		Function: fn,
		Args:     args,
	}
}

func (c *Call) Equal(other slang.Expression) bool {
	c2, ok := other.(*Call)
	if !ok {
		return false
	}

	if !c2.Function.Equal(c.Function) {
		return false
	}

	if !slices.EqualFunc(c.Args, c2.Args, argsEqual) {
		return false
	}

	return true
}

func argsEqual(a, b *Arg) bool {
	return a.Name == b.Name && a.Value.Equal(b.Value)
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

func expressionsEqual(a, b slang.Expression) bool {
	return a.Equal(b)
}

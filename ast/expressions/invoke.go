package expressions

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
)

type Invoke struct {
	Subject *Expression  `json:"subject"`
	Args    []*InvokeArg `json:"args"`
}

type InvokeArg struct {
	Name  string      `json:"name,omitempty"`
	Value *Expression `json:"value"`
}

func NewInvoke(
	subject *Expression,
	args ...*InvokeArg,
) *Expression {
	return NewExpression(slang.ExprINVOKE, &Invoke{
		Subject: subject,
		Args:    args,
	})
}

func NewInvokeArgKW(name string, val *Expression) *InvokeArg {
	return &InvokeArg{Name: name, Value: val}
}

func NewInvokeArgPos(val *Expression) *InvokeArg {
	return &InvokeArg{Name: "", Value: val}
}

func (i *Invoke) IsEqual(other Core) bool {
	i2, ok := other.(*Invoke)
	if !ok {
		return false
	}

	if !i2.Subject.IsEqual(i.Subject) {
		return false
	}

	if !slices.EqualFunc(i.Args, i2.Args, invokeArgsEqual) {
		return false
	}

	return true
}

func (i *Invoke) Render(level int, w slang.CodeWriter) {
	i.Subject.Render(level, w)

	switch len(i.Args) {
	case 0:
		w.WriteString("()")
	case 1:
		w.WriteString("(")

		i.Args[0].Render(level, w)

		w.WriteString(")")
	default:
		w.WriteString("[")

		subLevel := level + 1
		for _, arg := range i.Args {
			w.WriteNewline()
			w.WriteIndent(subLevel)

			arg.Render(subLevel, w)
		}

		w.WriteIndent(level)
		w.WriteString(")")
	}
}

func (a *InvokeArg) Render(level int, w slang.CodeWriter) {
	if a.Name != "" {
		w.WriteString(a.Name)
		w.WriteString(":")

	}

	a.Value.Render(level, w)
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

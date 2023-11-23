package expressions

import (
	"github.com/jamestunnell/slang"
)

type Negative struct {
	Value slang.Expression
}

func NewNegative(val slang.Expression) slang.Expression {
	return &Negative{Value: val}
}

func (b *Negative) Type() slang.ExprType { return slang.ExprNEGATIVE }

func (b *Negative) Equal(other slang.Expression) bool {
	b2, ok := other.(*Negative)
	if !ok {
		return false
	}

	return b2.Value.Equal(b.Value)
}

// func (expr *Negative) Eval(env *slang.Environment) (slang.Object, error) {
// 	obj, err := expr.Value.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	return obj.Send(slang.MethodNEG)
// }

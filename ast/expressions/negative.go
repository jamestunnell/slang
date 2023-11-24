package expressions

import (
	"github.com/jamestunnell/slang"
)

type Negative struct {
	*Base

	Value slang.Expression `json:"value"`
}

func NewNegative(val slang.Expression) slang.Expression {
	return &Negative{
		Base:  NewBase(slang.ExprNEGATIVE),
		Value: val,
	}
}

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

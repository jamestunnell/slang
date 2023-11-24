package expressions

import (
	"github.com/jamestunnell/slang"
)

type Not struct {
	*Base

	Value slang.Expression `json:"value"`
}

func NewNot(val slang.Expression) slang.Expression {
	return &Not{
		Base:  NewBase(slang.ExprNOT),
		Value: val,
	}
}

func (b *Not) Equal(other slang.Expression) bool {
	b2, ok := other.(*Not)
	if !ok {
		return false
	}

	return b2.Value.Equal(b.Value)
}

// func (expr *Not) Eval(env *slang.Environment) (slang.Object, error) {
// 	obj, err := expr.Value.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	return obj.Send(slang.MethodNOT)
// }

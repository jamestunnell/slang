package expressions

import (
	"github.com/jamestunnell/slang"
)

type Bool struct {
	*Base

	Value bool `json:"value"`
}

func NewBool(val bool) *Bool {
	return &Bool{
		Base:  NewBase(slang.ExprBOOL),
		Value: val,
	}
}

func (b *Bool) Equal(other slang.Expression) bool {
	b2, ok := other.(*Bool)
	if !ok {
		return false
	}

	return b2.Value == b.Value
}

// func (expr *Bool) Eval(env *slang.Environment) (slang.Object, error) {
// 	return objects.NewBool(expr.Value), nil
// }

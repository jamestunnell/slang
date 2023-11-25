package expressions

import (
	"github.com/jamestunnell/slang"
)

type Integer struct {
	*Base

	Value int64 `json:"value"`
}

func NewInteger(val int64) *Integer {
	return &Integer{
		Base:  NewBase(slang.ExprINTEGER),
		Value: val,
	}
}

func (i *Integer) Equal(other slang.Expression) bool {
	i2, ok := other.(*Integer)
	if !ok {
		return false
	}

	return i2.Value == i.Value
}

// func (expr *Integer) Eval(env *slang.Environment) (slang.Object, error) {
// 	return objects.NewInteger(expr.Value), nil
// }

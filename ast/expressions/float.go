package expressions

import (
	"github.com/jamestunnell/slang"
)

type Float struct {
	*Base

	Value float64 `json:"value"`
}

func NewFloat(val float64) *Float {
	return &Float{
		Base:  NewBase(slang.ExprFLOAT),
		Value: val,
	}
}

func (f *Float) Equal(other slang.Expression) bool {
	f2, ok := other.(*Float)
	if !ok {
		return false
	}

	return f2.Value == f.Value
}

// func (expr *Float) Eval(env *slang.Environment) (slang.Object, error) {
// 	return objects.NewFloat(expr.Value), nil
// }

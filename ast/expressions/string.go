package expressions

import (
	"github.com/jamestunnell/slang"
)

type String struct {
	*Base

	Value string `json:"value"`
}

func NewString(val string) *String {
	return &String{
		Base:  NewBase(slang.ExprSTRING),
		Value: val,
	}
}

func (i *String) Equal(other slang.Expression) bool {
	i2, ok := other.(*String)
	if !ok {
		return false
	}

	return i2.Value == i.Value
}

// func (expr *String) Eval(env *slang.Environment) (slang.Object, error) {
// 	return objects.NewString(expr.Value), nil
// }

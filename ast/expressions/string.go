package expressions

import (
	"github.com/jamestunnell/slang"
)

type String struct {
	Value string
}

func NewString(val string) *String {
	return &String{Value: val}
}

func (i *String) Type() slang.ExprType { return slang.ExprSTRING }

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

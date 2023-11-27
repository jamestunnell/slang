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

func (s *String) Equal(other slang.Expression) bool {
	s2, ok := other.(*String)
	if !ok {
		return false
	}

	return s2.Value == s.Value
}

// func (expr *String) Eval(env *slang.Environment) (slang.Object, error) {
// 	return objects.NewString(expr.Value), nil
// }

package expressions

import "github.com/jamestunnell/slang"

type Const[T comparable] struct {
	*Base

	Value T
}

func NewConst[T comparable](typ slang.ExprType, val T) *Const[T] {
	return &Const[T]{
		Base:  NewBase(typ),
		Value: val,
	}
}

func (c *Const[T]) Equal(other slang.Expression) bool {
	c2, ok := other.(*Const[T])
	if !ok {
		return false
	}

	return c2.Value == c.Value
}

package expressions

import (
	"fmt"

	"github.com/jamestunnell/slang"
)

type Const[T comparable] struct {
	Value T
}

func NewConst[T comparable](typ slang.ExprType, val T) *Expression {
	return NewExpression(typ, &Const[T]{Value: val})
}

func (c *Const[T]) IsEqual(other Core) bool {
	c2, ok := other.(*Const[T])
	if !ok {
		return false
	}

	return c2.Value == c.Value
}

func (c *Const[T]) Render(level int, w slang.CodeWriter) {
	w.WriteString(fmt.Sprintf("%v", c.Value))
}

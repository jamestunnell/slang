package expressions

import (
	"github.com/jamestunnell/slang"
)

type AccessElem struct {
	*Base

	Container slang.Expression `json:"container"`
	Key       slang.Expression `json:"key"`
}

func NewAccessElem(container, key slang.Expression) slang.Expression {
	return &AccessElem{
		Base:      NewBase(slang.ExprACCESSELEM),
		Container: container,
		Key:       key,
	}
}

func (c *AccessElem) Equal(other slang.Expression) bool {
	c2, ok := other.(*AccessElem)
	if !ok {
		return false
	}

	return c2.Container.Equal(c.Container) && c2.Key.Equal(c.Key)
}

// func (expr *AccessElem) Eval(env *slang.Environment) (slang.Object, error) {
// 	ary, err := expr.Array.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	idx, err := expr.AccessElem.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	return ary.Send(slang.MethodINDEX, idx)
// }

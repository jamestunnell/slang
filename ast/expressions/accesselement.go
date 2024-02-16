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

func (c *AccessElem) IsEqual(other slang.Expression) bool {
	c2, ok := other.(*AccessElem)
	if !ok {
		return false
	}

	return c2.Container.IsEqual(c.Container) && c2.Key.IsEqual(c.Key)
}

func (expr *AccessElem) Eval(env slang.Environment) (slang.Object, error) {
	container, err := expr.Container.Eval(env)
	if err != nil {
		return nil, err
	}

	key, err := expr.Key.Eval(env)
	if err != nil {
		return nil, err
	}

	return container.Send(slang.MethodELEM, key)
}

package expressions

import (
	"github.com/jamestunnell/slang"
)

type Key struct {
	Map slang.Expression `json:"map"`
	Key slang.Expression `json:"key"`
}

func NewKey(mapVal, key slang.Expression) slang.Expression {
	return &Key{
		Map: mapVal,
		Key: key,
	}
}

func (c *Key) Type() slang.ExprType {
	return slang.ExprKEY
}

func (c *Key) Equal(other slang.Expression) bool {
	c2, ok := other.(*Key)
	if !ok {
		return false
	}

	return c2.Map.Equal(c.Map) && c2.Key.Equal(c.Key)
}

// func (expr *Key) Eval(env *slang.Environment) (slang.Object, error) {
// 	ary, err := expr.Array.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	idx, err := expr.Key.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	return ary.Send(slang.MethodINDEX, idx)
// }

package expressions

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/objects"
)

type Index struct {
	Array slang.Expression
	Index slang.Expression
}

func NewIndex(ary, idx slang.Expression) slang.Expression {
	return &Index{
		Array: ary,
		Index: idx,
	}
}

func (c *Index) Type() slang.ExprType {
	return slang.ExprINDEX
}

func (c *Index) Equal(other slang.Expression) bool {
	c2, ok := other.(*Index)
	if !ok {
		return false
	}

	return c2.Array.Equal(c.Array) && c2.Index.Equal(c.Index)
}

func (expr *Index) Eval(env *slang.Environment) (slang.Object, error) {
	ary, err := expr.Array.Eval(env)
	if err != nil {
		return objects.NULL(), err
	}

	idx, err := expr.Index.Eval(env)
	if err != nil {
		return objects.NULL(), err
	}

	return ary.Send(slang.MethodINDEX, idx)
}

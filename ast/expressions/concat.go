package expressions

import "github.com/jamestunnell/slang"

type Concat struct {
	*Base

	StringExprs []slang.Expression `json:"stringExpressions"`
}

func NewConcat(exprs []slang.Expression) *Concat {
	return &Concat{
		Base:        NewBase(slang.ExprCONCAT),
		StringExprs: exprs,
	}
}

func (c *Concat) Equal(other slang.Expression) bool {
	c2, ok := other.(*Concat)
	if !ok {
		return false
	}

	return slang.ExpressionsEqual(c.StringExprs, c2.StringExprs)
}

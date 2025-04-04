package expressions

import "github.com/jamestunnell/slang"

type Concat struct {
	StringExprs []slang.Expression `json:"stringExpressions"`
}

func NewConcat(exprs ...slang.Expression) *Expression {
	return NewExpression(slang.ExprCONCAT, &Concat{StringExprs: exprs})
}

func (c *Concat) IsEqual(other Core) bool {
	c2, ok := other.(*Concat)
	if !ok {
		return false
	}

	return slang.ExpressionsEqual(c.StringExprs, c2.StringExprs)
}

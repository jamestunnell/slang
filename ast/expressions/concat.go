package expressions

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type Concat struct {
	StringExprs []*Expression `json:"stringExpressions"`
}

func NewConcat(exprs ...*Expression) *Expression {
	return NewExpression(slang.ExprCONCAT, &Concat{StringExprs: exprs})
}

func (c *Concat) IsEqual(other Core) bool {
	c2, ok := other.(*Concat)
	if !ok {
		return false
	}

	return slices.EqualFunc(c.StringExprs, c2.StringExprs, expressionsEqual)
}

func (c *Concat) Render(level int, w slang.CodeWriter) {
	w.WriteString(`"`)

	for _, expr := range c.StringExprs {
		if expr.Type == slang.ExprSTR {
			core, ok := expr.Core.(*Const[string])
			if ok {
				w.WriteString(core.Value)
			}
		} else {
			w.WriteString("${")

			expr.Render(level, w)

			w.WriteString("}")
		}
	}

	w.WriteString(`"`)
}

func expressionsEqual(a, b *Expression) bool {
	return a.IsEqual(b)
}

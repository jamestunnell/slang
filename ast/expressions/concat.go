package expressions

import (
	"reflect"
	"strings"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
	"github.com/jamestunnell/slang/objects"
)

type Concat struct {
	*Base

	StringExprs []slang.Expression `json:"stringExpressions"`
}

func NewConcat(exprs ...slang.Expression) *Concat {
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

func (c *Concat) Eval(env slang.Environment) (slang.Object, error) {
	var sb strings.Builder

	for _, stringExpr := range c.StringExprs {
		obj, err := stringExpr.Eval(env)
		if err != nil {
			return nil, err
		}

		str, ok := obj.(*objects.String)
		if !ok {
			err = customerrs.NewErrObjectType("*String", reflect.TypeOf(obj).String())

			return nil, err
		}

		sb.WriteString(str.Value)
	}

	return objects.NewString(sb.String()), nil
}

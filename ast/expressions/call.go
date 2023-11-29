package expressions

import (
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
)

type Call struct {
	*Base

	Function       slang.Expression            `json:"function"`
	PositionalArgs []slang.Expression          `json:"positionalArgs"`
	KeywordArgs    map[string]slang.Expression `json:"keywordArgs"`
}

func NewCall(
	fn slang.Expression,
	posArgs []slang.Expression,
	kwArgs map[string]slang.Expression,
) slang.Expression {
	return &Call{
		Base:           NewBase(slang.ExprCALL),
		Function:       fn,
		PositionalArgs: posArgs,
		KeywordArgs:    kwArgs,
	}
}

func (c *Call) Equal(other slang.Expression) bool {
	c2, ok := other.(*Call)
	if !ok {
		return false
	}

	if !c2.Function.Equal(c.Function) {
		return false
	}

	if !slices.EqualFunc(c.PositionalArgs, c2.PositionalArgs, expressionsEqual) {
		return false
	}

	if !maps.EqualFunc(c.KeywordArgs, c2.KeywordArgs, expressionsEqual) {
		return false
	}

	return true
}

// func (expr *Call) Eval(env *slang.Environment) (slang.Object, error) {
// 	obj, err := expr.Function.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	vals := make([]slang.Object, len(expr.Arguments))
// 	for i := 0; i < len(expr.Arguments); i++ {
// 		val, err := expr.Arguments[i].Eval(env)
// 		if err != nil {
// 			return objects.NULL(), err
// 		}

// 		vals[i] = val
// 	}

// 	return obj.Send(slang.MethodCALL, vals...)
// }

func expressionsEqual(a, b slang.Expression) bool {
	return a.Equal(b)
}

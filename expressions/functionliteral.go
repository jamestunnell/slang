package expressions

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/objects"
)

type FunctionLiteral struct {
	Params []*slang.Param
	Body   slang.Statement
}

func NewFunctionLiteral(
	params []*slang.Param, body slang.Statement) *FunctionLiteral {
	return &FunctionLiteral{
		Params: params,
		Body:   body,
	}
}

func (f *FunctionLiteral) Type() slang.ExprType {
	return slang.ExprFUNCTIONLITERAL
}

func (f *FunctionLiteral) Equal(other slang.Expression) bool {
	f2, ok := other.(*FunctionLiteral)
	if !ok {
		return false
	}

	return f.Body.Equal(f2.Body) &&
		slices.EqualFunc(f.Params, f2.Params, paramsEqual)

}

func paramsEqual(a, b *slang.Param) bool {
	return a.Name == b.Name && a.Type == b.Type
}

func (expr *FunctionLiteral) Eval(env *slang.Environment) (slang.Object, error) {
	f := objects.NewFunction(expr.Params, expr.Body, env)

	return f, nil
}

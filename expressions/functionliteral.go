package expressions

import (
	"golang.org/x/exp/slices"

	"github.com/akrennmair/slice"
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/objects"
)

type FunctionLiteral struct {
	Params []*Identifier
	Body   slang.Statement
}

func NewFunctionLiteral(
	params []*Identifier, body slang.Statement) *FunctionLiteral {
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
		slices.EqualFunc(f.Params, f2.Params, indentifiersEqual)

}

func indentifiersEqual(a, b *Identifier) bool {
	return a.Equal(b)
}

func (expr *FunctionLiteral) Eval(env *slang.Environment) (slang.Object, error) {
	params := slice.Map(expr.Params, func(ident *Identifier) string {
		return ident.Name
	})
	f := objects.NewFunction(params, expr.Body, env)

	return f, nil
}

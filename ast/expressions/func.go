package expressions

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
)

type Func struct {
	*Base
	*ast.Function
}

func NewFunc(fn *ast.Function) *Func {
	return &Func{
		Base:     NewBase(slang.ExprFUNC),
		Function: fn,
	}
}

func (f *Func) Equal(other slang.Expression) bool {
	f2, ok := other.(*Func)
	if !ok {
		return false
	}

	return f.Function.Equal(f2.Function)
}

// func (expr *Func) Eval(env *slang.Environment) (slang.Object, error) {
// 	f := objects.NewFunction(expr.Params, expr.Body, env)

// 	return f, nil
// }

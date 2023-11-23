package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
)

type Func struct {
	Name     string
	Function *ast.Function
}

func NewFunc(name string, f *ast.Function) *Func {
	return &Func{Name: name, Function: f}
}

func (f *Func) Type() slang.StatementType {
	return slang.StatementFUNC
}

func (f *Func) Equal(other slang.Statement) bool {
	f2, ok := other.(*Func)
	if !ok {
		return false
	}

	return f.Name == f2.Name && f.Function.Equal(f2.Function)
}

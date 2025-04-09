package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
)

type Var struct {
	Name         string                  `json:"name"`
	InitialValue *expressions.Expression `json:"initialValue"`
}

func NewVar(name string, initialVal *expressions.Expression) *Statement {
	core := &Var{Name: name, InitialValue: initialVal}

	return NewStatement(slang.StatementVAR, core)
}

func (f *Var) IsEqual(other Core) bool {
	f2, ok := other.(*Var)
	if !ok {
		return false
	}

	if !f.InitialValue.IsEqual(f2.InitialValue) {
		return false
	}

	return f.Name == f2.Name
}

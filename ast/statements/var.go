package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/tokens"
)

type Var struct {
	Name         string                  `json:"name"`
	InitialValue *expressions.Expression `json:"initialValue"`
}

func NewVar(name string, initialVal *expressions.Expression) *Statement {
	core := &Var{Name: name, InitialValue: initialVal}

	return NewStatement(slang.StatementVAR, core)
}

func (v *Var) GetName() (string, bool) {
	return v.Name, true
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

func (v *Var) Render(level int, w slang.CodeWriter) {
	w.WriteString(tokens.StrVAR)
	w.WriteString(" ")
	w.WriteString(v.Name)
	w.WriteString(" ")

	v.InitialValue.Render(level, w)
}

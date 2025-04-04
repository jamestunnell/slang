package statements

import (
	"github.com/jamestunnell/slang"
)

type Var struct {
	Name      string     `json:"name"`
	ValueType slang.Type `json:"type"`
}

func NewVar(name string, valueType slang.Type) *Statement {
	core := &Var{Name: name, ValueType: valueType}

	return NewStatement(slang.StatementVAR, core)
}

func (f *Var) IsEqual(other Core) bool {
	f2, ok := other.(*Var)
	if !ok {
		return false
	}

	if !slang.TypesEqual(f.ValueType, f2.ValueType) {
		return false
	}

	return f.Name == f2.Name
}

package statements

import (
	"github.com/jamestunnell/slang"
)

type Var struct {
	*Base

	Name      string     `json:"name"`
	ValueType slang.Type `json:"valueType"`
}

func NewVar(name string, valueType slang.Type) *Var {
	return &Var{
		Base:      NewBase(slang.StatementVAR),
		Name:      name,
		ValueType: valueType,
	}
}

func (f *Var) Equal(other slang.Statement) bool {
	f2, ok := other.(*Var)
	if !ok {
		return false
	}

	if !f.ValueType.IsEqual(f2.ValueType) {
		return false
	}

	return f.Name == f2.Name
}

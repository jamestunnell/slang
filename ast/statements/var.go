package statements

import (
	"github.com/jamestunnell/slang"
)

type Var struct {
	*Base

	Name      string `json:"name"`
	ValueType string `json:"valueType"`
}

func NewVar(name, valueType string) *Var {
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

	return f.Name == f2.Name && f.ValueType == f2.ValueType
}

package statements

import (
	"github.com/jamestunnell/slang"
)

type Const struct {
	Name  string           `json:"name"`
	Value slang.Expression `json:"value"`
}

func NewConst(name string, val slang.Expression) *Statement {
	core := &Const{Name: name, Value: val}

	return NewStatement(slang.StatementCONST, core)
}

func (f *Const) IsEqual(other Core) bool {
	f2, ok := other.(*Const)
	if !ok {
		return false
	}

	if !f.Value.Equal(f2.Value) {
		return false
	}

	return f.Name == f2.Name
}

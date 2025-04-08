package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
)

type Const struct {
	Name  string                  `json:"name"`
	Value *expressions.Expression `json:"value"`
}

func NewConst(name string, val *expressions.Expression) *Statement {
	core := &Const{Name: name, Value: val}

	return NewStatement(slang.StatementCONST, core)
}

func (f *Const) IsEqual(other Core) bool {
	f2, ok := other.(*Const)
	if !ok {
		return false
	}

	if !f.Value.IsEqual(f2.Value) {
		return false
	}

	return f.Name == f2.Name
}

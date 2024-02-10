package statements

import (
	"github.com/jamestunnell/slang"
)

type Const struct {
	*Base

	Name  string           `json:"name"`
	Value slang.Expression `json:"value"`
}

func NewConst(name string, val slang.Expression) *Const {
	return &Const{
		Base:  NewBase(slang.StatementCONST),
		Name:  name,
		Value: val,
	}
}

func (f *Const) Equal(other slang.Statement) bool {
	f2, ok := other.(*Const)
	if !ok {
		return false
	}

	if !f.Value.Equal(f2.Value) {
		return false
	}

	return f.Name == f2.Name
}

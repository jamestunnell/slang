package statements

import (
	"github.com/jamestunnell/slang"
)

type Field struct {
	*Base

	Name      string `json:"name"`
	ValueType string `json:"valueType"`
}

func NewField(name, valueType string) *Field {
	return &Field{
		Base:      NewBase(slang.StatementFIELD),
		Name:      name,
		ValueType: valueType,
	}
}

func (f *Field) Equal(other slang.Statement) bool {
	f2, ok := other.(*Field)
	if !ok {
		return false
	}

	return f.Name == f2.Name && f.ValueType == f2.ValueType
}

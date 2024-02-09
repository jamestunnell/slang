package statements

import (
	"github.com/jamestunnell/slang"
)

type Field struct {
	*Base

	Name      string     `json:"name"`
	ValueType slang.Type `json:"valueType"`
}

func NewField(name string, valueType slang.Type) *Field {
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

	if !f.ValueType.IsEqual(f2.ValueType) {
		return false
	}

	return f.Name == f2.Name
}

package statements

import (
	"github.com/jamestunnell/slang"
)

type ClassField struct {
	*Base

	Name      string     `json:"name"`
	ValueType slang.Type `json:"valueType"`
}

func NewClassField(name string, valueType slang.Type) *ClassField {
	return &ClassField{
		Base:      NewBase(slang.StatementCLASSFIELD),
		Name:      name,
		ValueType: valueType,
	}
}

func (f *ClassField) Equal(other slang.Statement) bool {
	f2, ok := other.(*ClassField)
	if !ok {
		return false
	}

	if !f.ValueType.IsEqual(f2.ValueType) {
		return false
	}

	return f.Name == f2.Name
}

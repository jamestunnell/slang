package statements

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type StructField struct {
	*Base

	Names     []string   `json:"names"`
	ValueType slang.Type `json:"valueType"`
}

func NewStructField(names []string, valueType slang.Type) *StructField {
	return &StructField{
		Base:      NewBase(slang.StatementCLASSFIELD),
		Names:     names,
		ValueType: valueType,
	}
}

func (f *StructField) Equal(other slang.Statement) bool {
	f2, ok := other.(*StructField)
	if !ok {
		return false
	}

	if !f.ValueType.IsEqual(f2.ValueType) {
		return false
	}

	if len(f.Names) != len(f2.Names) {
		return false
	}

	for _, name := range f.Names {
		if !slices.Contains(f2.Names, name) {
			return false
		}
	}

	return true
}

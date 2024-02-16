package types

import (
	"fmt"

	"github.com/jamestunnell/slang"
)

type ArrayType struct {
	ValueType slang.Type
}

func NewArrayType(valType slang.Type) *ArrayType {
	return &ArrayType{ValueType: valType}
}

func (typ *ArrayType) String() string {
	return fmt.Sprintf("[]%s", typ.ValueType)
}

func (typ *ArrayType) IsEqual(other slang.Type) bool {
	typ2, ok := other.(*ArrayType)
	if !ok {
		return false
	}

	return typ.ValueType.IsEqual(typ2.ValueType)
}

func (typ *ArrayType) IsArray() bool {
	return true
}

func (typ *ArrayType) IsMap() bool {
	return false
}

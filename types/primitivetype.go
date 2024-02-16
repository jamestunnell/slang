package types

import (
	"github.com/jamestunnell/slang"
)

type PrimitiveType struct {
	Name string
}

func NewPrimitiveType(name string) *PrimitiveType {
	return &PrimitiveType{
		Name: name,
	}
}

func (typ *PrimitiveType) String() string {
	return typ.Name
}

func (typ *PrimitiveType) IsEqual(other slang.Type) bool {
	typ2, ok := other.(*PrimitiveType)
	if !ok {
		return false
	}

	return typ.Name == typ2.Name
}

func (typ *PrimitiveType) IsArray() bool {
	return false
}

func (typ *PrimitiveType) IsMap() bool {
	return false
}

package types

import (
	"strings"

	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
)

type BasicType struct {
	NameParts []string
}

func NewBasicType(nameParts ...string) *BasicType {
	return &BasicType{
		NameParts: nameParts,
	}
}

func (typ *BasicType) String() string {
	return strings.Join(typ.NameParts, ".")
}

func (typ *BasicType) IsEqual(other slang.Type) bool {
	typ2, ok := other.(*BasicType)
	if !ok {
		return false
	}

	return slices.Equal(typ.NameParts, typ2.NameParts)
}

func (typ *BasicType) IsArray() bool {
	return false
}

func (typ *BasicType) IsMap() bool {
	return false
}

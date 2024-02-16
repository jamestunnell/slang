package types

import (
	"fmt"

	"github.com/jamestunnell/slang"
)

type MapType struct {
	KeyType, ValueType slang.Type
}

func NewMapType(keyType, valType slang.Type) *MapType {
	return &MapType{KeyType: keyType, ValueType: valType}
}

func (typ *MapType) String() string {
	return fmt.Sprintf("[%s]%s", typ.KeyType, typ.ValueType)
}

func (typ *MapType) IsEqual(other slang.Type) bool {
	typ2, ok := other.(*MapType)
	if !ok {
		return false
	}

	return typ.KeyType.IsEqual(typ2.KeyType) && typ.ValueType.IsEqual(typ2.ValueType)
}

func (typ *MapType) IsArray() bool {
	return false
}

func (typ *MapType) IsMap() bool {
	return true
}

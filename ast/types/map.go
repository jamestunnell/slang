package types

import (
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/tokens"
)

type Map struct {
	KeyType, ValueType slang.Type
}

func NewMap(keyType, valType slang.Type) *Type {
	core := &Map{ValueType: valType}

	return NewType(slang.TypeMAP, core)
}

func (typ *Map) String() string {
	return fmt.Sprintf("%s<%s,%s>", tokens.StrMAP, typ.KeyType, typ.ValueType)
}

func (typ *Map) IsEqual(other Core) bool {
	typ2, ok := other.(*Map)
	if !ok {
		return false
	}

	if !typ.KeyType.IsEqual(typ2.KeyType) {
		return false
	}

	if !typ.ValueType.IsEqual(typ2.ValueType) {
		return false
	}

	return true
}

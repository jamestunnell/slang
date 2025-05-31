package types

import (
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/tokens"
)

type Array struct {
	ValueType slang.Type
}

func NewArray(valType slang.Type) *Type {
	core := &Array{ValueType: valType}

	return NewType(slang.TypeARRAY, core)
}

func (typ *Array) String() string {
	return fmt.Sprintf("%s<%s>", tokens.StrARY, typ.ValueType)
}

func (typ *Array) IsEqual(other Core) bool {
	typ2, ok := other.(*Array)
	if !ok {
		return false
	}

	if !typ.ValueType.IsEqual(typ2.ValueType) {
		return false
	}

	return true
}

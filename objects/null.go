package objects

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/types"
)

type Null struct {
	*Base
}

const (
	StrNULL       = "null"
	nullClassName = "Null"
)

var nullClass *BuiltInClass

func init() {
	nullClass = NewBuiltInClass(
		types.NewPrimitiveType(nullClassName),
		map[string]slang.MethodFunc{},
	)
}

var null = &Null{
	Base: NewBase(nullClass),
}

func NULL() *Null {
	return null
}

func (obj *Null) IsEqual(other slang.Object) bool {
	_, ok := other.(*Null)

	return ok
}

func (obj *Null) Inspect() string {
	return StrNULL
}

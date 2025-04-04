package types

import (
	"github.com/jamestunnell/slang"
)

type Bool struct{}
type Err struct{}
type Flt struct{}
type Int struct{}
type Str struct{}

func NewBool() *Type {
	return NewType(slang.TypeBOOLEAN, &Bool{})
}

func NewErr() *Type {
	return NewType(slang.TypeERROR, &Err{})
}

func NewFlt() *Type {
	return NewType(slang.TypeFLOAT, &Flt{})
}

func NewInt() *Type {
	return NewType(slang.TypeINTEGER, &Int{})
}

func NewStr() *Type {
	return NewType(slang.TypeSTRING, &Str{})
}

func (t *Bool) String() string {
	return slang.StrTypeBOOLEAN
}

func (t *Bool) IsEqual(other Core) bool {
	_, ok := other.(*Bool)

	return ok
}

func (t *Err) String() string {
	return slang.StrTypeERROR
}

func (t *Err) IsEqual(other Core) bool {
	_, ok := other.(*Err)

	return ok
}

func (t *Flt) String() string {
	return slang.StrTypeFLOAT
}

func (t *Flt) IsEqual(other Core) bool {
	_, ok := other.(*Flt)

	return ok
}

func (t *Int) String() string {
	return slang.StrTypeINTEGER
}

func (t *Int) IsEqual(other Core) bool {
	_, ok := other.(*Int)

	return ok
}

func (t *Str) String() string {
	return slang.StrTypeSTRING
}

func (t *Str) IsEqual(other Core) bool {
	_, ok := other.(*Str)

	return ok
}

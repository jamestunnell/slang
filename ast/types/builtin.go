package types

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/tokens"
)

type Bool struct{}
type Empty struct{}
type Err struct{}
type Flt struct{}
type Int struct{}
type Str struct{}

func NewBool() *Type {
	return NewType(slang.TypeBOOLEAN, &Bool{})
}

func NewEmpty() *Type {
	return NewType(slang.TypeEMPTY, &Empty{})
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

func (t *Bool) Render(w slang.CodeWriter) {
	w.WriteString(tokens.StrBOOL)
}

func (t *Empty) String() string {
	return "???"
}

func (t *Empty) IsEqual(other Core) bool {
	_, ok := other.(*Empty)

	return ok
}

func (t *Err) String() string {
	return tokens.StrERR
}

func (t *Err) IsEqual(other Core) bool {
	_, ok := other.(*Err)

	return ok
}

func (t *Flt) String() string {
	return tokens.StrFLOAT
}

func (t *Flt) IsEqual(other Core) bool {
	_, ok := other.(*Flt)

	return ok
}

func (t *Int) String() string {
	return tokens.StrINT
}

func (t *Int) IsEqual(other Core) bool {
	_, ok := other.(*Int)

	return ok
}

func (t *Str) String() string {
	return tokens.StrSTR
}

func (t *Str) IsEqual(other Core) bool {
	_, ok := other.(*Str)

	return ok
}

package ast

import "github.com/jamestunnell/slang"

type Param = slang.NameType

func NewParam(name string, typ slang.Type) *Param {
	return &slang.NameType{
		Name: name,
		Type: typ,
	}
}

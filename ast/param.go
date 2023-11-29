package ast

import "github.com/jamestunnell/slang"

type Param = slang.NameType

func NewParam(name, typ string) *Param {
	return &slang.NameType{
		Name: name,
		Type: typ,
	}
}

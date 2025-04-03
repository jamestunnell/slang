package ast

import "github.com/jamestunnell/slang"

func NewParam(name string, typ slang.Type) slang.Param {
	return NewNameType(name, typ)
}

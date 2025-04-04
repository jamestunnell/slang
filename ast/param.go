package ast

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/types"
)

func NewParam(name string, typ *types.Type) slang.Param {
	return NewNameType(name, typ)
}

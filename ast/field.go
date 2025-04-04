package ast

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/types"
)

func NewField(name string, typ *types.Type) slang.Field {
	return NewNameType(name, typ)
}

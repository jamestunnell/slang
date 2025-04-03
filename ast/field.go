package ast

import "github.com/jamestunnell/slang"

func NewField(name string, typ slang.Type) slang.Field {
	return NewNameType(name, typ)
}

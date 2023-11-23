package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
)

type Class struct {
	Name  string
	Class *ast.Class
}

func NewClass(name string, c *ast.Class) *Class {
	return &Class{Name: name, Class: c}
}

func (c *Class) Type() slang.StatementType {
	return slang.StatementFUNC
}

func (c *Class) Equal(other slang.Statement) bool {
	c2, ok := other.(*Class)
	if !ok {
		return false
	}

	return c.Name == c2.Name && c.Class.Equal(c2.Class)
}

package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
)

type Class struct {
	*Base

	Name  string     `json:"name"`
	Class *ast.Class `json:"class"`
}

func NewClass(name string, c *ast.Class) *Class {
	return &Class{
		Base:  NewBase(slang.StatementFUNC),
		Name:  name,
		Class: c,
	}
}

func (c *Class) Equal(other slang.Statement) bool {
	c2, ok := other.(*Class)
	if !ok {
		return false
	}

	return c.Name == c2.Name && c.Class.Equal(c2.Class)
}

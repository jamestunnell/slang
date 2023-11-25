package statements

import (
	"github.com/jamestunnell/slang"
)

type Class struct {
	*Base

	Name       string            `json:"name"`
	Comment    string            `json:"comment"`
	Statements []slang.Statement `json:"statements"`
}

func NewClass(name, comment string, stmts ...slang.Statement) *Class {
	return &Class{
		Base:       NewBase(slang.StatementCLASS),
		Name:       name,
		Comment:    comment,
		Statements: stmts,
	}
}

func (c *Class) Equal(other slang.Statement) bool {
	c2, ok := other.(*Class)
	if !ok {
		return false
	}

	return c.Name == c2.Name && slang.StatementsEqual(c.Statements, c2.Statements)
}

package statements

import (
	"github.com/jamestunnell/slang"
)

type Struct struct {
	*Base

	Name       string            `json:"name"`
	Comment    string            `json:"comment"`
	Statements []slang.Statement `json:"statements"`
}

func NewStruct(name, comment string, stmts ...slang.Statement) *Struct {
	return &Struct{
		Base:       NewBase(slang.StatementCLASS),
		Name:       name,
		Comment:    comment,
		Statements: stmts,
	}
}

func (c *Struct) Equal(other slang.Statement) bool {
	c2, ok := other.(*Struct)
	if !ok {
		return false
	}

	return c.Name == c2.Name && slang.StatementsEqual(c.Statements, c2.Statements)
}

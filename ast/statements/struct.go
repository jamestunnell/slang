package statements

import (
	"github.com/jamestunnell/slang"
)

type Struct struct {
	*Base

	Name       string            `json:"name"`
	Statements []slang.Statement `json:"statements"`
}

func NewStruct(name string, stmts ...slang.Statement) *Struct {
	return &Struct{
		Base:       NewBase(slang.StatementSTRUCT),
		Name:       name,
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

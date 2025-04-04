package statements

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
)

type Struct struct {
	Name   string        `json:"name"`
	Fields []slang.Field `json:"fields"`
}

func NewStruct(name string, fields ...slang.Field) *Statement {
	core := &Struct{Name: name, Fields: fields}

	return NewStatement(slang.StatementSTRUCT, core)
}

func (c *Struct) IsEqual(other Core) bool {
	c2, ok := other.(*Struct)
	if !ok {
		return false
	}

	if c.Name != c2.Name {
		return false
	}

	return slices.EqualFunc(c.Fields, c2.Fields, slang.NameTypesEqual)
}

func (s *Struct) GetName() string {
	return s.Name
}

func (s *Struct) GetFields() []slang.Field {
	return s.Fields
}

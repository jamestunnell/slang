package statements

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type Struct struct {
	*Base

	Name   string        `json:"name"`
	Fields []slang.Field `json:"fields"`
}

func NewStruct(name string, fields ...slang.Field) *Struct {
	return &Struct{
		Base:   NewBase(slang.StatementSTRUCT),
		Name:   name,
		Fields: fields,
	}
}

func (c *Struct) Equal(other slang.Statement) bool {
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

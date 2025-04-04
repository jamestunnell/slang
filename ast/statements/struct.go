package statements

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/types"
)

type Struct struct {
	Name   string            `json:"name"`
	Fields []*types.NameType `json:"fields"`
}

func NewStruct(name string, fields ...*types.NameType) *Statement {
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

	return slices.EqualFunc(c.Fields, c2.Fields, nameTypesEqual)
}

func (s *Struct) GetName() string {
	return s.Name
}

func (s *Struct) GetFields() []slang.Field {
	fields := make([]slang.Field, len(s.Fields))

	for i, f := range s.Fields {
		fields[i] = f
	}

	return fields
}
